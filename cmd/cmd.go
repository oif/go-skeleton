package main

import (
	"github.com/oif/go-skeleton/cmd/echo"
	"github.com/oif/go-skeleton/pkg/meta"
	"github.com/oif/gokit/logs"
	"github.com/spf13/cobra"
)

func main() {
	setupRuntime()

	_, err := logs.Setup(logs.EnableSourceHook(), logs.LogLevel("info"))
	if err != nil {
		panic(err)
	}

	commands := &cobra.Command{Use: meta.App}
	for _, cmd := range []*cobra.Command{
		echo.NewCommand(),
	} {
		commands.AddCommand(cmd)
	}
	if err := commands.Execute(); err != nil {
		panic(err)
	}
}
