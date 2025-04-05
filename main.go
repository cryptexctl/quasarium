package main

import (
	//	"fmt"
	"quasarium/cmd"
)

var Version = "v1.0"

func main() {
	cmd.Version = Version
	cmd.Execute()
}
