package medium

type TimeMap struct {
	store map[string]MapEntry
}

type MapEntry struct {
	timestamp []int
	value     []string
}

func ConstructorTimeMap() TimeMap {
	return TimeMap{
		store: map[string]MapEntry{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {

	mapEntry, ok := this.store[key]

	if !ok {
		mapEntry = MapEntry{
			timestamp: []int{},
			value:     []string{},
		}
	}
	values := mapEntry.value
	values = append(values, value)
	mapEntry.value = values
	this.store[key] = mapEntry
}

func (this *TimeMap) Get(key string, timestamp int) string {

	mapEntry, ok := this.store[key]

	if !ok {
		return ""
	}

	timestamps := mapEntry.timestamp
	l := 0
	r := len(timestamps)

	for l < r {
		mid := (l + r) / 2
		if timestamps[mid] == timestamp {
			l = mid
			break
		}
		if timestamps[mid] < timestamp {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if timestamps[l] == timestamp {
		return mapEntry.value[l]
	}

	if timestamps[l] > timestamp {
		l--
	}

	if l >= 0 {
		return mapEntry.value[l]
	} else {
		return ""
	}

}
