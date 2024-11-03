/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"bufio"
	"os"
    "strings"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task to the database with a specified deadline",
	Long: `Adds a new task to the database with a specified deadline`,
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

		defer sqliteTaskDatabase.Close()

		//Ask for task input from the user
		reader := bufio.NewReader(os.Stdin)
        fmt.Print("Define the task: ")
        task, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal("Error reading task:", err)
        }

		// Remove newline and trailing spaces        
		Task = strings.TrimSpace(task) 

		fmt.Print("\nAdd deadline date YYYY/MM/DD: ")
		fmt.Scanf("%d/%d/%d", &year, &month, &day)

		var _deadline time.Time = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		Deadline := _deadline.Format("2006/01/02")

		// The id is with autoincrement
		addTaskSQL := `INSERT INTO taskTable(Task, Deadline) VALUES(?, ?)`
		statement, err := sqliteTaskDatabase.Prepare(addTaskSQL)
		if err != nil {
			log.Fatal(err)
		}

		statement.Exec(Task, Deadline)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Task added succesfully")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
