package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type employee struct {
	id      uint
	name    string
	phone   string
	address string
}

func connetcDB() (*sql.DB, error) {
	//template connection DB
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/golang-basic1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Connection Success")
	return db, nil
}
func execUpdate() {
	db, err := connetcDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()
	id := 1
	name := "sopan mukti"
	phone := "340i34-3"

	_, err = db.Exec("UPDATE employees set name = ?,phone=? where id = ?", name, phone, id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("updated succes")

}

func execDelete() {
	db, err := connetcDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	id := 3
	_, err = db.Exec("DELETE from employees where id =?", id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("data berhasil di hapus")
}

func prepare() {
	db, err := connetcDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	q, err := db.Prepare("select id,name,phone,address from employees where id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var employe1 = employee{}
	q.QueryRow(2).Scan(&employe1.id, &employe1.name, &employe1.phone, &employe1.address)
	var employe2 = employee{}
	q.QueryRow(1).Scan(&employe2.id, &employe2.name, &employe2.phone, &employe2.address)

	fmt.Println("succes")
}
func queryRows() {
	db, err := connetcDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	employee := employee{}
	id := 3
	err = db.QueryRow("SELECT id,name,phone,address from employees where id = ?", id).Scan(&employee.id, &employee.name, &employee.phone, &employee.address)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(employee.id, employee.name, employee.phone, employee.address)
}

func queryGet() {
	db, err := connetcDB()

	if err != nil {
		fmt.Println(err.Error())
		return

	}
	defer db.Close()
	rows, err := db.Query("select id,name,phone,address from employees")

	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	var resp []employee

	for rows.Next() {
		var data = employee{}
		var err = rows.Scan(&data.id, &data.name, &data.phone, &data.address)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		resp = append(resp, data)
	}

	for _, d := range resp {
		fmt.Println(d)
	}
}

func execInsert() {
	db, err := connetcDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	name := "Sopan"
	phone := "0892382329"
	address := "jakarta"

	_, err = db.Exec("Insert into employees (name,phone,address)values(?,?,?)", name, phone, address)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success")

}
func main() {
	fmt.Println("SQL")
	// execInsert()
	//queryGet()
	// queryRows()
	// execUpdate()
	// execDelete()
	prepare()

}
