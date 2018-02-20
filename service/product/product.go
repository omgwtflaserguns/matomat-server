package product

import (
	"context"
	"github.com/omgwtflaserguns/matomat-server/db"
	pb "github.com/omgwtflaserguns/matomat-server/generated"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("log")

type Service struct{}

func (s *Service) ListProducts(ctx context.Context, in *pb.ProductRequest) (*pb.ProductList, error) {
	rows, err := db.DbCon.Query("SELECT id, name, price FROM PRODUCT")

	if err != nil {
		logger.Panic(err)
	}

	products := []*pb.Product{}
	for rows.Next() {

		var id int32
		var name string
		var price float32

		err = rows.Scan(&id, &name, &price)
		if err != nil {
			logger.Panicf("Scan failed: %v", err)
		}
		products = append(products, &pb.Product{Id: id, Name: name, Price: price})
	}
	return &pb.ProductList{Products: products}, nil
}