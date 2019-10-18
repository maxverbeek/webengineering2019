package api

import (
	"encoding/json"
	"log"
	"mime"
	"net/http"
	"strings"

	"webeng/api/model"

	"github.com/gocarina/gocsv"
)

type HttpResponse struct {
	status  int
	payload interface{}
}

// Determine whether the request `content-type` includes a
// server-acceptable mime-type
//
// Failure should yield an HTTP 415 (`http.StatusUnsupportedMediaType`)
func HasContentType(r *http.Request, mimetype string) bool {
	contentType := r.Header.Get("Content-type")
	if contentType == "" {
		return mimetype == "application/json"
	}

	for _, v := range strings.Split(contentType, ",") {
		t, _, err := mime.ParseMediaType(v)
		if err != nil {
			break
		}
		if t == mimetype {
			return true
		}
	}
	return false
}

func (response *HttpResponse) Render(w http.ResponseWriter, r *http.Request) {
	if HasContentType(r, "application/json") {
		response.RenderJSON(w, r)
	} else if r.Header.Get("Content-Type") == "text/csv" {
		response.RenderCSV(w, r)
	} else {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		log.Print(r.Header.Get("Content-Type"))
	}
}

func (response *HttpResponse) RenderJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(response.status)
	json.NewEncoder(w).Encode(response.payload)
}

func (response *HttpResponse) RenderCSV(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/csv")
	w.WriteHeader(response.status)

	var err error

	switch response.payload.(type) {

	// type assertions for multi object repsonses
	case []model.Song:
		err = gocsv.Marshal(response.payload.([]model.Song), w)
		break

	case []model.Artist:
		err = gocsv.Marshal(response.payload.([]model.Artist), w)
		break

	// some type assertions for single types
	case model.Song:
		err = gocsv.Marshal([]model.Song{response.payload.(model.Song)}, w)
		break

	case model.Artist:
		err = gocsv.Marshal([]model.Artist{response.payload.(model.Artist)}, w)
		break

	// we fucked
	default:
		log.Println("idk what to reply with")
		break
	}

	if err != nil {
		log.Print(err)
	}
}
