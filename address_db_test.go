package peer

import (
	"testing"
	// "log"
)


func TestNewAddressStore(t *testing.T) {
	db, err := NewAddressStore()
	if err != nil { t.Errorf("Error opening bolt db: %s\n", err.Error()) }

	db.Set("foo", "boo")
	_, err = db.Get("foo")
	if err != nil { t.Errorf("Error getting from bolt db: %s\n", err.Error()) }
}