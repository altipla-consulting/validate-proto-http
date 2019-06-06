package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	log "github.com/sirupsen/logrus"
	"google.golang.org/genproto/googleapis/api/annotations"
	"libs.altipla.consulting/errors"
)

var pathRe = regexp.MustCompile(
	// Start anchor.
	`^` +

		// Repeat every path segment at least one time.
		`(` +

		// Start with a slash.
		`/(` +

		// Simple segment with text only.
		`([a-z0-9-]+)|` +

		// Parameter with optional text only in both sides.
		`(([a-z0-9-]+:)?{[a-z0-9_]+}(:[a-zA-Z0-9-]+)?)|` +

		// Text only with a colon in the middle.
		`([a-z0-9-]+:[a-zA-Z0-9-]+)` +

		`)` +

		`)+` +

		// End anchor.
		`$`,
)

func main() {
	if err := run(); err != nil {
		log.Fatal(errors.Stack(err))
	}
}

func init() {
	if os.Getenv("BUILD_ID") != "" {
		// We don't need colors in Jenkins
		log.SetFormatter(&log.TextFormatter{DisableColors: true})
	} else {
		log.SetFormatter(&log.TextFormatter{ForceColors: true})
	}
}

func run() error {
	return errors.Trace(checkFiles("."))
}

func checkFiles(root string) error {
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
			mappings := make(map[string]bool)
			for _, service := range file.Service {
				log.WithFields(log.Fields{
					"service": service.GetName(),
					"file":    file.GetName(),
				}).Info("... validate service")

				for _, method := range service.GetMethod() {
					ext, err := proto.GetExtension(method.GetOptions(), annotations.E_Http)
					if err != nil {
						if err == proto.ErrMissingExtension {
							continue
						}

						return errors.Trace(err)
					}
					rule := ext.(*annotations.HttpRule)

					var path, methodName string
					switch r := rule.Pattern.(type) {
					case *annotations.HttpRule_Get:
						path = r.Get
						methodName = "get"
					case *annotations.HttpRule_Put:
						path = r.Put
						methodName = "put"
					case *annotations.HttpRule_Post:
						path = r.Post
						methodName = "post"
					case *annotations.HttpRule_Delete:
						path = r.Delete
						methodName = "delete"
					default:
						return errors.Errorf("unknown http rule pattern: %#v", rule.Pattern)
					}

					log.WithFields(log.Fields{
						"name": method.GetName(),
						"path": path,
					}).Info("...... validate method")

					if !pathRe.MatchString(path) {
						return errors.Errorf("bad http mapping: %v", path)
					}

					fullpath := methodName + "::" + path
					if mappings[fullpath] {
						return errors.Errorf("repeated http mapping: %v", path)
					}
					mappings[fullpath] = true
				}
			}
		}

		return nil
	}
	if err := filepath.Walk(root, fn); err != nil {
		return errors.Trace(err)
	}

	return nil
}
