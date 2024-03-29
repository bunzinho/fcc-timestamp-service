package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type unixUTCTime struct {
	Unix int64  `json:"unix"`
	UTC  string `json:"utc"`
}

func timestamp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	t := time.Now()

	unix := &unixUTCTime{
		t.UnixMilli(),
		t.UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT"),
	}

	err := json.NewEncoder(w).Encode(unix)
	if err != nil {
		log.Println(err)
	}
}

func convert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	timeFromURL := strings.Split(r.URL.Path, "/")[3]
	writeJSON := json.NewEncoder(w)

	if timeFromURL == "" {
		t := time.Now()
		writeJSON.Encode(&unixUTCTime{
			t.UnixMilli(),
			t.UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT"),
		})
		return
	}

	ms, err := strconv.ParseInt(timeFromURL, 10, 64)
	if err == nil {
		t := time.UnixMilli(ms)
		err = writeJSON.Encode(&unixUTCTime{
			t.UnixMilli(),
			t.UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT"),
		})
		if err != nil {
			log.Println(err)
		}
		return
	}
	t, err := parseTime(timeFromURL)

	if err != nil {
		log.Println(err)
		w.Write([]byte(`{"error":"Invalid Date"}`))
		return
	}

	err = writeJSON.Encode(&unixUTCTime{
		t.UnixMilli(),
		t.UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT"),
	})
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.Handle("/", http.RedirectHandler("/api/timestamp", http.StatusTemporaryRedirect))
	http.HandleFunc("/api/timestamp", timestamp)
	http.HandleFunc("/api/timestamp/", convert)
	log.Println("Timestamp service started port 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
