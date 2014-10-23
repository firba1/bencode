package bencode

import (
	"bytes"
	"reflect"
	"testing"
)

type valueTestcase struct {
	in  string
	out interface{}
}

func TestConsumeValue(t *testing.T) {
	tests := []valueTestcase{
		valueTestcase{"i2e", 2},
		valueTestcase{"1:a", "a"},
		//valueTestcase{"li2ei42ei666ee", []int{2, 42, 666}},
	}
	for _, test := range tests {
		out := reflect.New(reflect.TypeOf(test.out)).Elem()
		consumeValue(out, bytes.NewBuffer([]byte(test.in)))
		if i := out.Interface(); !reflect.DeepEqual(i, test.out) {
			t.Error("Expecting", test.out, "got", i)
		}
	}
}

type intTestcase struct {
	in  string
	out int
}

func TestConsumeInt(t *testing.T) {
	tests := []intTestcase{
		intTestcase{"i2e", 2},
		intTestcase{"i42e", 42},
		intTestcase{"i666e", 666},
	}
	for _, test := range tests {
		var o int
		out := reflect.ValueOf(&o).Elem()
		consumeInt(out, bytes.NewBuffer([]byte(test.in)))
		if i := out.Interface(); !reflect.DeepEqual(i, test.out) {
			t.Error("Expecting", test.out, "got", i)
		}
	}
}

type stringTestcase struct {
	in  string
	out string
}

func TestConsumeString(t *testing.T) {
	tests := []stringTestcase{
		stringTestcase{"1:s", "s"},
		stringTestcase{"4:butt", "butt"},
		stringTestcase{"11:buttfartass", "buttfartass"},
	}
	for _, test := range tests {
		var o string
		out := reflect.ValueOf(&o).Elem()
		consumeString(out, bytes.NewBuffer([]byte(test.in)))
		if i := out.Interface(); !reflect.DeepEqual(i, test.out) {
			t.Error("Expecting", test.out, "got", i)
		}
	}
}

/*
type listTestcase struct {
	in  string
	out interface{}
}

func TestConsumeList(t *testing.T) {
	tests := []listTestcase{
		listTestcase{"li2ei42ei666ee", []int{2, 42, 666}},
		listTestcase{"l1:a1:b1:ce", []string{"a", "b", "c"}},
		listTestcase{"lli2eeli42eee", [][]int{[]int{2}, []int{42}}},
	}

	for _, test := range tests {
		out := consumeList(bytes.NewBuffer([]byte(test.in)))
		if i := out.Interface(); !reflect.DeepEqual(i, test.out) {
			t.Error("Expecting", test.out, "got", i)
		}
	}
}

type dictTestcase struct {
	in  string
	out interface{}
}

type testStruct struct {
	a int
	b string
}

func TestConsumeDict(t *testing.T) {
	tests := []dictTestcase{
		dictTestcase{"d1:Ai42e1:B3:xyze", struct {
			A int
			B string
		}{A: 42, B: "xyz"}},
	}

	for _, test := range tests {
		out := consumeDict(bytes.NewBuffer([]byte(test.in)), reflect.TypeOf(test.out))
		if i := out.Interface(); !reflect.DeepEqual(i, test.out) {
			t.Error("Expecting", test.out, "got", i)
		}
	}
}
*/
