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
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/weaviate/weaviate/entities/models"
)

// ObjectsReferencesDeleteHandlerFunc turns a function with the right signature into a objects references delete handler
type ObjectsReferencesDeleteHandlerFunc func(ObjectsReferencesDeleteParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ObjectsReferencesDeleteHandlerFunc) Handle(params ObjectsReferencesDeleteParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ObjectsReferencesDeleteHandler interface for that can handle valid objects references delete params
type ObjectsReferencesDeleteHandler interface {
	Handle(ObjectsReferencesDeleteParams, *models.Principal) middleware.Responder
}

// NewObjectsReferencesDelete creates a new http.Handler for the objects references delete operation
func NewObjectsReferencesDelete(ctx *middleware.Context, handler ObjectsReferencesDeleteHandler) *ObjectsReferencesDelete {
	return &ObjectsReferencesDelete{Context: ctx, Handler: handler}
}

/*
	ObjectsReferencesDelete swagger:route DELETE /objects/{id}/references/{propertyName} objects objectsReferencesDelete

Delete a single reference from the list of references.

Delete the single reference that is given in the body from the list of references that this property has. <br/><br/>**Note**: This endpoint is deprecated and will be removed in a future version. Use the `/objects/{className}/{id}/references/{propertyName}` endpoint instead.
*/
type ObjectsReferencesDelete struct {
	Context *middleware.Context
	Handler ObjectsReferencesDeleteHandler
}

func (o *ObjectsReferencesDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewObjectsReferencesDeleteParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
