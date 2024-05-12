package entity

import "time"

type ProductEntity struct {
	ID          int64     `bson:"_id"`
	Title       string    `bson:"title"`
	Price       float64   `bson:"price"`
	Category    string    `bson:"category"`
	Brand       string    `bson:"brand"`
	URL         string    `bson:"url"`
	Description string    `bson:"description"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}
