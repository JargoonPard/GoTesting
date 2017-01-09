package mystuff

// Foo is a test structure
type Foo struct {
	count int
	names []string
}

// AddName is a simple method to fake out
func AddName(name string) (int, []string) {

	names := []string{}

	names = append(names, name)

	return len(names), names
}
