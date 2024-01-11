package service

import (
	"sharebox/dao"
	"sharebox/model"
	"sharebox/utils"
	"time"

	"github.com/kataras/iris/v12"
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
		utils.Log.Error("Read JSON failed", err)
		return
	}
	result, err := dao.QuerySQL("SELECT * FROM text WHERE code = ?", request.Code)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{
			"code":    50001,
			"message": "Internal server error",
		})
		utils.Log.Error("Select SQL failed", err)
		return
	}
	if result.Next() {
		var textData model.Text
		err := result.Scan(&textData.ID, &textData.UID, &textData.Code, &textData.Content, &textData.Expires, &textData.PickupCount, &textData.CreateTime, &textData.UpdateTime)
		result.Close()
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{
				"code":    50001,
				"message": "Internal server error",
			})
			utils.Log.Error("Scan SQL failed", err)
			return
		}
		if textData.PickupCount != -1 {
			if textData.PickupCount == 0 {
				DeleteText(request.Code)
				ctx.StatusCode(iris.StatusOK)
				ctx.JSON(iris.Map{
					"code":    20001,
					"message": "Not found",
				})
				return
			} else if textData.PickupCount == 1 {
				DeleteText(request.Code)
			}
			_, err := dao.ExecSQL("UPDATE text SET pickup_count = pickup_count - 1, update_time = ? WHERE code = ?", utils.GetTimeNow(), request.Code)
			if err != nil {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.JSON(iris.Map{
					"code":    50001,
					"message": "Internal server error",
				})
				utils.Log.Error("Update SQL failed", err)
				return
			}
		} else if textData.Expires != -1 && textData.Expires-time.Now().Unix() < 0 {
			DeleteText(request.Code)
			ctx.StatusCode(iris.StatusOK)
			ctx.JSON(iris.Map{
				"code":    20001,
				"message": "Not found",
			})
			return
		}
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"message": "OK",
			"code":    20000,
			"data": iris.Map{
				"content": textData.Content,
			},
		})
	} else {
		result.Close()
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"code":    20001,
			"message": "Not found",
		})
	}
}
