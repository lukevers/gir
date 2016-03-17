/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package storage

import (
	"github.com/boltdb/bolt"
	"fmt"
	"errors"
)

type Storage struct {
	db *bolt.DB
	table string
}

func New(dsn string) (*Storage, error) {
	db, err := bolt.Open(dsn, 0600, nil)
	if err != nil {
		return &Storage{}, fmt.Errorf("Error opening boltdb database: %s", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("default"))
		if err != nil {
			return fmt.Errorf("Error creating boltdb bucket: %s", err)
		}

		return nil
	})

	if err != nil {
		return &Storage{}, err
	}

	return &Storage{
		db: db,
		table: "default",
	}, nil
}

func (s *Storage) Use(table string) error {
	s.table = table
	return s.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(table))
		if err != nil {
			return fmt.Errorf("Error creating boltdb bucket: %s", err)
		}

		return nil
	})
}

func (s *Storage) Get(key string) (string, error) {
	var value string
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(s.table))
		if key, _ := b.Cursor().First(); key == nil {
			return errors.New("There is currently no data.")
		}

		v := b.Get([]byte(key))
		if v == nil {
			return errors.New("The key does not exist.")
		}

		value = string(v)

		return nil
	})

	return value, err
}

func (s *Storage) Put(key, value string) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(s.table))
		return b.Put([]byte(key), []byte(value))
	})
}

func (s *Storage) Remove(key string) error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(s.table))
		return b.Delete([]byte(key))
	})
	return err
}

