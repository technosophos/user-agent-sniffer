# User-Agent Sniffer

This is a simple user agent sniffer. It displays what your user agent string looks like from the server's side.

## Build

To build this, you will need:

* [Spin](https://spin.fermyon.dev)
* [TinyGo](https://tinygo.org/getting-started/install/), but not version 0.26, which has a known issue compiling to Wasm
* [Go](https://go.dev/)

Build instructions:

1. Clone this repo
1. Run `spin build`

To deploy, configure your deployment target, and then run:

```
$spin deploy
```

The above should give you a URL.
