package server

import (
	"slices"
	"testing"
)

var cases = []Record{
	{
		Value: []byte("hello world!"),
	},
}

var l = NewLog()

func TestAppend(t *testing.T) {

	for i, c := range cases {
		offset, err := l.Append(c)
		if err != nil {
			t.Fatalf("error appending record")
		}
		if uint64(i) != offset {
			t.Fatalf("wrong offset")
		}
	}
}

func TestRead(t *testing.T) {
	for _, c := range cases {
		record, err := l.Read(c.Offset)
		if err != nil {
			t.Fatalf("error reading record")
		}
		if !slices.Equal(c.Value, record.Value) {
			t.Fatalf("wrong value")
		}
	}
}
