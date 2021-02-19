package dao

import (
	"context"
	"go_mongodb_demo/core"
	"go_mongodb_demo/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func Distinct(t string, field string, query bson.M) []interface{} {
	c := core.GetmgoCollT(t)
	idata, err := c.Distinct(context.TODO(), field, query)
	if err != nil {
		utils.LogError("Distinct:%v\n", err)
	}

	return idata
}
