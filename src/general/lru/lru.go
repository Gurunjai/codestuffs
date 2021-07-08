package lru

import (
	// "fmt"
	"gkrish/list"
)

var lCache *lruCache

type lruCache struct {
	list.DoubleList
}

func addToCache(key int) {
	
}