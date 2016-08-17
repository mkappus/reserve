package store

import (
	"encoding"
	"encoding/binary"

	"github.com/boltdb/bolt"
)

type (
	Key []byte

	Storer interface {
		Key() Key
		encoding.BinaryMarshaler
		encoding.BinaryUnmarshaler
	}

	bucket struct {
		name []byte
		st   Storer
		vals []Storer
	}

	Store struct {
		db     *bolt.DB
		bucket *bucket
	}
)

var db *bolt.DB

func getDB() (*bolt.DB, error) {
	if db != nil {
		return db, nil
	}

	return bolt.Open("madison.db", 0600, nil)
}

func makeBucket(db *bolt.DB, name []byte) error {
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}

	if _, err = tx.CreateBucketIfNotExists(name); err != nil {
		return err
	}
	return tx.Commit()

}

func New(bucketName string, st Storer) (*Store, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	bn := []byte(bucketName)

	if err = makeBucket(db, bn); err != nil {
		return nil, err
	}

	return &Store{
		db: db,
		bucket: &bucket{
			name: bn,
			st:   st,
		},
	}, nil
}

// Proto object ids are uint64
func IdToKey(id uint64) (k Key) {
	binary.BigEndian.PutUint64(k, id)
	return
}

func nextKey(b *bolt.Bucket) Key {
	c := b.Cursor()
	_, val := c.Next()
	return Key(val)
}

func (s *Store) Create(st Storer) (k Key, err error) {
	data, err := st.MarshalBinary()
	if err != nil {
		return k, err
	}
	tx, err := s.db.Begin(true)
	if err != nil {
		return k, err
	}
	bucket := tx.Bucket(s.bucket.name)
	k = nextKey(bucket)
	if err = bucket.Put(k, data); err != nil {
		return k, err
	}
	return k, tx.Commit()
}

func (s *Store) Read(k Key) (Storer, error) {
	var data []byte
	err := s.db.View(func(tx *bolt.Tx) error {
		data = tx.Bucket(s.bucket.name).Get(k)
		return nil
	})
	if err != nil {
		return nil, err
	}
	err = s.bucket.st.UnmarshalBinary(data)
	return s.bucket.st, nil
}

func (s *Store) readAll
		
