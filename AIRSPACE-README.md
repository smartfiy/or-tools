# Airspace Golang Wrapper Repo for OR-Tools

This is a fork of [Google's OR-Tools repo](https://github.com/google/or-tools).
It has Go bindings and binaries for use with Go projects.

## Mac Setup for Local Go Development
 1. Download binaries for Mac:
    `https://github.com/AirspaceTechnologies/or-tools/releases/download/v7.6-go1.14.1/or-tools_MacOsX-10.15.4_v7.6.7700.tar.gz`
 1. Install/extract to rpath (e.g. `/usr/local`)

## Mac Setup for Local OR-tools Development (Extending/Wrapping OR-tools)
 1. Install XCode:
    `xcode-select install`
 1. Install C++ tools:
    `brew install cmake wget pkg-config`
 1. Install SWIG 4.0.1:
    `brew install swig`
 1. Install protobuf for Go:
    `go get github.com/golang/protobuf/protoc-gen-go@v1.3`
 1. Clone Airspace OR-tools:
    `git clone git@github.com:AirspaceTechnologies/or-tools.git`
 1. Make third party:
    `make third_party`
 1. Make go:
    `make clean_go go`
 1. Make go tests:
    `make test_go`

## Build and Release
 1. Follow steps above (`Mac Setup for Local OR-tools Development`)
 1. Make Debian binary archive (takes ~45 mins, uses Docker to build everything from scratch):
    `make --directory=makefiles debian_go_export`
 1. Make local Mac binary archive:
    `make golib_archive`
 1. Log into Github and create a release with the resulting binaries

## TODO
 1. Make `IntVar`->`IntExpr` casting cleaner
