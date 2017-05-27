# Go-kit Foundation

This is an example base implementation of a go-kit project.

## Pre-requisites

### Docker Dev (preferred)
For Docker, we have chosen to employ a two-step build and run container procedure
for our Go projects.

To use the gokit-base docker container you'll first need to build a container
using the scripts supplied in the `docker/` folder in this repository. On your first
run you'll need to execute `./docker/build.sh -i -v latest`.

Finally, `docker-compose up`

### Local Dev
To get the gokit-base project up and running you'll need to have a few things installed beforehand:
1. [Install Go](https://golang.org/doc/install)
2. [Install Glide](https://github.com/Masterminds/glide#install)

Now simply do a git clone of the project into your GOPATH and run `glide install` to get all of the
required dependencies. Finally, run `go build` and `./gokit-base` to start the listening server!

## Repository and Project Structure

We are using Glide to manage our vendor dependencies. This unfortunately forces us
into a non-standard Go repository layout. Granted, we believe we have established some
best practices that can be referenced in this sample application and our other projects in
the "Mentat" namespace.

```
app
│   .gitignore
│   docker-compose.yml
│   glide.yaml
│   glide.lock
│   Jenkinsfile
│   main.go
│   README.md
│
└───config
│   │   config.go
└───docker
│   │   build.sh
│   └───app
│   │    │   Dockerfile
│   │    └───bin
│   │    │   │   .gitkeep
│   │    │   │   app_binary
│   │    │   │   ...
│   │    └───resources
│   │    │   │   env.default
│   │    │   │   init.sh
│   │    │   │   ...
│
└───domain_object
│   │   domain_object.go
│   │   endpoint.go
│   │   instrumenting.go
│   │   logging.go
│   │   service.go
│   │   transport.go
└───health
│   │   endpoint.go
│   │   instrumenting.go
│   │   transport.go
└───vendor
```

### Enumeration of the Project Structure

- [docker-compose.yml](https://docs.docker.com/compose) describes the composition of application container(s) and
dependencies as well as any networking considerations.
- [glide.yml and glide.lock](https://glide.readthedocs.io/en/latest/glide.yaml/) provide a manifest of library dependencies for your project.
- [Jenkinsfile](https://jenkins.io/doc/book/pipeline/jenkinsfile) contains Jenkins 2 pipeline configuration scripts written in Groovy for CI/CD.
- The `config/` folder contains your specific application configuration code. This may be something as simple
as an Environment struct read from Consul or MySQL connection details (also read from Consul). We are currently
using [Viper]() to manage configuration and an example can be found in this project.
- The `docker/` folder contains your application container [Dockerfile](), container build and runtime configuration in `build.sh` and
`init.sh`, and an ignored binary folder that will be the target of the container build script.
- The `domain_object/` folder is a sample folder structure for an application business object. In the example app
this can be seen in the `users/` folder. Further discussion on go-kit idioms such as `endpoint.go` will follow.
- The `health/` folder includes an example basic HTTP health endpoint. This folder may grow to contain other application
health checks including canaries or integration health checks.
- The `vendor/` folder is not committed to source control, but shown here to demonstrate the location of installed
vendor libraries.

## Go-kit Fundamentals

This section will be a brief, high level discussion of the basic go-kit idioms employed in this example repository.
A better (and more thorough) resource can be found at Peter Bourgon's website or the go-kit repo itself which includes
a thorough example application structured from the Domain Driven Design book "Shipping" topic:

- [Bourgon's presentation](https://peter.bourgon.org/applied-go-kit/#1)
- [UK Gophercon talk](https://www.youtube.com/watch?v=JXEjAwNWays)
- [Go-kit on Github](https://github.com/go-kit/kit)

**Endpoints** are "the building block of Go-kit components". They are "implemented by servers, and called by
clients". An endpoint ingests an application service and decodes and encodes transport requests/responses. Requests
and responses should have a defined `struct` that will contain data parsed from the client request or from the application
service for the response. **Endpoints** are chainable and may be wrapped by **Middleware**. An example of this kind of
decoration may be observed in the `logging.go` or `instrumenting.go` files, although these wrap **Service** methods.

**Transports** are "bindings" to "concrete" transport methods. Simply put, this means that your `transport.go` file should
contain any HTTP, gRPC, or socket transport logic. If you are coming from another micro framework, this may be a
unique practice. Typically we see business service logic (the stuff in `service.go`) in a handler function. Go-kit
encourages us to extract that logic and place it in another module and keep our transport specific code pristine.

**Loggers** like the one you may find in `logging/logger.go` implement Go-kit's application logging conventions. This package
"may be wrapped to encode conventions, enforce type-safety, provide leveled logging, and so on". In the example application
we wrap the Go-kit logger to enable Monolog-style formatted logs for our Heka decoders to consume. In this manner we are able
to define a single safe format for all of our applications' logs. **Telemetry** works in the same way.

## Final Notes and Caveats

This example app is meant to be a template for hardened, production ready Go microservices. By no means
is this a "one size fits all" template. There may be very small services that don't require the structure defined
here BUT we have an answer for that too! All of these constructs are easily chainable and reasonable to place
in a simple `main.go` file alongside some other helper modules and a Dockerfile.

Additionally, we don't claim to have all of the answers! If you have a better idea or practice please open an Issue
against this repository with your suggestion and we will do our best to facilitate a discussion about your concern or suggestion.

Thanks!
