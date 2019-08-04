package main

import (
	cloudkms "cloud.google.com/go/kms/apiv1"
	"context"
	"encoding/base64"
	"fmt"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
	"os"
)

// Kms manages kms decryption
type Kms struct {
	KeyringKeyName string
}

// GetFromEnvOrKms returns value either env or KMS
func (k *Kms) GetFromEnvOrKms(key string) (string, error) {
	if os.Getenv(key) != "" {
		return os.Getenv(key), nil
	}

	kmsKey := "KMS_" + key

	if os.Getenv(kmsKey) != "" {
		value, err := k.decrypt(os.Getenv(kmsKey))

		if err != nil {
			return "", err
		}

		return value, nil
	}

	return "", fmt.Errorf("either %s or %s is required", key, kmsKey)
}

func (k *Kms) decrypt(base64Value string) (string, error) {
	ctx := context.Background()
	client, err := cloudkms.NewKeyManagementClient(ctx)
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
