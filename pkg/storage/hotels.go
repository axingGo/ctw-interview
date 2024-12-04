// storage/storage.go

package storage

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type HotelInfo struct {
	QueueName      string
	HotelName      string
	Star           int
	Price          float64
	PriceBeforeTax float64
	CheckInDate    string
	CheckOutDate   string
	Guests         int
}

type Database struct {
	db *sql.DB
}

func NewDatabase(dsn string) (*Database, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}

func (d *Database) SaveResults(queueName string, hotels []*HotelInfo) error {
	for _, result := range hotels {
		_, err := d.db.Exec(`INSERT INTO hotel 
            (hotel_name, star, price, price_before_tax, check_in_date, check_out_date, guests,create_time,update_time) 
            VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, queueName,
			result.HotelName, result.Star, result.Price, result.PriceBeforeTax,
			result.CheckInDate, result.CheckOutDate, result.Guests, time.Now().Unix(), time.Now().Unix(),
		)
		if err != nil {
			return err
		}
	}
	return nil
}
