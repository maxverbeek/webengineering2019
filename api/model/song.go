package model

type SongShort struct {
	ArtistId     string `json:"artist_id"`
	ReleaseId    int `json:"release_id"`
	SongId       string  `csv:"song.id" json:"id"`
	SongTitle    string  `csv:"song.title" json:"title"`
	SongYear     int     `csv:"song.year" json:"year"`
	SongDuration float64 `csv:"song.duration" json:"duration"`
}

// The Song model
// swagger:model Song
type Song struct {
	SongShort
	SongArtistMbtags            float64 `json:"artist_mbtags" csv:"song.artist_mbtags"`
	SongArtistMbtagsCount       float64 `json:"artist_mdtags_count" csv:"song.artist_mbtags_count"`
	SongBarsConfidence          float64 `json:"bars_confidence" csv:"song.bars_confidence"`
	SongBarsStart               float64 `json:"bars_start" csv:"song.bars_start"`
	SongBeatsConfidence         float64 `json:"beats_confidence" csv:"song.beats_confidence"`
	SongBeatsStart              float64 `json:"beats_start" csv:"song.beats_start"`
	SongEndFadeIn               float64 `json:"end_of_fade_in" csv:"song.end_of_fade_in"`
	SongHotttnesss              float64 `json:"hotttnesss" csv:"song.hotttnesss"`
	SongKey                     float64 `json:"key" csv:"song.key"`
	SongKeyConfidence           float64 `json:"key_confidence" csv:"song.key_confidence"`
	SongLoudness                float64 `json:"loudness" csv:"song.loudness"`
	SongMode                    int     `json:"mode" csv:"song.mode"`
	SongModeConfidence          float64 `json:"mode_confidence" csv:"song.mode_confidence"`
	SongStartFadeOut            float64 `json:"start_of_fade_out" csv:"song.start_of_fade_out"`
	SongTatumsConfidence        float64 `json:"tatums_confidence" csv:"song.tatums_confidence"`
	SongTatumsStart             float64 `json:"tatums_start" csv:"song.tatums_start"`
	SongTempo                   float64 `json:"tempo" csv:"song.tempo"`
	SongTimeSignature           float64 `json:"time_signature" csv:"song.time_signature"`
	SongTimeSignatureConfidence float64 `json:"time_signature_confidence" csv:"song.time_signature_confidence"`
}
