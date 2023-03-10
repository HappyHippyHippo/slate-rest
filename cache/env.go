package cache

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/env"
)

const (
	// EnvID defines the cache package base environment variable name.
	EnvID = slate.EnvID + "_CACHE"
)

var (
	// StoresConfigPath contains the configuration path that holds the
	// cache stores connection configurations.
	StoresConfigPath = env.String(EnvID+"_STORES_CONFIG_PATH", "slate.cache.stores")

	// ObserveConfig defines the store pool cfg observing flag
	// used to register in the cfg object an observer of the store
	// cfg entries list, so it can reset the stores pool.
	ObserveConfig = env.Bool(EnvID+"_OBSERVE_CONFIG", true)

	// DefaultExpiration @todo doc.
	DefaultExpiration = env.Int(EnvID+"_DEFAULT_EXPIRATION", 60000)
)
