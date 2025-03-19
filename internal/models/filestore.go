package models

type File struct {
	Name string `json:"name" bson:"name"`
	Size int64  `json:"size" bson:"size"`
	Path string `json:"path" bson:"path"`
}