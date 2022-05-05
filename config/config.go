package config

import (
	"encoding/json"
	"fmt"
	"github.com/sierrasoftworks/humane-errors-go"
	"os"
	"path/filepath"
)

type Config struct {
	Aliases map[string]string `json:"aliases"`
}

func Save(config *Config) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return humane.Wrap(
			err,
			"Could not determine user home directory",
		)
	}

	err = os.MkdirAll(filepath.Join(home, ".hue"), 0755)
	if err != nil {
		return humane.Wrap(
			err,
			fmt.Sprintf("Could not create configuration directory '%s'", filepath.Join(home, ".hue")),
			"Make sure that you have permission to write to the configuration directory.",
		)
	}

	f, err := os.OpenFile(filepath.Join(home, ".hue", "config.json"), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return humane.Wrap(
			err,
			fmt.Sprintf("Failed to open configuration file '%s'", filepath.Join(home, ".hue", "config.json")),
			"Make sure that the file exists and is readable by your current user.",
		)
	}

	//goland:noinspection GoUnhandledErrorResult
	defer f.Close()

	err = f.Truncate(0)
	if err != nil {
		return humane.Wrap(err, "Failed to truncate configuration file.", "Make sure that you have permission to write to the configuration file.")
	}
	err = json.NewEncoder(f).Encode(config)
	if err != nil {
		return humane.Wrap(
			err,
			"Failed to write configuration file.",
			"Make sure that you have permission to write to the configuration file.",
		)
	}

	return nil
}

func Load() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, humane.Wrap(
			err,
			"Could not determine user home directory",
		)
	}

	config := Config{
		Aliases: make(map[string]string),
	}

	f, err := os.Open(filepath.Join(home, ".hue", "config.json"))
	if err != nil {
		if os.IsNotExist(err) {
			return &config, nil
		}

		return nil, humane.Wrap(
			err,
			fmt.Sprintf("Failed to open configuration file '%s'", filepath.Join(home, ".hue", "config.json")),
			"Make sure that the file exists and is readable by your current user.",
		)
	}

	//goland:noinspection GoUnhandledErrorResult
	defer f.Close()

	err = json.NewDecoder(f).Decode(&config)
	if err != nil {
		return nil, humane.Wrap(
			err,
			"Failed to parse configuration file.",
			"Make sure that your configuration file contains valid JSON and try again.",
		)
	}

	return &config, nil
}
