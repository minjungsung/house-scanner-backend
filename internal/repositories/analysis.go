package repositories

import (
	"context"
	"house-scanner-backend/internal/db"
	"house-scanner-backend/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AnalysisRepository struct {
	db *mongo.Database
}

func NewAnalysisRepository() *AnalysisRepository {
	db := db.GetMongoDB()
	return &AnalysisRepository{db: db.Database("house_scanner")}
}

func (r *AnalysisRepository) CreateAnalysis(analysis *models.Analysis) error {
	_, err := r.db.Collection("analysis").InsertOne(context.Background(), analysis)
	return err
}

func (r *AnalysisRepository) GetAnalysis(id int) (*models.Analysis, error) {
	var analysis models.Analysis
	if err := r.db.Collection("analysis").FindOne(context.Background(), bson.M{"id": id}).Decode(&analysis); err != nil {
		return nil, err
	}
	return &analysis, nil
}

func (r *AnalysisRepository) UpdateAnalysis(id int, analysis *models.Analysis) error {
	_, err := r.db.Collection("analysis").UpdateOne(context.Background(), bson.M{"id": id}, bson.M{"$set": analysis})
	return err
}

func (r *AnalysisRepository) DeleteAnalysis(id int) error {
	_, err := r.db.Collection("analysis").DeleteOne(context.Background(), bson.M{"id": id})
	return err
}

func (r *AnalysisRepository) GetAnalyses() ([]models.Analysis, error) {
	var analyses []models.Analysis
	cursor, err := r.db.Collection("analysis").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var analysis models.Analysis
		if err := cursor.Decode(&analysis); err != nil {
			return nil, err
		}
		analyses = append(analyses, analysis)
	}
	return analyses, nil
}
