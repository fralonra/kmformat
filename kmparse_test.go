package kmparse

import (
	"testing"
)

type testList []struct {
	name, got, exp string
}

func (tl testList) validate(t *testing.T) {
	for _, test := range tl {
		if test.got != test.exp {
			t.Errorf("On %v, expected '%v', but got '%v'",
				test.name, test.exp, test.got)
		}
	}
}

func TestKFormat(t *testing.T)  {
  testList{
		{"100", KFormat(100), "100"},
		{"1010", KFormat(1010), "1.01k"},
		{"2566", KFormat(2566), "2.57k"},
  }.validate(t)
}

func TestMFormat(t *testing.T)  {
  testList{
		{"1000", MFormat(1000), "1000"},
		{"1010000", MFormat(1010000), "1.01m"},
		{"2566000", MFormat(2566000), "2.57m"},
  }.validate(t)
}

func TestFormat(t *testing.T)  {
  testList{
		{"100", Format(100), "100"},
		{"1010", Format(1010), "1.01k"},
		{"2566000", Format(2566000), "2.57m"},
  }.validate(t)
}
