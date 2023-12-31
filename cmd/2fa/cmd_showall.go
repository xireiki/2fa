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
	disableHotp bool
)

var commandShowAll = &cobra.Command {
	Use:   "showall",
	Short: "Print all two-step verification codes",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := cmd.Flags().GetString("config")
		if err != nil {
			log.Fatal(err)
		}
		configDirectory, err = cmd.Flags().GetString("config-directory")
		if err != nil {
			log.Fatal(err)
		}
		err = showall(config, configDirectory)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	commandShowAll.PersistentFlags().BoolVarP(&disableHotp, "disable-hotp", "d", false, "disable hotp output")
	mainCommand.AddCommand(commandShowAll)
}

func showall(ConfigDirectory string, ConfigPath string) error {
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	defer signal.Stop(osSignals)

	var option config.ConfigYaml
	err := option.ReadConfig(filepath.Join(ConfigPath, ConfigDirectory))
	if err != nil {
		return err
	}
	for i, v := range option.Data {
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
				err = option.WriteConfig(filepath.Join(ConfigPath, ConfigDirectory))
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