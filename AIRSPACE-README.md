# Airspace Golang Wrapper Repo for OR-Tools

This is a fork of [Google's OR-Tools repo](https://github.com/google/or-tools).
It has Go bindings and binaries for use with Go projects.

## Mac Setup for Local Go Development
 1. Download binaries for Mac:
    `https://github.com/AirspaceTechnologies/or-tools/releases/download/v7.6-go1.14.1/or-tools_MacOsX-10.15.4_v7.6.7700.tar.gz`
 1. Install/extract to rpath:
    `tar -xf or-tools_MacOsX-10.15.4_v7.6.7700.tar.gz --strip 1 -C /usr/local`
 1. Export `DYLD_LIBRARY_PATH` if necessary:
    `export DYLD_LIBRARY_PATH=$DYLD_LIBRARY_PATH:/usr/local/lib`

## Mac Setup for Local OR-tools Development (Extending/Wrapping OR-tools)
 1. Install XCode:
    `xcode-select install`
 1. Install C++ tools:
    `brew install cmake wget pkg-config`
 1. Install SWIG 4.0.1:
    `brew install swig`
 1. Install protobuf for Go:
    `go get -u github.com/golang/protobuf/protoc-gen-go@v1.5.2`
 1. Clone Airspace OR-tools:
    `git clone git@github.com:AirspaceTechnologies/or-tools.git`
 1. Make third party:
    `make clean_third_party third_party`
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

## Update Fork from Upstream
 1. Configure git remote pointing to upstream or-tools repo:
    `git remote add upstream git@github.com:google/or-tools.git`
 1. Fetch upstream:
    `git fetch upstream`
 1. Checkout fork's `stable` branch and pull:
    `git checkout stable && git pull`
 1. Merge changes from upstream `stable` branch:
    `git merge upstream/stable`
 1. Push fork's `stable` branch:
    `git push`
 1. Checkout fork's `airspace` branch and pull:
    `git checkout airspace && git pull`
 1. Merge changes from `stable` to `airspace` branch:
    `git merge stable`
 1. Push fork's `airspace` branch:
    `git push`
 1. Optionally build and release using steps above

## TODO
 1. Make `IntVar`->`IntExpr` casting cleaner
