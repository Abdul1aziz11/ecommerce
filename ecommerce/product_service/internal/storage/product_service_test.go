package storage

import (
	"product_service/proto"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCreateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewProductRepo(db)
	product := &proto.Product{
		ProductImg:   "image.jpg",
		ProductName:  "Test Product",
		ProductPrice: "10.99",
		ProductDesc:  "This is a test product",
	}

	mock.ExpectExec(`INSERT INTO products`).
		WithArgs(product.ProductImg, product.ProductName, product.ProductPrice, product.ProductDesc).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := repo.CreateProduct(product)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Success created product", resp.Message)
}

func TestUpdateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewProductRepo(db)
	product := &proto.Product{
		Id:           "1",
		ProductImg:   "image.jpg",
		ProductName:  "Test Product",
		ProductPrice: "10.99",
		ProductDesc:  "This is a test product",
	}

	mock.ExpectExec(`UPDATE products`).
		WithArgs(product.ProductImg, product.ProductName, product.ProductPrice, product.ProductDesc, product.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := repo.UpdateProduct(product)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Success updated product", resp.Message)
}

func TestDeleteProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewProductRepo(db)
	req := &proto.ProductRequest{
		Id: "1",
	}

	mock.ExpectExec(`UPDATE products`).
		WithArgs(sqlmock.AnyArg(), req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := repo.DeleteProduct(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Success Deleted product", resp.Message)
}

func TestGetProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewProductRepo(db)
	req := &proto.ProductRequest{
		Id: "1",
	}

	rows := sqlmock.NewRows([]string{"id", "product_img", "product_name", "product_price", "product_desc", "created_at", "updated_at", "deleted_at"}).
		AddRow("1", "image.jpg", "Test Product", "10.99", "This is a test product", time.Now(), time.Now(), 0)

	mock.ExpectQuery(`SELECT \* FROM products WHERE id = \$1`).
		WithArgs(req.Id).
		WillReturnRows(rows)

	product, err := repo.GetProduct(req)
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "1", product.Id)
	assert.Equal(t, "image.jpg", product.ProductImg)
	assert.Equal(t, "Test Product", product.ProductName)
	assert.Equal(t, "10.99", product.ProductPrice)
	assert.Equal(t, "This is a test product", product.ProductDesc)
	assert.IsType(t, &timestamppb.Timestamp{}, product.CreatedAt)
	assert.IsType(t, &timestamppb.Timestamp{}, product.UpdatedAt)
	assert.Equal(t, int64(0), product.DeletedAt)
}
