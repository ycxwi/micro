package main

//go:generate ./scripts/generate.sh

import (
	"github.com/ycxwi/micro/v3/cmd"

	_ "github.com/ycxwi/micro/v3/cmd/usage"
	// load packages so they can register commands
	_ "github.com/ycxwi/micro/v3/client/cli"
	_ "github.com/ycxwi/micro/v3/cmd/server"
	_ "github.com/ycxwi/micro/v3/cmd/service"
	_ "github.com/ycxwi/micro/v3/cmd/usage"
)

func main() {
	cmd.Run()
}
