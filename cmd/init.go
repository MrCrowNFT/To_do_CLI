/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"	
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize to do on CLI by creating the database",
	Long: `To do on CLI wil create a new sqlite database to store all the tasks and 
	deadlines, each with it own id for better manegment.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	exists, err := dbExists()
	if err != nil{
		log.Fatal(err)
	}
	if exists == true {
		fmt.Print("Already initialized\n")
		return
	}

	//Create task database
	file, err := os.Create("sqlite-task-database.db")
	fmt.Print("Tasks database created\n")
	if err != nil {
		log.Fatal(err)
	}
	
	file.Close()

	//open the database
	sqliteTaskDatabase, err := sql.Open("sqlite3", "./sqlite-task-database.db")
	if err != nil {
		log.Fatal(err)
	}

	//close the database once the main function finish executing

	//Create table
	taskDbTable := `CREATE TABLE taskTable(
		Id INTEGER NOT NULL AUTO_INCREMENT,
		Task TEXT NOT NULL,
		Deadline TEXT NOT NULL,
		PRIMARY KEY (Id));`
	
	fmt.Println("Creating Task Table")

	_, err = sqliteTaskDatabase.Exec(taskDbTable)
    if err != nil {
    log.Fatal(err)
    }

	//close the database
	sqliteTaskDatabase.Close()

	fmt.Println("Task Table Created Id/Task/Deadline")
}

func dbExists()(bool, error){
	//try openeing database
	_, err := os.Stat("sqlite-task-database.db")
	//if it opens returns true
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err){
		return false, nil
	}
	return false, err
}
