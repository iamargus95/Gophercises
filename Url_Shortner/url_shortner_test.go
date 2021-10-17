package main

import (
	"reflect"
	"testing"
)

func TestYamlReader(t *testing.T) {

	var testYaml = []struct {
		filename string
		data     []byte
		err      error
	}{
		{
			filename: "default.yaml",
			data: []byte(`- path: "/github"
			url: "https://github.com/iamargus95"
			- path: "/personal"
			url: "https://iamargus95.github.io/"
			- path: "/resume"
			url: "https://iamargus95.github.io/assets/documents/Suraj_Kamath.pdf"`),
		},
		{
			filename: "emptyfile.yaml",
			data:     nil, // []byte(``) does not work
		},
	}

	for _, v := range testYaml {
		got, _ := yamlReader(v.filename)
		if reflect.DeepEqual(got, v.data) {
			t.Fail()
		}
	}
}
