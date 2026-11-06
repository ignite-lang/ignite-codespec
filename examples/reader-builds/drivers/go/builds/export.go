package builds

import "fmt"

// Lines calls each case of the written package, whose functions keep Ignite's lowercase names.
func Lines() []string {
	xs := []int64{1, 2}
	return []string{
		fmt.Sprintf("size %d", size(xs)),
		fmt.Sprintf("name %s", name(Green)),
		fmt.Sprintf("bump %d", bump(1)),
		fmt.Sprintf("same %t", same(Pair{xs: xs}, Pair{xs: xs})),
	}
}
