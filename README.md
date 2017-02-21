# go-essentials

The web app starter pack for Go. Development environment included.

## What you get in the containers

* Go (1.8)
* Profiling
* Automatic code reloads
* Automatic test running
* Vendored dependency management
* Postgres by default (easily switchable)

## Motivation and how it works

The motivation was simple: Invest your time on executing and not configuring.

This is how it works.

The development environment runs inside of a Docker container. Your codebase is loaded
via a volume, so anytime you make a file change, your application is rebuilt and ran inside the container. In addition to the development tools being containerized, your dependencies are vendored. There are pros and cons to vendoring, but one major pro is having dependable reproducible builds and that outweighs the cons. For more please see the "Q&A section". Last but not least, your test suite is ran inside the container.

## What you need before getting started

* Docker - To run a Go development container and a database container

> Tip: Download the latest and greatest: https://www.docker.com/products/overview#/install_the_platform

That's all you need to run and start coding.

If you don't already have a local install of Go running. Install it. You'll
eventually need it to vendor your dependencies:

* GB - To vendor any dependencies your application will be relying on.
* Go - A local install so you can install GB.

> Tip: One of the easiest ways to install Go is via GVM (https://github.com/moovweb/gvm).


## How to get started

Assuming you have **Docker**, **Go**, and **GB** installed on your machine, all you need to do is:

```
$ docker-compose up
```

## Special thanks

* GB
* Reflex
* GoConvey
