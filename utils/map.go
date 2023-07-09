package utils

func DeleteKv(m map[string]any, keys ...string) {
	if m == nil || len(keys) == 0 {
		return
	}

	for _, k := range keys {
		delete(m, k)
	}
}

func DeleteKvWhenUpdate(m map[string]any) {
	DeleteKv(m, "create_by", "create_time", "update_time")
}
