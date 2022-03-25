package limiter

// WhiteList represent white list of keys
type WhiteList struct {
	Keys []string
}

// HasKey implements basic look-up of given key in WhiteList
func (w *WhiteList) HasKey(key string) bool {
	for _, k := range w.Keys {
		if k == key {
			return true
		}
	}
	return false
}
