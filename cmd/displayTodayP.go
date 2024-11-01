/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// displayTodayPCmd represents the displayTodayP command
var displayTodayPCmd = &cobra.Command{
	Use:   "displayTodayP",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("displayTodayP called")
	},
}

func init() {
	rootCmd.AddCommand(displayTodayPCmd)

	sqliteTaskDatabase, err := sql.Open("sqlite3", "./sqlite-task-database.db")
	if err != nil{
		log.Fatal(err)
	}

	rows, err := sqliteTaskDatabase.Query("SELECT * FROM taskTable")
	if err != nil{
		log.Fatal(err)
	}

	for rows.Next(){
		var Id int
		var Task string
		var Deadline string

		var _today time.Time = time.Now()
		today := _today.Format("2006/01/02")

		rows.Scan(&Id, &Task, &Deadline)
		if Deadline <= today {
			log.Println(Id, Task, Deadline)
		}
	}

	rows.Close()
}
