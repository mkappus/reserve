package users

import (
	"github.com/golang/protobuf/proto"
	"github.com/mkappus/reserve/src/store"
)

// Teacher, Student and Course are Storers
func (t *Teacher) Key() store.Key {
	return store.IdToKey(t.Id)
}

func (t *Teacher) MarshalBinary() ([]byte, error) {
	return proto.Marshal(t)
}

func (t *Teacher) UnmarshalBinary(data []byte) error {
	return proto.Unmarshal(data, t)
}

func (s *Student) Key() store.Key {
	return store.IdToKey(s.Id)
}

func (s *Student) MarshalBinary() ([]byte, error) {
	return proto.Marshal(s)
}

func (s *Student) UnmarshalBinary(data []byte) error {
	return proto.Unmarshal(data, s)
}

func (c *Course) Key() store.Key {
	return store.IdToKey(c.Id)
}

func (c *Course) MarshalBinary() ([]byte, error) {
	return proto.Marshal(c)
}

func (c *Course) UnmarshalBinary(data []byte) error {
	return proto.Unmarshal(data, c)
}
