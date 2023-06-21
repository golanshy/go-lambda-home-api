package home_db_repo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golanshy/go-lambda-home-api/config"
	"github.com/golanshy/go-lambda-home-api/data_models"
	"github.com/golanshy/go-lambda-home-api/repositories/dtos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

var (
	mongoClient           *mongo.Client
	sensorsCollection     *mongo.Collection
	unitsCollection       *mongo.Collection
	temperatureCollection *mongo.Collection
	database              = "home-api-database"
	unitsCollName         = "units"
	sensorsCollName       = "sensors"
	temperatureCollName   = "temperature"
)

type Homes interface {
	GetHome(ctx context.Context, homeID string) (*data_models.Home, error)
	UpdateHome(ctx context.Context, home *data_models.Home) error
	InsertUnitData(ctx context.Context, data *data_models.Unit) error
	GetUnitData(ctx context.Context, unitId string) (*data_models.Unit, error)
	GetTempForSensor(ctx context.Context, unitId string, sensorId string, timeStamp time.Time) (float32, error)
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

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	log.Printf("Connecting to mongo")
	var connectErr error
	mongoClient, connectErr = mongo.Connect(ctx, options.Client().ApplyURI(secret.MongodbURI))
	if connectErr != nil {
		log.Printf("Failed to connect to mongo %s", connectErr.Error())
		log.Fatal(connectErr)
	}
	log.Printf("mongoClient connected")

	unitsCollection = mongoClient.Database(database).Collection(unitsCollName)
	sensorsCollection = mongoClient.Database(database).Collection(sensorsCollName)
	temperatureCollection = mongoClient.Database(database).Collection(temperatureCollName)
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

	unitDto := &dtos.UnitDTO{
		UnitId:    data.UnitId,
		HomeId:    data.HomeId,
		Sensors:   dtos.FromSensors(data.Sensors),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	log.Printf("InsertUnitData data: %+v", data)

	filter := bson.D{{"unit_id", data.UnitId}}
	update := bson.D{{"$set", bson.D{{"home_id", unitDto.HomeId}}}}
	opts := options.Update().SetUpsert(true)
	if _, err := unitsCollection.UpdateOne(ctx, filter, update, opts); err != nil {
		log.Printf("failed upserting units document: %+v", err)
		return fmt.Errorf("error upserting units document: %w", err)
	}

	for _, sensor := range data.Sensors {
		if sensor != nil && sensor.TempData != nil {

			filter := bson.D{{"sensor_id", sensor.SensorId}}
			update := bson.D{{"$set", bson.D{{"sensor_id", sensor.SensorId}, {"unit_id", data.UnitId}, {"updated_at", time.Now()}}}}
			opts := options.Update().SetUpsert(true)
			if _, err := sensorsCollection.UpdateOne(ctx, filter, update, opts); err != nil {
				log.Printf("failed upserting sensors document: %+v", err)
				return fmt.Errorf("error upserting sensors document: %w", err)
			}

			for _, temp := range sensor.TempData.Data {
				tempDTO := &dtos.TempRecordDTO{
					HomeId:      data.HomeId,
					UnitId:      data.UnitId,
					SensorId:    sensor.SensorId,
					Temperature: temp.TempInC,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}
				log.Printf("temperatureCollection InsertOne: %+v", tempDTO)
				doc, err := temperatureCollection.InsertOne(ctx, tempDTO)
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

func (s *StoreRepository) GetUnitData(ctx context.Context, unitId string) (*data_models.Unit, error) {

	if strings.TrimSpace(unitId) == "" {
		return nil, errors.New("invalid unit id")
	}

	unitDto := &dtos.UnitDTO{}

	filter := bson.M{"unit_id": unitId}
	if err := unitsCollection.FindOne(ctx, filter).Decode(unitDto); err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, errors.New("not found")
	}

	filter = bson.M{"unit_id": unitId}
	o := options.Find().SetSort(bson.M{"updated_at": -1})
	cur, err := sensorsCollection.Find(ctx, filter, o)
	if err != nil {
		log.Printf("failed to find sensors: %+v", err)
		return nil, fmt.Errorf("failed to find sensors: %w", err)
	}
	var results []*dtos.SensorDTO
	if err = cur.All(ctx, &results); err != nil {
		log.Printf("failed reading sensors documents: %+v", err)
		return nil, fmt.Errorf("failed reading sensors documents: %w", err)
	}
	sensors := dtos.FromSensorDTOs(results)
	if sensors != nil {
		for _, sensor := range sensors {
			if sensor != nil {
				filter := bson.D{{"unit_id", unitId}, {"home_id", unitDto.HomeId}, {"sensor_id", sensor.SensorId}, {"created_at", bson.M{
					"$gte": primitive.NewDateTimeFromTime(time.Now().Add(-time.Hour * 8)),
				},
				}}
				o := options.Find().SetSort(bson.M{"created_at": 1})

				cur, err := temperatureCollection.Find(ctx, filter, o)
				if err != nil {
					log.Printf("failed to find temperature documents: %+v", err)
					return nil, fmt.Errorf("failed to find temperature documents: %w", err)
				}
				var results []*dtos.TempRecordDTO
				if err = cur.All(ctx, &results); err != nil {
					log.Printf("failed reading temperature documents: %+v", err)
					return nil, fmt.Errorf("failed reading temperature documents: %w", err)
				}

				sort.Slice(results, func(i, j int) bool {
					return results[i].CreatedAt.Before(results[j].CreatedAt)
				})

				sensor.TempData = &data_models.TempData{
					Data: dtos.FromTempDTOs(results),
				}
			}
		}
	}

	unit := unitDto.ToUnit()
	unit.Sensors = sensors

	return unit, nil
}

func (s *StoreRepository) GetTempForSensor(ctx context.Context, unitId string, sensorId string, timeStamp time.Time) (float32, error) {

	filter := bson.D{
		{"unit_id", unitId},
		{"sensor_id", sensorId},
		{"created_at", bson.M{
			"$gte": primitive.NewDateTimeFromTime(timeStamp.Add(-time.Hour)),
		},
		},
		{"created_at", bson.M{
			"$lte": primitive.NewDateTimeFromTime(timeStamp.Add(time.Hour)),
		},
		}}

	singleResult := temperatureCollection.FindOne(ctx, filter)
	var result *dtos.TempRecordDTO
	findErr := singleResult.Decode(&result)
	if findErr != nil {
		log.Printf("failed reading temperature documents: %+v", findErr)
		return 0, findErr
	}

	return result.Temperature, nil
}
