package cache

import "github.com/fandrien/book-cabin/model"

type Cache interface {
	Get(key string) (*model.SearchResponse, bool)
	Set(key string, value *model.SearchResponse)
}
