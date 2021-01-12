package grabexpress

import "time"

// ServiceType ...
type ServiceType string

// ServiceType enum
const (
	ServiceTypeInstant ServiceType = "INSTANT"
	ServiceTypeSameDay ServiceType = "SAME_DAY"
	ServiceTypeBulk    ServiceType = "BULK"
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
	// OrderStatusQueueing -
	OrderStatusQueueing OrderStatus = "QUEUING"
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

// Waypoint ...
type Waypoint struct {
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

// Service ...
type Service struct {
	ID   int64
	Type ServiceType
	Name string
}

// Currency ...
type Currency struct {
	Code     string
	Symbol   string
	Exponent int64
}

// Timeline ...
type Timeline struct {
	Create    *time.Time
	Allocate  *time.Time
	Pickup    *time.Time
	DropOff   *time.Time
	Completed *time.Time
	Cancel    *time.Time
	Return    *time.Time
	Fail      *time.Time
}

// QuoteBase ...
type QuoteBase struct {
	Service           Service
	Currency          Currency
	Amount            float64
	EstimatedTimeline Timeline
	Distance          int64
}

// Quote ...
type Quote struct {
	QuoteBase
	Packages    []Package
	Origin      Waypoint
	Destination Waypoint
}

// CashOnDelivery ...
type CashOnDelivery struct {
	Amount float64
}

// Contact ...
type Contact struct {
	FirstName    string
	LastName     *string
	Title        *string
	CompanyName  *string
	Email        string
	Phone        string
	IsSmsEnabled bool
	Instruction  *string
}

// Schedule ...
type Schedule struct {
	PickupTimeFrom *time.Time
	PickupTimeTo   *time.Time
}

// Vehicle ...
type Vehicle struct {
	LicensePlate        string
	Model               string
	PhysicalVehicleType string
}

// Courier ...
type Courier struct {
	Name        string
	Phone       string
	PictureURL  string
	Rating      float64
	Coordinates Coordinates
	Vehicle     Vehicle
}

// AdvanceInfo ...
type AdvanceInfo struct {
	FailedReason string
}

// Delivery ...
type Delivery struct {
	DeliveryID      string
	MerchantOrderID string
	Quote           Quote
	PaymentMethod   PaymentMethod
	Status          OrderStatus
	TrackingURL     string
	Courier         Courier
	Timeline        Timeline
	Schedule        Schedule
	CashOnDelivery  CashOnDelivery
	InvoiceNumber   string
	PickupPin       string
	AdvanceInfo     AdvanceInfo
	Sender          Contact
	Recipient       Contact
}

// CreateQuotesRequest ...
type CreateQuotesRequest struct {
	ServiceType *ServiceType
	Packages    []Package
	Origin      Waypoint
	Destination Waypoint
}

// CreateQuotesResponse ...
type CreateQuotesResponse struct {
	Quotes      []QuoteBase
	Packages    []Package
	Origin      Waypoint
	Destination Waypoint
}

// CreateDeliveryRequest ...
type CreateDeliveryRequest struct {
	MerchantOrderID string
	ServiceType     ServiceType
	PaymentMethod   *PaymentMethod
	Packages        []Package
	CashOnDelivery  CashOnDelivery
	Sender          Contact
	Recipient       Contact
	Origin          Waypoint
	Destination     Waypoint
	Schedule        *Schedule
}

// CreateDeliveryResponse ...
type CreateDeliveryResponse struct {
	Delivery
}

// GetDeliveryResponse ...
type GetDeliveryResponse struct {
	Delivery
}

// ErrorResponse ...
type ErrorResponse struct {
	Error string `json:"message"`
}
