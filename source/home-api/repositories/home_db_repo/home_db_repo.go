package home_db_repo

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/golanshy/go-lambda-home-api/config"
	"github.com/golanshy/go-lambda-home-api/data_models"
	"github.com/golanshy/go-lambda-home-api/repositories/dtos"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var (
	mongoClient       *mongo.Client
	sensorsCollection *mongo.Collection
	mongoURI          string

	database        = "home-api-database"
	sensorsCollName = "sensors"
)

type Homes interface {
	GetHome(ctx context.Context, homeID string) (*data_models.Home, error)
	UpdateHome(ctx context.Context, home *data_models.Home) error
	InsertUnitData(ctx context.Context, data *data_models.Unit) error
}

func init() {

	type secrets struct {
		MongodbURI string `json:"mongodbURI"`
	}
	var secret secrets
	err := json.Unmarshal([]byte(os.Getenv("SECRETS")), &secret)
	if err != nil {
		log.Fatal(err)
	}
	mongoURI = secret.MongodbURI

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	var connectErr error
	mongoClient, connectErr = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if connectErr != nil {
		log.Fatal(connectErr)
	}
	sensorsCollection = mongoClient.Database(database).Collection(sensorsCollName)
}

type StoreRepository struct {
}

func NewRepository(c *config.Config) *StoreRepository {
	return &StoreRepository{}
}

func (s *StoreRepository) GetHome(ctx context.Context, homeID string) (*data_models.Home, error) {
	log.Printf("GetHome %s", homeID)
	return nil, errors.New("not_implemented")
}

func (s *StoreRepository) UpdateHome(ctx context.Context, home *data_models.Home) error {
	log.Printf("UpdateHome %+v", home)
	return errors.New("not_implemented")
}

func (s *StoreRepository) InsertUnitData(ctx context.Context, data *data_models.Unit) error {

	if data == nil {
		return errors.New("no unit data")
	}

	for _, sensor := range data.Sensors {
		if sensor != nil && sensor.TempData != nil {
			for _, temp := range sensor.TempData.Data {
				tempDTO := &dtos.TempRecordDTO{
					HomeId:      data.HomeId,
					UnitId:      data.Id,
					SensorId:    sensor.Id,
					Temperature: temp.TempInC,
					CreatedAt:   time.Now(),
				}
				log.Printf("sensorsCollection InsertOne: %+v", tempDTO)
				doc, err := sensorsCollection.InsertOne(ctx, tempDTO)
				if err != nil {
					log.Printf("failed inserting temperature document: %+v", err)
				} else {
					log.Printf("Successfully inserted temperature document: %+v", doc.InsertedID)
				}
			}
		}
	}
	return nil
}
