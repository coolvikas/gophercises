package main

import (
	"fmt"
	"os"
	"path/filepath"
	cmd "task/cobracmd"
    "github.com/mitchellh/go-homedir"
	"task/db"
)

func must(err error)  {
	if err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
func main()  {
	//fmt.Println("using cobra for this exercise")
	home,_ := homedir.Dir()
	dbPath:=filepath.Join(home,"tasks.db")
    must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}
