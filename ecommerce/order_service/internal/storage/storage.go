package storage

import (
	"database/sql"
	"order_service/internal/service"
	"order_service/proto/order"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (o *OrderRepo) CreateOrder(req *order.Order) (*order.OrderResponse, error) {
	dbQuery := `
		INSERT INTO orders
			(user_id, product_id, location)
		VALUES
			($1, $2, $3)
	`
	_, err := o.db.Exec(dbQuery, req.UserId, req.ProductId, req.Locate)
	if err != nil {
		return nil, err
	}
	return &order.OrderResponse{
		Message: "Success Create order",
	}, nil
}

func (o *OrderRepo) UpdateOrder(req *order.Order) (*order.OrderResponse, error) {
	dbQuery := `
		UPDATE orders
		SET
			user_id=$1, product_id=$2, location=$3, update_at=CURRENT_TIMESTAMP
		WHERE
			id = $4
	`
	_, err := o.db.Exec(dbQuery, req.UserId, req.ProductId, req.Locate, req.Id)
	if err != nil {
		return nil, err
	}
	return &order.OrderResponse{
		Message: "Success Update order",
	}, nil
}

func (o *OrderRepo) DeleteOrder(req *order.OderRequest) (*order.OrderResponse, error) {
	dbQuery := `
		UPDATE orders
		SET
			deleted_at=$1
		WHERE
			id = $2
	`
	_, err := o.db.Exec(dbQuery, time.Now().Unix(), req.Id)
	if err != nil {
		return nil, err
	}
	return &order.OrderResponse{
		Message: "Success Delete order",
	}, nil
}

func (o *OrderRepo) GetOrder(req *order.OderRequest) (*order.FullInfo, error) {
	var (
		id         string
		userId     string
		productId  string
		locate     string
		created_at time.Time
		updated_at time.Time
		deleted_at int64
	)

	dbQuery := `
		SELECT * FROM orders WHERE id = $1
	`
	row := o.db.QueryRow(dbQuery, req.Id)
	err := row.Scan(
		&id,
		&userId,
		&productId,
		&locate,
		&created_at,
		&updated_at,
		&deleted_at,
	)
	if err != nil {
		return nil, err
	}

	respUser, err := service.UserService(userId)
	if err != nil {
		return nil, err
	}
	respProduct, err := service.ProductService(productId)
	if err != nil {
		return nil, err
	}

	user := order.UserST{
		Id:        respUser.Id,
		FirstName: respUser.FirstName,
		LastName:  respUser.LastName,
		Email:     respUser.Email,
		CreatedAt: timestamppb.New(respUser.CreatedAt.AsTime()),
		UpdatedAt: timestamppb.New(respUser.UpdatedAt.AsTime()),
		DeletedAt: respUser.DeletedAt,
	}

	product := order.ProductST{
		Id:           respProduct.Id,
		ProductImg:   respProduct.ProductImg,
		ProductName:  respProduct.ProductName,
		ProductPrice: respProduct.ProductPrice,
		ProductDesc:  respProduct.ProductDesc,
		CreatedAt:    timestamppb.New(respProduct.CreatedAt.AsTime()),
		UpdatedAt:    timestamppb.New(respProduct.UpdatedAt.AsTime()),
		DeletedAt:    respProduct.DeletedAt,
	}

	fullInfo := &order.FullInfo{
		Id:        id,
		Product:   &product,
		User:      &user,
		Locate:    locate,
		CreatedAt: timestamppb.New(created_at),
		UpdatedAt: timestamppb.New(updated_at),
		DeletedAt: deleted_at,
	}
	return fullInfo, nil
}
