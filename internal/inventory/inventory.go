package inventory

import (
	"sync"
)

type Product struct {
	ID       int64
	Title    string
	Category string
	Kind     string
	Vendor   string
	Mfd      string
}

var productMap sync.Map

type ProductRepo interface {
	Create(p *Product) (*Product, error)
	LookupByName(name string) (*Product, error)
	LookupByID(id int64) (*Product, error)
}

func NewProductRepo() ProductRepo {
	return new(productRepoImpl)
}

type productRepoImpl struct {
}

func (p2 productRepoImpl) Create(p *Product) (*Product, error) {

	pi, _ := productMap.LoadOrStore(p.ID, *p)
	product := pi.(Product)
	return &product, nil
}

func (p2 productRepoImpl) LookupByName(name string) (*Product, error) {
	var p Product
	productMap.Range(func(key, value interface{}) bool {
		p = value.(Product)
		if p.Title == name {
			return false
		}
		return true
	})

	return &p, nil
}

func (p2 productRepoImpl) LookupByID(id int64) (*Product, error) {
	pi, ok := productMap.Load(id)
	if !ok {
		return nil, nil
	}

	product := pi.(Product)
	return &product, nil
}
