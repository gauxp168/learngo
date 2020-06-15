package services

import (
	"context"
)

type ProdService struct {

}

func (this *ProdService) GetProdStock(ctx context.Context, in *ProdRequset) (*ProdResponse, error)  {
	return &ProdResponse{ProdStock: 20}, nil
}
