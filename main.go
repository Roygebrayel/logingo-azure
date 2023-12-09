package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"login/database"
	"login/templates"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"text/template"
	"time"

	_ "github.com/denisenkom/go-mssqldb"

	// "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

const (
	accountName   = "goroystorage"
	accountKey    = "7ptxtHcYKu3ih7LFrhJBHGpju6g0ah6Otc9mIxHP4sy2njfPFPFK6PY1GY2iNyHq3kJifJfucAFd+AStKFiXdA=="
	containerName = "uploadedfiles"
	uploadPath    = "storage"
	// azureBlobEndpoint = "https://" + accountName + ".blob.core.windows.net"
	azureBlobEndpoint = "https://goroystorage.blob.core.windows.net/uploadedfiles"
)

func main() {

	// Serve static files (e.g., stylesheets) from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handle GET requests to the root endpoint ("/")
	http.HandleFunc("/", templates.LoginHandleTmpl)

	// http.HandleFunc("/api/CredentialsPost",templates.LoginSuccessHandleTmpl)

	// http.HandleFunc("/login-failed", templates.LoginFailedHandleTmpl)
	http.HandleFunc("/upload-file", UploadFileHandleTmpl)
	http.HandleFunc("/users", GetUsersHandler)

	// Handle POST requests to the root endpoint ("/")
	// http.HandleFunc("CredentialsPost", CredentialsPosting)
	http.HandleFunc("/callapi", CallingApi)

	http.HandleFunc("/upload", StorageUpload)

	// Run the HTTP server
	fmt.Println("Server is running on :3777")
	listenAddr := ":3777"

	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.ListenAndServe(listenAddr, nil)
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	// Add any other fields as needed
}
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func insertUser(username, password string) error {


	// Define credentials as a struct
	credentials := struct {
		port     int
		user     string
		password string
		server   string
		database string
	}{
		port:     1433,
		user:     "roy",
		password: "ROge313313313313",
		server:   "sqlservergo.database.windows.net",
		database: "roysqldb",
	}

	// Construct connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		credentials.server, credentials.user, credentials.password, credentials.port, credentials.database)

	// Open a connection to the database
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error opening database:", err.Error())
	}
	defer db.Close()

	// Hash the password (use the existing hashPassword function)
	hashedPassword, err := database.HashPassword(password)
	if err != nil {
		return err
	}

	// Insert user into the database
	_, err = db.Exec("INSERT INTO users (username, password) VALUES (@p1, @p2)", username, hashedPassword)
	fmt.Println("it worked")

	return err
}

func CallingApi(w http.ResponseWriter, r *http.Request) {
	// Replace these values with actual credentials
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Create a Credentials struct with the provided values
	credentials := Credentials{
		Username: username,
		Password: password,
		// Add any other fields as needed
	}

	// Convert the Credentials struct to JSON
	payload, err := json.Marshal(credentials)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Create an HTTP client with timeout
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// Replace the URL with your actual API endpoint
	apiURL := "https://functionnodeappp.azurewebsites.net/api/CredentialsPosting?code=t-i5WYKe8O5cYIrWwH4b-6E1b6O_VwmcZVHy2V3RqurcAzFu79QDVA=="

	// Send a POST request to the API endpoint with the JSON payload
	resp, err := client.Post(apiURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code:", resp.StatusCode)
		return
	}
	// Insert user into the database (replace with your database logic)
	insertUser(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("welcome user")

}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Make a GET request to the API
	apiURL := "https://functionnodeappp.azurewebsites.net/api/GetRequest?code=WRH5BeKMYyDY_rJkWlhIoqtO_4bszOcqFizpP1fPp9DiAzFuBVbx2A=="
	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, "Error fetching data from API: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check the HTTP response status
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Unexpected status code: "+resp.Status, http.StatusInternalServerError)
		return
	}

	// Decode the JSON response
	var users []User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		http.Error(w, "Error decoding JSON response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Render HTML page with user data
	renderHTML(w, users)
}

func renderHTML(w http.ResponseWriter, users []User) {

	// Parse the HTML template
	tmpl, err := template.New("userList").Parse(htmlTemplate)
	if err != nil {
		http.Error(w, "Error parsing HTML template", http.StatusInternalServerError)
		return
	}

	// Execute the template with the user data
	err = tmpl.Execute(w, users)
	if err != nil {
		http.Error(w, "Error executing HTML template", http.StatusInternalServerError)
		return
	}
}

// }

const loginSuccessfulTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Login Successful</title>
</head>
<body>
    <h1>Welcome, {{.username}}!</h1>
</body>
</html>
`

const loginFailedTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Login Failed</title>
</head>
<body>
    <h1>Login failed. Please try again.</h1>
</body>
</html>
`

const htmlTemplate = `
		<!DOCTYPE html>
		<html>
		<head>
			<title>User List</title>
		</head>
		<body>
			<h1>User List</h1>
			<table border="1">
				<tr>
					<th>Username</th>
					<th>Password</th>
				</tr>
				{{range .}}
				<tr>
					<td>{{.Username}}</td>
					<td>{{.Password}}</td>
				</tr>
				{{end}}
			</table>
		</body>
		</html>
	`

func UploadFileHandleTmpl(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/upload-file.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}




func StorageUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error retrieving file from form", http.StatusBadRequest)
			return
		}
		defer file.Close()

		filePath := filepath.Join(uploadPath, header.Filename)
		out, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Error creating file", http.StatusInternalServerError)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			http.Error(w, "Error saving file", http.StatusInternalServerError)
			return
		}

		// Upload to Azure Storage
		err = uploadToAzureStorage(filePath, header.Filename)
		if err != nil {
			// Log the error message for debugging
			fmt.Println("Error uploading to Azure Storage:", err)

			// Return a more detailed error response to the client
			http.Error(w, fmt.Sprintf("Error uploading to Azure Storage: %v", err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "File uploaded successfully")
	}
}

func uploadToAzureStorage(filePath, filename string) error {
	// Log the start of the function
	log.Printf("Uploading file %s to Azure Storage\n", filePath)

	// Create SharedKeyCredential
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Println("Error creating SharedKeyCredential:", err)
		return err
	}

	// Create Pipeline
	pipeline := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	// Parse Azure Blob URL
	azureBlobURL, err := url.Parse(azureBlobEndpoint)
	if err != nil {
		log.Println("Error parsing Azure Blob URL:", err)
		return err
	}

	// Create Service URL
	serviceURL := azblob.NewServiceURL(*azureBlobURL, pipeline)

	// Create Container URL
	containerURL := serviceURL.NewContainerURL(containerName)

	// Create Blob URL
	blobURL := containerURL.NewBlockBlobURL(filename)

	// Open the local file
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening local file:", err)
		return err
	}
	defer file.Close()

	// Upload file to Azure Storage
	_, err = azblob.UploadFileToBlockBlob(
		context.TODO(),
		file,
		blobURL,
		azblob.UploadToBlockBlobOptions{
			BlockSize:   4 * 1024 * 1024, // 4MB block size
			Parallelism: 16,              // 16 workers
		},
	)

	// Log the result of the upload
	if err != nil {
		log.Println("Error uploading file to Azure Storage:", err)
	} else {
		log.Println("File uploaded successfully to Azure Storage")
	}

	return err
}
