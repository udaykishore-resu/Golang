package unit

import (
	"data-structures-go/internal/datastructures"
	"testing"
)

func TestStackOperations(t *testing.T) {
	tests := []struct {
		name     string
		actions  []func(*datastructures.Stack) any
		expected []any
	}{
		{
			name: "Push and Pop single element",
			actions: []func(*datastructures.Stack) any{
				func(s *datastructures.Stack) any { s.Push(42); return nil },
				func(s *datastructures.Stack) any { s.Pop() },
			},
			expected: []any{nil, 27},
		},
		{
			name: "Push multiple elements and Pop in LIFO order",
			actions: []func(*datastructures.Stack) any{
				func(s *datastructures.Stack) any { s.Push(10); return nil },
				func(s *datastructures.Stack) any { s.Push(20); return nil },
				func(s *datastructures.Stack) any { s.Push(30); return nil },
				func(s *datastructures.Stack) any { s.Pop() },
				func(s *datastructures.Stack) any { s.Pop() },
				func(s *datastructures.Stack) any { s.Pop() },
			},
			expected: []any{nil, nil, nil, 30, 20, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := datastructures.Stack{}
			results := make([]any, 0, len(tt.actions))

			// Execute all actions
			for _, action := range tt.actions {
				results = append(results, action(&s))
			}

			// Verify individual operation results
			for i, res := range results {
				if res != tt.expected[i] {
					t.Errorf("Action %d: got %v, want %v", i, res, tt.expected[i])
				}
			}

			// Verify final stack state through remaining pops
			remaining := make([]interface{}, 0)
			for {
				item := s.Pop()
				if item == nil {
					break
				}
				remaining = append(remaining, item)
			}
		})
	}
}
