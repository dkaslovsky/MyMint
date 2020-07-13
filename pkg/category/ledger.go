package category

import (
	"io/ioutil"
	"sort"
	"strings"
)

// Categories contains string categories
type Categories map[string]struct{}

// LoadCategories reads strings from a newline delimited file
func LoadCategories(path string) (c Categories, err error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return c, err
	}
	categories := strings.Split(string(bytes), "\n")

	c = Categories{}
	for _, category := range categories {
		if category == "" {
			continue
		}
		c[category] = struct{}{}
	}
	return c, err
}

// Contains evaluates whether Categories contains a specified string
func (c Categories) Contains(category string) bool {
	_, found := c[category]
	return found
}

// Add adds a string to Categories and functions as a no-op if the string already exists
func (c Categories) Add(category string) {
	c[category] = struct{}{}
}

// Delete deletes a string from Categories and is a no-op if the string does not exist
func (c Categories) Delete(category string) {
	delete(c, category)
}

// Write writes the string representation of Categories to a file
func (c Categories) Write(path string) (err error) {
	bytes := []byte(c.String())
	return ioutil.WriteFile(path, bytes, 0644)
}

func (c Categories) String() (s string) {
	categories := []string{}
	for category := range c {
		categories = append(categories, category)
	}
	sort.Strings(categories)
	return strings.Join(categories, "\n")
}
