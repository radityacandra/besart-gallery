package types

type CreateOrderInput struct {
	UserId       string `db:"user_id"`
	OrderItems   []OrderItem
	ShippingInfo ShippingInfo
}

type OrderItem struct {
	ProductId string
	Quantity  int
}

type ShippingInfo struct {
	FullName string
	Address  string
	PhoneNo  string
	Notes    *string
}

type OrderItemRepo struct {
	Id        string `db:"id"`
	OrderId   string `db:"order_id"`
	ProductId string `db:"product_id"`
	Quantity  int    `db:"quantity"`
	CreatedAt int64  `db:"created_at"`
	CreatedBy string `db:"created_by"`
}

type CreateOrderRepoInput struct {
	Id         string          `db:"id"`
	UserId     string          `db:"user_id"`
	OrderItems []OrderItemRepo `db:"-"`
	Status     string          `db:"status"`
	FullName   string          `db:"receiver_full_name"`
	Address    string          `db:"receiver_address"`
	PhoneNo    string          `db:"receiver_phone_no"`
	Notes      *string         `db:"shipping_notes"`
	CreatedAt  int64           `db:"created_at"`
	CreatedBy  string          `db:"created_by"`
}
