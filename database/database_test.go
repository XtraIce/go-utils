package database

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func Test_GetDatabaseCredentials(t *testing.T) {
	home := os.Getenv("HOME")
	if cred := GetCredentials(path.Join(home, "/go/src/go-utils/database/test_creds.json")); cred == false {
		t.Fatalf("Failed To Get Credentials from file")
	}

	for i, c := range credentials_.Database {
		if c.Type == "" || c.Name == "" || c.Username == "" || c.Password == "" {
			t.Fatalf("Failed to get Any/All Credentials from struct %d \n Type=%s\tName=%s\tUsername=%s\tPassword=%s\n",
				i, c.Type, c.Name, c.Username, c.Password)
		}
		fmt.Printf("Type=%s\tName=%s\tUsername=%s\tPassword=%s\n", c.Type, c.Name, c.Username, c.Password)
	}
}

func Test_DatabaseConnection(t *testing.T) {
	home := os.Getenv("HOME")
	if cred := GetCredentials(path.Join(home, "/go/src/creds.json")); cred == false {
		t.Fatalf("Failed To Get Credentials from file")
	}

	if conn := Connect("golang"); conn == false {
		t.Failed()
	}
}
