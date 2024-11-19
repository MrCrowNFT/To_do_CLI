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
	Short: "Complete task Id",
	Long: `Completes the given task Id by deleting it from database.`,
	Run: func(cmd *cobra.Command, args []string) {
		var finished_id int

		// Open database
		sqliteTaskDatabase, err := sql.Open("sqlite3", "./sqlite-task-database.db")
		if err != nil {
			log.Fatal(err)
		}

		defer sqliteTaskDatabase.Close()

		// Ask for Id to complete from the user
		fmt.Printf("Enter the completed task id: ")
		fmt.Scanf("%v", &finished_id)

		// Delete selected task Id from the database
		result, err := sqliteTaskDatabase.Exec(`DELETE FROM taskTable WHERE Id = ?`, finished_id)
		if err != nil{
			fmt.Println("Error executing delete")
			log.Fatal(err)
		}

		// Check affected rows 
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
}
