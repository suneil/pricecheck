package store

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

// Item item
type Item struct {
	ID    string    `bson:"id"`
	Name  string    `bson:"name"`
	Price float64   `bson:"price"`
	Ts    time.Time `bson:"ts"`
}

// NewItem creates a new item object
func NewItem(id string, name string, price float64) *Item {
	item := &Item{
		id,
		name,
		price,
		time.Now(),
	}

	return item
}

// Store stores item into database
func Store(item *Item) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("track").C("prices")
	err = c.Insert(item)
	if err != nil {
		log.Fatal(err)
	}
}
