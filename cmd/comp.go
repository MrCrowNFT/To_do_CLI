/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// compCmd represents the comp command
var compCmd = &cobra.Command{
	Use:   "comp",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var finished_id int

		sqliteTaskDatabase, err := sql.Open("sqlite3", "./sqlite-task-database.db")
		if err != nil {
			log.Fatal(err)
		}

		defer sqliteTaskDatabase.Close()

		fmt.Printf("Enter the completed task id: ")
		fmt.Scanf("%v", &finished_id)

		result, err := sqliteTaskDatabase.Exec(`DELETE FROM taskTable WHERE Id = ?`, finished_id)
		if err != nil{
			fmt.Println("Error executing delete")
			log.Fatal(err)
		}

		rowsAffected, err := result.RowsAffected()
		if err !=nil {
			fmt.Println("Error checking affected rows")
			log.Fatal(err)
		}
		if rowsAffected == 0 {
			fmt.Println("No task with that Id exist on database")
		} else{
			fmt.Println("Task Finished succesfully")
		}

		fmt.Println("comp called")
	},
}

func init() {
	rootCmd.AddCommand(compCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
