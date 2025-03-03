package persistencache

import (
	"bytes"
	"crypto/sha256"
	"database/sql"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	_ "modernc.org/sqlite" // SQLite driver
)

// Cache handles SQLite-based persistent caching with Gob encoding
type Cache struct {
	db *sql.DB
}

// NewCache initializes the cache database and table
func NewCache(dbFile string) *Cache {
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		log.Fatal(err)
	}

	// Create table for caching function results
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS cache (
		key TEXT PRIMARY KEY,
		value BLOB,
		expiration INTEGER
	);`)
	if err != nil {
		log.Fatal(err)
	}

	return &Cache{db: db}
}

// serialize encodes a Go object using Gob
func serialize(value interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(value)
	return buffer.Bytes(), err
}

// deserialize decodes a Gob-encoded byte array back into an object
func deserialize(data []byte, value interface{}) error {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	return decoder.Decode(value)
}

// Set stores a value in the cache with expiration
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	expiration := time.Now().Add(ttl).Unix()
	serializedValue, err := serialize(value)
	if err != nil {
		log.Println("Serialization error:", err)
		return
	}

	_, err = c.db.Exec("INSERT OR REPLACE INTO cache (key, value, expiration) VALUES (?, ?, ?)", key, serializedValue, expiration)
	if err != nil {
		log.Println("Cache Set Error:", err)
	}
}

// Get retrieves a value from the cache and deserializes it
func (c *Cache) Get(key string, value interface{}) bool {
	var data []byte
	var expiration int64

	err := c.db.QueryRow("SELECT value, expiration FROM cache WHERE key = ?", key).Scan(&data, &expiration)
	if err != nil {
		return false
	}

	// Check if expired
	if expiration < time.Now().Unix() {
		_, _ = c.db.Exec("DELETE FROM cache WHERE key = ?", key)
		return false
	}

	err = deserialize(data, value)
	return err == nil
}

// GenerateKey hashes function arguments into a unique cache key
func GenerateKey(args ...interface{}) string {
	hash := sha256.New()
	for _, arg := range args {
		data, _ := serialize(arg)
		hash.Write(data)
	}
	return hex.EncodeToString(hash.Sum(nil))
}

// Memoize decorates a function with caching
func Memoize(cache *Cache, ttl time.Duration, fn func(args ...interface{}) interface{}) func(...interface{}) interface{} {
	return func(args ...interface{}) interface{} {
		key := GenerateKey(args...)
		var cachedResult interface{}

		if cache.Get(key, &cachedResult) {
			fmt.Println("Cache hit:", args)
			return cachedResult
		}

		fmt.Println("Cache miss:", args)
		result := fn(args...)
		cache.Set(key, result, ttl)
		return result
	}
}

// Example struct to cache
type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	cache := NewCache("cache.db")
	defer cache.db.Close()

	// Example function returning a struct
	getUser := func(args ...interface{}) interface{} {
		time.Sleep(2 * time.Second) // Simulate slow database query
		return User{ID: args[0].(int), Name: "John Doe", Age: 30}
	}

	memoizedGetUser := Memoize(cache, 10*time.Second, getUser)

	// First call (slow, fetches data)
	fmt.Println(memoizedGetUser(1))

	// Second call (fast, retrieves from cache)
	fmt.Println(memoizedGetUser(1))
}
