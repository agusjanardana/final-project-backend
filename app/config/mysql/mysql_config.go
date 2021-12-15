package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"vaccine-app-be/app/config"
	"vaccine-app-be/drivers/records"
	"vaccine-app-be/exceptions"
)

type Client interface {
	Conn() *gorm.DB
	Close()
}

func New(configuration config.Config) Client {
	dsn := newMySQLConfig(configuration).String()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	exceptions.PanicIfError(err)

	err = db.AutoMigrate(records.Citizens{})
	exceptions.PanicIfError(err)

	log.Println("MySql Connected")
	return &client{db}
}

type client struct {
	db *gorm.DB
}

func (c *client) Conn() *gorm.DB {
	return c.db
}

func (c *client) Close() {
	sqlDB, err := c.db.DB()
	exceptions.PanicIfError(err)

	err = sqlDB.Close()
	exceptions.PanicIfError(err)
}

type mySQLConfig struct {
	Host string
	Port string
	User string
	Password string
	DBName string
}

func newMySQLConfig(configuration config.Config) *mySQLConfig {
	dbConfig := mySQLConfig{
		Host: configuration.Get("DB_HOST"),
		Port: configuration.Get("DB_PORT"),
		User: configuration.Get("DB_USER"),
		Password: configuration.Get("DB_PASSWORD"),
		DBName: configuration.Get("DB_NAME"),
	}

	return &dbConfig
}

func (dbConfig *mySQLConfig) String() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}