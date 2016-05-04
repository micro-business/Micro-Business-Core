package config

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hashicorp/consul/api"
)

type ConsulHelper struct {
	ConsulAddress string
	ConsulScheme  string
}

func (consulHelper ConsulHelper) GetInt(keyPath string) (int, error) {
	keyPair, err := getKeyPair(consulHelper, keyPath)

	if err != nil {
		return 0, err
	}

	if keyPair == nil {
		return 0, errors.New(fmt.Sprintf("Consul key %s does not exist.", keyPath))

	}

	valueInString := string(keyPair.Value)

	if len(valueInString) == 0 {
		return 0, errors.New(fmt.Sprintf("Consul key %s is empty.", keyPath))

	}

	if value, err := strconv.Atoi(valueInString); err != nil {
		return 0, err
	} else {
		if value == 0 {
			return 0, errors.New(fmt.Sprintf("Consul key %s is zero.", keyPath))
		}

		return value, nil
	}
}

func (consulHelper ConsulHelper) GetString(keyPath string) (string, error) {
	keyPair, err := getKeyPair(consulHelper, keyPath)

	if err != nil {
		return "", err
	}

	if keyPair == nil {
		return "", errors.New(fmt.Sprintf("Consul key %s does not exist.", keyPath))

	}

	return string(keyPair.Value), nil

}

func getKV(consulHelper ConsulHelper) (*api.KV, error) {
	config := api.DefaultConfig()

	if len(consulHelper.ConsulAddress) != 0 && len(consulHelper.ConsulScheme) != 0 {
		config.Address = consulHelper.ConsulAddress
		config.Scheme = consulHelper.ConsulScheme
	}

	if client, err := api.NewClient(config); err != nil {
		return nil, err
	} else {
		return client.KV(), nil
	}
}

func getKeyPair(consulHelper ConsulHelper, configKeyPath string) (*api.KVPair, error) {
	kv, err := getKV(consulHelper)

	if err != nil {
		return nil, err
	}

	if keyPair, _, err := kv.Get(configKeyPath, nil); err != nil {
		return nil, err
	} else {
		return keyPair, nil
	}
}
