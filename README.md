# godebug

Simple debug logger with colorful namespaces and log time.

The color of namespace is calculated from it's hex representation.

## Motivation

Inspired by [debug](https://www.npmjs.com/package/debug).

## Install

`go get github.com/SlavaBatig/godebug`

or

`dep ensure --add github.com/SlavaBatig/godebug`

## Example

### Code

```
package main

import "github.com/SlavaBatig/godebug"

func main() {
	debug, err := utils.Debug("test:namespace")

	debug("Printed to stdout")

	err("Printed to stderr")
}
```

### Result printed to terminal

![](https://i.imgur.com/sJxokfb.png)
