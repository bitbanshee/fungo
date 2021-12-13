package collections

import "github.com/bitbanshee/fungo/primitives"

type Set interface {
	primitives.Functor
	Iterate() Iterator
}

func NewSet(items ...interface{}) Set {
	data := make(map[interface{}]bool)
	for _, item := range items {
		data[item] = true
	}
	return &set{data, nil}
}

type set struct {
	data map[interface{}]bool
	mfs  []primitives.MapFunc
}

func (set *set) Iterate() Iterator {
	var data []interface{}
	for key := range set.data {
		data = append(data, key)
	}
	if len(set.mfs) == 0 {
		return IteratorFrom(data...)
	}
	iterator := IteratorFrom(data...).(primitives.Functor)
	for _, mf := range set.mfs {
		iterator = iterator.Map(mf)
	}
	return iterator.(Iterator)
}

func (s *set) Map(mf primitives.MapFunc) primitives.Functor {
	data := make(map[interface{}]bool)
	for key, value := range s.data {
		data[key] = value
	}
	var mfs []primitives.MapFunc
	if len(s.mfs) > 0 {
		mfs = make([]primitives.MapFunc, len(s.mfs))
		copy(mfs, s.mfs)
	}
	mfs = append(mfs, mf)
	return &set{data, mfs}
}
