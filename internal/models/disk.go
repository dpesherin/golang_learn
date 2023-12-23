package models

type DiskItem struct {
	ID     int64  `json:"id"`
	Type   string `json:"type"`
	Path   string `json:"path"`
	Parent int64  `json:"parent"`
}
