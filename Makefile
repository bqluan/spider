package_root=github.com/bqluan/spider
mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_path := $(dir $(mkfile_path))

build:
	go build $(package_root)/...

install:
	go install $(package_root)/...

run: install
	spider
