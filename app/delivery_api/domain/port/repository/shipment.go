package repository

import (
	"context"
	"github.com/tkame123/ddd-sample/app/delivery_api/domain/model"
)

type Shipment interface {
	FindOne(ctx context.Context, id *model.ShipmentID) (*model.Shipment, error)
	Save(ctx context.Context, order *model.Shipment) error
}
