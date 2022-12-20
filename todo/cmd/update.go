// /*
// Copyright © 2022 NAME HERE <EMAIL ADDRESS>
// */
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(cmd.Flags())
		if len(args) == 1 {
			todoName := args[0]
			for index, todo := range todos.Todos {
				if todo.Name == todoName {
					// check if --descriptions is passed as flag or not
					// fmt.Println(cmd.Flags())
					// fmt.Println(index, todo)
					if cmd.Flags().Lookup("description").Changed {
						description, _ := cmd.Flags().GetString("description")
						todos.Todos[index].Description = description
					}
					// check if --deadline is passed as flag or not
					if cmd.Flags().Lookup("deadline").Changed {
						deadline, _ := cmd.Flags().GetString("deadline")
						todos.Todos[index].Deadline = deadline

					}
				}

			}

			viper.Set("todos", todos.Todos)
			err := viper.WriteConfig()
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	},
}

func init() {
	fmt.Println("Update  init")
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)
	updateCmd.Flags().String("description", "", "Todo description")
	updateCmd.Flags().String("deadline", tomorrow.String(), "Deadline for the todo")

	rootCmd.AddCommand(updateCmd)
	fmt.Println("Update  init2")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
