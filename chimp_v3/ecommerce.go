package gochimp

import (
	"errors"
	"fmt"
)

const (
	stores_path       = "/ecommerce/stores"
	single_store_path = stores + "/%s"

	carts_path       = single_store_path + "/carts"
	single_cart_path = carts_path + "/%s"

	customers_path       = single_store_path + "/customers"
	single_customer_path = customers_path + "/%s"

	products_path       = single_store_path + "/products"
	single_product_path = products_path + "/%s"

	orders_path       = single_store_path + "/orders"
	single_order_path = orders_path + "/%s"
)

// ------------------------------------------------------------------------------------------------
// Stores
// ------------------------------------------------------------------------------------------------

type StoreCreationRequest struct {
	ID            string  `json:"id"`
	ListID        string  `json:"list_id"`
	Name          string  `json:"name"`
	Platform      string  `json:"platform"`
	Domain        string  `json:"domain"`
	EmailAddress  string  `json:"email_address"`
	CurrencyCode  string  `json:"currency_code"`
	MoneyFormat   string  `json:"money_format"`
	PrimaryLocale string  `json:"primary_locale"`
	Timezone      string  `json:"timezone"`
	Phone         string  `json:"phone"`
	Address       Address `json:"address"`
}

type Store struct {
	StoreCreationRequest
	withLinks

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`

	api *ChimpAPI
}

func (store Store) CanMakeRequest() error {
	if store.ID == "" {
		return errors.New("ID is a required field")
	}

	if store.api == nil {
		return errors.New("No referenece to the API")
	}

	return nil
}

type ListOfStores struct {
	Stores []Store `json:"stores"`
	baseList
}

func (api ChimpAPI) GetStores(params *ExtendedQueryParams) (*ListOfStores, error) {
	response := new(ListOfStores)

	err := api.Request("GET", stores_path, params, nil, response)
	if err != nil {
		return nil, err
	}

	for _, entry := range response.Stores {
		entry.api = &api
	}

	return response, nil
}

func (api ChimpAPI) GetStore(id string, params QueryParams) (*Store, error) {
	endpoint := fmt.Sprintf(single_store_path, id)

	response := new(Store)
	response.api = &api

	return response, api.Request("GET", endpoint, params, nil, response)
}

func (api ChimpAPI) DeleteStore(id string) (bool, error) {
	endpoint := fmt.Sprintf(single_store_path, id)
	return api.RequestOk("DELETE", endpoint)
}

func (api ChimpAPI) CreateStore(body *StoreCreationRequest) (*Store, error) {
	response := new(Store)
	response.api = &api
	return response, api.Request("POST", stores_path, nil, body, response)
}

func (api ChimpAPI) UpdateStore(id string, body *StoreCreationRequest) (*Store, error) {
	endpoint := fmt.Sprintf(single_list_path, id)
	response := new(Store)
	response.api = &api
	return response, api.Request("PATCH", endpoint, nil, body, response)
}

// ------------------------------------------------------------------------------------------------
// Carts
// ------------------------------------------------------------------------------------------------

type ListOfCarts struct {
	Carts []Cart `json:"cart"`
	baseList
}

type CartCreationRequest struct {
	ID           string     `json:"id"`
	Customer     Customer   `json:"customer"`
	CampaignID   string     `json:"campaign_id"`
	CheckoutURL  string     `json:"checkout_url"`
	CurrencyCode string     `json:"currency_code"`
	OrderTotal   float64    `json:"order_total"` // Float?
	TaxTotal     float64    `json:"tax_total"`   // Float?
	Lines        []LineItem `json:"lines"`
}

type Cart struct {
	CartCreationRequest
	withLinks

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type LineItem struct {
	ProductID           string  `json:"product_id"`
	ProductTitle        string  `json:"product_title"`
	ProductVariantID    string  `json:"product_variant_id"`
	ProductVariantTitle string  `json:"product_variant_title"`
	Quantity            int     `json:"quantity"`
	Price               float64 `json:"price"`
}

func (store Store) GetCarts(params *ExtendedQueryParams) (*ListOfCarts, error) {
	if err := store.CanMakeRequest(); err != nil {
		return nil, err
	}

	return store.api.GetCarts(store.ID, params)
}

func (api ChimpAPI) GetCarts(storeID string, params *ExtendedQueryParams) (*ListOfCarts, error) {
	endpoint := fmt.Sprintf(carts_path, storeID)
	response := new(ListOfCarts)
	return response, api.Request("GET", endpoint, nil, nil, response)
}

func (store Store) GetCart(cartID string, params *BasicQueryParams) (*Cart, error) {
	if err := store.CanMakeRequest(); err != nil {
		return nil, err
	}

	return store.api.GetCart(store.ID, cartID, params)
}

func (api ChimpAPI) GetCart(storeID, cartID string, params *BasicQueryParams) (*Cart, error) {
	endpoint := fmt.Sprintf(single_cart_path, storeID, cartID)

	response := new(Cart)

	return response, api.Request("GET", endpoint, params, nil, response)
}

func (store Store) CreateCart(body *CartCreationRequest) (*Cart, error) {
	if err := store.CanMakeRequest(); err != nil {
		return nil, err
	}

	return store.api.CreateCart(store.ID, body)
}

func (api ChimpAPI) CreateCart(storeID string, body *CartCreationRequest) (*Cart, error) {
	endpoint := fmt.Sprintf(single_store_path, storeID)
	response := new(Cart)
	return response, api.Request("POST", endpoint, nil, body, response)
}

func (store Store) UpdateCart(id string, body *CartCreationRequest) (*Cart, error) {
	if err := store.CanMakeRequest(); err != nil {
		return nil, err
	}

	return store.api.UpdateCart(store.ID, id, body)
}

func (api ChimpAPI) UpdateCart(storeID, cartID string, body *CartCreationRequest) (*Cart, error) {
	endpoint := fmt.Sprintf(single_cart_path, storeID, cartID)
	response := new(Cart)
	return response, api.Request("PATCH", endpoint, nil, body, response)
}

// ------------------------------------------------------------------------------------------------
// Customers
// ------------------------------------------------------------------------------------------------

type ListOfCustomers struct {
	baseList

	StoreID  string   `json:""`
	Customer Customer `json:"customer"`
}

func (store Store) GetCustomers(params *ExtendedQueryParams) (*ListOfCustomers, error) {
	if err := store.CanMakeRequest(); err != nil {
		return nil, err
	}

	return store.api.GetCustomers(store.ID, params)
}

func (api ChimpAPI) GetCustomers(storeID string, params *ExtendedQueryParams) (*ListOfCustomers, error) {
	endpoint := fmt.Sprintf(customers_path, storeID)
	response := new(ListOfCustomers)
	return response, api.Request("GET", endpoint, params, nil, response)
}

func (store Store) GetCustomer(customerID string, params *BasicQueryParams) (*Customer, error) {
	if err := store.CanMakeRequest(); err != nil {
		return nil, err
	}

	return store.api.GetCustomer(store.ID, customerID, params)
}

func (api ChimpAPI) GetCustomer(storeID, customerID string, params *BasicQueryParams) (*Customer, error) {
	endpoint := fmt.Sprintf(single_customer_path, storeID, customerID)
	response := new(Customer)
	return response, api.Request("GET", endpoint, params, nil, response)
}

func (store Store) DeleteCustomer(customerID string) (bool, error) {
	if err := store.CanMakeRequest(); err != nil {
		return false, err
	}
	return store.api.DeleteCustomer(store.ID, customerID)
}

func (api ChimpAPI) DeleteCustomer(storeID, customerID string) (bool, error) {
	endpoint := fmt.Sprintf(single_customer_path, storeID, customerID)
	return api.RequestOk("DELETE", endpoint)
}

func (store Store) CreateCustomer(body *CustomerCreationRequest) (*Customer, error) {
	if err := store.CanMakeRequest(); err != nil {
		return nil, err
	}
	return store.api.CreateCustomer(store.ID, body)
}

func (api ChimpAPI) CreateCustomer(storeID string, body *CustomerCreationRequest) (*Customer, error) {
	endpoint := fmt.Sprintf(customers_path, storeID)
	response := new(Customer)
	return response, api.Request("POST", endpoint, nil, body, response)
}

func (store Store) UpdateCustomer(customerID string, body *CustomerCreationRequest) (*Customer, error) {
	if err := store.CanMakeRequest(); err != nil {
		return nil, err
	}
	return store.api.UpdateCustomer(store.ID, customerID, body)
}

func (api ChimpAPI) UpdateCustomer(storeID, customerID string, body *CustomerCreationRequest) (*Customer, error) {
	endpoint := fmt.Sprintf(single_customer_path, storeID, customerID)
	response := new(Customer)
	return response, api.Request("PATCH", endpoint, nil, body, response)
}

// ------------------------------------------------------------------------------------------------
// Orders
// ------------------------------------------------------------------------------------------------

type OrderQueryParams struct {
	ExtendedQueryParams
	CustomerID string
}

func (q OrderQueryParams) Params() map[string]string {
	m := q.ExtendedQueryParams.Params()
	m["customer_id"] = q.CustomerID
	return m
}

type OrderCreationRequest struct {
	ID                 string     `json:"id"`
	Customer           Customer   `json:"customer"`
	CampaignID         string     `json:"campaign_id"`
	FinancialStatus    string     `json:"financial_status"`
	FulfillmentStatus  string     `json:"fulfillment_status"`
	CurrencyCode       string     `json:"currency_code"`
	OrderTotal         float64    `json:"order_total"`
	TaxTotal           float64    `json:"tax_total"`
	ShippingTotal      float64    `json:"shipping_total"`
	ProcessedAtForeign string     `json:"processed_at_foreign"`
	CancelledAtForeign string     `json:"cancelled_at_foreign"`
	UpdatedAtForeign   string     `json:"updated_at_foreign"`
	ShippingAddress    Address    `json:"shipping_address"`
	BillingAddress     Address    `json:"billing_address"`
	Lines              []LineItem `json:"lines"`
}

type Order struct {
	OrderCreationRequest
	withLinks

	StoreID string
	api     *ChimpAPI
}

func (order Order) CanMakeRequest() error {
	if order.ID == "" {
		return errors.New("ID is required")
	}
	if order.api == nil {
		return errors.New("api is required")
	}
	if order.StoreID == "" {
		return errors.New("StoreID is required")
	}
	return nil
}

type ListOfOrders struct {
	baseList
	StoreID string  `json:"store_id"`
	Orders  []Order `json:"orders"`
}

func (store Store) GetOrders(params *OrderQueryParams) (*ListOfOrders, error) {
	if err := store.CanMakeRequest(); err != nil {
		return nil, err
	}
	return store.api.GetOrders(store.ID, params)
}

func (api ChimpAPI) GetOrders(storeID string, params *OrderQueryParams) (*ListOfOrders, error) {
	endpoint := fmt.Sprintf(orders_path, storeID)
	response := new(ListOfOrders)
	err := api.Request("GET", endpoint, params, nil, response)
	if err != nil {
		return nil, err
	}

	for _, r := range response.Orders {
		r.api = &api
		r.StoreID = storeID
	}

	return response, nil
}

func (store Store) GetOrder(orderID string, params *BasicQueryParams) (*Order, error) {
	if err := store.CanMakeRequest(); err != nil {
		return nil, err
	}
	return store.api.GetOrder(store.ID, orderID, params)
}

func (api ChimpAPI) GetOrder(storeID string, orderID string, params *BasicQueryParams) (*Order, error) {
	endpoint := fmt.Sprintf(single_order_path, storeID, orderID)
	response := new(Order)
	response.api = &api
	response.StoreID = storeID
	return response, api.Request("GET", endpoint, params, nil, response)
}

func (store Store) CreateOrder(body *OrderCreationRequest) (*Order, error) {
	if err := store.CanMakeRequest(); err != nil {
		return nil, err
	}
	return store.api.CreateOrder(store.ID, body)
}

func (api ChimpAPI) CreateOrder(storeID string, body *OrderCreationRequest) (*Order, error) {
	endpoint := fmt.Sprintf(orders_path, storeID)
	response := new(Order)
	response.api = &api
	response.StoreID = storeID
	return response, api.Request("POST", endpoint, nil, body, response)
}

func (store Store) UpdateOrder(customerID string, body *OrderCreationRequest) (*Order, error) {
	if err := store.CanMakeRequest(); err != nil {
		return nil, err
	}
	return store.api.UpdateOrder(store.ID, customerID, body)
}

func (api ChimpAPI) UpdateOrder(storeID, orderID string, body *OrderCreationRequest) (*Order, error) {
	endpoint := fmt.Sprintf(single_order_path, storeID, orderID)
	response := new(Order)
	response.api = &api
	response.StoreID = storeID
	return response, api.Request("PATCH", endpoint, nil, body, response)
}

// ------------------------------------------------------------------------------------------------
// Order Lines
// ------------------------------------------------------------------------------------------------

type ListOfOrderLines struct {
	baseList

	StoreID string      `json:"store_id"`
	OrderID string      `json:"order_id"`
	Lines   []OrderLine `json:"lines"`
}

type OrderLine struct {
}

func (order Order) GetLines(params *ExtendedQueryParams) (*ListOfOrderLines, error) {
	if err := order.CanMakeRequest(); err != nil {
		return nil, err
	}
	return order.api.GetLines(order.StoreID, order.ID, params)
}

func (api ChimpAPI) GetLines(storeID, orderID string, params *ExtendedQueryParams) (*ListOfOrderLines, error) {
}

func (order Order) GetLine(id string, params *BasicQueryParams) (*OrderLine, error) {
	if err := order.CanMakeRequest(); err != nil {
		return nil, err
	}
	return order.api.GetLine(order.StoreID, order.ID, id, params)
}

func (api ChimpAPI) GetLine(storeID, orderID, lineID string, params *BasicQueryParams) (*OrderLine, error) {
}

//func (order Order) CreateLine(
//func (api ChimpAPI) CreateLine(
//func (order Order) UpdateLine(
//func (api ChimpAPI) UpdateLine(storeID, orderID, lineID string

// ------------------------------------------------------------------------------------------------
// Products
// ------------------------------------------------------------------------------------------------
