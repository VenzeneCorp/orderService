package utils

import (
	"log"
	"time"

	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func init() {
	// Initialize Sonyflake with default settings
	sf = sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Unix(1609459200, 0), // Custom start time (optional, default is 2010-01-01)
		MachineID: func() (uint16, error) {
			// Return a unique machine ID (e.g., can be from a config or environment variable)
			return 1, nil
		},
	})
	if sf == nil {
		log.Fatal("sonyflake not initialized")
	}
}

func GenerateID() (uint64, error) {
	// Generate a new Snowflake ID
	return sf.NextID()
}
