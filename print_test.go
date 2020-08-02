package vik

import (
	"testing"
)

type cat struct {
	s int
	i string
	u []byte
}

func TestPrint(t *testing.T) {
	var a uint = 1
	b := []byte{'a', 'b'}
	c := "dfadsfasd"
	d := []string{"dfdf", "dfsafdsf"}
	var ca = &cat{s: 1, i: "string", u: []byte{'1', 'a'}}
	e := map[byte]byte{'a': 'b'}
	f := true
	Println(a, b, c, d, ca, e, f)
}

func TestLog(t *testing.T) {
	var a uint = 1
	b := []byte{'1', '2'}
	c := "dfadsfasd"
	d := []string{"dfdf", "dfsafdsf"}
	var ca = cat{s: 1, i: "string", u: []byte{'1', 'a'}}
	e := map[byte]byte{'a': 'b'}
	Logln(a, b, c, d, ca, e)
}
