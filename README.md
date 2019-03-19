# librespot-golang

## Introduction

This package is a rework of the fork of the [librespot-golang] library. All credits go to the original developers.

Main differences:

- Go [module](https://github.com/golang/go/wiki/Modules) support
- Proper import paths as per convention
- Better project layout wich adheres to the Go [project-layout](https://github.com/golang-standards/project-layout) community pattern
- Release tags per [SemVer 2.0](https://semver.org/) standard
- Small code and Protobuf refactoring, linter support, CI builds, etc.

## Description

[librespot-golang] is an opensource Golang library based on the [librespot](https://github.com/plietar/librespot) project,
allowing you to control Spotify Connect devices, get metadata, and play music. It has itself been based on
[SpotControl](https://github.com/badfortrains/spotcontrol), and its main goal is to provide a suitable replacement
for the defunct libspotify.

This is still highly experimental and in development. Do not use it in production projects yet, as the API is incomplete
and subject to heavy changes.

## Installation

This package can be installed using `go get`:

```shell
go get github.com/nezorflame/librespot-golang
```

or, if you're using Go 1.12+, using module-aware `go get`:

```shell
go get github.com/nezorflame/librespot-golang@VERSION
```

where `VERSION` is the desired version tag.

## Usage

To use the package look at the example micro-controller (for Spotify Connect), or micro-client (for audio playback).

## Building for mobile

The package `librespotmobile` contains bindings suitable for use with Gomobile, which lets you use a subset of the
librespot library on Android and iOS.

To get started, install gomobile, and simply run (for Android):

```shell
cd /path/to/librespot-golang
export GOPATH=$(pwd)
gomobile init -ndk /path/to/android-ndk
gomobile bind librespotmobile
```

This will build you a file called `librespotmobile.aar` which you can include in your Android Studio project.

## To-Do's

* Handling disconnections, timeouts, etc (overall failure tolerance)
* Playlist management
* Spotify Radio support

[librespot-golang]: https://github.com/librespot-org/librespot-golang
