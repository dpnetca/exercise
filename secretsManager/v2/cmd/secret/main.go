package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dpnetca/exercise/secretsManager/v2/secret"
)

func main() {
	var mode, key, value, file, cryptKey string
	flag.StringVar(&file, "file", ".secret", "Secrets File, default .secret")
	flag.StringVar(&mode, "mode", "", "SET, GET, DELETE secret key/value pair")
	flag.StringVar(&key, "key", "", "Key to set, if key already exists it will be updated")
	flag.StringVar(&value, "value", "", "Secret value to save, required for SET, ignored for GET and DELETE")
	flag.StringVar(&cryptKey, "crypt", "", "Required encryption Key as flag or environment variable 'SECRET_CRPYT_KEY")

	flag.Parse()

	secrets := secret.NewVault(file, cryptKey)

	if len(cryptKey) == 0 {
		os.Getenv("SECRET_CRPYT_KEY")
	}
	if len(cryptKey) == 0 {
		log.Fatalln("Required encryption key not provided")
	}

	switch strings.ToLower(mode) {
	case "set":
		err := secrets.Set(key, value)
		if err != nil {
			log.Fatalf("Error setting Key: %v\n", err)
		}
		fmt.Println("Key Set")
	case "get":
		value, err := secrets.Get(key)
		if err != nil {
			log.Fatalf("Error Getting Key: %v\n", err)
		}
		fmt.Printf("secret: %v\n", value)

	case "delete":
		err := secrets.Delete(key)
		if err != nil {
			log.Fatalf("Error Deleting Key: %v\n", err)
		}
		fmt.Println("Key Deleted")
	default:
		log.Fatalf("In valde Mode, please set mode to SET, GET, or DELETE")

	}
}
