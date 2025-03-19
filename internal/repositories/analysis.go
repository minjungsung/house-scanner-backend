package repositories

import (
	"context"
	"house-scanner-backend/internal/db"
	"house-scanner-backend/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	// Convert struct to bson.M, which will automatically omit empty fields
	doc, err := bson.Marshal(analysis)
	if err != nil {
		return err
	}

	var bsonDoc bson.M
	if err := bson.Unmarshal(doc, &bsonDoc); err != nil {
		return err
	}

	// Insert the document
	_, err = r.db.Collection("analysis").InsertOne(context.Background(), bsonDoc)
	return err
}

func (r *AnalysisRepository) GetAnalysis(id string) (*models.Analysis, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var analysis models.Analysis
	if err := r.db.Collection("analysis").FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&analysis); err != nil {
		return nil, err
	}
	return &analysis, nil
}

func (r *AnalysisRepository) UpdateAnalysis(id string, analysis *models.Analysis) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	doc, err := bson.Marshal(analysis)
	if err != nil {
		return err
	}

	var bsonDoc bson.M
	if err := bson.Unmarshal(doc, &bsonDoc); err != nil {
		return err
	}

	_, err = r.db.Collection("analysis").UpdateOne(context.Background(), bson.M{"_id": objectID}, bson.M{"$set": bsonDoc})
	return err
}

func (r *AnalysisRepository) DeleteAnalysis(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.db.Collection("analysis").DeleteOne(context.Background(), bson.M{"_id": objectID})
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
