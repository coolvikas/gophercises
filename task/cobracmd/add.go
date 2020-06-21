package cobracmd

import (
	"fmt"
	cobra "github.com/spf13/cobra"
	"os"
	"strings"
	"task/db"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {

		task := strings.Join(args," ")
		var err error
		if task != ""{
			_,err =db.CreateTask(task)
			if err != nil{
				fmt.Println("can't add task Something went wrong",err.Error())
				os.Exit(1)
			}
			fmt.Println("added task : ",task)

		}else {
			fmt.Println("please provide some valid task")
		}

	},
}

func init()  {
	RootCmd.AddCommand(addCmd)
}