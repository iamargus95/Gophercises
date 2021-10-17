package main

import (
	"reflect"
	"testing"

	"github.com/pkg/errors"
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
			err: nil,
		},
		{
			filename: "noexistent.yaml",
			data:     nil,
			err:      errors.New("could not read YAML file"),
			//Fix this. Find a way to test fmt.Errorf()
		},
	}

	for _, v := range testYaml {
		got, _ := yamlReader(v.filename)
		if reflect.DeepEqual(got, v.data) {
			t.Fail()
		}
	}

	for _, v := range testYaml {
		_, err := yamlReader(v.filename)
		if err != v.err {
			t.Fail()
		}
	}
}
