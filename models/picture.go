package models

import "fmt"

type Picture struct {
	Model
	AlbumId uint   `json:"album_id"`
	Album   *Album `json:"-"`
	Caption string `json:"caption"`
	Path    string `json:"-"`
}

func (p Picture) String() string {
	return fmt.Sprintf("<%d: %s>", p.Id, p.Caption)
}

func (p *Picture) Valid() error {
	if p.AlbumId == 0 {
		return fmt.Errorf("invalid album_id: %d", p.AlbumId)
	}
	if len(p.Caption) > 50 {
		return fmt.Errorf("caption too long. Max 50 chars")
	}
	return nil
}
