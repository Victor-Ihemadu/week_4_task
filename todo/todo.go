package todo

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type Item struct {
	Position int
	Text     string
	Done     bool
}

func SaveItems(items []Item) error {
	b, err := json.Marshal(items)
	err = ioutil.WriteFile("mytodolist.csv", b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadItems() ([]Item, error) {
	b, err := ioutil.ReadFile("mytodolist.csv")
	if err != nil {
		return []Item{}, err
	}
	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}
	for i, _ := range items {
		items[i].Position = i + 1
	}
	return items, nil

}

func (i *Item) Label() string {
	return strconv.Itoa(i.Position) + "."
}
