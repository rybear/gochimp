package carts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetEcommerceStoreStoreIDCartsParams creates a new GetEcommerceStoreStoreIDCartsParams object
// with the default values initialized.
func NewGetEcommerceStoreStoreIDCartsParams() *GetEcommerceStoreStoreIDCartsParams {
	var ()
	return &GetEcommerceStoreStoreIDCartsParams{}
}

/*GetEcommerceStoreStoreIDCartsParams contains all the parameters to send to the API endpoint
for the get ecommerce store store ID carts operation typically these are written to a http.Request
*/
type GetEcommerceStoreStoreIDCartsParams struct {

	/*StoreID*/
	StoreID string
}

// WithStoreID adds the storeId to the get ecommerce store store ID carts params
func (o *GetEcommerceStoreStoreIDCartsParams) WithStoreID(storeId string) *GetEcommerceStoreStoreIDCartsParams {
	o.StoreID = storeId
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetEcommerceStoreStoreIDCartsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param store_id
	if err := r.SetPathParam("store_id", o.StoreID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
