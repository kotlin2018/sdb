package main

import (
	"context"
	"github.com/yemingfeng/sdb/internal/store"
	"github.com/yemingfeng/sdb/internal/store/engine"
	"golang.org/x/sync/semaphore"
	"log"
	"math/rand"
	"strconv"
	"sync"
)

func randBytes1() []byte {
	return []byte("hello" + strconv.Itoa(rand.Int()%10000))
}

func main() {
	for i := 0; i < 10; i++ {
		_, _ = store.Set([]byte("h"+strconv.Itoa(i)), []byte("w"+strconv.Itoa(i)))
	}
	_ = store.Iterate(&engine.PrefixIteratorOption{Prefix: []byte("h"), Offset: -1, Limit: 3},
		func(key []byte, value []byte) error {
			log.Printf("key = %s, value = %s", key, value)
			return nil
		})
	log.Printf("=====")
	_ = store.Iterate(&engine.PrefixIteratorOption{Prefix: []byte("h"), Offset: -9, Limit: 3},
		func(key []byte, value []byte) error {
			log.Printf("key = %s, value = %s", key, value)
			return nil
		})
	log.Printf("=====")
	_ = store.Iterate(&engine.PrefixIteratorOption{Prefix: []byte("h"), Offset: 0, Limit: 3},
		func(key []byte, value []byte) error {
			log.Printf("key = %s, value = %s", key, value)
			return nil
		})
	log.Printf("=====")
	_ = store.Iterate(&engine.PrefixIteratorOption{Prefix: []byte("h"), Offset: 3, Limit: 5},
		func(key []byte, value []byte) error {
			log.Printf("key = %s, value = %s", key, value)
			return nil
		})
	s := semaphore.NewWeighted(100)

	var wg sync.WaitGroup
	for true {
		for i := 0; i < 100000; i++ {
			wg.Add(2)
			go func() {
				defer wg.Done()
				_ = s.Acquire(context.Background(), 1)
				_, _ = store.Set(randBytes1(), randBytes1())
				s.Release(1)
			}()
			go func() {
				defer wg.Done()
				_ = s.Acquire(context.Background(), 1)
				_, _ = store.Get(randBytes1())
				s.Release(1)
			}()
		}
		wg.Wait()
	}
}
