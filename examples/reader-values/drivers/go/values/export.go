package values

import "fmt"

// Lines calls each case of the written package, whose functions keep Ignite's lowercase names.
func Lines() []string {
	return []string{
		fmt.Sprintf("alias_list %d", alias_list()),
		fmt.Sprintf("alias_map %d", alias_map()),
		fmt.Sprintf("param_map %d", param_map()),
		fmt.Sprintf("record_map %d", record_map()),
		fmt.Sprintf("arm_wild %s", arm_wild(1)),
		fmt.Sprintf("arm_bind %s", arm_bind(1)),
	}
}
