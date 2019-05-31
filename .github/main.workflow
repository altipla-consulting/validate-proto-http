
workflow "Build on push" {
  on = "push"
  resolves = ["Build"]
}

action "Build" {
  uses = "actions-contrib/go@master"
  args = "build ./..."
}

workflow "Publish on release" {
  on = "release"
  resolves = ["Publish"]
}

action "Publish" {
  uses = "altipla-consulting/altipla.actions/go-release@master"
  env = {
    BINARY_NAME = "validate-proto-http"
  }
  secrets = ["GITHUB_TOKEN"]
}
