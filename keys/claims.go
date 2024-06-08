package keys

import (
	"context"
	"fmt"
)

// This file contains keys and helper methods used to retrieve and store
// the corresponding claims from the jwt token in the context.

const (
	// SubjectKey is the key used to store the subject claim in the context
	SubjectKey Key = "sub"
	// IssuerKey is the key used to store the issuer claim in the context
	IssuerKey Key = "iss"
	// AudienceKey is the key used to store the audience claim in the context
	AudienceKey Key = "aud"
	// ExpiryKey is the key used to store the expiry claim in the context
	ExpiryKey Key = "exp"
	// IssuedAtKey is the key used to store the issued at claim in the context
	IssuedAtKey Key = "iat"
	// IsRefreshTokenKey is the key used to store the is refresh token claim in the context
	IsRefreshTokenKey Key = "irf"
)

// IsRefreshToken returns true if the token is a refresh token.
func IsRefreshToken(ctx context.Context) bool {
	irf, err := GetClaim(ctx, IsRefreshTokenKey)
	if err != nil {
		return false
	}

	return irf == "true"
}

// GetClaim retrieves the claim from the context using the provided key.
func GetClaim(ctx context.Context, key Key) (string, error) {
	claims, err := GetClaimsFromContext(ctx)
	if err != nil {
		return "", err
	}

	claim, ok := claims[string(key)].(string)
	if !ok {
		return "", fmt.Errorf("claim %s not found in context", key)
	}

	return claim, nil
}

// GetClaimOrEmpty retrieves the claim from the context using the provided key.
// If the claim is not found, it returns an empty string.
func GetClaimOrEmpty(ctx context.Context, key Key) string {
	claim, err := GetClaim(ctx, key)
	if err != nil {
		return fmt.Sprintf("%s not found", key)
	}

	return claim
}
