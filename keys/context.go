package keys

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/metadata"
)

// This file contains the keys used to store values in the context.
// It is also used to access and store gRPC metadata.

const (
	// ClaimsKey is the key used to store the jwt claims in the context
	ClaimsKey Key = "claims"
	// TokenKey is the key used to store the jwt token in the context
	TokenKey Key = "token"
)

// ----------------- Context Functions -----------------

// GetValues retrieves the values from the context using the provided keys.
func GetContextValues(ctx context.Context, keys ...Key) ([]string, error) {
	var values []string

	for _, key := range keys {
		value, err := GetContextValue(ctx, key)
		if err != nil {
			return nil, err
		}

		values = append(values, value)
	}

	return values, nil
}

// GetContextValue retrieves the value from the context using the provided key.
// It returns the value as a string.
func GetContextValue(ctx context.Context, key Key) (string, error) {
	value, ok := ctx.Value(key).(string)
	if !ok {
		return "", fmt.Errorf("value %s not found in context", key)
	}

	return value, nil
}

// GetContextValueString retrieves the value from the context using the provided key.
func GetContextValueString(ctx context.Context, key Key) (string, error) {
	value, ok := ctx.Value(key).(string)
	if !ok {
		return "", fmt.Errorf("value %s not found in context", key)
	}

	return value, nil
}

// GetClaimsFromContext retrieves the claim from the context using the provided key.
func GetClaimsFromContext(ctx context.Context) (jwt.MapClaims, error) {
	claims, ok := ctx.Value(ClaimsKey).(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("claims not found in context")
	}

	return claims, nil
}

// ----------------- Metadata Functions -----------------

// GetMetadataValues retrieves the values from the metadata using the provided keys.
func GetMetadataValues(md metadata.MD, keys ...Key) ([]string, error) {
	var values []string

	for _, key := range keys {
		value, err := GetMetadataValue(md, key)
		if err != nil {
			return nil, err
		}

		values = append(values, value)
	}

	return values, nil
}

// GetMetadataValue retrieves the value from the metadata using the provided key.
// It returns the value as a string.
func GetMetadataValue(md metadata.MD, key Key) (string, error) {
	value := md.Get(string(key))
	if len(value) == 0 {
		return "", fmt.Errorf("value %s not found in metadata", key)
	}

	return value[0], nil
}
