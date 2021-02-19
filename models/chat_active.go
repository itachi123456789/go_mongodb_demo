package models

import (
	"go_mongodb_demo/core"
	"go_mongodb_demo/utils"

	"gopkg.in/mgo.v2/bson"
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

//Insert 添加单条数据
func (u *ChatActives) Insert() (err error) {
	c := core.GetmgoColl(u.TableName())
	err = c.Insert(u)
	if err != nil {
		utils.LogError("Insert:%v\n", err)
	}
	return
}

//All 获取全部
func (u *ChatActives) GetAll() []ChatActives {
	c := core.GetmgoColl(u.TableName())
	var mdata []ChatActives
	err := c.Find(bson.M{"lock": 1, "status": 99}).All(&mdata)
	if err != nil {
		utils.LogError("GetAll:%v\n", err)
	}

	return mdata
}

//Distinct 去重查询
func (u *ChatActives) GetByDistinct(query bson.M) []string {
	c := core.GetmgoColl(u.TableName())
	var (
		mdata []string
		err   error
	)
	err = c.Find(query).Select("mid").Distinct("mid", &mdata)
	if err != nil {
		utils.LogError("GetByDistinct:%v\n", err)
	}

	return mdata
}

//One  获取单条
func (u *ChatActives) GetOne(mid string) ChatActives {
	c := core.GetmgoColl(u.TableName())
	var mdata ChatActives
	err := c.Find(bson.M{"mid": mid, "lock": 1, "status": 99}).Limit(1).One(&mdata)
	if err != nil && err.Error() != "not found" {
		utils.LogError("GetOne:%v\n", err)
		return false
	}
	return mdata
}

//Update 更新
func (u *ChatActives) Update() error {
	c := core.GetmgoColl(u.TableName())

	err := c.Update(bson.M{"gid": u.Gid, "mid": u.Mid, "status": 99}, bson.M{"$set": bson.M{"lock": 1, "from": u.From, "uptime": utils.NstimeUnix(), "lotime": utils.NstimeUnix()}})

	if err != nil {
		utils.LogError("Update:%v\n", err)
	}
	return err
}
