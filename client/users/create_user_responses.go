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

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/weaviate/weaviate/entities/models"
)

// CreateUserReader is a Reader for the CreateUser structure.
type CreateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUserCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateUserConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateUserUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUserCreated creates a CreateUserCreated with default headers values
func NewCreateUserCreated() *CreateUserCreated {
	return &CreateUserCreated{}
}

/*
CreateUserCreated describes a response with status code 201, with default header values.

User created successfully
*/
type CreateUserCreated struct {
	Payload *models.UserAPIKey
}

// IsSuccess returns true when this create user created response has a 2xx status code
func (o *CreateUserCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create user created response has a 3xx status code
func (o *CreateUserCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user created response has a 4xx status code
func (o *CreateUserCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user created response has a 5xx status code
func (o *CreateUserCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create user created response a status code equal to that given
func (o *CreateUserCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create user created response
func (o *CreateUserCreated) Code() int {
	return 201
}

func (o *CreateUserCreated) Error() string {
	return fmt.Sprintf("[POST /users/db/{user_id}][%d] createUserCreated  %+v", 201, o.Payload)
}

func (o *CreateUserCreated) String() string {
	return fmt.Sprintf("[POST /users/db/{user_id}][%d] createUserCreated  %+v", 201, o.Payload)
}

func (o *CreateUserCreated) GetPayload() *models.UserAPIKey {
	return o.Payload
}

func (o *CreateUserCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserAPIKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserBadRequest creates a CreateUserBadRequest with default headers values
func NewCreateUserBadRequest() *CreateUserBadRequest {
	return &CreateUserBadRequest{}
}

/*
CreateUserBadRequest describes a response with status code 400, with default header values.

Malformed request.
*/
type CreateUserBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create user bad request response has a 2xx status code
func (o *CreateUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user bad request response has a 3xx status code
func (o *CreateUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user bad request response has a 4xx status code
func (o *CreateUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user bad request response has a 5xx status code
func (o *CreateUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create user bad request response a status code equal to that given
func (o *CreateUserBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create user bad request response
func (o *CreateUserBadRequest) Code() int {
	return 400
}

func (o *CreateUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/db/{user_id}][%d] createUserBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserBadRequest) String() string {
	return fmt.Sprintf("[POST /users/db/{user_id}][%d] createUserBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserUnauthorized creates a CreateUserUnauthorized with default headers values
func NewCreateUserUnauthorized() *CreateUserUnauthorized {
	return &CreateUserUnauthorized{}
}

/*
CreateUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type CreateUserUnauthorized struct {
}

// IsSuccess returns true when this create user unauthorized response has a 2xx status code
func (o *CreateUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user unauthorized response has a 3xx status code
func (o *CreateUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user unauthorized response has a 4xx status code
func (o *CreateUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user unauthorized response has a 5xx status code
func (o *CreateUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create user unauthorized response a status code equal to that given
func (o *CreateUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create user unauthorized response
func (o *CreateUserUnauthorized) Code() int {
	return 401
}

func (o *CreateUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /users/db/{user_id}][%d] createUserUnauthorized ", 401)
}

func (o *CreateUserUnauthorized) String() string {
	return fmt.Sprintf("[POST /users/db/{user_id}][%d] createUserUnauthorized ", 401)
}

func (o *CreateUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserForbidden creates a CreateUserForbidden with default headers values
func NewCreateUserForbidden() *CreateUserForbidden {
	return &CreateUserForbidden{}
}

/*
CreateUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateUserForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create user forbidden response has a 2xx status code
func (o *CreateUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user forbidden response has a 3xx status code
func (o *CreateUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user forbidden response has a 4xx status code
func (o *CreateUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user forbidden response has a 5xx status code
func (o *CreateUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create user forbidden response a status code equal to that given
func (o *CreateUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create user forbidden response
func (o *CreateUserForbidden) Code() int {
	return 403
}

func (o *CreateUserForbidden) Error() string {
	return fmt.Sprintf("[POST /users/db/{user_id}][%d] createUserForbidden  %+v", 403, o.Payload)
}

func (o *CreateUserForbidden) String() string {
	return fmt.Sprintf("[POST /users/db/{user_id}][%d] createUserForbidden  %+v", 403, o.Payload)
}

func (o *CreateUserForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserNotFound creates a CreateUserNotFound with default headers values
func NewCreateUserNotFound() *CreateUserNotFound {
	return &CreateUserNotFound{}
}

/*
CreateUserNotFound describes a response with status code 404, with default header values.

user not found
*/
type CreateUserNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create user not found response has a 2xx status code
func (o *CreateUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user not found response has a 3xx status code
func (o *CreateUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user not found response has a 4xx status code
func (o *CreateUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user not found response has a 5xx status code
func (o *CreateUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create user not found response a status code equal to that given
func (o *CreateUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create user not found response
func (o *CreateUserNotFound) Code() int {
	return 404
}

func (o *CreateUserNotFound) Error() string {
	return fmt.Sprintf("[POST /users/db/{user_id}][%d] createUserNotFound  %+v", 404, o.Payload)
}

func (o *CreateUserNotFound) String() string {
	return fmt.Sprintf("[POST /users/db/{user_id}][%d] createUserNotFound  %+v", 404, o.Payload)
}

func (o *CreateUserNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserConflict creates a CreateUserConflict with default headers values
func NewCreateUserConflict() *CreateUserConflict {
	return &CreateUserConflict{}
}

/*
CreateUserConflict describes a response with status code 409, with default header values.

User already exists
*/
type CreateUserConflict struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create user conflict response has a 2xx status code
func (o *CreateUserConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user conflict response has a 3xx status code
func (o *CreateUserConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user conflict response has a 4xx status code
func (o *CreateUserConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user conflict response has a 5xx status code
func (o *CreateUserConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create user conflict response a status code equal to that given
func (o *CreateUserConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create user conflict response
func (o *CreateUserConflict) Code() int {
	return 409
}

func (o *CreateUserConflict) Error() string {
	return fmt.Sprintf("[POST /users/db/{user_id}][%d] createUserConflict  %+v", 409, o.Payload)
}

func (o *CreateUserConflict) String() string {
	return fmt.Sprintf("[POST /users/db/{user_id}][%d] createUserConflict  %+v", 409, o.Payload)
}

func (o *CreateUserConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateUserConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserUnprocessableEntity creates a CreateUserUnprocessableEntity with default headers values
func NewCreateUserUnprocessableEntity() *CreateUserUnprocessableEntity {
	return &CreateUserUnprocessableEntity{}
}

/*
CreateUserUnprocessableEntity describes a response with status code 422, with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous.
*/
type CreateUserUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create user unprocessable entity response has a 2xx status code
func (o *CreateUserUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user unprocessable entity response has a 3xx status code
func (o *CreateUserUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user unprocessable entity response has a 4xx status code
func (o *CreateUserUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user unprocessable entity response has a 5xx status code
func (o *CreateUserUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create user unprocessable entity response a status code equal to that given
func (o *CreateUserUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create user unprocessable entity response
func (o *CreateUserUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateUserUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /users/db/{user_id}][%d] createUserUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateUserUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /users/db/{user_id}][%d] createUserUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateUserUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateUserUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserInternalServerError creates a CreateUserInternalServerError with default headers values
func NewCreateUserInternalServerError() *CreateUserInternalServerError {
	return &CreateUserInternalServerError{}
}

/*
CreateUserInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type CreateUserInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create user internal server error response has a 2xx status code
func (o *CreateUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user internal server error response has a 3xx status code
func (o *CreateUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user internal server error response has a 4xx status code
func (o *CreateUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user internal server error response has a 5xx status code
func (o *CreateUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create user internal server error response a status code equal to that given
func (o *CreateUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create user internal server error response
func (o *CreateUserInternalServerError) Code() int {
	return 500
}

func (o *CreateUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users/db/{user_id}][%d] createUserInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateUserInternalServerError) String() string {
	return fmt.Sprintf("[POST /users/db/{user_id}][%d] createUserInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateUserInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateUserBody create user body
swagger:model CreateUserBody
*/
type CreateUserBody struct {

	// EXPERIMENTAL, DONT USE. THIS WILL BE REMOVED AGAIN. - set the given time as creation time
	// Format: date-time
	CreateTime strfmt.DateTime `json:"createTime,omitempty"`

	// EXPERIMENTAL, DONT USE. THIS WILL BE REMOVED AGAIN. - import api key from static user
	Import *bool `json:"import,omitempty"`
}

// Validate validates this create user body
func (o *CreateUserBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateUserBody) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(o.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"createTime", "body", "date-time", o.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create user body based on context it is used
func (o *CreateUserBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateUserBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateUserBody) UnmarshalBinary(b []byte) error {
	var res CreateUserBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
