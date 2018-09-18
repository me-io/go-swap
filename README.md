
## Go Currency Exchange Library [WIP]

[![Build Status](https://travis-ci.org/me-io/go-swap.svg?branch=master)](https://travis-ci.org/me-io/go-swap)
[![Go Report Card](https://goreportcard.com/badge/github.com/me-io/go-swap)](https://goreportcard.com/report/github.com/me-io/go-swap)


Swap allows you to retrieve currency exchange rates from various services such as **[Fixer](https://fixer.io)**, **[CurrencyLayer](https://currencylayer.com)** or **[1Forge](https://1forge.com)** 
and optionally cache the results. 

> Inspired by [florianv/swap](https://github.com/florianv/swap).

## QuickStart

```bash
$ go get github.com/me-io/go-swap
```

```go
package main

import (
	"fmt"
	"github.com/me-io/go-swap"
	ex "github.com/me-io/go-swap/exchanger"
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
|Exchanger|type|#|
|:---|:----|:---|
|[Google][1]|HTML / Regex|:heavy_check_mark:|
|[Yahoo][2]|JSON / API|:heavy_check_mark:|
|[Currency Layer][3]|JSON / API| :heavy_check_mark: |
|[Fixer.io][4]|JSON / API| :heavy_check_mark: |
|[The Money Converter][5]|API| TODO |
|[Open Exchange Rates][6]|API| TODO |
|[1forge][7]|API| TODO |

[1]: //google.com
[2]: //yahoo.com
[3]: //currencylayer.com
[4]: //fixer.io
[5]: //themoneyconverter.com
[6]: //openexchangerates.org
[7]: //1forge.com

## TODO LIST
- [ ] cache integration
- [ ] Add Mocks
- [ ] Benchmark & Performance optimization ` memory leak`
- [ ] examples
- [ ] code coverage
- [ ] godoc 
- [ ] herokuapp demo with links
- [ ] contributors list 


## Contributing

Anyone is welcome to [contribute](CONTRIBUTING.md), however, if you decide to get involved, please take a moment to review the guidelines:

* [Only one feature or change per pull request](CONTRIBUTING.md#only-one-feature-or-change-per-pull-request)
* [Write meaningful commit messages](CONTRIBUTING.md#write-meaningful-commit-messages)
* [Follow the existing coding standards](CONTRIBUTING.md#follow-the-existing-coding-standards)

## License

The code is available under the [MIT license](LICENSE.md).
