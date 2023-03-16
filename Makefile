all: build
.PHONY: all

BINDIR ?= _output

generate_exes: $(BINDIR)/openapi-gen \

$(BINDIR)/openapi-gen:
	go build -o $@ k8s.io/code-generator/cmd/openapi-gen

# Regenerate all files if the gen exes changed or any "types.go" files changed
generate_files: generate_exes
  # generate apiserver deps
	hack/update-apiserver-gen.sh

build:
	go build -mod=vendor -o _output/server github.com/qiujian16/virtual-manifestwork/cmd/server