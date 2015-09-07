package static

import (
	"reflect"
	"testing"
)

var isvTests = []struct {
	status int
	valid  bool
}{
	{0, false},
	{99, false},
	{100, true},
	{599, true},
	{600, false},
}

func TestIsStatusValid(t *testing.T) {
	t.Logf("I can't believe I actually wrote tests for this.")
	for _, st := range isvTests {
		valid := isStatusValid(st.status)
		if st.valid != valid {
			t.Errorf("IsStatusValid failed. For %d, expected %t, got %t.", st.status, st.valid, valid)
		}
	}
}

var bwhTests = []struct {
	fullBody string
	headers  map[string]string
	body     string
	err      bool
	pass     bool
}{
	{"header: value\n\nbody", map[string]string{"header": "value"}, "body", false, true},
	{"header: value\n\nbody\nbody2", map[string]string{"header": "value"}, "body\nbody2", false, true},
	{"header: other\n\nbody", map[string]string{"header": "value"}, "body", false, false},
	{"header: value\nsecond: another\n\nbody", map[string]string{"header": "value", "second": "another"}, "body", false, true},

	// Error cases
	{"\n\nbody", map[string]string{}, "", true, false},
	{"header:tooclose\n\nbody", map[string]string{"header": "value"}, "", true, false},
	{"brokenheader\n\nbody", map[string]string{"header": "value"}, "", true, false},
}

func TestBodyWithHeaders(t *testing.T) {
	t.Logf("Running %d bodyWithHeader tests\n", len(bwhTests))
	for i, bh := range bwhTests {
		t.Logf("Running bodyWithHeader test %d\n", i+1)
		headers, body, err := parseBodyWithHeaders(bh.fullBody)
		if err != nil {
			if !bh.err {
				t.Logf("For input: %q\n", bh.fullBody)
				t.Errorf("Expected err %t, got %t", bh.err, (err != nil))
			}
			continue
		}
		if (reflect.DeepEqual(bh.headers, headers) && body == bh.body) != bh.pass {
			t.Logf("For input: %q\n", bh.fullBody)
			t.Logf("Expecting to pass: %t\n", bh.pass)
			t.Errorf("Expected headers %v, got %v\nExpected body %q, got %q", bh.headers, headers, bh.body, body)
		}
	}
}
