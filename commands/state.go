package commands

import (
	"fmt"

	"github.com/amimof/huego"
	"github.com/sierrasoftworks/hue/config"
	"github.com/zalando/go-keyring"
)

func GetBridge() (*huego.Bridge, error) {
	host, err := keyring.Get("com.sierrasoftworks.hue", "$default_host$")
	if err != nil {
		return nil, err
	}

	user, err := keyring.Get("com.sierrasoftworks.hue", host)
	if err != nil {
		return nil, err
	}

	bridge := huego.New(host, user)

	return bridge, nil
}

func GetConfig() (*config.Config, error) {
	return config.Load()
}

func ToHuman(state *huego.State) string {
	if state.On && state.Bri > 0 {
		return fmt.Sprintf("%.0f%%", (float32(state.Bri)/float32(0xff))*100.0)
	} else {
		return "off"
	}
}
