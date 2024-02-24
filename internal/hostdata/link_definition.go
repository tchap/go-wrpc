package hostdata

type LinkDefinition struct {
	ActorID    string            `json:"actorId"`
	ContractID string            `json:"contractId"`
	LinkName   string            `json:"linkName"`
	ProviderID string            `json:"providerId"`
	Values     map[string]string `json:"values"`
}
