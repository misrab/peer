/*
	An interface to get or set known addresses. Implemented in something 
	like LevelDB or redis
*/

package peer 

import (
	// "log"
	"time"
	"errors"

	"github.com/boltdb/bolt"
)


type AddressStore interface {
	Get(key string) (string, error)
	Set(key, value string) error
}

type dbWrapper struct {
	db *bolt.DB
}


func NewAddressStore() (AddressStore, error) {
	// Open the my.db data file in your current directory.
    // It will be created if it doesn't exist.
    db, err := bolt.Open("./db/addresses.db", 0600, &bolt.Options{Timeout: 5 * time.Second})
    if err != nil { return nil, err }

    store := new(dbWrapper)
    store.db = db

    return store, nil
    // // defer db.Close()

}

// Insert data into a bucket.
func (store *dbWrapper) Set(key, value string) error {
	db := store.db
	err := db.Update(func(tx *bolt.Tx) error {
	    b, err := tx.CreateBucketIfNotExists([]byte("addresses"))
	    if err != nil { return err }
	    // b := tx.Bucket([]byte("a"))
	    b.Put([]byte(key), []byte(value))
	    // b.Put([]byte("susy"), []byte("que"))
	    return nil
	})

	return err
}


func (store *dbWrapper) Get(key string) (string, error) {
	db := store.db
	var result string

	err := db.View(func(tx *bolt.Tx) error {
	    b := tx.Bucket([]byte("addresses"))
	    if b == nil {
	    	return errors.New("No bucket with given name.")
	    }

	    v := b.Get([]byte(key))
	    result = string(v)
	    return nil
	})

	return result, err
}