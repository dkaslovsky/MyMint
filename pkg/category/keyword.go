package category

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/doug-martin/goqu/v9"
)

// KeywordCatMap is a mapping of description keywords to a category
type KeywordCatMap map[string]string

// LoadKeywordCatMap loads a KeywordCatMap from a file
func LoadKeywordCatMap(path string) (c KeywordCatMap, err error) {
	c = KeywordCatMap{}
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

// GetFromRecord returns a category according to a matched keyword
func (c KeywordCatMap) GetFromRecord(rec goqu.Record, matchField string) (category string, err error) {
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
