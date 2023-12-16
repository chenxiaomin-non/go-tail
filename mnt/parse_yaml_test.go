package mnt_test

import (
	"testing"

	"github.com/chenxiaomin-non/go-tail/mnt"
)

func TestParseYaml(t *testing.T) {
	filepath := "test.yaml"
	obsMap, err := mnt.ParseYaml(filepath)

	if err != nil {
		t.Fatal(err)
	}

	for k, v := range obsMap {
		t.Log(k, v)
	}
}
