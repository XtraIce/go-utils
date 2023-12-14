package database

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

var username_, password_, dbName_ string
var credentials_ Credentials

type Credentials struct {
	Database []Database `json:"Database"`
}

type Database struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetCredentials(jsonFile string) bool {
	credentials_ = Credentials{}
	file, err := os.Open(jsonFile)
	if err != nil {
		fmt.Println("Error opening file, check if path is valid: ", err)
		return false
	}
	defer file.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := io.ReadAll(file)

	err = json.Unmarshal(byteValue, &credentials_)
	if err != nil {
		fmt.Println("Failed to unmarshal Credential type data, check json file: ", err)
		return false
	}

	if len(credentials_.Database) <= 0 {
		fmt.Printf("Credentials are empty. credential json incorrect.")
		return false
	}

	return true
}

func Connect(database string) bool {
	if len(credentials_.Database) == 0 {
		fmt.Printf("Cannot connect. Credentials empty.")
		return false
	}

	var dt Database
	for _, db := range credentials_.Database {
		if db.Name == database {
			dt = db
			break
		}
	}

	//object destructuring
	var dialector gorm.Dialector
	dsn := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", dt.Username, dt.Password, dt.Name)
	switch dt.Type {
	case "mysql":
		dialector = mysql.Open(dsn)
	case "sqlite":
		dialector = sqlite.Open(dsn)
	}

	d, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return false
	}
	db = d

	return true
}

func GetDB() *gorm.DB {
	return db
}
