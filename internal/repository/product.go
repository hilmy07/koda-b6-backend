package repository

import (
	"backend/internal/models"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type ProductRepository struct {
	db *pgx.Conn
}

func NewProductRepository(db *pgx.Conn) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetProductList() ([]models.ProductList, error) {

	rows, err := r.db.Query(
		context.Background(),
		`SELECT
		p.id,
		p.name_product,
		p.description,
		p.base_price,
		pi.path,
		COALESCE(AVG(pr.rating),0) AS rating
		FROM products p
		LEFT JOIN product_images pi ON pi.product_id = p.id
		LEFT JOIN product_reviews pr ON pr.product_id = p.id
		GROUP BY p.id, pi.path
		LIMIT 6`,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.ProductList])

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetRecommendedProduct() ([]models.ProductList, error) {

	rows, err := r.db.Query(
		context.Background(),
		`SELECT
		p.id,
		p.name_product,
		p.description,
		p.base_price,
		pi.path,
		COALESCE(AVG(pr.rating),0) AS rating
		FROM products p
		LEFT JOIN product_images pi ON pi.product_id = p.id
		LEFT JOIN product_reviews pr ON pr.product_id = p.id
		GROUP BY p.id, pi.path
		LIMIT 4`,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.ProductList])

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProductReview() ([]models.ProductReview, error) {
	
	rows, _ := r.db.Query(
		context.Background(), `SELECT pr.id, u.fullname, message, rating FROM product_reviews pr JOIN users u ON pr.user_id = u.id`,
	)

	reviews, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[models.ProductReview])

	return reviews, nil
}

func (r *ProductRepository) GetProductDetail(productID int) (*models.ProductDetail, error) {
	var product models.ProductDetail
	var images []string
	var sizes []string
	var variants []string

	query := `
	SELECT
		p.id,
		p.name_product,
		p.base_price,
		COALESCE(
			ARRAY(
				SELECT pi.path
				FROM product_images pi
				WHERE pi.product_id = p.id
				LIMIT 4
			),
			'{}'
		) AS images,
		COALESCE(
			(SELECT COUNT(*) FROM product_reviews pr WHERE pr.product_id = p.id),
			0
		) AS review_count,
		COALESCE(
			ARRAY(
				SELECT ps.size_name
				FROM product_sizes ps
				WHERE ps.product_id = p.id
				LIMIT 3
			),
			'{}'
		) AS sizes,
		COALESCE(
			ARRAY(
				SELECT pv.variant_name
				FROM product_variants pv
				WHERE pv.product_id = p.id
				LIMIT 2
			),
			'{}'
		) AS variants
	FROM products p
	WHERE p.id = $1
	`

	err := r.db.QueryRow(context.Background(), query, productID).Scan(
		&product.ID,
		&product.NameProduct,
		&product.BasePrice,
		&images,
		&product.ReviewCount,
		&sizes,
		&variants,
	)
	if err != nil {
		return nil, err
	}

	product.Images = images
	product.Sizes = sizes
	product.Variants = variants

	return &product, nil
}

// func (r *ProductRepository) GetProduct() ([]models.Product, error) {
// 	rows, err := r.db.Query(
// 		context.Background(),
// 		`SELECT id,name_product,description,base_price,stock,created_at,updated_at 
// 		FROM products LIMIT 6`,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	product, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Product])
// 	if err != nil {
// 		if err == pgx.ErrNoRows {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}

// 	return product, nil
// }

func (r *ProductRepository) CreateProduct(req models.Product) error {

	now := time.Now()

	_, err := r.db.Exec(
		context.Background(),
		`INSERT INTO products 
		(name_product, description, base_price, stock, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6)`,
		req.Name_product,
		req.Description,
		req.Base_price,
		req.Stock,
		now,
		now,
	)

	return err
}

func (r *ProductRepository) DeleteProduct(id int) error {
	// now := time.Now()

	// id, _ := strconv.Atoi(ctx.Param("id"))

	// _, err := r.db.Exec(
	// 	context.Background(),
	// 	`DELETE FROM products 
	// 	WHERE id=$1`,
	// 	id,
	// )

	// return err
	_, err := r.db.Exec(
		context.Background(),
		`DELETE FROM products WHERE id=$1`,
		id,
	)

	return err
}

func (r *ProductRepository) GetProductVariant() ([]models.ProductVariant, error){
	rows, err := r.db.Query(
		context.Background(),
		`SELECT id,product_id,variant_name,add_price 
		FROM product_variants`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	productVariant, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.ProductVariant])
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return productVariant, nil
}

func (r *ProductRepository) CreateProductVariant(req models.ProductVariant) error {
	
	_, err := r.db.Exec(
		context.Background(),
		`INSERT INTO product_variants 
		(product_id, variant_name, add_price)
		VALUES ($1,$2,$3)`,
		req.ProductID,
		req.VariantName,
		req.AddPrice,
	)

	return err
}

func (r *ProductRepository) GetProductSize() ([]models.ProductSize, error){
	rows, err := r.db.Query(
		context.Background(),
		`SELECT id,product_id,size_name,add_price 
		FROM product_sizes`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	productSize, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.ProductSize])
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return productSize, nil
}

func (r *ProductRepository) CreateProductSize(req models.ProductSize) error {
	
	_, err := r.db.Exec(
		context.Background(),
		`INSERT INTO product_sizes 
		(product_id, size_name, add_price)
		VALUES ($1,$2,$3)`,
		req.ProductID,
		req.SizeName,
		req.AddPrice,
	)

	return err
}



