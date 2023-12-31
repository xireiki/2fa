package config

import (
	"fmt"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type ConfigYaml struct {
	Data   []DataOption `yaml:"data"`
}

type _DataOption struct {
	Name       string    `yaml:"name"`
	Issuer     string    `yaml:"issuer,omitempty"`
	Type       string    `yaml:"type"`
	Secret     string    `yaml:"secret"`
	Digits     int       `yaml:"digits,omitempty"`
	TOTPOption TOTPOption `yaml:",inline"`
	HOTPOption HOTPOption `yaml:",inline"`
}

type DataOption _DataOption

type TOTPOption struct {
	Period int `yaml:"period,omitempty"`
}

type HOTPOption struct {
	Counter int `yaml:"counter,omitempty"`
}

func (d *DataOption) UnmarshalYAML(unmarshal func(interface{}) error) error {
	tmp := _DataOption{}
	err := unmarshal(&tmp)
	if err != nil {
		return err
	}
	*d = DataOption(tmp)
	if d.Digits < 6 {
		d.Digits = 6
	} else if d.Digits > 8 {
		return fmt.Errorf("Digits is an integer between 6 and 8.")
	}
	switch d.Type {
	case "totp":
		if d.TOTPOption.Period == 0 {
			d.TOTPOption.Period = 30
		}
	case "hotp":
	default:
		return fmt.Errorf("Unknown type: %s", d.Type)
	}
	return nil
}

func (c *ConfigYaml) ReadConfig(filePath string) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, &c)
	if err != nil {
		return err
	}
	return nil
}

func (c *ConfigYaml) WriteConfig(filePath string) error {
	content, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}
	file, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	ioutil.WriteFile(filePath, content, file.Mode().Perm())
	return nil
}
