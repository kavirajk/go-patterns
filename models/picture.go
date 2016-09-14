package models

type Picture struct {
	Model
	AlbumId uint   `json:"album_id"`
	Album   *Album `json:"-"`
	Caption string `json:"caption"`
	Path    string `json:"-"`
}
