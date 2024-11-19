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

// displayCmd represents the display command
var displayCmd = &cobra.Command{
	Use:   "display",
	Short: "Display all the pending tasks on the comand line",
	Long: `Display all the pending tasks on the comand line on a table showing the id, task and deadline`,
	Run: func(cmd *cobra.Command, args []string) {
		// Open database
		sqliteTaskDatabase, err := sql.Open("sqlite3", "./sqlite-task-database.db")
		if err != nil {
			log.Fatal(err)
		}

		// Select all rows from the database
		rows, err := sqliteTaskDatabase.Query("SELECT * FROM taskTable")
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		// Styling for the table on the command line
		fmt.Println("ID | Task                       | Deadline")
		fmt.Println("--------------------------------------------")

		// Print all the tasks on the command line
		for rows.Next() {
			var Id int
			var Task string
			var Deadline string
			err := rows.Scan(&Id, &Task, &Deadline)
			if err != nil {
				log.Fatal("Error scanning row:", err)
			}

			fmt.Printf("%-3d| %-25s | %s\n", Id, Task, Deadline)
		}

	},
}

func init() {
	rootCmd.AddCommand(displayCmd)

}
