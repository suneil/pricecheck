package store

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	session *mgo.Session
)

func init() {
	_session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	_session.SetMode(mgo.Monotonic, true)
	session = _session
}

func logError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func getCollection() *mgo.Collection {
	collection := session.DB("track").C("prices")

	return collection
}

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
		time.Now().UTC(),
	}

	return item
}

// Store stores item into database
func Store(item *Item) {
	collection := getCollection()
    
    lastItem := getLastPrice(item.ID)
    
    if lastItem.Price == item.Price {
        return
    }
    
	err := collection.Insert(item)
	if err != nil {
		log.Fatal(err)
	}
}

func getLastPrice(productID string) Item {
	collection := getCollection()
	var item Item

	query := collection.Find(bson.M{"id": productID}).Sort("-ts")
    
    count, err := query.Count() 
	logError(err)
    
    if count > 0 {
        query.One(&item)
    }

	return item
}
