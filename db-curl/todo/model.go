package todo

import (
	"database/sql"
	"fmt"
	"time"
)

func GetItems(db *sql.DB) ([]Items, error) {
	rows, err := db.Query("select * from test where is_completed=false")
	itm := []Items{}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("read row")
		for rows.Next() {
			data := Items{}
			rows.Scan(&data.Id, &data.Title, &data.IsCompleted, &data.CreatedAt, &data.CompletedAt)
			itm = append(itm, data)
			//fmt.Printf("%d - %s - %v - %v - %v\n", data.Id, data.Title, data.IsCompleted, data.CreatedAt, data.CompletedAt)
		}
	}
	return itm, nil
}

func InsertItems(db *sql.DB, title string) error {
	id, _ := GetNextId(db)
	CreatedAt := time.Now()
	CompletedAt := time.Time{}
	_, err := db.Exec("INSERT INTO test(id, title, is_completed, created_time, completed_time)VALUES ($1,$2,$3,$4,$5)", id, title, false, CreatedAt, CompletedAt)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func UpdateItems(db *sql.DB, id int, title string) error {
	_, err := db.Exec("update test set title=$1 where id=$2", title, id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func DeleteItems(db *sql.DB, id int) error {
	_, err := db.Exec("delete from test where id=$1", id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func GetNextId(db *sql.DB) (int, error) {
	rows, err := db.Query("SELECT MAX(id) FROM test")
	var maxId int = 0
	if err != nil {
		fmt.Println(err)
	} else {
		for rows.Next() {
			rows.Scan(&maxId)
		}
	}
	return maxId + 1, nil
}
