package main

import (
	"os"
	"log"

	"github.com/spf13/cobra"
)

var (
	configPath      string
	configDirectory string
	workingDir      string
)

var mainCommand = &cobra.Command{
	Use:              "2fa",
	PersistentPreRun: preRun,
}


func init() {
	mainCommand.PersistentFlags().StringVarP(&configPath, "config", "c", "", "set configuration file path")
	mainCommand.PersistentFlags().StringVarP(&configDirectory, "config-directory", "C", "", "set configuration directory path")
	mainCommand.PersistentFlags().StringVarP(&workingDir, "directory", "D", "", "set working directory")
}

func main() {
	if err := mainCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}

func preRun(cmd *cobra.Command, args []string) {
	if workingDir != "" {
		_, err := os.Stat(workingDir)
		if err != nil {
			os.Mkdir(workingDir, 0o755)
		}
		err = os.Chdir(workingDir)
		if err != nil {
			log.Fatal(err)
		}
	}
	if configDirectory == "" {
		configDirectory = os.Getenv("HOME")
	}
	if configPath == "" {
		configPath = ".2fa"
	}
}
