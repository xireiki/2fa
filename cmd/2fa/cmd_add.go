package main

import (
	"log"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"main/config"
)

var (
	Name      string
	Issuer    string
	Type      string
	Secret    string
	Digits    int
	Counter   int
	Period    int
)

var commandAdd = &cobra.Command{
	Use:   "add",
	Short: "Add 2fa",
	Run: func(cmd *cobra.Command, args []string) {
		err := add2fa(cmd)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	commandAdd.PersistentFlags().StringVarP(&Name, "name", "n", "", "Your account name")
	commandAdd.PersistentFlags().StringVarP(&Issuer, "issuer", "i", "", "Your issuer name")
	commandAdd.PersistentFlags().StringVarP(&Type, "type", "t", "totp", "Validator Type，default: totp")
	commandAdd.PersistentFlags().StringVarP(&Secret, "secret", "s", "", "Key")
	commandAdd.PersistentFlags().IntVarP(&Digits, "digits", "d", 6, "Length of code, default: 6")
	commandAdd.PersistentFlags().IntVarP(&Counter, "counter", "q", 0, "Counter")
	commandAdd.PersistentFlags().IntVarP(&Period, "period", "p", 30, "Refresh interval")
	mainCommand.AddCommand(commandAdd)
}

func add2fa(cmd *cobra.Command) error {
	if Name == "" {
		return fmt.Errorf("Name cannot be empty. Use -h to view help.")
	}
	if Secret == "" {
		return fmt.Errorf("Key cannot be empty. Use -h to view help.")
	}
	configPath, err := cmd.Flags().GetString("config")
	if err != nil {
		log.Fatal(err)
	}
	configDirectory, err = cmd.Flags().GetString("config-directory")
	if err != nil {
		log.Fatal(err)
	}
	var options config.ConfigYaml
	err = options.ReadConfig(filepath.Join(configDirectory, configPath))
	if err != nil {
		return err
	}
	data := config.DataOption {
		Name: Name,
		Issuer: Issuer,
		Type: Type,
		Digits: Digits,
		Secret: Secret,
	}
	switch Type {
	case "hotp":
		data.HOTPOption = config.HOTPOption {
			Counter: Counter,
		}
	case "totp":
		data.TOTPOption = config.TOTPOption {
			Period: Period,
		}
	default:
		return fmt.Errorf("Unknown type: %s", Type)
	}
	options.Data = append(options.Data, data)
	err = options.WriteConfig(filepath.Join(configDirectory, configPath))
	if err != nil {
		return err
	}
	OutName := Name
	if Issuer != "" {
		OutName += "（" + Issuer + "）"
	}
	fmt.Println("Added:", OutName)
	return nil
}