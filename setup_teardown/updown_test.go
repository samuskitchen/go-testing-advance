package setup_teardown

import (
	"os"
	"testing"
)

// Standard numbers map
var numbers map[string]int = map[string]int{"zero": 0, "three": 3}

// TestMain will exec each test, one by one
func TestMain(m *testing.M) {
	// exec setUp function
	setUp("one", 1)
	// exec test and this returns an exit code to pass to os
	retCode := m.Run()
	// exec tearDown function
	tearDown("one")
	// If exit code is distinct of zero,
	// the test will be failed (red)
	os.Exit(retCode)
}

// setUp function, add a number to numbers slice
func setUp(key string, value int) {
	numbers[key] = value
}

// tearDown function, delete a number to numbers slice
func tearDown(key string) {
	delete(numbers, key)
}

// First test
func TestOnePlusOne(t *testing.T) {
	numbers["one"] = numbers["one"] + 1

	if numbers["one"] != 2 {
		t.Errorf("1 plus 1 = 2, not %v", numbers["one"])
	}
}

// Second test
func TestOnePlusTwo(t *testing.T) {
	numbers["one"] = numbers["one"] + 2

	if numbers["one"] != 3 {
		t.Errorf("1 plus 2 = 3, not %v", numbers["one"])
	}
}
