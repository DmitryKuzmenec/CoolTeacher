package database

import (
	"errors"
	"fmt"

	"github.com/DmitryKuzmenec/CoolTeacher/model"

	"github.com/DmitryKuzmenec/CoolTeacher/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB - struct of database
type DB struct {
	*gorm.DB
}

// Dial - connection to database
func Dial(conf *config.Config) (*DB, error) {
	user := conf.Mysql.User
	if user == "" {
		return nil, errors.New("mysql user is not present")
	}
	pass := conf.Mysql.Password
	if pass == "" {
		return nil, errors.New("mysql password is not present")
	}
	host := conf.Mysql.Host
	if host == "" {
		return nil, errors.New("mysql host is not present")
	}
	port := conf.Mysql.Port
	if port == "" {
		return nil, errors.New("mysql port is not present")
	}
	database := conf.Mysql.Database
	if database == "" {
		return nil, errors.New("mysql database is not present")
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, pass, host, port, database,
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

// Migrate - autocreate tables
func (db *DB) Migrate() error {
	return db.AutoMigrate(model.User{})
}
