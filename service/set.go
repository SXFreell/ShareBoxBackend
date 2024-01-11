package service

import (
	"sharebox/dao"
	"sharebox/model"
	"sharebox/utils"

	"github.com/kataras/iris/v12"
)

func SetSomething(ctx iris.Context) {
	var request model.SetReq
	err := ctx.ReadJSON(&request)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"message": "Bad request",
		})
		utils.Log.Error("Read JSON failed", err)
		return
	}

	if request.Type == "TEXT" {
		createTime := utils.GetTimeNow()
		var code string
		for {
			code = utils.GenerateCode()
			row, err := dao.QuerySQL("SELECT * FROM text WHERE code = ?", code)
			if err != nil {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.JSON(iris.Map{
					"message": "Internal server error",
				})
				utils.Log.Error("Select SQL failed", err)
				return
			} else if row.Next() {
				row.Close()
				continue
			} else {
				break
			}
		}
		_, err := dao.ExecSQL("INSERT INTO text (uid, code, content, expires, pickup_count, create_time, update_time) VALUES (?, ?, ?, ?, ?, ?, ?)",
			"uid000000",
			code,
			request.SetTextContent.Content,
			request.SetTextContent.Expires,
			request.SetTextContent.PickupCount,
			createTime,
			createTime,
		)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{
				"message": "Internal server error",
			})
			utils.Log.Error("Insert SQL failed", err)
			return
		} else {
			ctx.StatusCode(iris.StatusOK)
			ctx.JSON(iris.Map{
				"message": "OK",
				"code":    code,
			})
			utils.Log.Info("Insert SQL success", nil)
			return
		}
	} else {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"message": "Bad request",
		})
		utils.Log.Error("Unknown type", err)
		return
	}
}
