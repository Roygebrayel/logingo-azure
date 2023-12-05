package main

// import (
	


// 	"testing"
// )

// func TestInsertUser(t *testing.T) {
//     // Your test implementation here
//     err := InsertUser("testuser", "testpassword")
//     if err != nil {
//         t.Errorf("InsertUser returned an error: %v", err)
//     }
// }

// 	// Additional assertions
// 	// Assuming you have access to the hashed password function used in the insertUser function
// 	expectedHashedPassword, err := database.HashPassword(password)
// 	if err != nil {
// 		t.Errorf("Error hashing the test password: %v", err)
// 	}

// 	// Assert that the password was hashed correctly
// 	hashedPassword :=
// 	if hashedPassword != expectedHashedPassword {
// 		t.Errorf("Password was not hashed correctly. Expected: %s, Got: %s", expectedHashedPassword, hashedPassword)
// 	}

// 	// Assert that the credentials were saved in the database
// 	fetchedUser, err := getUserFromDatabase(username)
// 	if err != nil {
// 		t.Errorf("Error fetching user from the database: %v", err)
// 	}

// 	// Compare the fetched user's information with the expected values
// 	if fetchedUser.Username != username || fetchedUser.Password != expectedHashedPassword {
// 		t.Errorf("User information does not match. Expected: %s, Got: %s", username, fetchedUser.Username)
// 	}

// 	// Note: This test assumes that the database is properly configured,
// 	// and the necessary Azure Key Vault setup is in place. Adjust the test
// 	// accordingly based on your specific environment and requirements.
// }

// // Example function to fetch user from the database
// func getUserFromDatabase(username string) (User, error) {
// 	// Implement your database query logic here
// 	// and return the fetched user information

// 	// Example query using the database/sql package
// 	var user User
// 	err := db.QueryRow("SELECT username, password FROM users WHERE username = ?", username).Scan(&user.Username, &user.Password)
// 	if err != nil {
// 		return User{}, err
// 	}

// 	return user, nil
// }

// // Example user struct
// type User struct {
// 	Username string
// 	Password string
// }
