package category

import (
	"io/ioutil"
	"sort"
	"strings"
)

// LedgerCategories contains string categories
type LedgerCategories map[string]struct{}

// LoadLedgerCategories reads strings from a newline delimited file
func LoadLedgerCategories(path string) (c LedgerCategories, err error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return c, err
	}
	categories := strings.Split(string(bytes), "\n")

	c = LedgerCategories{}
	for _, category := range categories {
		if category == "" {
			continue
		}
		c[category] = struct{}{}
	}
	return c, err
}

// Contains evaluates whether LedgerCategories contains a specified string
func (c LedgerCategories) Contains(category string) bool {
	_, found := c[category]
	return found
}

// Add adds a string to LedgerCategories and functions as a no-op if the string already exists
func (c LedgerCategories) Add(category string) {
	c[category] = struct{}{}
}

// Delete deletes a string from LedgerCategories and is a no-op if the string does not exist
func (c LedgerCategories) Delete(category string) {
	delete(c, category)
}

// Write writes the string representation of LedgerCategories to a file
func (c LedgerCategories) Write(path string) (err error) {
	bytes := []byte(c.String())
	return ioutil.WriteFile(path, bytes, 0644)
}

func (c LedgerCategories) String() (s string) {
	categories := []string{}
	for category := range c {
		categories = append(categories, category)
	}
	sort.Strings(categories)
	return strings.Join(categories, "\n")
}
