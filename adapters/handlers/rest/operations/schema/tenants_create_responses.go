//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// TenantsCreateOKCode is the HTTP code returned for type TenantsCreateOK
const TenantsCreateOKCode int = 200

/*
TenantsCreateOK Added new tenants to the specified class

swagger:response tenantsCreateOK
*/
type TenantsCreateOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Tenant `json:"body,omitempty"`
}

// NewTenantsCreateOK creates TenantsCreateOK with default headers values
func NewTenantsCreateOK() *TenantsCreateOK {

	return &TenantsCreateOK{}
}

// WithPayload adds the payload to the tenants create o k response
func (o *TenantsCreateOK) WithPayload(payload []*models.Tenant) *TenantsCreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the tenants create o k response
func (o *TenantsCreateOK) SetPayload(payload []*models.Tenant) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TenantsCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Tenant, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// TenantsCreateUnauthorizedCode is the HTTP code returned for type TenantsCreateUnauthorized
const TenantsCreateUnauthorizedCode int = 401

/*
TenantsCreateUnauthorized Unauthorized or invalid credentials.

swagger:response tenantsCreateUnauthorized
*/
type TenantsCreateUnauthorized struct {
}

// NewTenantsCreateUnauthorized creates TenantsCreateUnauthorized with default headers values
func NewTenantsCreateUnauthorized() *TenantsCreateUnauthorized {

	return &TenantsCreateUnauthorized{}
}

// WriteResponse to the client
func (o *TenantsCreateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// TenantsCreateForbiddenCode is the HTTP code returned for type TenantsCreateForbidden
const TenantsCreateForbiddenCode int = 403

/*
TenantsCreateForbidden Forbidden

swagger:response tenantsCreateForbidden
*/
type TenantsCreateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewTenantsCreateForbidden creates TenantsCreateForbidden with default headers values
func NewTenantsCreateForbidden() *TenantsCreateForbidden {

	return &TenantsCreateForbidden{}
}

// WithPayload adds the payload to the tenants create forbidden response
func (o *TenantsCreateForbidden) WithPayload(payload *models.ErrorResponse) *TenantsCreateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the tenants create forbidden response
func (o *TenantsCreateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TenantsCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TenantsCreateUnprocessableEntityCode is the HTTP code returned for type TenantsCreateUnprocessableEntity
const TenantsCreateUnprocessableEntityCode int = 422

/*
TenantsCreateUnprocessableEntity Invalid Tenant class

swagger:response tenantsCreateUnprocessableEntity
*/
type TenantsCreateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewTenantsCreateUnprocessableEntity creates TenantsCreateUnprocessableEntity with default headers values
func NewTenantsCreateUnprocessableEntity() *TenantsCreateUnprocessableEntity {

	return &TenantsCreateUnprocessableEntity{}
}

// WithPayload adds the payload to the tenants create unprocessable entity response
func (o *TenantsCreateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *TenantsCreateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the tenants create unprocessable entity response
func (o *TenantsCreateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TenantsCreateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TenantsCreateInternalServerErrorCode is the HTTP code returned for type TenantsCreateInternalServerError
const TenantsCreateInternalServerErrorCode int = 500

/*
TenantsCreateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response tenantsCreateInternalServerError
*/
type TenantsCreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewTenantsCreateInternalServerError creates TenantsCreateInternalServerError with default headers values
func NewTenantsCreateInternalServerError() *TenantsCreateInternalServerError {

	return &TenantsCreateInternalServerError{}
}

// WithPayload adds the payload to the tenants create internal server error response
func (o *TenantsCreateInternalServerError) WithPayload(payload *models.ErrorResponse) *TenantsCreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the tenants create internal server error response
func (o *TenantsCreateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TenantsCreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
