package services

import (
	"context"

	"grpc-go-server/pb/paginationpb"
	"grpc-go-server/pb/productpb"
)

type ProductService struct {
	productpb.UnimplementedProductServiceServer
}

func (p ProductService) GetProducts(context.Context, *productpb.Empty) (*productpb.Products, error) {

	pro := &productpb.Products{
		Pagination: &paginationpb.Pagination{Total: 10, PerPage: 4, CurrentPage: 1, LastPage: 2},
		Data:       []*productpb.Product{{Id: 1, Name: "Kurma Ajwa", Price: 2.3, Stock: 300, Category: &productpb.Category{Id: 1, Name: "food"}}},
	}

	return pro, nil
}
