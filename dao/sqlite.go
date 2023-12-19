package dao

import (
	"sharebox/utils"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

var ()

func init() {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Open database failed")
	} else {
		utils.Log.Info("Open database success")
	}
	defer db.Close()

	// Check tables
	tables := []string{"user", "text", "file"}
	for _, table := range tables {
		exists, err := ensureTableExists(db, table)
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"err": err,
			}).Error("Ensure table exists failed")
		} else if !exists {
			err := createTable(db, table)
			if err != nil {
				utils.Log.WithFields(logrus.Fields{
					"err": err,
				}).Error("Create table failed")
			}
		}
	}
}

func ensureTableExists(db *sql.DB, tableName string) (bool, error) {
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

func createTable(db *sql.DB, tableName string) error {
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
