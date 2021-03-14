// Code is generated. DO NOT EDIT.

package terraform

// IsType returns true if the given string is a Terraform AWS resource type.
func IsType(s string) bool {
	for _, t := range Types {
		if t == s {
			return true
		}
	}

	return false
}
