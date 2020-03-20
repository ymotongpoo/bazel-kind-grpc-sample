# Bazel, Kind and gRPC sample

This is a sample application that demonstrates:

* Bazel
  * Go project with gRPC directory and BUILD structure
  * Container creation and push to the container registory
  * Deploying containers to Kubernetes cluster
* gRPC
  * `protoc` for Go applications
  * simple gRPC server and client implementations

## prerequisites

* [Bazel](https://bazel.build/)
* [Kind](https://kind.sigs.k8s.io/) for local testing
  * [Docker](https://www.docker.com/)
  * Go 1.11+
* [Google Kubernetes Engine](https://cloud.google.com/kubernetes-engine) (optional)
  * In the case you'd like to run this sample on GKE

## How to run

### 1. Install Bazel into your system

This sample requires [Bazel](https://bazel.build/) to be installed in your environment.
Follow [the installation steps](https://docs.bazel.build/versions/2.2.0/install.html) and make sure to have Bazel.

The author's environment is as the followings:

* OS: Linux (Arch Linux and Debian Linux)
* Bazel: 2.2.0

### 2. Install Kind or prepare Kubernetes instance to deploy apps

This sample builds containers of a client and a server that communicate with each other in gRPC, and runs them on Kubernetes in coopereted manner.

For ease to use, you need one of the following environments:

* Kind (on Docker)
* Google Kubernetes Engine
* ...or whatever Kubernetes cluster you can prepare

#### 2.1 Install Kind

```
$ GO111MODULE="on" go get sigs.k8s.io/kind@v0.7.0 && kind create cluster
```"

### 3. Run Bazel target

To be documented...

```
$ bazel run :sample
```
