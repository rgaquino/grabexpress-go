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
	Height int64 `json:"height"`
	Weight int64 `json:"weight"`
	Width  int64 `json:"width"`
	Depth  int64 `json:"depth"`
}

// Coordinates ...
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Waypoint ...
type Waypoint struct {
	Address     string             `json:"address"`
	Keywords    *string            `json:"keywords,omitempty"`
	CityCode    *string            `json:"cityCode,omitempty"`
	Coordinates Coordinates        `json:"coordinates"`
	Extra       *map[string]string `json:"extra,omitempty"`
}

// Package ...
type Package struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Quantity    int64      `json:"quantity"`
	Price       float64    `json:"price"`
	Dimensions  Dimensions `json:"dimensions"`
}

// Service ...
type Service struct {
	ID   int64       `json:"id"`
	Type ServiceType `json:"type"`
	Name string      `json:"name"`
}

// Currency ...
type Currency struct {
	Code     string `json:"code"`
	Symbol   string `json:"symbol"`
	Exponent int64  `json:"exponent"`
}

// Timeline ...
type Timeline struct {
	Create    *time.Time `json:"create,omitempty"`
	Allocate  *time.Time `json:"allocate,omitempty"`
	Pickup    *time.Time `json:"pickup,omitempty"`
	DropOff   *time.Time `json:"dropoff,omitempty"`
	Completed *time.Time `json:"completed,omitempty"`
	Cancel    *time.Time `json:"cancel,omitempty"`
	Return    *time.Time `json:"return,omitempty"`
	Fail      *time.Time `json:"fail,omitempty"`
}

// QuoteBase ...
type QuoteBase struct {
	Service           Service   `json:"service"`
	Currency          Currency  `json:"currency"`
	Amount            float64   `json:"amount"`
	EstimatedTimeline *Timeline `json:"estimatedTimeline,omitempty"`
	Distance          int64     `json:"distance"`
}

// Quote ...
type Quote struct {
	QuoteBase
	Packages    []Package `json:"package,omitempty"`
	Origin      Waypoint  `json:"origin"`
	Destination Waypoint  `json:"destination"`
}

// CashOnDelivery ...
type CashOnDelivery struct {
	Amount float64 `json:"amount"`
}

// Contact ...
type Contact struct {
	FirstName    string  `json:"firstName"`
	LastName     *string `json:"lastName,omitempty"`
	Title        *string `json:"title,omitempty"`
	CompanyName  *string `json:"companyName,omitempty"`
	Email        string  `json:"email"`
	Phone        string  `json:"phone"`
	IsSmsEnabled bool    `json:"smsEnabled"`
	Instruction  *string `json:"instruction,omitempty"`
}

// Schedule ...
type Schedule struct {
	PickupTimeFrom *time.Time `json:"pickupTimeFrom,omitempty"`
	PickupTimeTo   *time.Time `json:"pickupTimeTo,omitempty"`
}

// Vehicle ...
type Vehicle struct {
	LicensePlate        string `json:"licensePlate"`
	Model               string `json:"model"`
	PhysicalVehicleType string `json:"physicalVehicleType"`
}

// Courier ...
type Courier struct {
	Name        string      `json:"name"`
	Phone       string      `json:"phone"`
	PictureURL  string      `json:"pictureURL"`
	Rating      float64     `json:"rating"`
	Coordinates Coordinates `json:"coordinates"`
	Vehicle     Vehicle     `json:"vehicle"`
}

// AdvanceInfo ...
type AdvanceInfo struct {
	FailedReason string `json:"failedReason"`
}

// Delivery ...
type Delivery struct {
	DeliveryID      string          `json:"deliveryID"`
	MerchantOrderID string          `json:"merchantOrderID"`
	Quote           Quote           `json:"quote"`
	PaymentMethod   PaymentMethod   `json:"paymentMethod"`
	Status          OrderStatus     `json:"status"`
	TrackingURL     string          `json:"trackingURL"`
	Courier         *Courier        `json:"courier,omitempty"`
	Timeline        *Timeline       `json:"timeline,omitempty"`
	Schedule        *Schedule       `json:"schedule,omitempty"`
	CashOnDelivery  *CashOnDelivery `json:"cashOnDelivery,omitempty"`
	InvoiceNumber   string          `json:"invoiceNumber"`
	PickupPin       string          `json:"pickupPin"`
	AdvanceInfo     *AdvanceInfo    `json:"advanceInfo,omitempty"`
	Sender          Contact         `json:"sender"`
	Recipient       Contact         `json:"recipient"`
}

// CreateQuotesRequest ...
type CreateQuotesRequest struct {
	ServiceType *ServiceType `json:"serviceType,omitempty"`
	Packages    []Package    `json:"packages,omitempty"`
	Origin      Waypoint     `json:"origin"`
	Destination Waypoint     `json:"destination"`
}

// CreateQuotesResponse ...
type CreateQuotesResponse struct {
	Quotes      []QuoteBase `json:"quotes,omitempty"`
	Packages    []Package   `json:"packages,omitempty"`
	Origin      Waypoint    `json:"origin"`
	Destination Waypoint    `json:"destination"`
}

// CreateDeliveryRequest ...
type CreateDeliveryRequest struct {
	MerchantOrderID string          `json:"merchantOrderID"`
	ServiceType     ServiceType     `json:"serviceType"`
	PaymentMethod   *PaymentMethod  `json:"paymentMethod,omitempty"`
	Packages        []Package       `json:"packages,omitempty"`
	CashOnDelivery  *CashOnDelivery `json:"cashOnDelivery,omitempty"`
	Sender          Contact         `json:"sender"`
	Recipient       Contact         `json:"recipient"`
	Origin          Waypoint        `json:"origin"`
	Destination     Waypoint        `json:"destination"`
	Schedule        *Schedule       `json:"schedule,omitempty"`
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
