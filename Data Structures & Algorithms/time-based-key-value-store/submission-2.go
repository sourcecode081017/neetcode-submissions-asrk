type Pair struct {
	timestamp int
	value string
}

type TimeMap struct {
	store map[string][]Pair
}

func Constructor() TimeMap {
	return TimeMap {
		store: make(map[string][]Pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key], Pair{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, exists := this.store[key]; !exists {
		return ""
	}
	pairs := this.store[key]
	idx := sort.Search(len(pairs), func(i int) bool {
		return pairs[i].timestamp > timestamp
	})
	if idx == 0 {
		return ""
	}
	return pairs[idx - 1].value
}
