package home_db_repo

import (
	"context"
	"github.com/golanshy/go-lambda-home-api/config"
	"github.com/golanshy/go-lambda-home-api/data_models"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
	collection  *mongo.Collection
	mongoURI    string

	database = "home-api-database"
	collName = "sensors"
)

type Homes interface {
	GetSensorsData(ctx context.Context, homeID string) (*data_models.HomeData, error)
	InsertSensorsData(ctx context.Context, homeID string, data *data_models.HomeData) error
}

type storeRepository struct {
}

func NewRepository(c *config.Config) Homes {
	mongoURI = c.MongoURI

	//ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	//defer cancel()
	//var connectErr error
	//mongoClient, connectErr = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	//if connectErr != nil {
	//	logger.Error("error error connecting to mongo accounts_db", nil)
	//	panic(connectErr)
	//}
	//collection = mongoClient.Database(database).Collection(collName)

	return &storeRepository{}
}

func (s storeRepository) GetSensorsData(ctx context.Context, homeID string) (*data_models.HomeData, error) {
	//TODO implement me
	panic("implement me")
}

func (s storeRepository) InsertSensorsData(ctx context.Context, homeID string, data *data_models.HomeData) error {
	//TODO implement me
	panic("implement me")
}
