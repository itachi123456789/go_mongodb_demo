package dao

import (
	"context"
	"go_mongodb_demo/core"
	"go_mongodb_demo/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ChatMessages struct {
	ID string `bson:"_id"`
}

func NewChatMessages() *ChatMessages {
	return new(ChatMessages)
}

func (u *ChatMessages) TableName() string {
	return dbPrefix + "chat_messages"
}

func (u *ChatMessages) GetChatMessagesList(query bson.M, sort string, SkipLimit []int) []ChatMessages {
	c := core.GetmgoCollT(u.TableName())
	skip := int64(SkipLimit[0])
	limit := int64(SkipLimit[1])
	ctx := context.TODO()
	cur, err := c.Find(ctx, query, &options.FindOptions{Sort: sort, Skip: &skip, Limit: &limit})
	if err != nil {
		utils.LogError("GetChatMessagesList:%v\n", err)
	}
	return CursorToChatMessages(ctx, cur)
}

func CursorToChatMessages(ctx context.Context, cur *mongo.Cursor) []ChatMessages {
	if ctx == nil {
		ctx = context.TODO()
	}

	var err error
	list := []ChatMessages{}
	for cur.Next(ctx) {
		var b ChatMessages
		if err = cur.Decode(&b); err != nil {
			utils.LogError("CursorToChatMessages:%v\n", err)
			return nil
		}

		list = append(list, b)
	}
	cur.Close(ctx)

	return list
}

type MessageMember struct {
	ID     string    `bson:"_id"`
	Member []Members `bson:"member"`
}

func (u *ChatMessages) GetMessageMemberList(query, sort bson.D, skipLimit []int) []MessageMember {
	lookup := bson.D{
		{
			"$lookup",
			bson.D{
				{"from", dbPrefix + "members"},
				{"localField", "uid"},
				{"foreignField", "_id"},
				{"as", "member"},
			},
		},
	}
	skip := bson.D{{"$skip", skipLimit[0]}}
	limit := bson.D{{"$limit", skipLimit[1]}}
	pipeline := mongo.Pipeline{query, lookup, sort, skip, limit}

	c := core.GetmgoCollT(u.TableName())
	ctx := context.TODO()
	cur, err := c.Aggregate(ctx, pipeline, nil)

	var odata []*MessageMember
	err = cur.All(ctx, &odata)
	if err != nil {
		utils.LogError("GetMessageMemberList:%v\n", err)
	}

	var mdata []MessageMember
	if err := cur.Err(); err != nil {
		utils.LogError("GetMessageMemberList:%v\n", err)
		return mdata
	}
	for cur.Next(ctx) {
		var b MessageMember
		if err = cur.Decode(&b); err != nil {
			utils.LogError("GetMessageMemberList:%v\n", err)
			return nil
		}
		mdata = append(mdata, b)
	}
	return mdata
}
