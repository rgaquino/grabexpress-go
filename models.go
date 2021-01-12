package grabexpress

// ServiceType ...
type ServiceType string

// ServiceType enum
const (
	ServiceTypeInstant ServiceType = "INSTANT"
	ServiceTypeSameDay ServiceType = "SAME_DAY"
)

// PaymentMethod ...
type PaymentMethod string

// PaymentMethod enum
const (
	PaymentMethodCashless PaymentMethod = "CASHLESS"
	PaymentMethodCash     PaymentMethod = "CASH"
)

// OrderStatus ...
type OrderStatus string

// OrderStatus enum
const (
	// OrderStatusAllocating -
	OrderStatusAllocating OrderStatus = "ALLOCATING"
	// OrderStatusPickingUp -
	OrderStatusPickingUp OrderStatus = "PICKING_UP"
	// OrderStatusInDelivery -
	OrderStatusInDelivery OrderStatus = "IN_DELIVERY"
	// OrderStatusInReturn -
	OrderStatusInReturn OrderStatus = "IN_RETURN"
	// OrderStatusCanceled -
	OrderStatusCanceled OrderStatus = "CANCELED"
	// OrderStatusReturned -
	OrderStatusReturned OrderStatus = "RETURNED"
	// OrderStatusFailed -
	OrderStatusFailed OrderStatus = "FAILED"
	// OrderStatusCompleted -
	OrderStatusCompleted OrderStatus = "COMPLETED"
)

// Dimensions ...
type Dimensions struct {
	Height int64
	Width  int64
	Depth  int64
	Weight int64
}

// Coordinates ...
type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// Address ...
type Address struct {
	Address     string
	Keywords    *string
	CityCode    *string
	Coordinates Coordinates
	Extra       *map[string]string
}

// Package ...
type Package struct {
	Name        string
	Description string
	Quantity    int64
	Price       float64
	Dimensions  Dimensions
}

// CreateQuotesRequest ...
type CreateQuotesRequest struct {
	ServiceType *ServiceType
	Packages    []Package
	Origin      Address
	Destination Address
}

// CreateQuotesResponse ...
type CreateQuotesResponse struct {

}

// ErrorResponse ...
type ErrorResponse struct {
	Error string `json:"message"`
}
