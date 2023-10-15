# FlexiRoute

FlexiRoute is a Golang library designed to provide developers with the flexibility when choosing or switching underlying web frameworks for their projects. This library simplifies the process of swapping out web frameworks without the need to alter the implementation of routes, handler signatures, middleware registration, and more.

## Features

- **Framework Agnostic**: FlexiRoute allows you to choose the web framework that suits your project's needs (currently only go-chi and native http supported).

- **Route Abstraction**: With FlexiRoute, you can define your routes and handlers using a consistent and framework-agnostic approach.

- **Middleware Compatibility**: Integrate middleware into your web application, regardless of the web framework in use.

- **Easy to Get Started**: Getting started with FlexiRoute is easy. You can quickly integrate it into your project and start enjoying the benefits of flexible web framework choices.

## Installation

You can install FlexiRoute using Go modules. Simply run:

```shell
go get github.com/nt-hellofresh/flexiroute
```

## Examples

See [configuration](internal/configure.go) how to bootstrap a FlexiRoute router.

See [here](cmd/default/main.go) for using the native go library as the underlying web framework.

See [here](cmd/chi/main.go) for using go chi as the underlying web framework.

## Contributing

This project is not actively maintained and was simply created out of a weekend of experimentation building a basic web framework of my own. It serves as a reference and a source of inspiration for other contributors whom wish to fork from it.

## License

FlexiRoute is licensed under the [MIT License](LICENSE)