package storage

import (
	"encoding/json"
	// "fmt"
	// "os"
	
	"go.etcd.io/bbolt"

	"termlock/internal/models"
)

var dbPath = "termlock.db"
var bucketName = []byte("entries")

func OpenDB() (*bbolt.DB, error) {
	return bbolt.Open(dbPath, 0600, nil)
}

func SaveEntries(db *bbolt.DB, entries []models.PasswordEntry, key []byte) error {
	return db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return err
		}
		for _, entry := range entries {
			encryptedPassword, err := Encrypt(entry.Password, key)
			if err != nil {
				return err
			}
			entry.Password = encryptedPassword
			val, err := json.Marshal(entry)
			if err != nil {
				return err
			}
			bucket.Put([]byte(entry.Title), val)
		}
		return nil
	})
}

func LoadEntries(db *bbolt.DB, key []byte) ([]models.PasswordEntry, error) {
	var entries []models.PasswordEntry
	err := db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return nil
		}

		return bucket.ForEach(func(k, v []byte) error {
			var entry models.PasswordEntry
			err := json.Unmarshal(v, &entry)
			if err != nil {
				return err
			}
			decryptedPassword, err := Decrypt(entry.Password, key)
			if err != nil {
				return err
			}
			entry.Password = decryptedPassword
			entries = append(entries, entry)
			return nil
		})
	})
	return entries, err
}
