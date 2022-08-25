package repository

import (
	"github.com/medium-stories/go-rabbitmq/order"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type sqliteRepo struct {
	Db *gorm.DB
}

type orderEntity struct {
	Identifier  string
	OrderStatus int
}

func NewSqlite(dbname string) *sqliteRepo {
	db, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logrus.Fatal(err)
	}

	if err = db.AutoMigrate(&orderEntity{}); err != nil {
		logrus.Fatal(err)
	}

	return &sqliteRepo{
		Db: db,
	}
}

func (storage *sqliteRepo) Save(bucket *order.Bucket) error {
	return storage.Db.Model(&orderEntity{}).Create(&orderEntity{
		Identifier:  bucket.Identifier,
		OrderStatus: bucket.OrderStatus,
	}).Error
}

func (storage *sqliteRepo) GetByIdentifier(identifier string) *order.Bucket {
	var orderData *orderEntity

	if result := storage.Db.Where("identifier", identifier).Find(&orderData); result.Error != nil {
		return nil
	}

	if orderData.Identifier == "" {
		return nil
	}

	return &order.Bucket{
		Identifier:  orderData.Identifier,
		OrderStatus: orderData.OrderStatus,
	}
}

func (storage *sqliteRepo) UpdateStatus(identifier string, status int) error {
	return storage.Db.Model(&orderEntity{}).
		Where("identifier", identifier).
		Update("order_status", status).Error
}
