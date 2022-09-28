package main

import "testing"

var testItems = []Item{
	{
		Id:            1,
		Name:          "panosit",
		ReservationId: 387,
	},
	{
		Id:            1,
		Name:          "tisonap",
		ReservationId: 783,
	},
}

type repoFake struct {
}

func (r *repoFake) ListItems() ([]Item, error) {
	return testItems, nil
}

func (r *repoFake) CreateItem(Item) error {
	return nil
}

func TestListItems(t *testing.T) {
	// given
	repo := new(repoFake)
	service = NewService(repo)

	// when
	service.repo.ListItems()

	// then

}
