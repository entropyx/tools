package tools

import (
	"context"
	"errors"
	"strconv"

	"github.com/entropyx/mango"
	"github.com/entropyx/mango/options"
	"github.com/entropyx/protos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var connection *mango.Connection

func Register(c context.Context, models ...interface{}) {
	connection.Register(c, models...)
}

func GetClientMongo() *mongo.Client {
	return connection.GetClient()
}

func GetConditions(selector *protos.Selector) (mango.M, error) {
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

func GetPage(selector *protos.Selector) (*options.Find, error) {
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
