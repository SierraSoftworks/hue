package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sierrasoftworks/hue/humanerrors"
)

type Config struct {
	Aliases map[string]string `json:"aliases"`
}

func Save(config *Config) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return humanerrors.NewWithCause(
			err,
			"Could not determine user home directory",
		)
	}

	err = os.MkdirAll(filepath.Join(home, ".hue"), 0755)
	if err != nil {
		return humanerrors.NewWithCause(
			err,
			fmt.Sprintf("Could not create configuration directory '%s'", filepath.Join(home, ".hue")),
			"Make sure that you have permission to write to the configuration directory.",
		)
	}

	f, err := os.OpenFile(filepath.Join(home, ".hue", "config.json"), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return humanerrors.NewWithCause(
			err,
			fmt.Sprintf("Failed to open configuration file '%s'", filepath.Join(home, ".hue", "config.json")),
			"Make sure that the file exists and is readable by your current user.",
		)
	}

	defer f.Close()

	f.Truncate(0)
	err = json.NewEncoder(f).Encode(config)
	if err != nil {
		return humanerrors.NewWithCause(
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
		return nil, humanerrors.NewWithCause(
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

		return nil, humanerrors.NewWithCause(
			err,
			fmt.Sprintf("Failed to open configuration file '%s'", filepath.Join(home, ".hue", "config.json")),
			"Make sure that the file exists and is readable by your current user.",
		)
	}

	defer f.Close()

	err = json.NewDecoder(f).Decode(&config)
	if err != nil {
		return nil, humanerrors.NewWithCause(
			err,
			"Failed to parse configuration file.",
			"Make sure that your configuration file contains valid JSON and try again.",
		)
	}

	return &config, nil
}
