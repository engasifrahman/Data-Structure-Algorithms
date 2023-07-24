package main

import "fmt"

// Entry represents a key-value pair in the hash map
type Entry struct {
	key   string
	value int
	next  *Entry
	prev  *Entry
}

// HashMap represents the separate chaining hash map
type HashMap struct {
	buckets []*Entry
	size    int
}

// NewHashMap creates a new instance of the hash map
func NewHashMap() *HashMap {
	return &HashMap{
		buckets: make([]*Entry, 16),
		size:    0,
	}
}

// hashFunction returns the hash value for a given key
func (hm *HashMap) hashFunction(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash<<5 - hash) + int(key[i])
	}
	
	return hash % len(hm.buckets)
}

// Put inserts a key-value pair into the hash map
func (hm *HashMap) Put(key string, value int) {
	hash := hm.hashFunction(key)
	newEntry := &Entry{
		key:   key,
		value: value,
	}

	if hm.buckets[hash] == nil {
		hm.buckets[hash] = newEntry
	} else {
		// Insert at the beginning of the doubly linked list
		newEntry.next = hm.buckets[hash]
		hm.buckets[hash].prev = newEntry
		hm.buckets[hash] = newEntry
	}
	hm.size++
}

// Get retrieves the value associated with a given key from the hash map
func (hm *HashMap) Get(key string) (int, bool) {
	hash := hm.hashFunction(key)
	current := hm.buckets[hash]
	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	return 0, false
}

// Remove removes a key-value pair from the hash map
func (hm *HashMap) Remove(key string) bool {
	hash := hm.hashFunction(key)
	current := hm.buckets[hash]
	for current != nil {
		if current.key == key {
			// Remove the key-value pair
			if current.prev == nil {
				// First node in the list
				hm.buckets[hash] = current.next
			} else {
				current.prev.next = current.next
			}
			if current.next != nil {
				current.next.prev = current.prev
			}
			hm.size--
			return true
		}
		current = current.next
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
