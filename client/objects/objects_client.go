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

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new objects API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for objects API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ObjectsClassDelete(params *ObjectsClassDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsClassDeleteNoContent, error)

	ObjectsClassGet(params *ObjectsClassGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsClassGetOK, error)

	ObjectsClassHead(params *ObjectsClassHeadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsClassHeadNoContent, error)

	ObjectsClassPatch(params *ObjectsClassPatchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsClassPatchNoContent, error)

	ObjectsClassPut(params *ObjectsClassPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsClassPutOK, error)

	ObjectsClassReferencesCreate(params *ObjectsClassReferencesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsClassReferencesCreateOK, error)

	ObjectsClassReferencesDelete(params *ObjectsClassReferencesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsClassReferencesDeleteNoContent, error)

	ObjectsClassReferencesPut(params *ObjectsClassReferencesPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsClassReferencesPutOK, error)

	ObjectsCreate(params *ObjectsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsCreateOK, error)

	ObjectsDelete(params *ObjectsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsDeleteNoContent, error)

	ObjectsGet(params *ObjectsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsGetOK, error)

	ObjectsHead(params *ObjectsHeadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsHeadNoContent, error)

	ObjectsList(params *ObjectsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsListOK, error)

	ObjectsPatch(params *ObjectsPatchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsPatchNoContent, error)

	ObjectsReferencesCreate(params *ObjectsReferencesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsReferencesCreateOK, error)

	ObjectsReferencesDelete(params *ObjectsReferencesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsReferencesDeleteNoContent, error)

	ObjectsReferencesUpdate(params *ObjectsReferencesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsReferencesUpdateOK, error)

	ObjectsUpdate(params *ObjectsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsUpdateOK, error)

	ObjectsValidate(params *ObjectsValidateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsValidateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ObjectsClassDelete deletes object based on its class and UUID

Delete an object based on its collection and UUID. <br/><br/>Note: For backward compatibility, beacons also support an older, deprecated format without the collection name. As a result, when deleting a reference, the beacon specified has to match the beacon to be deleted exactly. In other words, if a beacon is present using the old format (without collection name) you also need to specify it the same way. <br/><br/>In the beacon format, you need to always use `localhost` as the host, rather than the actual hostname. `localhost` here refers to the fact that the beacon's target is on the same Weaviate instance, as opposed to a foreign instance.
*/
func (a *Client) ObjectsClassDelete(params *ObjectsClassDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsClassDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewObjectsClassDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "objects.class.delete",
		Method:             "DELETE",
		PathPattern:        "/objects/{className}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ObjectsClassDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ObjectsClassDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for objects.class.delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ObjectsClassGet gets a specific object based on its class and UUID also available as websocket bus

Get a data object based on its collection and UUID.
*/
func (a *Client) ObjectsClassGet(params *ObjectsClassGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsClassGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewObjectsClassGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "objects.class.get",
		Method:             "GET",
		PathPattern:        "/objects/{className}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ObjectsClassGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ObjectsClassGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for objects.class.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ObjectsClassHead checks object s existence based on its class and uuid

Checks if a data object exists based on its collection and uuid without retrieving it. <br/><br/>Internally it skips reading the object from disk other than checking if it is present. Thus it does not use resources on marshalling, parsing, etc., and is faster. Note the resulting HTTP request has no body; the existence of an object is indicated solely by the status code.
*/
func (a *Client) ObjectsClassHead(params *ObjectsClassHeadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsClassHeadNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewObjectsClassHeadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "objects.class.head",
		Method:             "HEAD",
		PathPattern:        "/objects/{className}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ObjectsClassHeadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ObjectsClassHeadNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for objects.class.head: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ObjectsClassPatch updates an object based on its UUID using patch semantics

Update an individual data object based on its class and uuid. This method supports json-merge style patch semantics (RFC 7396). Provided meta-data and schema values are validated. LastUpdateTime is set to the time this function is called.
*/
func (a *Client) ObjectsClassPatch(params *ObjectsClassPatchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsClassPatchNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewObjectsClassPatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "objects.class.patch",
		Method:             "PATCH",
		PathPattern:        "/objects/{className}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ObjectsClassPatchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ObjectsClassPatchNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for objects.class.patch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ObjectsClassPut updates a class object based on its uuid

Update an object based on its uuid and collection. This (`put`) method replaces the object with the provided object.
*/
func (a *Client) ObjectsClassPut(params *ObjectsClassPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsClassPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewObjectsClassPutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "objects.class.put",
		Method:             "PUT",
		PathPattern:        "/objects/{className}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ObjectsClassPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ObjectsClassPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for objects.class.put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ObjectsClassReferencesCreate adds a single reference to a class property

Add a single reference to an object. This adds a reference to the array of cross-references of the given property in the source object specified by its collection name and id
*/
func (a *Client) ObjectsClassReferencesCreate(params *ObjectsClassReferencesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsClassReferencesCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewObjectsClassReferencesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "objects.class.references.create",
		Method:             "POST",
		PathPattern:        "/objects/{className}/{id}/references/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ObjectsClassReferencesCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ObjectsClassReferencesCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for objects.class.references.create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ObjectsClassReferencesDelete deletes a single reference from the list of references

Delete the single reference that is given in the body from the list of references that this property has.
*/
func (a *Client) ObjectsClassReferencesDelete(params *ObjectsClassReferencesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsClassReferencesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewObjectsClassReferencesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "objects.class.references.delete",
		Method:             "DELETE",
		PathPattern:        "/objects/{className}/{id}/references/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ObjectsClassReferencesDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ObjectsClassReferencesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for objects.class.references.delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ObjectsClassReferencesPut replaces all references to a class property

Replace **all** references in cross-reference property of an object.
*/
func (a *Client) ObjectsClassReferencesPut(params *ObjectsClassReferencesPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsClassReferencesPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewObjectsClassReferencesPutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "objects.class.references.put",
		Method:             "PUT",
		PathPattern:        "/objects/{className}/{id}/references/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ObjectsClassReferencesPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ObjectsClassReferencesPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for objects.class.references.put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ObjectsCreate creates a new object

Create a new object. <br/><br/>Meta-data and schema values are validated. <br/><br/>**Note: Use `/batch` for importing many objects**: <br/>If you plan on importing a large number of objects, it's much more efficient to use the `/batch` endpoint. Otherwise, sending multiple single requests sequentially would incur a large performance penalty. <br/><br/>**Note: idempotence of `/objects`**: <br/>POST /objects will fail if an id is provided which already exists in the class. To update an existing object with the objects endpoint, use the PUT or PATCH method.
*/
func (a *Client) ObjectsCreate(params *ObjectsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewObjectsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "objects.create",
		Method:             "POST",
		PathPattern:        "/objects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ObjectsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ObjectsCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for objects.create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ObjectsDelete deletes an object based on its UUID

Deletes an object from the database based on its UUID. <br/><br/>**Note**: This endpoint is deprecated and will be removed in a future version. Use the `/objects/{className}/{id}` endpoint instead.
*/
func (a *Client) ObjectsDelete(params *ObjectsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewObjectsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "objects.delete",
		Method:             "DELETE",
		PathPattern:        "/objects/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ObjectsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ObjectsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for objects.delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ObjectsGet gets a specific object based on its UUID

Get a specific object based on its UUID. Also available as Websocket bus. <br/><br/>**Note**: This endpoint is deprecated and will be removed in a future version. Use the `/objects/{className}/{id}` endpoint instead.
*/
func (a *Client) ObjectsGet(params *ObjectsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewObjectsGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "objects.get",
		Method:             "GET",
		PathPattern:        "/objects/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ObjectsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ObjectsGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for objects.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ObjectsHead checks object s existence based on its UUID

Checks if an object exists in the system based on its UUID. <br/><br/>**Note**: This endpoint is deprecated and will be removed in a future version. Use the `/objects/{className}/{id}` endpoint instead.
*/
func (a *Client) ObjectsHead(params *ObjectsHeadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsHeadNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewObjectsHeadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "objects.head",
		Method:             "HEAD",
		PathPattern:        "/objects/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ObjectsHeadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ObjectsHeadNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for objects.head: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ObjectsList gets a list of objects

Lists all Objects in reverse order of creation, owned by the user that belongs to the used token.
*/
func (a *Client) ObjectsList(params *ObjectsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewObjectsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "objects.list",
		Method:             "GET",
		PathPattern:        "/objects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ObjectsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ObjectsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for objects.list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ObjectsPatch updates an object based on its UUID using patch semantics

Update an object based on its UUID (using patch semantics). This method supports json-merge style patch semantics (RFC 7396). Provided meta-data and schema values are validated. LastUpdateTime is set to the time this function is called. <br/><br/>**Note**: This endpoint is deprecated and will be removed in a future version. Use the `/objects/{className}/{id}` endpoint instead.
*/
func (a *Client) ObjectsPatch(params *ObjectsPatchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsPatchNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewObjectsPatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "objects.patch",
		Method:             "PATCH",
		PathPattern:        "/objects/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ObjectsPatchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ObjectsPatchNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for objects.patch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ObjectsReferencesCreate adds a single reference to a class property

Add a cross-reference. <br/><br/>**Note**: This endpoint is deprecated and will be removed in a future version. Use the `/objects/{className}/{id}/references/{propertyName}` endpoint instead.
*/
func (a *Client) ObjectsReferencesCreate(params *ObjectsReferencesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsReferencesCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewObjectsReferencesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "objects.references.create",
		Method:             "POST",
		PathPattern:        "/objects/{id}/references/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ObjectsReferencesCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ObjectsReferencesCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for objects.references.create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ObjectsReferencesDelete deletes a single reference from the list of references

Delete the single reference that is given in the body from the list of references that this property has. <br/><br/>**Note**: This endpoint is deprecated and will be removed in a future version. Use the `/objects/{className}/{id}/references/{propertyName}` endpoint instead.
*/
func (a *Client) ObjectsReferencesDelete(params *ObjectsReferencesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsReferencesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewObjectsReferencesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "objects.references.delete",
		Method:             "DELETE",
		PathPattern:        "/objects/{id}/references/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ObjectsReferencesDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ObjectsReferencesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for objects.references.delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ObjectsReferencesUpdate replaces all references to a class property

Replace all references in cross-reference property of an object. <br/><br/>**Note**: This endpoint is deprecated and will be removed in a future version. Use the `/objects/{className}/{id}/references/{propertyName}` endpoint instead.
*/
func (a *Client) ObjectsReferencesUpdate(params *ObjectsReferencesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsReferencesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewObjectsReferencesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "objects.references.update",
		Method:             "PUT",
		PathPattern:        "/objects/{id}/references/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ObjectsReferencesUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ObjectsReferencesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for objects.references.update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ObjectsUpdate updates an object based on its UUID

Updates an object based on its UUID. Given meta-data and schema values are validated. LastUpdateTime is set to the time this function is called. <br/><br/>**Note**: This endpoint is deprecated and will be removed in a future version. Use the `/objects/{className}/{id}` endpoint instead.
*/
func (a *Client) ObjectsUpdate(params *ObjectsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewObjectsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "objects.update",
		Method:             "PUT",
		PathPattern:        "/objects/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ObjectsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ObjectsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for objects.update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ObjectsValidate validates an object based on a schema

Validate an object's schema and meta-data without creating it. <br/><br/>If the schema of the object is valid, the request should return nothing with a plain RESTful request. Otherwise, an error object will be returned.
*/
func (a *Client) ObjectsValidate(params *ObjectsValidateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ObjectsValidateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewObjectsValidateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "objects.validate",
		Method:             "POST",
		PathPattern:        "/objects/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ObjectsValidateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ObjectsValidateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for objects.validate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
