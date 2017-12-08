package registers

import (
	"strings"
	"testing"
)

var testData = `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`

func TestTestData(t *testing.T) {
	s := NewScanner(strings.NewReader(testData))
	s.Scan()
	s.Run()

	actual := s.FindLargestValue()
	expected := 1

	if actual != expected {
		t.Errorf("Expected largest value of %d, got %d", expected, actual)
	}
}
