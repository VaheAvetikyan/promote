package repository

import (
	"example/promote/model"
	"testing"
)

func TestAdd(t *testing.T) {
	repo := New()
	repo.Add(model.Record{})
	if len(repo.Records) == 0 {
		t.Error("Not Added")
	}
}

func TestGetAll(t *testing.T) {
	repo := New()
	repo.Add(model.Record{})
	results := repo.GetAll()
	if len(results) != 1 {
		t.Error("Something wrong")
	}
}

func TestGetById(t *testing.T) {
	const PRICE = 4.5
	repo := New()
	repo.Add(model.Record{Id: "1", Price: PRICE, ExpirationDate: "12/12/2022"})
	var result = repo.GetById(1)
	if result.Price != PRICE {
		t.Error("Something wrong")
	}
}
