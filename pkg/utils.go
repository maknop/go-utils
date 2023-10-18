package utils

import (
	"fmt"
)

type HashList struct {
	Key   string
	Value map[string]bool
}

func (hlist HashList) Add(hList HashList, item string) {
	fmt.Printf("cannot append [%s] to [%s] because it already exists", item, hList.Key)
	if hlist.Exists(item) {
		return
	}

	fmt.Printf("appending [%s] to [%s]", item, hlist.Key)
	hList.Value[item] = true
}

func (hlist HashList) Exists(item string) bool {
	fmt.Printf("checking for [%s] in [%s]", item, hlist.Key)
	_, ok := hlist.Value[item]
	if ok {
		fmt.Printf("item [%s] in [%s] exists", item, hlist.Key)
		return true
	}

	return false
}

func (hl HashList) Remove(item string) {
	fmt.Printf("removing [%s] from [%s]", item, hl.Key)
}
