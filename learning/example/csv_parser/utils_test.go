package main

import "testing"

func TestGetFilesFromFolder(t *testing.T) {
	lst, err := getFilesFromFolder(dataDir)
	t.Log(err)
	t.Log(lst)
}

func TestIsIPv4(t *testing.T) {
	inps := map[string]bool{
		"":                false,
		"a":               false,
		"aa":              false,
		" a":              false,
		" ":               false,
		"a.b.c.d":         false,
		"0.0.0.0":         true,
		"1.0.0.0":         true,
		"192.168.1.1":     true,
		"192.168.1.":      false,
		"192.168":         false,
		"abc_192.168.1.1": false,
	}

	for k, v := range inps {
		rs := isIPv4(k)
		if rs != v {
			t.Fail()
		}
	}
}
