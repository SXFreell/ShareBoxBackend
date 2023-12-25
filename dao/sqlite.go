package dao

import (
	"os"
	"sharebox/utils"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
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
				utils.Log.WithFields(logrus.Fields{
					"err": err,
				}).Error("Create data folder failed")
			} else {
				utils.Log.Info("Create data folder success")
			}
		} else {
			utils.Log.WithFields(logrus.Fields{
				"err": err,
			}).Error("Stat data folder failed")
		}
	} else {
		utils.Log.Info("Data folder exists")
	}

	db, err = sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Open database failed")
	} else {
		utils.Log.Info("Open database success")
	}

	// Check tables
	tables := []string{"user", "text", "file"}
	for _, table := range tables {
		exists, err := ensureTableExists(table)
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"err": err,
			}).Error("Ensure table exists failed")
		} else if !exists {
			err := createTable(table)
			if err != nil {
				utils.Log.WithFields(logrus.Fields{
					"err": err,
				}).Error("Create table failed")
			}
		}
	}
}

func ensureTableExists(tableName string) (bool, error) {
	query := "SELECT name FROM sqlite_master WHERE type='table' AND name=?"
	var name string
	err := db.QueryRow(query, tableName).Scan(&name)

	if err != nil && err != sql.ErrNoRows {
		utils.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Query table failed")
		return false, err
	} else if err == sql.ErrNoRows {
		utils.Log.WithFields(logrus.Fields{
			"tableName": tableName,
		}).Info("Table not exists")
		return false, nil
	} else {
		utils.Log.WithFields(logrus.Fields{
			"tableName": tableName,
		}).Info("Table exists")
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
		utils.Log.WithFields(logrus.Fields{
			"tableName": tableName,
		}).Error("Table not exists")
		return nil
	}
	_, err := db.Exec(query, tableName)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Create table failed")
		return err
	} else {
		utils.Log.WithFields(logrus.Fields{
			"tableName": tableName,
		}).Info("Create table success")
		return nil
	}
}

func QuerySQL(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Query failed")
		return nil, err
	} else {
		utils.Log.Info("Query success")
		return rows, nil
	}
}

func ExecSQL(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Exec failed")
		return nil, err
	} else {
		utils.Log.Info("Exec success")
		return result, nil
	}
}
