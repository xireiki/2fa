package main

import (
	"os"
	"os/signal"
	"syscall"
	"log"
	"fmt"
	"path/filepath"

	"main/config"
	"main/otp"
	"github.com/spf13/cobra"
)

var (
	limitIssuer string
)

var commandShow = &cobra.Command {
	Use:   "show",
	Short: "Print two-step verification codes",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := show(cmd, args)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	commandShow.PersistentFlags().StringVarP(&limitIssuer, "issuer", "i", "", "Your issuer name")
	mainCommand.AddCommand(commandShow)
}

func show(cmd *cobra.Command, args []string) error {
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	defer signal.Stop(osSignals)

	configPath, err := cmd.Flags().GetString("config")
	if err != nil {
		return err
	}
	configDirectory, err = cmd.Flags().GetString("config-directory")
	if err != nil {
		return err
	}

	var option config.ConfigYaml
	err = option.ReadConfig(filepath.Join(configDirectory, configPath))
	if err != nil {
		return err
	}

	for i, v := range option.Data {
		if v.Name != args[0] {
			continue
		}
		if limitIssuer != "" && v.Issuer != limitIssuer {
			continue
		}
		name := v.Name + "（" + v.Issuer + "）"
		switch v.Type {
			case "totp":
				code, err := otp.TotpStr(v.Secret, v.Digits, v.TOTPOption.Period)
				if err != nil {
					return err
				}
				fmt.Printf("%-*s\t\t====\t%s\n", v.Digits, code, name)
				break;
			case "hotp":
				if disableHotp {
					continue
				}
				code, err := otp.HotpStr(v.Secret, v.HOTPOption.Counter, v.Digits)
				if err != nil {
					return err
				}
				fmt.Printf("%-*s\t\t====\t%s\n", v.Digits, code, name)
				option.Data[i].HOTPOption.Counter += 1
				err = option.WriteConfig(filepath.Join(configDirectory, configPath))
				if err != nil {
					return err
				}
				break;
			default:
				break;
		}
	}
	return nil
}