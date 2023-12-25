package service

import (
	"sharebox/dao"
	"sharebox/model"
	"sharebox/utils"

	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

func GetSomething(ctx iris.Context) {
	var request model.GetReq
	err := ctx.ReadJSON(&request)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"code":    40001,
			"message": "Bad request",
		})
		utils.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Read JSON failed")
		return
	}
	result, err := dao.QuerySQL("SELECT * FROM text WHERE code = ?", request.Code)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{
			"code":    50001,
			"message": "Internal server error",
		})
		utils.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Select SQL failed")
		return
	}
	if result.Next() {
		var id int
		var uid string
		var code string
		var content string
		var expires int
		var pickupCount int
		var createTime string
		var updateTime string
		err := result.Scan(&id, &uid, &code, &content, &expires, &pickupCount, &createTime, &updateTime)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{
				"code":    50001,
				"message": "Internal server error",
			})
			utils.Log.WithFields(logrus.Fields{
				"err": err,
			}).Error("Scan SQL failed")
			return
		}
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"message": "OK",
			"code":    20000,
			"data": iris.Map{
				"content": content,
			},
		})
	} else {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{
			"code":    40401,
			"message": "Not found",
		})
	}

}
