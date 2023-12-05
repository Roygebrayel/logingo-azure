# We strongly recommend using the required_providers block to set the

terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=2.0.0"
    }
  }
}

provider "azurerm" { 
  subscription_id = "189af0e7-fa98-4751-b5a8-0d39f4d49d5f"
  client_id       = "14581ffd-6ddc-44de-90f3-e3afc4569daa"
  client_secret   = "54cae9d3-42b7-4f90-8a31-73384f5329ec"  
  tenant_id       = "e731378a-11b1-405c-bb1c-c047ae5a5d54" 
      
 features {} 
 }

# Create a resource group
resource "azurerm_resource_group" "example" {
  name     = "example-resources"
  location = "West Europe"
}

# Create a virtual network within the resource group
resource "azurerm_virtual_network" "example" {
  name                = "example-network"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location
  address_space       = ["10.0.0.0/16"]
}