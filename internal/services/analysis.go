package services

import (
	"context"
	"fmt"
	"house-scanner-backend/internal/db"
	"house-scanner-backend/internal/models"
	"house-scanner-backend/internal/repositories"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AnalysisService struct {
	repo *repositories.AnalysisRepository
}

func NewAnalysisService() *AnalysisService {
	return &AnalysisService{repo: repositories.NewAnalysisRepository()}
}

func (s *AnalysisService) CreateAnalysis(analysis *models.Analysis, fileName string) error {
	return s.repo.CreateAnalysis(analysis, fileName)
}

func (s *AnalysisService) GetAnalysis(id string) (*models.Analysis, error) {
	return s.repo.GetAnalysis(id)
}

func (s *AnalysisService) UpdateAnalysis(id string, analysis *models.Analysis) error {
	return s.repo.UpdateAnalysis(id, analysis)
}

func (s *AnalysisService) DeleteAnalysis(id string) error {
	return s.repo.DeleteAnalysis(id)
}

func (s *AnalysisService) GetAnalyses(name string, phone string) ([]models.Analysis, error) {
	return s.repo.GetAnalyses(name, phone)
}

// SubscribeToAnalysisUpdates subscribes to real-time updates for a specific analysis using MongoDB Change Streams
func (s *AnalysisService) SubscribeToAnalysisUpdates(analysisID string, callback func(*models.Analysis)) error {
	// Get MongoDB collection
	collection := db.GetMongoDB().Database("house_scanner").Collection("analysis")

	// Create a pipeline to filter changes for the specific analysis
	pipeline := mongo.Pipeline{
		{{"$match", bson.D{{"documentKey._id", analysisID}}}},
	}

	// Create a change stream
	changeStream, err := collection.Watch(context.Background(), pipeline)
	if err != nil {
		return fmt.Errorf("failed to create change stream: %v", err)
	}
	defer changeStream.Close(context.Background())

	// Start a goroutine to handle change events
	go func() {
		for changeStream.Next(context.Background()) {
			var changeEvent struct {
				OperationType string          `bson:"operationType"`
				DocumentKey   bson.D          `bson:"documentKey"`
				FullDocument  models.Analysis `bson:"fullDocument"`
			}

			if err := changeStream.Decode(&changeEvent); err != nil {
				log.Printf("Error decoding change event: %v", err)
				continue
			}

			// Call the callback function with the updated analysis
			callback(&changeEvent.FullDocument)
		}

		if err := changeStream.Err(); err != nil {
			log.Printf("Change stream error: %v", err)
		}
	}()

	// Keep the subscription alive
	<-context.Background().Done()

	return nil
}
