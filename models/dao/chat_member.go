package dao

import (
	"context"
	"go_mongodb_demo/core"
	"go_mongodb_demo/utils"

	"go.mongodb.org/mongo-driver/bson"
)

type ChatMembers struct {
	ID string `bson:"_id"`
}

func NewChatMember() *ChatMembers {
	return new(ChatMembers)
}

func (u *ChatMembers) TableName() string {
	return dbPrefix + "chat_members"
}

//DeleteOne删除
func (u *ChatMembers) DeleteOne(query bson.M) error {
	c := core.GetmgoCollT(u.TableName())
	_, err := c.DeleteOne(context.TODO(), query)
	if err != nil {
		utils.LogError("DeleteOne:%v\n", err)
		return err
	}
	return nil
}

//InsertMany 插入多条
func (u *ChatMembers) InsertBatch(udata []ChatMembers) error {
	c := core.GetmgoCollT(u.TableName())
	var idata = make([]interface{}, len(udata))
	for _, v := range udata {
		idata = append(idata, v)
	}
	_, err := c.InsertMany(context.TODO(), idata)
	return err
}
