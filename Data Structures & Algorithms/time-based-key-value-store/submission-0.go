type TimeMap struct {
    m map[string][]pair
}

type pair struct {
    value string
    timestamp int
}

func Constructor() TimeMap {
    return TimeMap{
        m: make(map[string][]pair),
    }
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
    this.m[key] = append(this.m[key], pair{ timestamp: timestamp, value: value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
    pairs, ok := this.m[key]
    if !ok {
        return ""
    }

    index := sort.Search(len(pairs), func(i int) bool {
       return pairs[i].timestamp > timestamp
   }) 

    if index == 0 {
        return ""
    }

    return pairs[index-1].value
}