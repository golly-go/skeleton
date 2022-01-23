package main

import (
	"github.com/slimloans/golly"
	"github.com/slimloans/golly-skeleton/app/initializers"
	"github.com/slimloans/golly/orm/migrate"
	"github.com/spf13/cobra"
)

var commands = append(golly.AppCommands,
	[]*cobra.Command{
		addNestedChildCommand(&cobra.Command{
			Use:              "migration",
			Short:            "Migration commands",
			TraverseChildren: true,
		}, migrate.Commands),
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

func addNestedChildCommand(root *cobra.Command, cmds []*cobra.Command) *cobra.Command {
	root.AddCommand(cmds...)
	return root
}
