package models

import (
	"go_mongodb_demo/core"
	"go_mongodb_demo/utils"

	"gopkg.in/mgo.v2/bson"
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

//All  分页获取列表
func (u *ChatMessages) GetChatMessagesList(query bson.M, sort string, SkipLimit []int) []ChatMessages {
	c := core.GetmgoColl(u.TableName())
	var mdata []ChatMessages
	err := c.Find(query).Sort(sort).Skip(SkipLimit[0]).Limit(SkipLimit[1]).All(&mdata)

	if err != nil {
		utils.LogError("GetChatMessagesList:%v\n", err)
	}

	return mdata
}

//Count 总数
func (u *ChatMessages) Count(query bson.M) int {
	c := core.GetmgoColl(u.TableName())
	count, err := c.Find(query).Count()

	if err != nil {
		utils.LogError("Count:%v\n", err)
	}

	return count
}
