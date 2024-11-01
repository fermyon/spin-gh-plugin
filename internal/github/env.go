package github

import (
	"fmt"
	"strings"
)

type EnvVar struct {
	Key   string
	Value string
}

func ParseEnvVars(values []string) ([]*EnvVar, error) {
	results := make([]*EnvVar, len(values))
	for idx, value := range values {
		ev, err := ParseEnvVar(value)
		if err != nil {
			return nil, err
		}
		results[idx] = ev
	}
	return results, nil
}

func ParseEnvVar(value string) (*EnvVar, error) {
	if value == "" {
		return nil, fmt.Errorf("value cannot be empty")
	}

	parts := strings.SplitN(value, "=", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid format, expected key=value")
	}

	key := parts[0]
	if !isValidEnvVarName(key) {
		return nil, fmt.Errorf("invalid environment variable name: %s", key)
	}

	return &EnvVar{
		Key:   key,
		Value: parts[1],
	}, nil
}

func isValidEnvVarName(name string) bool {
	if name == "" {
		return false
	}

	for _, char := range name {
		if !(char == '_' || char == '-' || ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || ('0' <= char && char <= '9')) {
			return false
		}
	}

	return true
}
