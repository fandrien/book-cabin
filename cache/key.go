package cache

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"

	"github.com/fandrien/book-cabin/model"
)

func BuildKey(req model.SearchRequest) string {

	bytes, _ := json.Marshal(req)

	hash := sha1.Sum(bytes)

	return hex.EncodeToString(hash[:])
}
