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

// finishTaskCmd represents the finishTask command
var finishTaskCmd = &cobra.Command{
	Use:   "finishTask",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("finishTask called")
	},
}

func init() {
	rootCmd.AddCommand(finishTaskCmd)

	var finished_id int

	sqliteTaskDatabase, err := sql.Open("sqlite3", "./sqlite-task-database.db")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Enter the finished task id")
	fmt.Scanf("%v", &finished_id)

	sqliteTaskDatabase.Query(`DELETE FROM taskTable WHERE Id = finished_id`)
}
