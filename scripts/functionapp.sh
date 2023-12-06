
# this script is to automatically run the functionapp while giving its triggers

# Change directory to Downloads/azure-functions-cli/
cd Downloads/azure-functions-cli/

# Run the 'func' command
./func

# Add the current directory to the PATH
export PATH=$(pwd):$PATH

# Run 'func' again
func

cd

cd go-test/functionnodeapp
# Initialize a new Azure Functions project
func init

# Create new Azure Functions with specified templates
echo "n" | func new --name CredentialsPosting --template "HTTP trigger"

cd CredentialsPosting


func settings add AzureWebJobsStorage "DefaultEndpointsProtocol=https;AccountName=functionnodeappa1d7e6;AccountKey=j/3Lb3aaOIFWSSyHuPg+Gu9Lk0uzwjCRFOXVkpf0/aV00VutHceLrR647tA3rZqoitd56s33OBFE+AStAbTvFw==;EndpointSuffix=core.windows.net"
func settings add MyHttpFunction_AUTH_LEVEL "function"
func settings add MyHttpFunction_KEY "313"

cd ..

echo "n" | func new --name GetRequest --template "HTTP trigger"

cd GetRequest

func settings add AzureWebJobsStorage "DefaultEndpointsProtocol=https;AccountName=functionnodeappa1d7e6;AccountKey=j/3Lb3aaOIFWSSyHuPg+Gu9Lk0uzwjCRFOXVkpf0/aV00VutHceLrR647tA3rZqoitd56s33OBFE+AStAbTvFw==;EndpointSuffix=core.windows.net"
func settings add MyHttpFunction_AUTH_LEVEL "function"
func settings add MyHttpFunction_KEY "313"

cd ..

echo "n" | func new --name Timer --template "TimerTrigger"


# Publish the Azure Functions to a function app
func azure functionapp publish functionnodeapp
