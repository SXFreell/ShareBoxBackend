package dao

import (
	"os"
	"sharebox/utils"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func init() {
	// 检测根目录是否有data文件夹
	_, err := os.Stat("./data")
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll("./data", os.ModePerm)
			if err != nil {
				utils.Log.Error("创建data文件夹失败", err)
			} else {
				utils.Log.Info("创建data文件夹成功", nil)
			}
		} else {
			utils.Log.Error("Stat data folder failed", err)
		}
	} else {
		utils.Log.Info("data文件夹存在", nil)
	}

	db, err = sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		utils.Log.Error("打开数据库失败", err)
	} else {
		utils.Log.Info("打开数据库成功", nil)
	}

	// Check tables
	tables := []string{"user", "text", "file"}
	for _, table := range tables {
		exists, err := ensureTableExists(table)
		if err != nil {
			utils.Log.Error("判断表是否存在出现错误", err)
		} else if !exists {
			err := createTable(table)
			utils.Log.Info("表不存在", utils.Mapper{"tableName": table})
			if err != nil {
				utils.Log.Error("创建表失败", err)
			} else {
				utils.Log.Info("创建表成功", utils.Mapper{"tableName": table})
			}
		} else {
			utils.Log.Info("表存在", utils.Mapper{"tableName": table})
		}
	}
}

func ensureTableExists(tableName string) (bool, error) {
	query := "SELECT name FROM sqlite_master WHERE type='table' AND name=?"
	var name string
	err := db.QueryRow(query, tableName).Scan(&name)

	if err != nil && err != sql.ErrNoRows {
		return false, err
	} else if err == sql.ErrNoRows {
		return false, nil
	} else {
		return true, nil
	}
}

func createTable(tableName string) error {
	query := ""
	switch tableName {
	case "user":
		query = user
	case "text":
		query = text
	case "file":
		query = file
	default:
		utils.Log.ErrorP("不支持创建该表", utils.Mapper{"tableName": tableName})
		return nil
	}
	_, err := db.Exec(query, tableName)
	if err != nil {
		return err
	}
	return nil
}

func QuerySQL(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	} else {
		return rows, nil
	}
}

func ExecSQL(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
