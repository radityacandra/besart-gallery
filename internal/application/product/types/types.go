package types

type ProductListInput struct {
	Page     int
	PageSize int
}

type ProductListOutput []ProductListOutputItem

type ProductListOutputItem struct {
	Id              string   `db:"id"`
	Name            string   `db:"name"`
	Image           string   `db:"image"`
	OriginalPrice   int64    `db:"original_price"`
	DiscountedPrice *int64   `db:"discounted_price"`
	Rating          *float64 `db:"rating"`
}

type GetProductInput struct {
	Page     int
	PageSize int
}

type FindProductByIdOutput struct {
	Description     string `db:"description"`
	Dimension       string `db:"dimension"`
	DiscountedPrice int    `db:"discounted_price"`
	Id              string `db:"id"`
	Image           string `db:"image"`
	Medium          string `db:"medium"`
	Name            string `db:"name"`
	OriginalPrice   int    `db:"original_price"`
	Rating          int    `db:"rating"`
}
