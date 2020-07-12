package category

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/doug-martin/goqu/v9"
)

type AutoCategories map[string]string

func LoadAutoCategories(path string) (c AutoCategories, err error) {
	c = AutoCategories{}
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (c AutoCategories) Apply(rec goqu.Record, matchField string) (category string, err error) {
	toMatch, ok := rec[matchField].(string)
	if !ok {
		return category, fmt.Errorf("Record does not contain field [%s]", matchField)
	}
	toMatch = strings.ToLower(toMatch)
	for key := range c {
		if strings.Contains(toMatch, key) {
			return c[key], nil
		}
	}
	return "", nil
}
