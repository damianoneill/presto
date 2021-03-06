// Code generated by go-swagger; DO NOT EDIT.

package tapi_common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUIDHandlerFunc turns a function with the right signature into a get data context topology UUID node node UUID owned node edge point owned node edge point UUID connection end point connection end point UUID handler
type GetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUIDHandlerFunc func(GetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUIDHandlerFunc) Handle(params GetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUIDParams) middleware.Responder {
	return fn(params)
}

// GetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUIDHandler interface for that can handle valid get data context topology UUID node node UUID owned node edge point owned node edge point UUID connection end point connection end point UUID params
type GetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUIDHandler interface {
	Handle(GetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUIDParams) middleware.Responder
}

// NewGetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUID creates a new http.Handler for the get data context topology UUID node node UUID owned node edge point owned node edge point UUID connection end point connection end point UUID operation
func NewGetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUID(ctx *middleware.Context, handler GetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUIDHandler) *GetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUID {
	return &GetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUID{Context: ctx, Handler: handler}
}

/*GetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUID swagger:route GET /data/context/topology={uuid}/node={node-uuid}/owned-node-edge-point={owned-node-edge-point-uuid}/connection-end-point={connection-end-point-uuid}/ tapi-common getDataContextTopologyUuidNodeNodeUuidOwnedNodeEdgePointOwnedNodeEdgePointUuidConnectionEndPointConnectionEndPointUuid

returns tapi.connectivity.ConnectionEndPoint

*/
type GetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUID struct {
	Context *middleware.Context
	Handler GetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUIDHandler
}

func (o *GetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
