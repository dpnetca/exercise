package main

import (
	"fmt"
	"log"

	"github.com/dpnetca/exercise/secretsManager/v3/secret"
)

func main() {
	cryptKey := "replaceMe"
	secrets := secret.NewVault(".secret", cryptKey)

	secrets.Set("tweets", "abcd1234")
	secrets.Delete("tweets")
	secrets.Set("insta", "qwerty12345")
	api, err := secrets.Get("insta")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(api) // qwerty12345

	secrets.Set("insta", "qwerty1234500000")

	api, err = secrets.Get("insta")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(api) // qwerty1234500000

	secrets.Set("insta", "")
	api, err = secrets.Get("insta")
	if err != nil {
		log.Fatalln(err) // no value for key insta
	}

	fmt.Println(api) // never runs  ^^
}
