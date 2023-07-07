package enmu

type MenuCache int8

const (
	Menu_Cache     = 1
	Menu_Not_Cache = 2
)

var menuCacheMap map[int]string

func init() {
	menuCacheMap = make(map[int]string)
	menuCacheMap[Menu_Cache] = "缓存"
	menuCacheMap[Menu_Not_Cache] = "不缓存"
}

func (m MenuCache) Desc() string {
	return menuCacheMap[int(m)]
}

func (m MenuCache) Value() int8 {
	return int8(m)
}

func (m MenuCache) Size() int {
	return len(menuCacheMap)
}

func (m MenuCache) Equals(value int8) bool {
	return int8(m) == value
}
