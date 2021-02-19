package dao

import (
	"context"
	"go_mongodb_demo/core"
	"go_mongodb_demo/utils"
)

type Logs struct {
	ID string `bson:"_id"`
}

func NewLog() *Logs {
	return new(Logs)
}

func (u *Logs) TableName() string {
	return dbPrefix + "logs"
}

func (u *Logs) InsertBatch(udata []Logs) error {
	var idata = make([]interface{}, len(udata))
	for k, v := range udata {
		if v.Ctime == 0 {
			udata[k].Ctime = utils.NstimeUnix()
		}
		udata[k].Ts = utils.NstimeUnix() * 1000

		idata = append(idata, v)
	}

	c := core.GetmgoCollT(u.TableName())
	_, err := c.InsertMany(context.TODO(), idata)
	if err != nil {
		utils.LogError("InsertBatch:%v\n", err)
	}
	return err
}
