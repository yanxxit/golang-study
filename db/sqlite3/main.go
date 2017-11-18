package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	ID   int
	Name string
}

func (user User) String() string {
	format := `ID : %d Name: %s`
	return fmt.Sprintf(format, user.ID, user.Name)
}
func (user User) Save() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("insert into users(userId,uname) values('%d,%s'),user.ID,user.Name")
	//db.Exec("insert into users(userId,uname) values(4,'test')")
	db.Exec(fmt.Sprintf(`insert into users(userId,uname) values('%d','%s')`, user.ID, user.Name))
	if err != nil {
		log.Fatal(err)
	}
}
func saveUser(w http.ResponseWriter, req *http.Request) {
	userId, _ := strconv.Atoi(req.FormValue("Id"))
	name := req.FormValue("name")
	out := strconv.Itoa(userId) + "-" + name
	u := &User{ID: userId, Name: name}
	log.Println(out, ";", u.String())
	u.Save()
	//fmt.Fprintf(w, u.String())
	http.Redirect(w, req, "/", http.StatusFound)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl/addUser.html")
	t.Execute(w, nil)
}
func getUsers(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var id, name string
	users := ""
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err == nil {
			log.Println(id, name)
			users = users + id + "," + name + ";\n"
		}
	}
	fmt.Fprintf(w, users)
}
func main() {
	http.HandleFunc("/", addUser)
	http.HandleFunc("/user", saveUser)
	http.HandleFunc("/users", getUsers)
	err := http.ListenAndServe(":8880", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}


