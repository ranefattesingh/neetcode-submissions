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

    var res pair
    l, r := 0, len(pairs) - 1
    for l <= r {
        m := l + (r-l)/2

        if pairs[m].timestamp <= timestamp {
            res = pairs[m]
            l = l + 1
        } else {
            r = m - 1
        }
    }


    return res.value
}