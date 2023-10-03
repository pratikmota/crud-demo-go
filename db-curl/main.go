package main

import (
	"bufio"
	"db-crud/db"
	"db-crud/todo"
	"fmt"
	"os"
	"strings"
)

func main() {

	dbConn, _ := db.GetDBConnection()

	for {
		fmt.Println("1 -> List all items")
		fmt.Println("2 -> Insert item")
		fmt.Println("3 -> Delete item")
		fmt.Println("4 -> Update item")
		fmt.Println("5 -> Complete item")
		fmt.Println("6 -> EXIT")

		fmt.Println("Enter Options: ")
		var opt int
		fmt.Scanln(&opt)
		switch opt {
		case 1:
			itms, _ := todo.GetItems(dbConn)
			for _, v := range itms {
				fmt.Println(v)
			}
		case 2:
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter Title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			_ = todo.InsertItems(dbConn, title)

			itms, _ := todo.GetItems(dbConn)
			for _, v := range itms {
				fmt.Println(v)
			}
		case 3:
			itms, _ := todo.GetItems(dbConn)
			for _, v := range itms {
				fmt.Println(v)
			}

			fmt.Println("Enter id: ")
			var opt int
			fmt.Scanln(&opt)
			_ = todo.DeleteItems(dbConn, opt)
		case 4:
			itms, _ := todo.GetItems(dbConn)
			for _, v := range itms {
				fmt.Println(v)
			}
			fmt.Println("Enter id: ")
			var opt int
			fmt.Scanln(&opt)
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter Title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			_ = todo.UpdateItems(dbConn, opt, title)

			itms2, _ := todo.GetItems(dbConn)
			for _, v := range itms2 {
				fmt.Println(v)
			}

		case 6:
			return
		}
	}
}
