package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrianwudev/go-practice/task/cmd"
	"github.com/adrianwudev/go-practice/task/db"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
