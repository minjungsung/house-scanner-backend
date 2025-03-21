package models

import "time"

type Analysis struct {
	ID               string         `json:"id,omitempty" bson:"_id,omitempty"`
	Name             string         `json:"name,omitempty" bson:"name,omitempty"`
	Phone            string         `json:"phone,omitempty" bson:"phone,omitempty"`
	Email            string         `json:"email,omitempty" bson:"email,omitempty"`
	Address          string         `json:"address,omitempty" bson:"address,omitempty"`
	RequestFileId    string         `json:"requestFileId,omitempty" bson:"requestFileId,omitempty"`
	AnalysisFileId   string         `json:"analysisFileId,omitempty" bson:"analysisFileId,omitempty"`
	RequestType      string         `json:"requestType,omitempty" bson:"requestType,omitempty"`
	RegisterNumber   string         `json:"registerNumber,omitempty" bson:"registerNumber,omitempty"`
	Status           AnalysisStatus `json:"status,omitempty" bson:"status,omitempty"`
	File             *File          `json:"file,omitempty" bson:"file,omitempty"`
	Title            string         `json:"title,omitempty" bson:"title,omitempty"`
	Content          string         `json:"content,omitempty" bson:"content,omitempty"`
	CreatedTimestamp time.Time      `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedTimestamp time.Time      `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type AnalysisStatus string

const (
	Pending    AnalysisStatus = "pending"
	Approved   AnalysisStatus = "approved"
	Rejected   AnalysisStatus = "rejected"
	Completed  AnalysisStatus = "completed"
	Failed     AnalysisStatus = "failed"
	Canceled   AnalysisStatus = "canceled"
	Processing AnalysisStatus = "processing"
)
