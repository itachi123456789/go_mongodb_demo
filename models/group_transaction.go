package models

import (
	"context"
	"errors"
	"go_mongodb_demo/core"
	"go_mongodb_demo/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateGroupByTransaction(m5 *ChatGroups, m2 *ChatMembers, m3 *ChatActives, m4 *TeamMembers) error {
	//创建群组
	var (
		db  = core.GetmgoDBT()
		ctx = context.Background()
		err error
	)

	//第一个事务：成功执行
	db.Client().UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		err = sessionContext.StartTransaction()
		if err != nil {
			utils.LogError("CreateGroupByTransaction:%v\n", err)
			return err
		}

		c5 := db.Collection(m5.TableName())
		_, err = c5.InsertOne(sessionContext, m5)
		if err != nil {
			utils.LogError("CreateGroupByTransaction:%v\n", err)
			return err
		}

		c2 := db.Collection(m2.TableName())
		_, err = c2.InsertOne(sessionContext, m2)
		if err != nil {
			utils.LogError("CreateGroupByTransaction:%v\n", err)
			sessionContext.AbortTransaction(sessionContext)
			return err
		}

		c3 := db.Collection(m3.TableName())
		_, err = c3.InsertOne(sessionContext, m3)
		if err != nil {
			utils.LogError("CreateGroupByTransaction:%v\n", err)
			sessionContext.AbortTransaction(sessionContext)
			return err
		}

		c4 := db.Collection(m4.TableName())
		query := bson.M{"_id": m4.ID, "mid": m4.Mid}
		rsp, err := c4.UpdateOne(sessionContext, query, bson.M{"$set": bson.M{"gid": m4.Gid}})
		if err != nil {
			utils.LogError("CreateGroupByTransaction:%v\n", err)
			sessionContext.AbortTransaction(sessionContext)
			return err
		}

		if rsp.ModifiedCount == 0 {
			utils.LogError("CreateGroupByTransaction update 0")
			sessionContext.AbortTransaction(sessionContext)
			return errors.New("update 0")
		}

		sessionContext.CommitTransaction(sessionContext)
		return nil
	})

	return nil
}
