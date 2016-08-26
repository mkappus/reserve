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
	}

	Store struct {
		db      *bolt.DB
		buckets map[string]*bucket
	}
)

var db *bolt.DB

func getDB() (*bolt.DB, error) {
	if db != nil {
		return db, nil
	}

	return bolt.Open("madison.db", 0600, nil)
}

func Open(path string, mode int) (*Store, error) {
	var err error
	db, err = bolt.Open(path, mode, nil)
	if err != nil {
		return nil, error
	}
	return &Store{
		db:      db,
		buckets: make(map[string]string),
	}, nil
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

func (s Store) NewBucket(bucketName string, st Storer) error {
	if err = makeBucket(s.db, bn); err != nil {
		return err
	}

	store.buckets[bucketName] = &bucket{
		db: db,
		bucket: &bucket{
			name: []byte(bucketName),
			st:   st,
		},
	}
	return nil
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

func (s *Store) Read(k Key, bucket string) (Storer, error) {
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

func (s *Store) ReadAll(bucket string) (all []Storer, err error) {
	var data []byte

	b := s.buckets[bucket]
	s.db.View(func(tx *bolt.Tx) {
		bucket := s.db.Bucket(b.name)
		c := bucket.Cursor()
		for k, v := c.First(); k != nil; k, v := c.Next() {
			err = b.st.MarshalBinary(v)
			if err != nil {
				continue
			}
			all = append(out, st)
			b.st = nil
		}
		return nil
	})

	return out, nil
}
