package main

import "fmt"

// SemVer demonstrates semantic versioning concepts
type Version struct {
	Major, Mainor, Patch int
}

func (v Version) String() string {
	return fmt.Sprintf("v%d.%d.%d", v.Major, v.Mainor, v.Patch)
}

// IsCompatible checks if other is backward compatible with v
func (v Version) IsCompatible(other Version) bool {
	return v.Major == other.Major
}

func main() {
	current := Version{1, 3, 0}
	previous := Version{1, 2, 5}
	breaking := Version{2, 0, 0}

	fmt.Println("current:", current)
	fmt.Println("compatible with previous:", current.IsCompatible(previous))
	fmt.Println("compatible with v2:", current.IsCompatible(breaking))
}
