package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/dpnetca/exercise/secretsManager/secret"
)

func main() {
	var mode, key, value, file string
	flag.StringVar(&file, "file", ".secret", "Secrets File, default .secret")
	flag.StringVar(&mode, "mode", "", "SET, GET, DELETE secret key/balue pair")
	flag.StringVar(&key, "key", "", "Key to set, if key already exists it will be updated")
	flag.StringVar(&value, "value", "", "secret value to save, required for SET, ignored for GET and DELETE")

	flag.Parse()

	secrets, err := secret.NewVault(file)
	if err != nil {
		log.Fatalf("error opening secrets vault")
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
