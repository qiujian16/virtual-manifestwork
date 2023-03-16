//go:build tools
// +build tools

// go mod won't pull in code that isn't depended upon, but we have some code we don't depend on from code that must be included
// for our build to work.
package dependencymagnet

import (
	_ "k8s.io/code-generator"
	_ "k8s.io/code-generator/cmd/go-to-protobuf/protoc-gen-gogo"
)
