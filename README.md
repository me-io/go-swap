
## Go Currency Exchange Library [WIP]

[![Build Status](https://travis-ci.org/me-io/go-swap.svg?branch=master)](https://travis-ci.org/me-io/go-swap)
[![Go Report Card](https://goreportcard.com/badge/github.com/me-io/go-swap)](https://goreportcard.com/report/github.com/me-io/go-swap)
[![Coverage Status](https://coveralls.io/repos/github/me-io/go-swap/badge.svg?branch=master)](https://coveralls.io/github/me-io/go-swap?branch=master)
[![GoDoc](https://godoc.org/github.com/me-io/go-swap?status.svg)](https://godoc.org/github.com/me-io/go-swap)


[![](https://images.microbadger.com/badges/version/meio/go-swap-server.svg)](https://microbadger.com/images/meio/go-swap-server)
[![COMMIT](https://images.microbadger.com/badges/commit/meio/go-swap-server.svg)](https://microbadger.com/images/meio/go-swap-server)
[![SIZE-LAYERS](https://images.microbadger.com/badges/image/meio/go-swap-server.svg)](https://microbadger.com/images/meio/go-swap-server)
[![Pulls](https://shields.beevelop.com/docker/pulls/meio/go-swap-server.svg?style=flat-square)](https://hub.docker.com/r/meio/go-swap-server)
[![Release](https://shields.beevelop.com/github/release/me-io/go-swap.svg?style=flat-square)](https://github.com/me-io/go-swap/releases)

Swap allows you to retrieve currency exchange rates from various services such as **[Fixer](https://fixer.io)**, **[CurrencyLayer](https://currencylayer.com)** or **[1Forge](https://1forge.com)** 
and optionally cache the results. 

> Inspired by [florianv/swap](https://github.com/florianv/swap) 

## QuickStart 

<a href="https://heroku.com/deploy?template=https://github.com/me-io/go-swap">
  <img src="https://www.herokucdn.com/deploy/button.svg" alt="Deploy">
</a>
<br>


```bash
# Or using docker  
$ docker pull meio/go-swap-server:latest && \
  docker run --rm --name go-swap-server -p 5000:5000 -it meio/go-swap-server:latest
```

### Programmatically
```bash
$ go get github.com/me-io/go-swap
```

```go
package main

import (
	"fmt"
	ex "github.com/me-io/go-swap/pkg/exchanger"
	"github.com/me-io/go-swap/pkg/swap"
)

func main() {
	SwapTest := swap.NewSwap()

	SwapTest.
		AddExchanger(ex.NewGoogleApi(nil)).
		Build()

	euroToUsdRate := SwapTest.Latest("EUR/USD")
	fmt.Println(euroToUsdRate.GetValue())
}

```


## Documentation
The documentation for the current branch can be found [here](#documentation).


## Services
|Exchanger|type|#|$|
|:---|:----|:---|:---|
|[Google][1]|HTML / Regex|:heavy_check_mark:|Free|
|[Yahoo][2]|JSON / API|:heavy_check_mark:|Free|
|[Currency Layer][3]|JSON / API| :heavy_check_mark: |Paid - ApiKey|
|[Fixer.io][4]|JSON / API| :heavy_check_mark: |Paid - ApiKey|
|[The Money Converter][5]|HTML / Regex| TODO |Free|
|[Open Exchange Rates][6]|API| TODO |Freemium / Paid - ApiKey|
|[1forge][7]|API| TODO |Freemium / Paid - ApiKey|

[1]: //google.com
[2]: //yahoo.com
[3]: //currencylayer.com
[4]: //fixer.io
[5]: //themoneyconverter.com
[6]: //openexchangerates.org
[7]: //1forge.com

## TODO LIST
- [x] Add Test Mocks
- [X] cache integration
- [ ] verbose logging
- [ ] config file for default tokens
- [ ] code coverage
- [ ] herokuapp demo with links
- [ ] godoc 
- [ ] examples
- [ ] v 1.0.0 release binary
- [ ] v 1.0.0 release docker image
- [ ] docker image mac os example
- [ ] Server Postman API collection 
- [ ] Benchmark & Performance optimization ` memory leak`
- [ ] contributors list 

## Contributing

Anyone is welcome to [contribute](CONTRIBUTING.md), however, if you decide to get involved, please take a moment to review the guidelines:

* [Only one feature or change per pull request](CONTRIBUTING.md#only-one-feature-or-change-per-pull-request)
* [Write meaningful commit messages](CONTRIBUTING.md#write-meaningful-commit-messages)
* [Follow the existing coding standards](CONTRIBUTING.md#follow-the-existing-coding-standards)

## License

The code is available under the [MIT license](LICENSE.md).
