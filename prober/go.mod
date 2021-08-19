module github.com/devguyio/playground/prober

go 1.15

require (
	k8s.io/apimachinery v0.20.2
	knative.dev/networking v0.0.0-20210125050654-94433ab7f620
	knative.dev/pkg v0.0.0-20210119162123-1bbf0a6436c3
)

replace knative.dev/networking => /home/abd4lla/Workspace/redhat/src/knative.dev/networking
