package data_structures

import "fmt"

// Map implementation: https://yourbasic.org/golang/maps-explained/
func mapExample() {
	set := make(map[string]bool) // New empty set
	set["Foo"] = true            // Add
	for k := range set {         // Loop
		fmt.Println(k)
	}
	delete(set, "Foo")    	// Delete
	size := len(set)      		// Size
	exists := set["Foo"] 	 	// Membership

	fmt.Println(size)
	fmt.Println(exists)
}

type void struct{}
func mapExampleAlternative() {
	var member void

	set := make(map[string]void) // New empty set
	set["Foo"] = member          // Add
	for k := range set {         // Loop
		fmt.Println(k)
	}
	delete(set, "Foo")      // Delete
	size := len(set)        // Size
	_, exists := set["Foo"] // Membership
	fmt.Println(size)
	fmt.Println(exists)
}