package db

import (
	// import gorm
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Name     string
	Host     string
	Port     string
	UserName string
	PassWord string
}

const (
	_defaultConnectSytle = "sqlite"
	_defaultLogLevel     = "info"
)

type DB struct {
	DBConfig
	gorm.DB
	ConnectStyle string
	LogLevel     string
}

var logLevelSettings = map[string]int{
	"info":  4,
	"warn":  3,
	"error": 2,
}

func New(Option ...Option) (*DB, error) {
	db := &DB{}
	{
		db.ConnectStyle = _defaultConnectSytle
		db.LogLevel = _defaultLogLevel
	}

	for _, opt := range Option {
		opt(db)
	}

	if err := db.verify(); err != nil {
		return nil, err
	}
	// open  sqlite db with gorm
	logLevel := logLevelSettings[db.LogLevel]
	if db.ConnectStyle == "sqlite" {
		handler, err := gorm.Open(sqlite.Open(db.Name), &gorm.Config{
			Logger: logger.Default.LogMode(logger.LogLevel(logLevel)),
		})
		if err != nil {
			return nil, err
		}
		db.DB = *handler
		return db, nil
	}

	// open mysql db with gorm
	if db.ConnectStyle == "mysql" {
		dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db.UserName, db.PassWord, db.Host, db.Port, db.Name)
		handler, err := gorm.Open(mysql.Open(dns), &gorm.Config{
			Logger: logger.Default.LogMode(logger.LogLevel(logLevel)),
		})
		if err != nil {
			return nil, err
		}
		db.DB = *handler
		return db, nil
	}
	return db, nil
}

var (
	ErrDBNameEmpty      = errors.New("db name is empty")
	ErrDBPassWordEmpty  = errors.New("db password is empty")
	ErrDBPortEmpty      = errors.New("db port is empty")
	ErrDBUserNameEmpty  = errors.New("db user name is empty")
	ErrHostEmpty        = errors.New("db host is empty")
	ErrLogLevelEmpty    = errors.New("db log level is empty")
	ErrLogLevelNotFound = errors.New("db log level not found")
	ErrConnectSytle     = errors.New("db connect style is not found")
)

func (db *DB) verify() error {
	if db.ConnectStyle != "sqlite" && db.ConnectStyle != "mysql" {
		return ErrConnectSytle
	}
	if db.Name == "" {
		return ErrDBNameEmpty
	}

	if db.ConnectStyle != "mysql" {
		return nil
	}

	return db.mysqlVerify()
}

func (db DB) mysqlVerify() error {
	if db.Host == "" {
		return ErrDBNameEmpty
	}
	if db.Port == "" {
		return ErrDBPortEmpty
	}
	if db.UserName == "" {
		return ErrDBUserNameEmpty
	}
	if db.PassWord == "" {
		return ErrDBPassWordEmpty
	}
	if db.LogLevel == "" {
		return ErrLogLevelEmpty
	}
	if _, ok := logLevelSettings[db.LogLevel]; !ok {
		return ErrLogLevelNotFound
	}
	return nil
}
