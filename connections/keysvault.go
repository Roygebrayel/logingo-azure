package connections

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func GetKeyVaultClient() (client keyvault.BaseClient) {
	tenantID := "e731378a-11b1-405c-bb1c-c047ae5a5d54"
	clientID := "c138d455-83bc-4c49-ba1f-9a86b6c56655"
	clientSecret := "Ug98Q~DQBadbL0Q9liXXqnLABR1HSG9paXuXMcLb"
	keyvaultClient := keyvault.New()
	clientCredentialConfig := auth.NewClientCredentialsConfig(clientID, clientSecret, tenantID)

	// From SDK NewClientCredentialsConfig generates a object to azure control plane
	// (By default Resource is set to management.azure.net)
	// There below line was added to access the azure data plane
	// Which is required to access secrets in keyvault

	clientCredentialConfig.Resource = "https://vault.azure.net"
	authorizer, err := clientCredentialConfig.Authorizer()

	if err != nil {
		fmt.Printf("Error occured while creating azure KV authroizer %v ", err)

	}
	keyvaultClient.Authorizer = authorizer

	return keyvaultClient
}

func GetSecret(vaultURI string, client *keyvault.BaseClient, secretName, secretVersion string) (string, error) {
	secretBundle, err := (*client).GetSecret(context.TODO(), vaultURI, secretName, secretVersion)
	if err != nil {
		return "", fmt.Errorf("error retrieving secret from Key Vault: %v", err)
	}
	return *secretBundle.Value, nil
}
