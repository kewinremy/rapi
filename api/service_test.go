package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type repoFake struct {
	items []Item
	err   error
}

func (r repoFake) ListItems() ([]Item, error) {
	return r.items, nil
}

func (r repoFake) CreateItem(Item) error {
	return r.err
}

func TestListItems(t *testing.T) {
	tests := []struct {
		name string
		want []Item
	}{
		{
			name: "List with items",
			want: []Item{
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
			},
		},
		{
			name: "Empty List",
			want: []Item{},
		},
	}

	// given
	w := httptest.NewRecorder()
	gc, _ := gin.CreateTestContext(w)
	repoFake := new(repoFake)
	service = new(Service)
	service.repo = repoFake
	var got []Item

	// when
	for _, tt := range tests {
		repoFake.items = tt.want

		service.GetItems(gc)

		// then
		if w.Code != http.StatusOK {
			log.Printf("GetItems failed, expected %v got %v\n", http.StatusOK, w.Code)
			t.Error()
			return
		}

		body, _ := ioutil.ReadAll(w.Body)
		err := json.Unmarshal(body, &got)
		if err != nil {
			log.Printf("Test failed due bad decoding\n")
			t.Error()
			return
		}

		if len(got) != len(tt.want) {
			log.Printf("Test failed due bad length, expected %v got %v\n", len(tt.want), len(got))
			t.Error()
			return
		}
	}
}
