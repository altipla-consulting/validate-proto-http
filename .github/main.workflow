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
  uses = "ngs/go-release.action@latest"
  secrets = ["GITHUB_TOKEN"]
}
