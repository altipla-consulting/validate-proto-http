package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	log "github.com/sirupsen/logrus"
	"google.golang.org/genproto/googleapis/api/annotations"
	"libs.altipla.consulting/errors"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(errors.Stack(err))
	}
}

func run() error {
	if os.Getenv("BUILD_ID") != "" {
		// We don't need colors in Jenkins
		log.SetFormatter(&log.TextFormatter{DisableColors: true})
	}

	log.Info("Searching for services to validate")
	fn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Trace(err)
		}

		if info.IsDir() || filepath.Ext(path) != ".pb" {
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return errors.Trace(err)
		}

		files := new(descriptor.FileDescriptorSet)
		if err := proto.Unmarshal(content, files); err != nil {
			return errors.Trace(err)
		}

		log.WithField("path", path).Info("Validate services in file")

		for _, file := range files.File {
			for _, service := range file.Service {
				log.WithFields(log.Fields{
					"service": service.GetName(),
					"file":    file.GetName(),
				}).Info("... validate service")

				for _, method := range service.GetMethod() {
					ext, err := proto.GetExtension(method.GetOptions(), annotations.E_Http)
					if err != nil {
						return errors.Trace(err)
					}
					rule := ext.(*annotations.HttpRule)

					var path string
					switch r := rule.Pattern.(type) {
					case *annotations.HttpRule_Get:
						path = r.Get
					case *annotations.HttpRule_Put:
						path = r.Put
					case *annotations.HttpRule_Post:
						path = r.Post
					case *annotations.HttpRule_Delete:
						path = r.Delete
					default:
						return errors.Errorf("unknown http rule pattern: %#v", rule.Pattern)
					}

					log.WithFields(log.Fields{
						"name": method.GetName(),
						"path": path,
					}).Info("...... validate method")
				}
			}
		}

		return nil
	}
	if err := filepath.Walk(".", fn); err != nil {
		return errors.Trace(err)
	}

	return nil
}
