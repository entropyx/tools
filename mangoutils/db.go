package tools

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/entropyx/mango"
	"github.com/entropyx/mango/options"
	mongotrace "github.com/entropyx/mongo-opencensus"
	"github.com/entropyx/protos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opencensus.io/trace"
)

var client *mongo.Client
var connection *mango.Connection

func init() {
	var err error
	var count uint8
	port, err := strconv.Atoi(os.Getenv("MONGO_PORT"))
	if err != nil {
		// panic(err)
	}
	config := &mango.Config{
		Address:  os.Getenv("MONGO_HOST"),
		Port:     uint(port),
		Database: os.Getenv("MONGO_DB"),
		Username: os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PASS"),
		Source:   os.Getenv("MONGO_SOURCE"),
	}
	config.SetMonitor(mongotrace.NewMonitor(mongotrace.WithServiceName("mongo debora"), mongotrace.WithSampler(trace.AlwaysSample())))
	for {
		connection, err = mango.Connect(config)
		if err == nil {
			break
		}
		// fmt.Println(err)
		count++
		if count == 30 {
			panic("unable to connect to mongo: " + err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}

func Register(c context.Context, models ...interface{}) {
	connection.Register(c, models...)
}

func GetClientMongo() *mongo.Client {
	return connection.GetClient()
}

func getConditions(selector *protos.Selector) (mango.M, error) {
	predicates := selector.Predicates
	conditions := mango.M{}
	if len(predicates) == 0 {
		return conditions, errors.New("You must set at least one predicate")
	}

	for _, predicate := range predicates {

		values := bson.A{}
		if predicate.Field == "id" {
			for _, value := range predicate.Values {

				obj, err := primitive.ObjectIDFromHex(value)
				if err != nil {
					return conditions, err
				}

				values = append(values, obj)
			}
			conditions["_id"] = mango.M{"$in": values}
		} else {
			for _, value := range predicate.Values {
				values = append(values, value)
			}
			conditions[predicate.Field] = mango.M{"$in": values}
		}
	}
	return conditions, nil
}

func getPage(selector *protos.Selector) (*options.Find, error) {
	paging := selector.Paging
	if paging != nil {
		page, err := strconv.ParseInt(paging.StartIndex, 10, 64)
		if err != nil {
			return nil, err
		}
		return &options.Find{Limit: int64(paging.Limit), Page: page}, nil
	}
	return nil, nil
}
