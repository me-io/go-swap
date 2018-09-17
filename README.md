
## Go Currency Exchange Library [WIP]

[![Build Status](https://travis-ci.org/me-io/go-swap.svg?branch=master)](https://travis-ci.org/me-io/go-swap)
[![Go Report Card](https://goreportcard.com/badge/github.com/me-io/go-swap)](https://goreportcard.com/report/github.com/me-io/go-swap)


Swap allows you to retrieve currency exchange rates from various services such as **[Fixer](https://fixer.io)**, **[currencylayer](https://currencylayer.com)** or **[1Forge](https://1forge.com)** 
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
|Exchanger|type|
|:---|:----|
|[Google][1]|Regex|
|[Yahoo][2]|API|

[1]: google.com
[2]: yahoo.com

## License

The MIT License (MIT). Please see [LICENSE](LICENSE) for more information.
