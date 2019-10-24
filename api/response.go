package api

import (
	"encoding/json"
	"log"
	"net/http"

	"webeng/api/model"

	"github.com/gocarina/gocsv"
	"github.com/golang/gddo/httputil"
)

type HttpResponse struct {
	status  int
	payload interface{}
}

type RestResponse struct {
	Data    interface{}       `json:"data,omitempty" csv:",omitempty"`
	Links   map[string]string `json:"links,omitempty" csv:"-"`
	Success bool              `json:"success"`
	Message string            `json:"message,omitempty" csv:",omitempty"`
}

func (response *HttpResponse) Render(w http.ResponseWriter, r *http.Request) {
	// parse the Accept header, and finds the best fitting content type
	contentType := httputil.NegotiateContentType(
		r,                                        // request
		[]string{"application/json", "text/csv"}, // accepts these
		"application/json",                       // default content type
	)

	switch contentType {
	case "application/json":
		response.RenderJSON(w, r)
	case "text/csv":
		response.RenderCSV(w, r)
	default:
		w.WriteHeader(http.StatusUnsupportedMediaType)
		log.Printf("Unsupported request content type: %s")
	}
}

func (response *HttpResponse) RenderJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(response.status)

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	encoder.Encode(response.payload)
}

func (response *HttpResponse) RenderCSV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/csv")
	w.WriteHeader(response.status)

	var err error

	switch response.payload.(type) {
	// type assertions for multi object repsonses
	case []model.Song:
		err = gocsv.Marshal(response.payload.([]model.Song), w)
	case []model.Artist:
		err = gocsv.Marshal(response.payload.([]model.Artist), w)
	// some type assertions for single types
	case model.Song:
		err = gocsv.Marshal([]model.Song{response.payload.(model.Song)}, w)
	case model.Artist:
		err = gocsv.Marshal([]model.Artist{response.payload.(model.Artist)}, w)
	case RestResponse:
		log.Print(response.payload.(RestResponse).Data)
		err = gocsv.Marshal(response.payload.(RestResponse).Data, w)
	// we fucked
	default:
		log.Println("idk what to reply with")
	}

	if err != nil {
		log.Print(err)
	}
}
