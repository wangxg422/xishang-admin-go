package utils

import (
	"backend/model/vo/common"
)

// ListToSelectTree
/**
结构体列表转换成树形前端树状选择器数据结构
*/
func ListToSelectTree(in any, parentId, childName, itemId, itemName string) (common.TreeSelectVO, error) {
	if in == nil {
		return common.TreeSelectVO{}, nil
	}

	list, err := ObjsToMapList(in)
	if err != nil {
		return common.TreeSelectVO{}, err
	}

	item := make(map[string]any)
	item[itemId] = "0"
	item[itemName] = "root"
	vo := common.TreeSelectVO{Id: "0"}
	buildSelectTree(list, &item, &vo, parentId, childName, itemId, itemName)

	return vo, nil
}

func buildSelectTree(list []map[string]any, item *map[string]any, tree *common.TreeSelectVO, parentId, childName, itemId, itemName string) {
	children, voChildren := getSelectChildren(list, item, parentId, itemId, itemName)

	tree.Id = (*item)[itemId].(string)
	tree.Label = (*item)[itemName].(string)
	tree.Children = voChildren

	length := len(children)
	if length > 0 {
		for i := 0; i < length; i++ {
			buildSelectTree(list, &children[i], &voChildren[i], parentId, childName, itemId, itemName)
		}
	}
}

func getSelectChildren(list []map[string]any, item *map[string]any, parentId, itemId, itemName string) ([]map[string]any, []common.TreeSelectVO) {
	var cList []map[string]any
	var voList []common.TreeSelectVO

	for _, d := range list {
		if d[parentId] == (*item)[itemId] {
			cList = append(cList, d)
			voList = append(voList, common.TreeSelectVO{
				Id:    d[itemId].(string),
				Label: d[itemName].(string),
			})
		}
	}

	return cList, voList
}
