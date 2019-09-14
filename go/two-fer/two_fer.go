// Package twofer contains the logic for two for one exercise
package twofer

import "fmt"

// ShareWith shares all the things with me
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
