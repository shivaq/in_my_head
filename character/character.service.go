package character

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/shivaq/in_my_head/cors"
	"golang.org/x/net/websocket"
)

const charactersPath = "characters"

func handleCharacters(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		characterList, err := getCharacterList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(characterList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:

		// type character struct {
		// 	id            string
		// 	characterType string
		// 	cloth         string
		// 	skin          string
		// 	description   string
		// }

		// var chara character

		r.ParseForm()
		log.Print(r.FormValue("nickname"))
		log.Print(r.Form)
		log.Print(reflect.TypeOf(r.Form))
		log.Print(r.Form["character-selection"][0])
		postedValue := r.Form["character-selection"][0]
		log.Print(postedValue)
		log.Print(reflect.TypeOf(postedValue))
		// json.Unmarshal([]byte(postedValue), &chara)
		// log.Print(chara.description)
		// log.Printf("\nID is %s and description is %s", chara.id, chara.description)
		log.Print(r.Header)
		log.Print(r.Header["Content-Type"])
		t := template.Must(template.ParseFiles("./frontend/html/confirmed_character.html"))

		t.Execute(w, r.FormValue("nickname"))
		// var character Character
		// // リクエストをデコードする
		// // リクエストをデコードして character のアドレスに格納
		// err := json.NewDecoder(r.Body).Decode(&character)
		// if err != nil {
		// 	log.Print(err)
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	return
		// }
		// characterID, err := insertCharacter(character)
		// if err != nil {
		// 	log.Print(err)
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	return
		// }
		// w.WriteHeader(http.StatusCreated)
		// w.Write([]byte(fmt.Sprintf(`{"characterId":%d}`, characterID)))
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleCharacter(w http.ResponseWriter, r *http.Request) {

	greeting := os.Getenv("GREETING")

	host, _ := os.Hostname()
	fmt.Fprintf(w, "%s, 世界さま\n", greeting)
	fmt.Fprintf(w, "Hostname: %s\n", host)

	// URL の Path を characters/ で split させて、Slice にする
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", charactersPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	characterID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		character, err := getCharacter(characterID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if character == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(character)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}

	case http.MethodPut:
		var character Character
		// リクエストボディをデコードして、character struct に格納
		err := json.NewDecoder(r.Body).Decode(&character)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if *character.CharacterID != characterID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = updateCharacter(character)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		err := removeCharacter(characterID)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Handler を登録する
func SetupRoutes(apiBasePath string) {
	charactersHandler := http.HandlerFunc(handleCharacters)
	characterHandler := http.HandlerFunc(handleCharacter)
	reportHandler := http.HandlerFunc(handleCharacterReport)
	// websocket のパスに行けば、自動更新されるのか？
	http.Handle("/websocket", websocket.Handler(characterSocket))
	// このパターンの場合は、このハンドラーへ
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, charactersPath), cors.HtmlCors(charactersHandler))
	// http.Handle(fmt.Sprintf("%s/%s", apiBasePath, charactersPath), cors.Middleware(charactersHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, charactersPath), cors.Middleware(characterHandler))
	http.Handle(fmt.Sprintf("%s/%s/reports", apiBasePath, charactersPath), cors.Middleware(reportHandler))
}
