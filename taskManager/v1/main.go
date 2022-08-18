package main

import (
	"log"

	"github.com/dpnetca/exercise/taskManager/cmd"
	"github.com/dpnetca/exercise/taskManager/db"
)

func main() {
	cmd.Rootcmd.Execute()
	// course uses a package called "go-homedir" to create the db in the user homedir, but I want to keep it contained to local location for now
	err := db.Init("my.db")
	if err != nil {
		log.Fatalln(err)
	}
}
