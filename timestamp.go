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
	json, _ := json.Marshal(&unixUTCTime{
		unixNanoToMilliseconds(t.UnixNano()),
		t.UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT"),
	})
	w.Write(json)
}

func convert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	timeFromURL := strings.Split(r.URL.Path, "/")[3]
	writeJSON := json.NewEncoder(w)

	if timeFromURL == "" {
		t := time.Now()
		writeJSON.Encode(&unixUTCTime{
			unixNanoToMilliseconds(t.UnixNano()),
			t.UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT"),
		})
		return
	}

	milliseconds, err := strconv.ParseInt(timeFromURL, 10, 64)
	if err == nil {
		t := unixMillisecondsToTime(milliseconds)
		writeJSON.Encode(&unixUTCTime{
			unixNanoToMilliseconds(t.UnixNano()),
			t.UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT"),
		})
		if err != nil {
			log.Panicln(err)
		}
		return
	}
	t, err := parseTime(timeFromURL)

	if err != nil {
		log.Println(err)
		w.Write([]byte("{\"error\":\"Invalid Date\"}"))
		return
	}

	writeJSON.Encode(&unixUTCTime{
		unixNanoToMilliseconds(t.UnixNano()),
		t.UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT"),
	})
	if err != nil {
		log.Panicln(err)
	}
}

func main() {
	http.HandleFunc("/api/timestamp", timestamp)
	http.HandleFunc("/api/timestamp/", convert)
	log.Println("Timestamp service started port 9000")
	log.Fatalln(http.ListenAndServe(":9000", nil))
}
