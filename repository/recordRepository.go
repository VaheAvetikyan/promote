package repository

import "example/promote/model"

type Repo struct {
	Records []model.Record
}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) Add(record model.Record) {
	r.Records = append(r.Records, record)
}

func (r *Repo) GetAll() []model.Record {
	return r.Records
}

func (r *Repo) GetById(id int) model.Record {
	return r.Records[id-1]
}

func init() {
	records := New().Records
	records = append(records, model.Record{Id: "1", Price: 4.5, ExpirationDate: "12/12/2022"})
	records = append(records, model.Record{Id: "2", Price: 5.1, ExpirationDate: "15/12/2022"})
	records = append(records, model.Record{Id: "3", Price: 11, ExpirationDate: "21/12/2022"})
	records = append(records, model.Record{Id: "4", Price: 7.5, ExpirationDate: "02/12/2022"})
	records = append(records, model.Record{Id: "5", Price: 8, ExpirationDate: "09/12/2022"})
}
