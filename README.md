
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
|[Google][1]|Regex|:heavy_check_mark:|
|[Yahoo][2]|API|:heavy_check_mark:|
|[Currency Layer][3]|API| :heavy_check_mark: |
|[Fixer.io][4]|API| TODO |
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

## License

The MIT License (MIT). Please see [LICENSE](LICENSE) for more information.
