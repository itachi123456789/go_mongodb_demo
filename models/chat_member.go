package models

import (
	"go_mongodb_demo/core"
	"go_mongodb_demo/utils"

	"gopkg.in/mgo.v2/bson"
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

//Remove
func (u *ChatMembers) Remove(query bson.M) (err error) {
	c := core.GetmgoColl(u.TableName())
	err = c.Remove(query)
	return
}

//Insert 插入多条数据
func (u *ChatMembers) InsertBatch(udata []ChatMembers) (err error) {
	c := core.GetmgoColl(u.TableName())
	err = c.Insert(udata)
	if err != nil {
		utils.LogError("InsertBatch:%v\n", err)
	}
	return
}

//UpdateAll 更新多条
func (u *ChatMembers) UpdateAll(query, udata bson.M) error {
	c := core.GetmgoColl(u.TableName())
	_, err := c.UpdateAll(query, bson.M{"$set": udata})
	if err != nil {
		utils.LogError("UpdateAll:%v\n", err)
	}
	return err
}
