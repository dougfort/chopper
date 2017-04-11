package config

import "github.com/dougfort/chopper/chopserv/types"

// Load populates the config struct from a stored file (.spindrift)
func Load() (types.Config, error) {

	// start with hard coded values, worry about the file later
	cfg := types.Config{
		Address: "127.0.0.1:9000",
	}

	return cfg, nil
}
