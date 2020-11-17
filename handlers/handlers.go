package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	_ "github.com/mainak90/simple-urlshortner/driver"
	base62 "github.com/mainak90/simple-urlshortner/utils"
	"io/ioutil"
	"log"
	"net/http"
)

type DBClient struct {
	Db *sql.DB
}

type Record struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

func (driver *DBClient) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	var url string
	vars := mux.Vars(r)
	id := base62.ToBase10(vars["encoded string"])
	err := driver.Db.QueryRow("Select url from web_url where id = $1", id).Scan(&url)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		responseMap := map[string]interface{}{"url": url}
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

func (driver *DBClient) GenerateShortURL(w http.ResponseWriter, r *http.Request) {
	var record Record
	var id int
	url, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(url, &record)
	err := driver.Db.QueryRow("Insert into web_url(url) values $1 returning id", record.URL).Scan(&id)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		responseMap := map[string]interface{}{"encoded_string": base62.ToBase62(id)}
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

func Heartbeat(w http.ResponseWriter, r *http.Request) {
	log.Printf("Heartbeat looks ok...")
	w.Write([]byte("Heartbeat looks ok"))
}
