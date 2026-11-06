package builds

import "fmt"

// Lines calls each case of the written package, whose functions keep Ignite's lowercase names.
func Lines() []string {
	xs := []int{1, 2}
	return []string{
		fmt.Sprintf("size %d", size(xs)),
		fmt.Sprintf("name %s", name(Green)),
	}
}
