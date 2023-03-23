package types

type File struct {
	Name     string `json:"name"`
	Hash     string `json:"hash"`
	Size     int    `json:"size"`
	Mime     string `json:"mime"`
	IsPublic bool   `json:"is_public"`
	Parent   string `json:"parent"`
}
