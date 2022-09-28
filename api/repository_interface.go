package main

type ItemRepository interface {
	ListItems() ([]Item, error)
	CreateItem(Item) error
}
