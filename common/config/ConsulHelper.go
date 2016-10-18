package config

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/consul/api"
	"github.com/microbusinesses/Micro-Businesses-Core/common/diagnostics"
)

// ConsulHelper makes working with Consul distributed service discovery easier by providing function to read keys/values and convert them to int, string and etc.
type ConsulHelper struct {
	ConsulAddress string
	ConsulScheme  string
}

// GetKeyPair returns key pair value.
func (consulHelper ConsulHelper) GetKeyPair(keyPath string) (*api.KVPair, error) {
	diagnostics.IsNotNilOrEmptyOrWhitespace(keyPath, "keyPath", "keyPath cannot be null, empty or contains whitespace only.")

	kv, err := getKV(consulHelper)

	if err != nil {
		return nil, err
	}

	keyPair, _, err := kv.Get(keyPath, nil)
	if err != nil {
		return nil, err
	}

	return keyPair, nil
}

// GetInt converts the value identified by the given key path to an integer and returns it.
func (consulHelper ConsulHelper) GetInt(keyPath string) (int, error) {
	diagnostics.IsNotNilOrEmptyOrWhitespace(keyPath, "keyPath", "keyPath cannot be null, empty or contains whitespace only.")

	keyPair, err := consulHelper.GetKeyPair(keyPath)

	if err != nil {
		return 0, err
	}

	if keyPair == nil {
		return 0, fmt.Errorf(fmt.Sprintf("Consul key %s does not exist.", keyPath))

	}

	valueInString := string(keyPair.Value)

	if len(valueInString) == 0 {
		return 0, fmt.Errorf(fmt.Sprintf("Consul key %s is empty.", keyPath))

	}

	value, err := strconv.Atoi(valueInString)

	if err != nil {
		return 0, err
	}

	if value == 0 {
		return 0, fmt.Errorf(fmt.Sprintf("Consul key %s is zero.", keyPath))
	}

	return value, nil
}

// GetString converts the value identified by the given key path to an string and returns it.
func (consulHelper ConsulHelper) GetString(keyPath string) (string, error) {
	diagnostics.IsNotNilOrEmptyOrWhitespace(keyPath, "keyPath", "keyPath cannot be null, empty or contains whitespace only.")

	keyPair, err := consulHelper.GetKeyPair(keyPath)

	if err != nil {
		return "", err
	}

	if keyPair == nil {
		return "", fmt.Errorf(fmt.Sprintf("Consul key %s does not exist.", keyPath))

	}

	return string(keyPair.Value), nil

}

func getKV(consulHelper ConsulHelper) (*api.KV, error) {
	config := api.DefaultConfig()

	if len(consulHelper.ConsulAddress) != 0 && len(consulHelper.ConsulScheme) != 0 {
		config.Address = consulHelper.ConsulAddress
		config.Scheme = consulHelper.ConsulScheme
	}

	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return client.KV(), nil
}
