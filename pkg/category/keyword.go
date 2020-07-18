package category

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/doug-martin/goqu/v9"
)

// KeywordCategories is a mapping of description keywords to a category
type KeywordCategories map[string]string

// LoadKeywordCategories loads a KeywordCatMap from a file
func LoadKeywordCategories(path string) (c KeywordCategories, err error) {
	c = KeywordCategories{}
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
func (c KeywordCategories) GetFromRecord(rec goqu.Record, matchField string) (category string, err error) {
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

// Contains evaluates whether KeywordCategories contains a specified key
func (c KeywordCategories) Contains(key string) bool {
	_, found := c[key]
	return found
}

// Add adds a key/value pair to KeywordCategories (overwriting an existing key/value pair)
func (c KeywordCategories) Add(key string, val string) {
	c[key] = val
}

// Delete deletes a key from KeywordCategories and is a no-op if the key does not exist
func (c KeywordCategories) Delete(key string) {
	delete(c, key)
}

// Write writes the string representation of KeywordCategories to a file
func (c KeywordCategories) Write(path string) (err error) {
	bytes := []byte(c.String())
	return ioutil.WriteFile(path, bytes, 0644)
}

func (c KeywordCategories) String() (s string) {
	bytes, _ := json.MarshalIndent(c, "", "\t") // safe to ignore err since struct came from json
	return string(bytes)
}
