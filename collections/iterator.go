package collections

import (
	"errors"

	"github.com/bitbanshee/fungo/primitives"
)

type Iterator interface {
	primitives.Functor
	primitives.Closeable
	Next() primitives.Monad
}

type KeyValue struct {
	key, value interface{}
}

func IteratorFrom(items ...interface{}) Iterator {
	var data []interface{}
	data = append(data, items...)
	itch := make(chan interface{})
	closech := make(chan bool)
	go func() {
		for _, item := range data {
			select {
			case <-closech:
				close(itch)
				return
			case itch <- item:
			}
		}
		close(itch)
		close(closech)
	}()
	return &iterator{
		data,
		itch,
		closech,
	}
}

func MapToKeyValueSlice(m map[interface{}]interface{}) []interface{} {
	var data []interface{}
	for key, value := range m {
		data = append(data, KeyValue{key, value})
	}
	return data
}

type iterator struct {
	data    []interface{}
	itch    chan interface{}
	closech chan bool
}

func (it *iterator) Next() primitives.Monad {
	item, ok := <-it.itch
	if !ok {
		return &primitives.Nothing{}
	}
	return primitives.Just(item)
}

func (it *iterator) Close() (rerror error) {
	defer func() {
		if perror := recover(); perror != nil {
			rerror = errors.New("iterator already closed")
			return
		}
	}()
	it.closech <- true
	return nil
}

func (it *iterator) Map(mf primitives.MapFunc) primitives.Functor {
	return &daisyChainIterator{it, mf}
}

type daisyChainIterator struct {
	from Iterator
	mf   primitives.MapFunc
}

func (d *daisyChainIterator) Next() primitives.Monad {
	v := d.from.Next()
	if n, ok := v.(*primitives.Nothing); ok {
		return n
	}
	r := d.mf(v.Unit())
	if r == nil {
		return &primitives.Nothing{}
	}
	if n, ok := r.(*primitives.Nothing); ok {
		return n
	}
	return primitives.Just(r)
}

func (d *daisyChainIterator) Close() (rerror error) {
	return d.from.Close()
}

func (d *daisyChainIterator) Map(mf primitives.MapFunc) primitives.Functor {
	return &daisyChainIterator{d, mf}
}
