package box

import "testing"

func TestBoxGet(t *testing.T) {
	testCases := []struct {
		name         string
		initialState []Item
		key          string

		expectedErrorType ErrorType
		expectedItem      Item
	}{
		{
			name:              "key does not exist",
			initialState:      []Item{},
			expectedErrorType: KeyNotFound,
			expectedItem:      Item{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			box := New(WithInitialState(tc.initialState))
			_, err := box.Get(tc.key)
			if err != nil && err.(OperationError).errType != tc.expectedErrorType {
				t.Errorf("case %q: unexpected error type: %v, expected: %v", tc.name, err.(OperationError).errType, tc.expectedErrorType)
			}
		})
	}

}
