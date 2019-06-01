
workflow "Test on push" {
  on = "push"
  resolves = ["Test"]
}

action "Test" {
  uses = "actions-contrib/go@master"
  args = "test -race ./..."
}

workflow "Publish on release" {
  on = "release"
  resolves = ["Publish"]
}

action "Publish" {
  needs = [
    "Test",
  ]
  uses = "altipla-consulting/altipla.actions/go-release@master"
  env = {
    BINARY_FOLDER = "./cmd/validate-proto-http"
  }
  secrets = ["GITHUB_TOKEN"]
}
