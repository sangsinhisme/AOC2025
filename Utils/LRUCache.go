package Utils

type LRUCache struct {
	cache    map[string]int
	capacity int
	order    []string // Tracks the order of access
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		cache:    make(map[string]int),
		capacity: capacity,
		order:    []string{},
	}
}

func (c *LRUCache) Get(key string) (int, bool) {
	if val, found := c.cache[key]; found {
		// Move the key to the end of the order to mark it as recently used
		c.markAsUsed(key)
		return val, true
	}
	return 0, false
}

func (c *LRUCache) Put(key string, value int) {
	if _, found := c.cache[key]; found {
		// Update value and mark as recently used
		c.cache[key] = value
		c.markAsUsed(key)
		return
	}

	// Add new key-value pair
	if len(c.cache) >= c.capacity {
		// Evict the least recently used item
		lru := c.order[0]
		c.order = c.order[1:]
		delete(c.cache, lru)
	}

	c.cache[key] = value
	c.order = append(c.order, key)
}

func (c *LRUCache) markAsUsed(key string) {
	// Remove the key from its current position in the order slice
	for i, k := range c.order {
		if k == key {
			c.order = append(c.order[:i], c.order[i+1:]...)
			break
		}
	}
	// Append the key to the end of the order slice
	c.order = append(c.order, key)
}
