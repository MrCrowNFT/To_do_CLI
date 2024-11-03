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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		sqliteTaskDatabase, err := sql.Open("sqlite3", "./sqlite-task-database.db")
		if err != nil {
			log.Fatal(err)
		}

		rows, err := sqliteTaskDatabase.Query("SELECT * FROM taskTable")
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		fmt.Println("ID | Task                       | Deadline")
		fmt.Println("--------------------------------------------")

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// displayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// displayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
