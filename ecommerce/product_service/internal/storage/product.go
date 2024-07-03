package storage

import (
	"database/sql"
	"time"

	"product_service/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (p *ProductRepo) CreateProduct(req *proto.Product) (*proto.ProductResponse, error) {
	dbQuery := `
		INSERT INTO products(
			product_img,
			product_name,
			product_price,
			product_desc
		) VALUES (
			$1, $2, $3, $4 
		)
	`
	_, err := p.db.Exec(dbQuery, req.ProductImg, req.ProductName, req.ProductPrice, req.ProductDesc)
	if err != nil {
		return nil, err
	}

	return &proto.ProductResponse{
		Message: "Success created product",
	}, nil
}

func (p *ProductRepo) UpdateProduct(req *proto.Product) (*proto.ProductResponse, error) {
	dbQuery := `
		UPDATE products
		SET
			product_img = $1,
			product_name = $2,
			product_price = $3,
			product_desc = $4,
			updated_at = CURRENT_TIMESTAMP
		WHERE
			id = $5
	`
	_, err := p.db.Exec(dbQuery, req.ProductImg, req.ProductName, req.ProductPrice, req.ProductDesc, req.Id)
	if err != nil {
		return nil, err
	}

	return &proto.ProductResponse{
		Message: "Success updated product",
	}, nil
}

func (p *ProductRepo) GetProduct(req *proto.ProductRequest) (*proto.Product, error) {
	var (
		product   proto.Product
		createdAt time.Time
		updatedAt time.Time
	)

	dbQuery := `
		SELECT id, product_img, product_name, product_price, product_desc, created_at, updated_at, deleted_at 
		FROM products 
		WHERE id = $1
	`
	row := p.db.QueryRow(dbQuery, req.Id)
	err := row.Scan(
		&product.Id,
		&product.ProductImg,
		&product.ProductName,
		&product.ProductPrice,
		&product.ProductDesc,
		&createdAt,
		&updatedAt,
		&product.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	product.CreatedAt = timestamppb.New(createdAt)
	product.UpdatedAt = timestamppb.New(updatedAt)

	return &product, nil
}

func (p *ProductRepo) DeleteProduct(req *proto.ProductRequest) (*proto.ProductResponse, error) {
	dbQuery := `
		UPDATE products
		SET 
			deleted_at = $1
		WHERE
			id = $2
	`
	_, err := p.db.Exec(dbQuery, time.Now().Unix(), req.Id)
	if err != nil {
		return nil, err
	}

	return &proto.ProductResponse{
		Message: "Success Deleted product",
	}, nil
}
