package main

import (
	cloudkms "cloud.google.com/go/kms/apiv1"
	"context"
	"encoding/base64"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
	"os"
	"strings"
)

// c.f. https://godoc.org/google.golang.org/api/cloudkms/v1#pkg-constants
const (
	// View and manage your data across Google Cloud Platform services
	cloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View and manage your keys and secrets stored in Cloud Key Management
	// Service
	cloudkmsScope = "https://www.googleapis.com/auth/cloudkms"
)

// Kms manages kms decryption
type Kms struct {
	KeyringKeyName string
}

// GetFromEnvOrKms returns value either env or KMS
func (k *Kms) GetFromEnvOrKms(key string, required bool) (string, error) {
	if os.Getenv(key) != "" {
		return strings.TrimSpace(os.Getenv(key)), nil
	}

	kmsKey := "KMS_" + key

	if os.Getenv(kmsKey) != "" {
		value, err := k.decrypt(os.Getenv(kmsKey))

		if err != nil {
			return "", err
		}

		return strings.TrimSpace(value), nil
	}

	if required {
		return "", fmt.Errorf("either %s or %s is required", key, kmsKey)
	}

	return "", nil
}

func (k *Kms) decrypt(base64Value string) (string, error) {
	ctx := context.Background()

	creds, err := google.FindDefaultCredentials(ctx, cloudkmsScope)
	if err != nil {
		return "", err
	}

	client, err := cloudkms.NewKeyManagementClient(ctx, option.WithCredentials(creds))
	if err != nil {
		return "", err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(base64Value)
	if err != nil {
		return "", err
	}

	if k.KeyringKeyName == "" {
		return "", fmt.Errorf("KMS_KEYRING_KEY_NAME is required")
	}

	// Build the request.
	req := &kmspb.DecryptRequest{
		Name:       k.KeyringKeyName,
		Ciphertext: ciphertext,
	}
	// Call the API.
	resp, err := client.Decrypt(ctx, req)
	if err != nil {
		return "", err
	}
	return string(resp.Plaintext), nil
}
