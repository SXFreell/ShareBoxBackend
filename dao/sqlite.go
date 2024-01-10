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

type Mapper map[string]interface{}

func init() {
	// 检测根目录是否有data文件夹
	_, err := os.Stat("./data")
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll("./data", os.ModePerm)
			if err != nil {
				utils.Log.Error("Create data folder failed", err)
			} else {
				utils.Log.Info("Create data folder success", nil)
			}
		} else {
			utils.Log.Error("Stat data folder failed", err)
		}
	} else {
		utils.Log.Info("Data folder exists", nil)
	}

	db, err = sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		utils.Log.Error("Open database failed", err)
	} else {
		utils.Log.Info("Open database success", nil)
	}

	// Check tables
	tables := []string{"user", "text", "file"}
	for _, table := range tables {
		exists, err := ensureTableExists(table)
		if err != nil {
			utils.Log.Error("Ensure table exists failed", err)
		} else if !exists {
			err := createTable(table)
			if err != nil {
				utils.Log.Error("Create table failed", err)
			}
		}
	}
}

func ensureTableExists(tableName string) (bool, error) {
	query := "SELECT name FROM sqlite_master WHERE type='table' AND name=?"
	var name string
	err := db.QueryRow(query, tableName).Scan(&name)

	if err != nil && err != sql.ErrNoRows {
		utils.Log.Error("Query table failed", err)
		return false, err
	} else if err == sql.ErrNoRows {
		utils.Log.Info("Table not exists", Mapper{"tableName": tableName})
		return false, nil
	} else {
		utils.Log.Info("Table exists", Mapper{"tableName": tableName})
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
		utils.Log.ErrorP("Table not exists", Mapper{"tableName": tableName})
		return nil
	}
	_, err := db.Exec(query, tableName)
	if err != nil {
		utils.Log.Error("Create table failed", err)
		return err
	} else {
		utils.Log.Info("Create table success", Mapper{"tableName": tableName})
		return nil
	}
}

func QuerySQL(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		utils.Log.Error("Query failed", err)
		return nil, err
	} else {
		utils.Log.Info("Query success", nil)
		return rows, nil
	}
}

func ExecSQL(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		utils.Log.Error("Exec failed", err)
		return nil, err
	} else {
		utils.Log.Info("Exec success", nil)
		return result, nil
	}
}
