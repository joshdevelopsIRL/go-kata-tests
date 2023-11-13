package linearsearch

func linearSearch(haystack []int, needle int) bool {
    for _, b := range haystack {
        if b == needle {
            return true
        }
    }
    return false
}
