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

type UpdateStatusInput struct {
	Status  string
	UserId  string
	OrderId string
}

type UpdateOrderStatusInput struct {
	OrderId   string `db:"id"`
	Status    string `db:"status"`
	UpdatedAt int64  `db:"updated_at"`
	UpdatedBy string `db:"updated_by"`
}

type ListOrderInput struct {
	UserId   string
	Page     int
	PageSize int
}

type ListOrderOutput struct {
	Data       []OrderOutput
	Pagination Pagination
}

type OrderOutput struct {
	Id          string `db:"id"`
	Status      string `db:"status"`
	TotalAmount int64  `db:"total_amount"`
}

type Pagination struct {
	Page      int
	PageSize  int
	TotalData int64
}

type DetailOrderOutput struct {
	Id           string
	OrderTime    int64
	Status       string
	ShippingInfo ShippingInfo
	OrderItems   []OrderItemDetail
}

type OrderItemDetail struct {
	Id           string
	ProductImage string
	ProductName  string
	Quantity     int
	ProductPrice int64
}
