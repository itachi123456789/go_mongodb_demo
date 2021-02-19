package models

import (
	"go_mongodb_demo/core"
	"go_mongodb_demo/utils"

	"gopkg.in/mgo.v2/bson"
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

func (u *ChatGroups) Count(gid string) int {
	c := core.GetmgoColl(u.TableName())
	iCount, err := c.Find(bson.M{"_id": gid, "status": 99}).Count()
	if err != nil && err.Error() != "not found" {
		utils.LogError("Count:%v\n", err)
	}

	return iCount
}

type ChatGroupFile struct {
	ID   string             `bson:"_id"`
	File []CollaborateFiles `bson:"file"`
}

func (u *ChatGroups) GetChatGroupAndFileList(query bson.M) []ChatGroupFile {
	pipeline := []bson.M{
		bson.M{"$match": query},
		bson.M{"$lookup": bson.M{"from": dbPrefix + "collaborate_files", "localField": "fid", "foreignField": "_id", "as": "file"}},
		bson.M{"$sort": bson.M{"ctime": -1}},
	}
	c := core.GetmgoColl(u.TableName())
	var mdata []ChatGroupFile
	err := c.Pipe(pipeline).All(&mdata)
	if err != nil {
		utils.LogError("GetChatGroupAndFileList:%v\n", err)
	}
	return mdata
}

type ChatGroupMember struct {
	ID     string    `bson:"_id"`
	Member []Members `bson:"member"`
}

func (u *ChatGroups) GetChatGroupMember(query bson.M) ChatGroupMember {
	pipeline := []bson.M{
		bson.M{"$match": query},
		bson.M{"$lookup": bson.M{"from": dbPrefix + "members", "localField": "primary", "foreignField": "_id", "as": "member"}},
	}
	c := core.GetmgoColl(u.TableName())
	var mdata ChatGroupMember
	err := c.Pipe(pipeline).One(&mdata)
	if err != nil && err.Error() != "not found" {
		utils.LogError("GetChatGroupMember:%v\n", err)
	}
	return mdata
}
