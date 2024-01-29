package main

import (
	"fmt"
	"github.com/kholodmv/GophKeeper/internal/client/commands"
)

var (
	buildVersion = "0.01"
)

func main() {
	fmt.Println("Client version: " + buildVersion)
	commands.Execute()
}
