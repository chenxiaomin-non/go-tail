package mnt

import (
	"fmt"
	"os"
	"strings"

	tail "github.com/chenxiaomin-non/go-tail/tail"
	yaml "gopkg.in/yaml.v2"
)

// parse yaml file to a map of tail.Observer object
// first parse the yaml file to a map[string]interface{}
// then parse the map to a map[string]tail.Observer
func ParseYaml(filePath string) (map[string]*tail.Observer, error) {
	// parse yaml file to a map[string]interface{}
	m, err := parseYamlToMap(filePath)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func FilterEnum(content []byte) []byte {
	enum := map[string]int{
		"ReadFromHead": 1,
		"ReadFromTail": 2,
	}

	str := string(content)
	for k, v := range enum {
		str = strings.ReplaceAll(str, k, fmt.Sprintf("%d", v))
	}

	return []byte(str)
}

// parse yaml file to a map[string]interface{}
func parseYamlToMap(filePath string) (map[string]*tail.Observer, error) {
	// read yaml file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	filtered_content := FilterEnum(content)

	// parse yaml file
	m := make(map[string]*tail.Observer)
	err = yaml.Unmarshal(filtered_content, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
