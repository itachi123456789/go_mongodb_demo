package dao

import (
	"context"
	"go_mongodb_demo/core"
	"go_mongodb_demo/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatGroups struct {
	ID string `bson:"_id"`
}

func NewChatGroup() *ChatGroups {
	return new(ChatGroups)
}

func (u *ChatGroups) TableName() string {
	return dbPrefix + "chat_groups"
}

type ChatGroupFile struct {
	ID     string             `bson:"_id"`
	Status int                `bson:"status"`
	Name   string             `bson:"name"`
	Tid    string             `bson:"tid"`
	Tpid   string             `bson:"tpid"`
	File   []CollaborateFiles `bson:"file"`
}

func (u *ChatGroups) GetChatGroupAndFileList(query bson.D) []ChatGroupFile {
	lookup := bson.D{
		{
			"$lookup",
			bson.D{
				{"from", dbPrefix + "collaborate_files"},
				{"localField", "fid"},
				{"foreignField", "_id"},
				{"as", "file"},
			},
		},
	}

	sort := bson.D{{"$sort", bson.D{{"ctime", -1}}}}
	pipeline := mongo.Pipeline{query, lookup, sort}

	c := core.GetmgoCollT(u.TableName())
	ctx := context.TODO()
	cur, err := c.Aggregate(ctx, pipeline, nil)

	var odata []*ChatGroupFile
	err = cur.All(ctx, &odata)
	if err != nil {
		utils.LogError("GetChatGroupAndFileList:%v\n", err)
	}

	var mdata []ChatGroupFile
	if err := cur.Err(); err != nil {
		utils.LogError("GetChatGroupAndFileList:%v\n", err)
		return mdata
	}
	for cur.Next(ctx) {
		var b ChatGroupFile
		if err = cur.Decode(&b); err != nil {
			utils.LogError("GetChatGroupAndFileList:%v\n", err)
			return nil
		}
		mdata = append(mdata, b)
	}

	return mdata
}

func (u *ChatGroups) GetChatGroupAndFile(query bson.D) ChatGroupFile {
	lookup := bson.D{
		{
			"$lookup",
			bson.D{
				{"from", dbPrefix + "collaborate_files"},
				{"localField", "fid"},
				{"foreignField", "_id"},
				{"as", "file"},
			},
		},
	}

	var mdata ChatGroupFile
	sort := bson.D{{"$sort", bson.D{{"ctime", -1}}}}
	pipeline := mongo.Pipeline{query, lookup, sort}

	c := core.GetmgoCollT(u.TableName())
	ctx := context.TODO()
	cur, err := c.Aggregate(ctx, pipeline, nil)

	if err = cur.Decode(&mdata); err != nil {
		utils.LogError("GetChatGroupAndFile:%v\n", err)
		return mdata
	}

	return mdata
}
