package cobracmd

import (
	"fmt"
	cobra "github.com/spf13/cobra"
	"strconv"
	"task/db"
)

// listCmd represents the list command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
			var ids []int
			for _,arg := range args{
				id,err := strconv.Atoi(arg)
				if err != nil{
					fmt.Println("failed to parse arg: ",arg)
				}
				ids = append(ids,id )
			}
			tasks, err := db.AllTasks()
			if err != nil{
				fmt.Println("something went wrong !!",err.Error())
				return
			}

			for _,id := range ids{
				if id <= 0 || id > len(tasks){
					fmt.Println("Invalid task number : ",id)
					continue
				}
				task := tasks[id-1]
				err := db.DeleteTask(task.Key)
				if err != nil{
					fmt.Println("failed to mark",id,"complete. Error: ",err)
				} else{
					fmt.Println("marked",id,"as completed")
				}
			}

			//fmt.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}

