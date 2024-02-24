package hostdata

import (
	"encoding/base64"
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

type HostData struct {
	ClusterIssuers     []string          `json:"clusterIssuers"`
	EnvValues          map[string]string `json:"envValues"`
	HostID             string            `json:"hostId"`
	InstanceID         string            `json:"instanceId"`
	InvocationSeed     string            `json:"invocationSeed"`
	LatticeRPCPrefix   string            `json:"latticeRpcPrefix"`
	LatticeRPCURL      string            `json:"latticeRpcUrl"`
	LatticeRPCUserJWT  string            `json:"latticeRpcUserJwt"`
	LatticeRpcUserSeed string            `json:"latticeRpcUserSeed"`
	LinkDefinitions    []LinkDefinition  `json:"linkDefinitions"`
	LinkName           string            `json:"linkName"`
	ProviderKey        string            `json:"providerKey"`
}

func Read(r io.Reader) (*HostData, error) {
	var data HostData
	if err := json.NewDecoder(base64.NewDecoder(base64.StdEncoding, r)).Decode(&data); err != nil {
		return nil, errors.Wrap(err, "failed to read/decode host data")
	}
	return &data, nil
}
