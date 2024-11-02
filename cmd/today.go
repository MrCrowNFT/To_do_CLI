/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("today called")
	},
}

func init() {
	displayCmd.AddCommand(todayCmd)

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

		var _today time.Time = time.Now()
		today := _today.Format("2006/01/02")

		rows.Scan(&Id, &Task, &Deadline)
		if Deadline == today {
			log.Println(Id, Task, Deadline)
		}
	}

	rows.Close()
}
