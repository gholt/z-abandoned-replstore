package replstore

// ReplGroupStoreConfig defines the settings when calling NewGroupStore.
type ReplGroupStoreConfig struct {
	// AddressIndex indicates which of the ring node addresses to use when
	// connecting to a node (see github.com/gholt/ring/Node.Address).
	AddressIndex int
}

func resolveReplGroupStoreConfig(c *ReplGroupStoreConfig) *ReplGroupStoreConfig {
	cfg := &ReplGroupStoreConfig{}
	if c != nil {
		*cfg = *c
	}
	// AddressIndex default of 0 is fine.
	return cfg
}
