package gobstore_test

import (
	"testing"
	"time"

	"github.com/timshannon/gobstore"
)

func TestGet(t *testing.T) {
	testWrap(t, func(store *gobstore.Store, t *testing.T) {
		key := "testKey"
		data := &TestData{
			Name: "Test Name",
			Time: time.Now(),
		}
		err := store.Insert(key, data)
		if err != nil {
			t.Fatalf("Error creating data for get test: %s", err)
		}

		result := &TestData{}

		err = store.Get(key, result)
		if err != nil {
			t.Fatalf("Error getting data from gobstore: %s", err)
		}

		if !data.equal(result) {
			t.Fatalf("Got %s wanted %s.", result, data)
		}
	})
}

func TestFind(t *testing.T) {
	testWrap(t, func(store *gobstore.Store, t *testing.T) {
		key := "testKey"
		data := &TestData{
			Name: "Test Name",
			Time: time.Now(),
		}

		err := store.Insert(key, data)
		if err != nil {
			t.Fatalf("Error creating data for get test: %s", err)
		}

		var result []*TestData
		//var result []TestData
		// handle pointer conversion

		err = store.Find(&result, gobstore.Where("Name").Eq("Test Name"))

		if err != nil {
			t.Fatalf("Error finding data from gobstore: %s", err)
		}

	})
}
