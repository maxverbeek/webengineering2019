package model

// The artist model
// swagger:model Artist
type Artist struct {
	ArtistFamiliarity float64 `json:"familiarity" csv:"artist.familiarity"`
	ArtistHotttnesss  float64 `json:"hotttnesss" csv:"artist.hotttnesss"`
	ArtistId          string  `json:"id" csv:"artist.id"`
	ArtistLatitude    float64 `json:"latitude" csv:"artist.latitude"`
	ArtistLocation    int     `json:"location" csv:"artist.location"`
	ArtistLongitude   float64 `json:"longitude" csv:"artist.longitude"`
	ArtistName        string  `json:"name" csv:"artist.name"`
	ArtistSimilar     float64 `json:"similar" csv:"artist.similar"`
	ArtistTerms       string  `json:"terms" csv:"artist.terms"`
	ArtistTermsFreq   float64 `json:"terms_freq" csv:"artist.terms_freq"`
}
