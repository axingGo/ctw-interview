// storage/storage.go

package storage

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type HotelInfo struct {
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

func (d *Database) SaveResults(hotels []*HotelInfo) error {
	for _, result := range hotels {
		_, err := d.db.Exec(`INSERT INTO hotel 
            (hotel_name, star, price, price_before_tax, check_in_date, check_out_date, guests) 
            VALUES (?, ?, ?, ?, ?, ?, ?)`,
			result.HotelName, result.Star, result.Price, result.PriceBeforeTax,
			result.CheckInDate, result.CheckOutDate, result.Guests,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
