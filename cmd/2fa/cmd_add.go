package main

import (
	"log"

	"github.com/spf13/cobra"
)

var commandRun = &cobra.Command{
	Use:   "add",
	Short: "Add 2fa",
	Run: func(cmd *cobra.Command, args []string) {
		err := add2fa()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	mainCommand.AddCommand(commandRun)
}

func add2fa() error {
	
	return nil
}