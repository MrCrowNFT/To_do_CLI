/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"database/sql"
	"log"
	"time"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var year int
	var month time.Month
	var day int
	var Task string

	// Open the database
	sqliteTaskDatabase, err := sql.Open("sqlite3", "./sqlite-task-database.db")
	if err != nil {
		fmt.Println("You must first call 'todo init'")
		log.Fatal(err)
	}

	//Ask for task input from the user
	fmt.Print("Define the task: ")
	fmt.Scan(&Task)

	fmt.Printf("Add deadline date YYYY/MM/DD: ")
	fmt.Scanf("%v/%v/%v", &year, &month, &day)

	var _deadline time.Time = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	Deadline := _deadline.Format("2006/01/02")

	// The id is with autoincrement
	addTaskSQL := `INSERT INTO taskTable(Task, Deadline) VALUES(?, ?)`
	statement, err := sqliteTaskDatabase.Prepare(addTaskSQL)
	if err != nil {
		log.Fatal(err)
	}

	statement.Exec(Task, Deadline)
	sqliteTaskDatabase.Close()
		fmt.Println("Task added succesfully")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)	
}
