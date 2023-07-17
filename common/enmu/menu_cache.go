package enmu

type MenuCache string

const (
	Menu_Cache     = "1"
	Menu_Not_Cache = "2"
)

var menuCacheMap map[string]string

func init() {
	menuCacheMap = make(map[string]string)
	menuCacheMap[Menu_Cache] = "缓存"
	menuCacheMap[Menu_Not_Cache] = "不缓存"
}

func (m MenuCache) Desc() string {
	return menuCacheMap[m.Value()]
}

func (m MenuCache) Value() string {
	return string(m)
}

func (m MenuCache) Size() int {
	return len(menuCacheMap)
}

func (m MenuCache) Equals(value string) bool {
	return string(m) == value
}
