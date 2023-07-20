package utils

import "encoding/json"

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

func ObjsToMapList(s any) ([]map[string]any, error) {
	var newMap []map[string]any
	data, _ := json.Marshal(s)
	err := json.Unmarshal(data, &newMap)

	return newMap, err
}
