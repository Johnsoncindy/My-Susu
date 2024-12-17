package config

import (
	"fmt"
	"os"
)

type MoMoConfig struct {
	SubscriptionKey string
	TargetEnvironment string
	CallbackHost string
	MomoHost string
	DisbursementKey string
}

func LoadMoMoConfig() (*MoMoConfig, error) {
	// Load configuration from environment variables
	config := &MoMoConfig{
		SubscriptionKey: os.Getenv("MOMO_SUBSCRIPTION_KEY"),
		TargetEnvironment: os.Getenv("MOMO_TARGET_ENVIRONMENT"),
		CallbackHost: os.Getenv("MOMO_CALLBACK_HOST"),
        MomoHost: os.Getenv("MOMO_MOMO_HOST"),
        DisbursementKey: os.Getenv("MOMO_DISBURSEMENT_KEY"),
	}
	
    // Validate required environment variables
	if config.SubscriptionKey == "" {
		return nil, fmt.Errorf("MOMO_SUBSCRIPTION_KEY is required")
	}
	if config.TargetEnvironment == "" {
		return nil, fmt.Errorf("MOMO_TARGET_ENVIRONMENT is required")
	}
	if config.CallbackHost == "" {
		return nil, fmt.Errorf("MOMO_CALLBACK_HOST is required")
	}
	if config.MomoHost == "" {
		return nil, fmt.Errorf("MOMO_MOMO_HOST is required")
	}
	if config.DisbursementKey == "" {
        return nil, fmt.Errorf("MOMO_DISBURSEMENT_KEY is required")
    }

	return config, nil
}
