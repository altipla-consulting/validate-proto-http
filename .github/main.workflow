
workflow "Build on push" {
  on = "push"
  resolves = ["Build"]
}

action "Build" {
  uses = "actions-contrib/go@master"
  args = "build ./..."
}
