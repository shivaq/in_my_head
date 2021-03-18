package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	// 下記は明示的には使わないので、_ をエイリアス？として指定している
	_ "github.com/go-sql-driver/mysql"
	"github.com/shivaq/in_my_head/character"
)

const basePath = "/api"

func main() {
	// database.SetupDatabase()
	character.SetupRoutes(basePath)

	// use PORT environment variable, or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	// . がないと、images が取得できない
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("./frontend/html"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./frontend/images"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./frontend/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./frontend/js"))))
	http.HandleFunc("/", home)

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func tesuto() string {
	return "世界マンですか？"
}

func random_name(num int) string {
	name := [10]string{"JoJo", "Son Gokuu", "Eto Ranze", "Janet Jackson", "Yo Kimiko",
		"Jim Morison", "Bob Dylan", "Kate Bush", "Joni Mitchel", "Ben Johnson"}
	return name[num]
}

func home(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	rand_num := rand.Intn(10)
	// サイト内リンクではない場合、下記のように返す？？
	t := template.Must(template.ParseFiles("./frontend/html/index.html"))

	t.Execute(w, random_name(rand_num))
}
