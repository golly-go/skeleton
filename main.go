package main

import (
	"github.com/slimloans/golly"
	"github.com/slimloans/golly-skeleton/app/initializers"
	"github.com/spf13/cobra"
)

var commands = append(golly.AppCommands,
	[]*cobra.Command{
		// {
		// 	Use:   "consumers",
		// 	Short: "Run Consumers",
		// 	Run:   func(cmd *cobra.Command, args []string) { redis.Run() },
		// },
	}...)

func main() {
	rootCMD := cobra.Command{}
	rootCMD.AddCommand(commands...)

	golly.RegisterPreboot(initializers.Preboots...)
	golly.RegisterInitializer(initializers.Initializers...)

	if err := rootCMD.Execute(); err != nil {
		panic(err)
	}
}
