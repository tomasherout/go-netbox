// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/tomasherout/go-netbox/models"
)

// NewIpamPrefixesUpdateParams creates a new IpamPrefixesUpdateParams object
// with the default values initialized.
func NewIpamPrefixesUpdateParams() *IpamPrefixesUpdateParams {
	var ()
	return &IpamPrefixesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIpamPrefixesUpdateParamsWithTimeout creates a new IpamPrefixesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIpamPrefixesUpdateParamsWithTimeout(timeout time.Duration) *IpamPrefixesUpdateParams {
	var ()
	return &IpamPrefixesUpdateParams{

		timeout: timeout,
	}
}

// NewIpamPrefixesUpdateParamsWithContext creates a new IpamPrefixesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIpamPrefixesUpdateParamsWithContext(ctx context.Context) *IpamPrefixesUpdateParams {
	var ()
	return &IpamPrefixesUpdateParams{

		Context: ctx,
	}
}

// NewIpamPrefixesUpdateParamsWithHTTPClient creates a new IpamPrefixesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIpamPrefixesUpdateParamsWithHTTPClient(client *http.Client) *IpamPrefixesUpdateParams {
	var ()
	return &IpamPrefixesUpdateParams{
		HTTPClient: client,
	}
}

/*IpamPrefixesUpdateParams contains all the parameters to send to the API endpoint
for the ipam prefixes update operation typically these are written to a http.Request
*/
type IpamPrefixesUpdateParams struct {

	/*Data*/
	Data *models.WritablePrefix
	/*ID
	  A unique integer value identifying this prefix.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam prefixes update params
func (o *IpamPrefixesUpdateParams) WithTimeout(timeout time.Duration) *IpamPrefixesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes update params
func (o *IpamPrefixesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes update params
func (o *IpamPrefixesUpdateParams) WithContext(ctx context.Context) *IpamPrefixesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes update params
func (o *IpamPrefixesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes update params
func (o *IpamPrefixesUpdateParams) WithHTTPClient(client *http.Client) *IpamPrefixesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes update params
func (o *IpamPrefixesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam prefixes update params
func (o *IpamPrefixesUpdateParams) WithData(data *models.WritablePrefix) *IpamPrefixesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam prefixes update params
func (o *IpamPrefixesUpdateParams) SetData(data *models.WritablePrefix) {
	o.Data = data
}

// WithID adds the id to the ipam prefixes update params
func (o *IpamPrefixesUpdateParams) WithID(id int64) *IpamPrefixesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam prefixes update params
func (o *IpamPrefixesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamPrefixesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
