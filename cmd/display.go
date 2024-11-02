/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"database/sql"
	"fmt"
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
		fmt.Println("display called")
	},
}

func init() {
	rootCmd.AddCommand(displayCmd)

	sqliteTaskDatabase, err := sql.Open("sqlite3", "./sqlite-task-database.db")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := sqliteTaskDatabase.Query("SELECT * FROM taskTable")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var Id int
		var Task string
		var Deadline string
		rows.Scan(&Id, &Task, &Deadline)
		log.Println(Id, Task, Deadline)
	}

	rows.Close()
}
