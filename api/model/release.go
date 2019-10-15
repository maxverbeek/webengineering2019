package model

// The release (album) model
// swagger:model Release
type Release struct {
	gorm.Model
	ReleaseId   int `csv:"release.id"`
	ReleaseName int `csv:"release.name"`
}
