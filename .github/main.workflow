workflow "Push" {
  on = "push"
  resolves = ["Push: Go 1.12", "Push: Go master"]
}

workflow "PR" {
  on = "pull_request"
  resolves = ["PR: Go 1.12", "PR: Go master"]
}

action "Push: Go 1.12" {
  uses = "./.github/tests-go-1.12"
}

action "Push: Go master" {
  uses = "./.github/tests-go-master"
}

action "PR: Go 1.12" {
  uses = "./.github/tests-go-1.12"
}

action "PR: Go master" {
  uses = "./.github/tests-go-master"
}
