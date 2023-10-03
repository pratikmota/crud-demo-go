package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Items struct {
	Title       string
	IsCompleted bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

func main() {

	itemList := []Items{}

	for {

		fmt.Println("1 -> List all items")
		fmt.Println("2 -> Insert item")
		fmt.Println("3 -> Delete item")
		fmt.Println("4 -> Update item")
		fmt.Println("5 -> Complete item")
		fmt.Println("6 -> EXIT")

		fmt.Println("Enter Options: ")
		var opt string
		fmt.Scanln(&opt)

		switch opt {
		//List all items
		case "1":
			{
				if len(itemList) == 0 {
					fmt.Println("No Items available")
				}
				for i, v := range itemList {
					fmt.Println(i, "-", v)
				}

			}
		// Insert item
		case "2":
			{
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Enter Title: ")
				title, _ := reader.ReadString('\n')
				title = strings.TrimSpace(title)

				itm := Items{
					Title:       title,
					IsCompleted: false,
					CreatedAt:   time.Now(),
					CompletedAt: time.Time{},
				}
				itemList = append(itemList, itm)
				fmt.Println("Item Inserted")
				for i, v := range itemList {
					fmt.Println(i, "-", v)
				}

			}
			// Delete item
		case "3":
			{
				for i, v := range itemList {
					fmt.Println(i, "-", v)
				}

				fmt.Println("Enter Item number: ")
				var num int
				fmt.Scanln(&num)
				newitem := []Items{}
				newitem = append(newitem, itemList[:num]...)
				newitem = append(newitem, itemList[num+1:]...)
				itemList = newitem
			}
			// Update item
		case "4":
			{

				for i, v := range itemList {
					fmt.Println(i, "-", v)
				}

				fmt.Println("Enter Item number to UPDATE: ")
				var num int
				fmt.Scanln(&num)

				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Enter New Title: ")
				title, _ := reader.ReadString('\n')
				title = strings.TrimSpace(title)
				itemList[num].Title = title
				for i, v := range itemList {
					fmt.Println(i, "-", v)
				}
			}
			// Complete item
		case "5":
			{
				for i, v := range itemList {
					fmt.Println(i, "-", v)
				}

				fmt.Println("Enter Item number: ")
				var num int
				fmt.Scanln(&num)
				itemList[num].IsCompleted = true
				itemList[num].CompletedAt = time.Now()
				for i, v := range itemList {
					fmt.Println(i, "-", v)
				}

			}
			// Exit
		case "6":
			{
				return
			}

		}

	}

}
