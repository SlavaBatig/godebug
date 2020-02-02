package debug

import "github.com/SlavaBatig/godebug"

func main() {
	debug, err := utils.Debug("test:namespace")

	debug("Printed to stdout")

	err("Printed to stderr")
}