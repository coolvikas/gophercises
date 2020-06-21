package cobracmd

import (
	"fmt"
	cobra "github.com/spf13/cobra"
	"os"
	"task/db"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the tasks todo",
	Run: func(cmd *cobra.Command, args []string) {
		tasks,err := db.AllTasks()
		if err != nil{
			fmt.Println("Something went wrong",err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks todo ! Take a break !!")
			return
		}
		fmt.Println("You have the following tasks:")
		for i,task := range tasks{
			fmt.Println(i+1,":",task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
