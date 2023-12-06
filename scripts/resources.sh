#!/bin/bash
    #   $servicePrincipalId = '14581ffd-6ddc-44de-90f3-e3afc4569daa'
    #   $servicePrincipalKey = '54cae9d3-42b7-4f90-8a31-73384f5329ec'
    #  $tenantId = 'e731378a-11b1-405c-bb1c-c047ae5a5d54'
    #   $subscription_Id = '14581ffd-6ddc-44de-90f3-e3afc4569daa'
      


# Usage: ./start_webapp.sh

# Check if all required environment variables are set
if [ -z "$SERVICE_PRINCIPAL_ID" ] || [ -z "$SERVICE_PRINCIPAL_KEY" ] || [ -z "$TENANT_ID" ] || [ -z "$SUBSCRIPTION_ID" ] || [ -z "$RESOURCE_GROUP" ] || [ -z "$WEBAPP_NAME" ]; then
    echo "Error: Missing required environment variables."
    exit 1
fi

# Log in to Azure
az login --service-principal -u $SERVICE_PRINCIPAL_ID -p $SERVICE_PRINCIPAL_KEY --tenant $TENANT_ID

# Set your subscription
az account set --subscription $SUBSCRIPTION_ID

# Start the web app
az webapp start --name $WEBAPP_NAME --resource-group $RESOURCE_GROUP
