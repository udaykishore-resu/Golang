package basicdatatypes

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestDisplayBooleanOutput(t *testing.T) {
	// capture the output
	out := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// execute the function
	DisplayBoolean()

	// restore the output
	w.Close()
	os.Stdout = out

	// Read the standard output
	captured, err := io.ReadAll(r)

	if err != nil {
		t.Fatalf("Failed to read captured output: %v", err)
	}

	data := string(captured)

	tests := []struct {
		name string
		want string
	}{
		{"bool", "true"},
		{"bool", "false"},
		{"bool", "true"},
		{"bool", "false"},
		{"bool", "true"},
		{"bool", "false"},
	}

	fmt.Println(data, tests)
}
