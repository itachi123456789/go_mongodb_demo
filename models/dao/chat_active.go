package dao

import (
	"context"
	"errors"
	"go_mongodb_demo/core"
	"go_mongodb_demo/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ChatActives struct {
	ID string `bson:"_id"`
}

func NewChatActive() *ChatActives {
	return new(ChatActives)
}

func (u *ChatActives) TableName() string {
	return dbPrefix + "chat_actives"
}

func (u *ChatActives) Insert() (string, error) {
	c := core.GetmgoCollT(u.TableName())
	res, err := c.InsertOne(context.TODO(), u)
	if err != nil {
		utils.LogError("Insert:%v\n", err)
		return "", err
	}
	return res.InsertedID.(string), nil
}

func (u *ChatActives) Find() []ChatActives {
	c := core.GetmgoCollT(u.TableName())
	var mdata []ChatActives
	ctx := context.TODO()
	cur, err := c.Find(ctx, bson.M{"lock": 1, "status": 99}, &options.FindOptions{})
	if err != nil {
		utils.LogError("Find:%v\n", err)
		return mdata
	}
	if err := cur.Err(); err != nil {
		utils.LogError("Find:%v\n", err)
		return mdata
	}

	return CursorToChatActives(ctx, cur)
}

func CursorToChatActives(ctx context.Context, cur *mongo.Cursor) []ChatActives {
	if ctx == nil {
		ctx = context.TODO()
	}

	var err error
	list := []ChatActives{}
	for cur.Next(ctx) {
		var b ChatActives
		if err = cur.Decode(&b); err != nil {
			utils.LogError("CursorToChatActives:%v\n", err)
			return nil
		}

		list = append(list, b)
	}
	cur.Close(ctx)

	return list
}

func (u *ChatActives) Distinct(Mids []string) []string {
	query := bson.M{"mid": bson.M{"$in": Mids}, "lock": 1, "status": 99}
	idata := Distinct(u.TableName(), "mid", query)

	return utils.IntersToStrs(idata)
}

func (u *ChatActives) FindOne() ChatActives {
	c := core.GetmgoCollT(u.TableName())
	var mdata ChatActives
	err := c.FindOne(context.TODO(), bson.M{"gid": u.Gid, "mid": u.Mid, "lock": 1, "status": 99}).Decode(&mdata)
	if err != nil && err != mongo.ErrNoDocuments {
		utils.LogError("FindOne:%v\n", err)
	}
	return mdata
}

func (u *ChatActives) CountDocuments(mid string) int {
	c := core.GetmgoCollT(u.TableName())
	iCount, _ := c.CountDocuments(context.TODO(), bson.M{"mid": mid, "lock": 1, "status": 99})
	return iCount
}

func (u *ChatActives) UpdateOne() error {
	c := core.GetmgoCollT(u.TableName())
	rsp, err := c.UpdateOne(context.TODO(), bson.M{"gid": u.Gid, "mid": u.Mid, "status": 99}, bson.M{"$set": bson.M{"lock": 1, "from": u.From, "uptime": utils.NstimeUnix(), "lotime": utils.NstimeUnix()}})
	if err != nil {
		utils.LogError("UpdateOne:%v\n", err)
		return err
	}
	if rsp.ModifiedCount == 0 {
		return errors.New("update 0")
	}
	return err
}

func (u *ChatActives) UpdateMany(gid string, unix int64) error {
	c := core.GetmgoCollT(u.TableName())
	rsp, err := c.UpdateMany(context.TODO(), bson.M{"gid": gid, "status": 99, "lock": 1}, bson.M{"$set": bson.M{"status": 4, "lock": 0, "uptime": unix}})
	if err != nil {
		utils.LogError("UpdateMany:%v\n", err)
		return err
	}
	if rsp.ModifiedCount == 0 {
		return errors.New("update 0")
	}
	return err
}
