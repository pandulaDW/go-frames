package helpers

func get64(x interface{}) interface{} {
	switch x := x.(type) {
	case uint8:
		return int64(x)
	case int8:
		return int64(x)
	case uint16:
		return int64(x)
	case int16:
		return int64(x)
	case uint32:
		return int64(x)
	case int32:
		return int64(x)
	case uint64:
		return int64(x)
	case int64:
		return int64(x)
	case int:
		return int64(x)
	case float32:
		return float64(x)
	case float64:
		return float64(x)
	}
	panic("invalid input")
}
