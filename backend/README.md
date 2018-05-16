# C4 Backend API Server

The backend API server is a go webserver that leverages gRPC.

## Requirements

*   go (version 1.10)
*   [dep](https://github.com/golang/dep) for dependency management

## Installation

*   Make sure the project is in your $GOPATH and is located @ github.com/maxmcd/c4
*   Initialize all git submodules
    *   `git submodule update --init --recursive`
*   Build all protos, binaries and docker images
    *   `make all`
*   Bring up the API
    *   `make up`

## Directories

*   `/pkg` for reusable packages
*   `/cmd` for binaries
*   `/third_party` for third party dependencies
*   `.docker/` for docker related files
