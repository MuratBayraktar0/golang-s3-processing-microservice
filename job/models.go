package main

type Product struct {
	ID          int     `json:"id" bson:"_id"`
	Title       string  `json:"title" bson:"title"`
	Price       float64 `json:"price" bson:"price"`
	Category    string  `json:"category" bson:"category"`
	Brand       string  `json:"brand" bson:"brand"`
	URL         string  `json:"url" bson:"url"`
	Description string  `json:"description" bson:"description"`
}
