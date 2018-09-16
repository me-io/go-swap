
## Go Currency Exchange Library [WIP]

[![Build Status](https://travis-ci.org/meabed/go-swap.svg?branch=master)](https://travis-ci.org/meabed/go-swap)
[![Go Report Card](https://goreportcard.com/badge/github.com/meabed/go-swap)](https://goreportcard.com/report/github.com/meabed/go-swap)


Swap allows you to retrieve currency exchange rates from various services such as **[Fixer](https://fixer.io)**, **[currencylayer](https://currencylayer.com)** or **[1Forge](https://1forge.com)** 
and optionally cache the results.

## QuickStart

```bash
$ go get github.com/meabed/go-swap
```

```go
package main

import (
	"fmt"
	"github.com/meabed/go-swap"
	ex "github.com/meabed/go-swap/exchanger"
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
