// Twofer package is used to get the desired string
package twofer

// Return string in the 'One for X, one for me.' format.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
