package service

import (
	"sharebox/dao"
	"sharebox/utils"
)

func ClearExpiredFile() {
	result, err := dao.ExecSQL(`DELETE FROM text
		WHERE (expires < strftime('%s', 'now') AND expires != -1)
		OR (pickup_count <= 0 AND pickup_count != -1)`)
	if err != nil {
		utils.Log.Error("删除过期文本出错", err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		utils.Log.Error("获取删除过期文本数量出错", err)
	}
	utils.Log.Info("删除过期文本成功", utils.Mapper{"count": count})

	// // 查询过期文件
	// result, err := dao.QuerySQL(`SELECT * FROM file
	// 	WHERE (expires < strftime('%s', 'now') AND expires != -1)
	// 	OR (pickup_count <= 0 AND pickup_count != -1)`)
	// if err != nil {
	// 	Log.Error("查询过期文件出错", err)
	// }
	// for result.Next() {
	// 	var id int
	// 	var uid string
	// 	var code string
	// 	var name string
	// 	var path string
	// 	var expires int
	// 	var pickupCount int
	// 	var createTime string
	// 	var updateTime string
	// 	err := result.Scan(&id, &uid, &code, &name, &path, &expires, &pickupCount, &createTime, &updateTime)
	// 	if err != nil {
	// 		Log.Error("Scan SQL failed", err)
	// 		continue
	// 	}
	// 	// 删除文件
	// 	err = dao.DeleteFile(path)
	// 	if err != nil {
	// 		Log.Error("删除文件出错", err)
	// 		continue
	// 	}
	// 	// 删除数据库记录
	// 	_, err = dao.ExecSQL("DELETE FROM file WHERE id = ?", id)
	// 	if err != nil {
	// 		Log.Error("删除数据库记录出错", err)
	// 		continue
	// 	}
	// }
}
