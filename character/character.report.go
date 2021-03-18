package character

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"
)

type CharacterReportFilter struct {
	UserIDFilter          string `json:"userID"`
	CharacterTypeFilter   string `json:"characterType"`
	CircleColorNameFilter string `json:"circleColorName"`
	IsPicUsedFilter       bool   `json:"isPicUsed"`
	NickNameFilter        string `json:"nickName"`
	FirstNameFilter       string `json:"firstName"`
	MiddleNameFilter      string `json:"middleName"`
	LastNameFilter        string `json:"lastName"`
	OriginalNameFilter    string `json:"originalName"`
	IsImaginaryFilter     bool   `json:"isImaginary"`
	BirthDateFilter       int    `json:"birthDate"`
	LevelFilter           int    `json:"level"`
	IsLikedFilter         bool   `json:"isLiked"`
	RegisteredDateFilter  int    `json:"registeredDate"`
	UpdatedDateFilter     int    `json:"updatedDate"`
	OnDeleteLockFilter    bool   `json:"onDeleteLock"`
	InLimboFilter         bool   `json:"inLimbo"`
}

func handleCharacterReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var characterFilter CharacterReportFilter
		err := json.NewDecoder(r.Body).Decode(&characterFilter)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		characters, err := searchForCharacterData(characterFilter)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		t := template.New("report.gotmpl").Funcs(template.FuncMap{"mod": func(i, x int) bool { return i%x == 0 }})
		t, err = t.ParseFiles(path.Join("templates", "report.gotmpl"))
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Disposition", "Attachment")

		var tpl bytes.Buffer
		err = t.Execute(&tpl, characters)
		rdr := bytes.NewReader(tpl.Bytes())
		http.ServeContent(w, r, "report.html", time.Now(), rdr)

	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
