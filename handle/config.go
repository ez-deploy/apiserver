package handle

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type config struct {
	GRPCIdentityHost string `json:"grpc_identity_host,omitempty" yaml:"grpc_identity_host,omitempty"`
	GRPCIdentityPort string `json:"grpc_identity_port,omitempty" yaml:"grpc_identity_port,omitempty"`
}

func getConfigFromFile(filename string) (*config, error) {
	rawCfg, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	cfg := &config{}
	if err := yaml.Unmarshal(rawCfg, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
