package main

import "fmt"

// Entry represents a key-value pair in the hash map
type Entry struct {
	key   string
	value int
}

// HashMap represents the hash map using quadratic probing
type HashMap struct {
	buckets   []Entry
	capacity  int
	size      int
	deleted   int
	loadLimit float64
}

// NewHashMap creates a new instance of the hash map
func NewHashMap() *HashMap {
	return &HashMap{
		buckets:   make([]Entry, 16),
		capacity:  16,
		loadLimit: 0.75,
	}
}

// hashFunction returns the hash value for a given key
func (hm *HashMap) hashFunction(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash<<5 - hash) + int(key[i])
	}
	return hash & (hm.capacity - 1)
}

// rehash increases the capacity and rehashes all elements
func (hm *HashMap) rehash() {
	hm.capacity *= 2
	newBuckets := make([]Entry, hm.capacity)
	hm.deleted = 0

	for i := 0; i < len(hm.buckets); i++ {
		if hm.buckets[i].key != "" {
			hm.insertEntry(newBuckets, hm.buckets[i])
		}
	}

	hm.buckets = newBuckets
}

// insertEntry inserts an entry into the hash map's buckets using quadratic probing
func (hm *HashMap) insertEntry(buckets []Entry, entry Entry) {
	index := hm.hashFunction(entry.key)
	step := 1

	for buckets[index].key != "" {
		index = (index + step*step) & (hm.capacity - 1) // Quadratic probing to find the next available slot
		step++
	}
	buckets[index] = entry
}

// Put inserts a key-value pair into the hash map
func (hm *HashMap) Put(key string, value int) {
	loadFactor := float64(hm.size) / float64(hm.capacity)

	if loadFactor >= hm.loadLimit {
		hm.rehash()
	}

	entry := Entry{
		key:   key,
		value: value,
	}

	hm.insertEntry(hm.buckets, entry)
	hm.size++
}

// Get retrieves the value associated with a given key from the hash map
func (hm *HashMap) Get(key string) (int, bool) {
	index := hm.hashFunction(key)
	step := 1

	for hm.buckets[index].key != "" {
		if hm.buckets[index].key == key {
			return hm.buckets[index].value, true
		}
		index = (index + step*step) & (hm.capacity - 1) // Quadratic probing to find the next slot
		step++
	}
	return 0, false
}

// Remove removes a key-value pair from the hash map
func (hm *HashMap) Remove(key string) bool {
	index := hm.hashFunction(key)
	step := 1

	for hm.buckets[index].key != "" {
		if hm.buckets[index].key == key {
			// Mark the entry as deleted instead of clearing it
			hm.buckets[index].key = "DELETED"
			hm.deleted++
			hm.size--
			return true
		}
		index = (index + step*step) & (hm.capacity - 1) // Quadratic probing to find the next slot
		step++
	}
	return false
}

// ContainsKey checks if the hash map contains a specific key
func (hm *HashMap) ContainsKey(key string) bool {
	_, exists := hm.Get(key)
	return exists
}

func main() {
	hm := NewHashMap()

	hm.Put("apple", 5)
	hm.Put("banana", 2)
	hm.Put("orange", 8)

	val, exists := hm.Get("apple")
	if exists {
		fmt.Println("Value for 'apple':", val)
	}

	hm.Remove("banana")

	fmt.Println("Contains 'banana':", hm.ContainsKey("banana"))
	fmt.Println("Contains 'orange':", hm.ContainsKey("orange"))
	fmt.Println("Size of the hash map:", hm.size)
}
