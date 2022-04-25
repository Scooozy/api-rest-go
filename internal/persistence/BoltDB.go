package persistence

import (
	"fmt"

	"log"

	"github.com/boltdb/bolt"
)

type BoltDbb struct {
	db *bolt.DB
}

func (bdb *BoltDbb) dbOpen(fileName string) {
	var err error
	bdb.db, err = bolt.Open(fileName, 0600, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

}

func (bdb *BoltDbb) dbClose() {
	defer bdb.db.Close()
}

func (bdb *BoltDbb) dbPath() string {
	return bdb.db.Path()
}

func (bdb *BoltDbb) createBucket(bucketName string) {
	bdb.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}

func (bdb *BoltDbb) dbPut(bucketName string, key string, value string) {
	bdb.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("get bucket: FAILED")
		}
		return bucket.Put([]byte(key), []byte(value))
	})
}

func (bdb *BoltDbb) dbGet(bucketName string, key string) string {
	var data string
	bdb.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("get bucket: FAILED")
		}
		value := b.Get([]byte(key))
		data = string(value)

		return nil
	})
	return data

}

func (bdb *BoltDbb) dbGetAll(bucketName string) []string {
	var datas []string
	bdb.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("get bucket: FAILED")
		}
		b.ForEach(func(k, v []byte) error {
			datas = append(datas, string(v))
			return nil
		})

		return nil
	})
	return datas
}

func (bdb *BoltDbb) dbDelete(bucketName string, key string) {
	bdb.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("get bucket: FAILED")
		}
		b.Delete([]byte(key))

		return nil
	})
}
