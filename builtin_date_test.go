package goja

import (
	"testing"
)

func TestIntArg(t *testing.T) {
	r := New()
	for _, tc := range []struct {
		description   string
		argNum        int
		funcCall      FunctionCall
		expectedValue int64
		expectedOk    bool
	}{
		{
			description:   "should return time value properly",
			argNum:        0,
			funcCall:      FunctionCall{ctx: r.vm.ctx, Arguments: []Value{valueInt(2024)}},
			expectedValue: 2024,
			expectedOk:    true,
		},
		{
			description:   "should return 0 when one argument is nil",
			argNum:        0,
			funcCall:      FunctionCall{ctx: r.vm.ctx, Arguments: []Value{nil}},
			expectedValue: 0,
			expectedOk:    false,
		},
		{
			description:   "should return 0 when argument is empty",
			argNum:        0,
			funcCall:      FunctionCall{ctx: r.vm.ctx, Arguments: []Value{}},
			expectedValue: 0,
			expectedOk:    false,
		},
		{
			description:   "should return 0 when the args field is nil",
			argNum:        0,
			funcCall:      FunctionCall{ctx: r.vm.ctx},
			expectedValue: 0,
			expectedOk:    false,
		},
	} {
		t.Run(tc.description, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("should not have panic: %v", r)
				}
			}()

			res, ok := _intArg(tc.funcCall, tc.argNum)
			if ok != tc.expectedOk {
				t.Fatalf("Expected ok value: %t, actual ok value: %t", tc.expectedOk, ok)
			}
			if res != tc.expectedValue {
				t.Fatalf("Expected time value: %d, actual time value: %d", tc.expectedValue, res)
			}
		})
	}

}
