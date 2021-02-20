package utils

import (
	"errors"
	"fmt"
	"os"
	"sync"

	ticketmodels "github.com/srgrcp/dvp-api/modules/ticket/models"
	"github.com/srgrcp/dvp-api/modules/user/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once     sync.Once
	db       *gorm.DB
	user     string
	password string
	host     string
	port     string
	database string
	debug    bool
)

func getEnvVariables() {
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	database = os.Getenv("DB_DATABASE")
	debug = os.Getenv("DB_DEBUG") == "true"
}

// MySQLConnection get a mysql connection
func MySQLConnection() *gorm.DB {
	getEnvVariables()
	once.Do(func() {
		var err error
		dsn := fmt.Sprintf(
			"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, password, host, port, database,
		)
		fmt.Println(dsn)
		db, err = gorm.Open(
			mysql.Open(dsn),
			&gorm.Config{},
		)
		if err != nil {
			panic(fmt.Sprintf("could not connect to the database: %s", err))
		}
	})
	return db
}

func generateUsers(db *gorm.DB) error {
	result := db.First(&models.User{})
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	if result.RowsAffected != 0 {
		return nil
	}
	users := []*models.User{
		{Name: "Andres", LastName: "Valencia", Email: "avalencia@test.com"},
		{Name: "Rafael", LastName: "Perez", Email: "rperez@test.com"},
		{Name: "Jose", LastName: "Rodriguez", Email: "jrodrigues@test.com"},
	}
	err := db.Create(&users).Error
	if err != nil {
		return err
	}
	return nil
}

// MigrateDatabase migrate database and create test data
func MigrateDatabase() error {
	db := MySQLConnection()
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{}, &ticketmodels.Ticket{})
	if err != nil {
		return err
	}
	err = generateUsers(db)
	if err != nil {
		return err
	}
	return nil
}
