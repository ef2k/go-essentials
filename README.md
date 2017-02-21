# go-essentials

This is a containerized development environment for Go. It's best suited for those wanting to get up and running quickly. You should have a basic understanding of Docker and Go to make it your own.

## How to use this repo

The easiest way to get started is to clone this repo, use it as a starting point, and make it your own.

## What you get in the containers

* Go (1.8)
* Automatic code reloads ([Reflex](https://github.com/cespare/reflex))
* Automatic test running ([GoConvey](https://github.com/smartystreets/goconvey))
* Vendored dependency management ([GB](http://getgb.io))
* Postgres by default (easily switchable)

## Motivation and how it works

The motivation was simple: *More executing and less configuring.*. The other was having reproducible builds that you can easily ship and share with others.

This is how it works.

The development environment runs inside of a Docker container. Your codebase is loaded
via a volume, so anytime you make a file change, your application is rebuilt and ran inside the container. In addition to the development tools being containerized, your dependencies are vendored. There are pros and cons to vendoring, but one major pro is having dependable reproducible builds and that outweighs the cons. Last but not least, your test suite is ran inside it's own container.

## What you need before getting started

* **Docker** - To run the Go development, testing, and database containers.
> Tip: Download the latest and greatest: https://www.docker.com/products/overview#/install_the_platform

* **Go** - A local install so you can install GB..
```
$ gvm install go1.8 -B
```
> Tip: One of the easiest ways to install Go is via GVM (https://github.com/moovweb/gvm).

* **GB** - To vendor any dependencies your application will be relying on.
```
$ go get github.com/constabulary/gb/...
```
> Tip: Once you have `GB` installed, start using `gb vendor fetch` instead of `go get`. This will update your list of dependencies and add it to `vendor/`. Here's an example:
```
$ gb vendor fetch github.com/pkg/profile
```

## How to get started

Assuming you have **Docker**, **Go**, and **GB** installed on your machine, all you need to do is:

```
$ source config/environments/dev
$ docker-compose up
```

This will load up environment variables found in `config/environments/dev` and will then `docker-compose up`. Leave the terminal open and you'll be able to see all log info.

**The server will run on http://localhost:8000**

**The test suite runs on http://localhost:8001**

> Tip: Learn all of the other docker-compose commands for greater control.

## Special thanks

* GB (https://getgb.io)
* Reflex (https://github.com/cespare/reflex)
* GoConvey (https://goconvey.co)

## License

Released under the MIT license.
