package product

import (
	"github.com/ckalagara/go-grpc-server/internal/inventory"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type ServiceImpl struct {
}

func (s ServiceImpl) Create(ctx context.Context, productDto *Product) (*Product, error) {
	if productDto == nil {

		log.WithFields(log.Fields{
			"code":    "PRODUCT_CREATION_ERR",
			"message": "error while creating productDto record",
		}).Error("nil productDto")

		return productDto, errors.New("productDto cannot be nil")
	}
	log.Debugf("processing productDto creation request with %v", productDto)
	productItem, err := inventory.NewProductRepo().Create(getProduct(productDto))
	if err != nil {
		log.WithFields(log.Fields{
			"code":    "PRODUCT_CREATION_ERR",
			"message": "error while creating productDto record",
			"product": *productDto,
		}).Error(err.Error())
		return productDto, errors.WithStack(err)
	}
	return getProductDto(productItem), nil
}

func (s ServiceImpl) LookupByName(ctx context.Context, name *ProductName) (*Product, error) {
	product, err := inventory.NewProductRepo().LookupByName(name.Name)
	if err != nil {
		log.WithFields(log.Fields{
			"code":         "PRODUCT_LOOKUP_ERR",
			"message":      "error while searching product record",
			"product_name": name.Name,
		}).Error(err.Error())
		return nil, errors.WithStack(err)
	}
	return getProductDto(product), nil
}

func (s ServiceImpl) LookupByID(ctx context.Context, id *ProductID) (*Product, error) {
	product, err := inventory.NewProductRepo().LookupByID(id.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"code":       "PRODUCT_LOOKUP_ERR",
			"message":    "error while searching product record",
			"product_id": id.Id,
		}).Error(err.Error())
		return nil, errors.WithStack(err)
	}
	return getProductDto(product), nil
}

func getProductDto(p *inventory.Product) *Product {
	if p == nil {
		return nil
	}
	return &Product{
		Id:           p.ID,
		Name:         p.Title,
		Category:     p.Category,
		Type:         p.Kind,
		Manufactured: p.Vendor,
		Manufacturer: p.Mfd,
	}
}

func getProduct(p *Product) *inventory.Product {
	if p == nil {
		return nil
	}

	return &inventory.Product{
		ID:       p.Id,
		Title:    p.Name,
		Category: p.Category,
		Kind:     p.Type,
		Vendor:   p.Manufacturer,
		Mfd:      p.Manufactured,
	}

}
