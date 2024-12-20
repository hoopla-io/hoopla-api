package dto

type ImageDTO struct {
	FileName string `json:"filename"`
	FilePath string `json:"path"`
	FileExt  string `json:"ext"`
}