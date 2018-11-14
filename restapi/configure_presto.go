// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/damianoneill/presto/restapi/operations"
	"github.com/damianoneill/presto/restapi/operations/tapi_common"
	"github.com/damianoneill/presto/restapi/operations/tapi_connectivity"
	"github.com/damianoneill/presto/restapi/operations/tapi_path_computation"
	"github.com/damianoneill/presto/restapi/operations/tapi_topology"
)

//go:generate swagger generate server --target .. --name presto --spec ../presto-nrp.yaml

func configureFlags(api *operations.PrestoAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.PrestoAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.XMLConsumer = runtime.XMLConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.XMLProducer = runtime.XMLProducer()

	api.TapiCommonGetDataContextHandler = tapi_common.GetDataContextHandlerFunc(func(params tapi_common.GetDataContextParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.GetDataContext has not yet been implemented")
	})
	api.TapiCommonGetDataContextConnectionUUIDHandler = tapi_common.GetDataContextConnectionUUIDHandlerFunc(func(params tapi_common.GetDataContextConnectionUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.GetDataContextConnectionUUID has not yet been implemented")
	})
	api.TapiCommonGetDataContextConnectionUUIDSwitchControlSwitchControlUUIDHandler = tapi_common.GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDHandlerFunc(func(params tapi_common.GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.GetDataContextConnectionUUIDSwitchControlSwitchControlUUID has not yet been implemented")
	})
	api.TapiCommonGetDataContextConnectivityServiceUUIDHandler = tapi_common.GetDataContextConnectivityServiceUUIDHandlerFunc(func(params tapi_common.GetDataContextConnectivityServiceUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.GetDataContextConnectivityServiceUUID has not yet been implemented")
	})
	api.TapiCommonGetDataContextNwTopologyServiceHandler = tapi_common.GetDataContextNwTopologyServiceHandlerFunc(func(params tapi_common.GetDataContextNwTopologyServiceParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.GetDataContextNwTopologyService has not yet been implemented")
	})
	api.TapiCommonGetDataContextPathCompServiceUUIDHandler = tapi_common.GetDataContextPathCompServiceUUIDHandlerFunc(func(params tapi_common.GetDataContextPathCompServiceUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.GetDataContextPathCompServiceUUID has not yet been implemented")
	})
	api.TapiCommonGetDataContextPathUUIDHandler = tapi_common.GetDataContextPathUUIDHandlerFunc(func(params tapi_common.GetDataContextPathUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.GetDataContextPathUUID has not yet been implemented")
	})
	api.TapiCommonGetDataContextServiceInterfacePointUUIDHandler = tapi_common.GetDataContextServiceInterfacePointUUIDHandlerFunc(func(params tapi_common.GetDataContextServiceInterfacePointUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.GetDataContextServiceInterfacePointUUID has not yet been implemented")
	})
	api.TapiCommonGetDataContextTopologyUUIDHandler = tapi_common.GetDataContextTopologyUUIDHandlerFunc(func(params tapi_common.GetDataContextTopologyUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.GetDataContextTopologyUUID has not yet been implemented")
	})
	api.TapiCommonGetDataContextTopologyUUIDLinkLinkUUIDHandler = tapi_common.GetDataContextTopologyUUIDLinkLinkUUIDHandlerFunc(func(params tapi_common.GetDataContextTopologyUUIDLinkLinkUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.GetDataContextTopologyUUIDLinkLinkUUID has not yet been implemented")
	})
	api.TapiCommonGetDataContextTopologyUUIDNodeNodeUUIDHandler = tapi_common.GetDataContextTopologyUUIDNodeNodeUUIDHandlerFunc(func(params tapi_common.GetDataContextTopologyUUIDNodeNodeUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.GetDataContextTopologyUUIDNodeNodeUUID has not yet been implemented")
	})
	api.TapiCommonGetDataContextTopologyUUIDNodeNodeUUIDNodeRuleGroupNodeRuleGroupUUIDHandler = tapi_common.GetDataContextTopologyUUIDNodeNodeUUIDNodeRuleGroupNodeRuleGroupUUIDHandlerFunc(func(params tapi_common.GetDataContextTopologyUUIDNodeNodeUUIDNodeRuleGroupNodeRuleGroupUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.GetDataContextTopologyUUIDNodeNodeUUIDNodeRuleGroupNodeRuleGroupUUID has not yet been implemented")
	})
	api.TapiCommonGetDataContextTopologyUUIDNodeNodeUUIDNodeRuleGroupNodeRuleGroupUUIDInterRuleGroupInterRuleGroupUUIDHandler = tapi_common.GetDataContextTopologyUUIDNodeNodeUUIDNodeRuleGroupNodeRuleGroupUUIDInterRuleGroupInterRuleGroupUUIDHandlerFunc(func(params tapi_common.GetDataContextTopologyUUIDNodeNodeUUIDNodeRuleGroupNodeRuleGroupUUIDInterRuleGroupInterRuleGroupUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.GetDataContextTopologyUUIDNodeNodeUUIDNodeRuleGroupNodeRuleGroupUUIDInterRuleGroupInterRuleGroupUUID has not yet been implemented")
	})
	api.TapiCommonGetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDHandler = tapi_common.GetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDHandlerFunc(func(params tapi_common.GetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.GetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUID has not yet been implemented")
	})
	api.TapiCommonGetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUIDHandler = tapi_common.GetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUIDHandlerFunc(func(params tapi_common.GetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUIDParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.GetDataContextTopologyUUIDNodeNodeUUIDOwnedNodeEdgePointOwnedNodeEdgePointUUIDConnectionEndPointConnectionEndPointUUID has not yet been implemented")
	})
	api.TapiPathComputationPostOperationsComputeP2PPathHandler = tapi_path_computation.PostOperationsComputeP2PPathHandlerFunc(func(params tapi_path_computation.PostOperationsComputeP2PPathParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_path_computation.PostOperationsComputeP2PPath has not yet been implemented")
	})
	api.TapiConnectivityPostOperationsCreateConnectivityServiceHandler = tapi_connectivity.PostOperationsCreateConnectivityServiceHandlerFunc(func(params tapi_connectivity.PostOperationsCreateConnectivityServiceParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_connectivity.PostOperationsCreateConnectivityService has not yet been implemented")
	})
	api.TapiConnectivityPostOperationsDeleteConnectivityServiceHandler = tapi_connectivity.PostOperationsDeleteConnectivityServiceHandlerFunc(func(params tapi_connectivity.PostOperationsDeleteConnectivityServiceParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_connectivity.PostOperationsDeleteConnectivityService has not yet been implemented")
	})
	api.TapiPathComputationPostOperationsDeleteP2PPathHandler = tapi_path_computation.PostOperationsDeleteP2PPathHandlerFunc(func(params tapi_path_computation.PostOperationsDeleteP2PPathParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_path_computation.PostOperationsDeleteP2PPath has not yet been implemented")
	})
	api.TapiConnectivityPostOperationsGetConnectionDetailsHandler = tapi_connectivity.PostOperationsGetConnectionDetailsHandlerFunc(func(params tapi_connectivity.PostOperationsGetConnectionDetailsParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_connectivity.PostOperationsGetConnectionDetails has not yet been implemented")
	})
	api.TapiConnectivityPostOperationsGetConnectivityServiceDetailsHandler = tapi_connectivity.PostOperationsGetConnectivityServiceDetailsHandlerFunc(func(params tapi_connectivity.PostOperationsGetConnectivityServiceDetailsParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_connectivity.PostOperationsGetConnectivityServiceDetails has not yet been implemented")
	})
	api.TapiConnectivityPostOperationsGetConnectivityServiceListHandler = tapi_connectivity.PostOperationsGetConnectivityServiceListHandlerFunc(func(params tapi_connectivity.PostOperationsGetConnectivityServiceListParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_connectivity.PostOperationsGetConnectivityServiceList has not yet been implemented")
	})
	api.TapiTopologyPostOperationsGetLinkDetailsHandler = tapi_topology.PostOperationsGetLinkDetailsHandlerFunc(func(params tapi_topology.PostOperationsGetLinkDetailsParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_topology.PostOperationsGetLinkDetails has not yet been implemented")
	})
	api.TapiTopologyPostOperationsGetNodeDetailsHandler = tapi_topology.PostOperationsGetNodeDetailsHandlerFunc(func(params tapi_topology.PostOperationsGetNodeDetailsParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_topology.PostOperationsGetNodeDetails has not yet been implemented")
	})
	api.TapiTopologyPostOperationsGetNodeEdgePointDetailsHandler = tapi_topology.PostOperationsGetNodeEdgePointDetailsHandlerFunc(func(params tapi_topology.PostOperationsGetNodeEdgePointDetailsParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_topology.PostOperationsGetNodeEdgePointDetails has not yet been implemented")
	})
	api.TapiCommonPostOperationsGetServiceInterfacePointDetailsHandler = tapi_common.PostOperationsGetServiceInterfacePointDetailsHandlerFunc(func(params tapi_common.PostOperationsGetServiceInterfacePointDetailsParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.PostOperationsGetServiceInterfacePointDetails has not yet been implemented")
	})
	api.TapiCommonPostOperationsGetServiceInterfacePointListHandler = tapi_common.PostOperationsGetServiceInterfacePointListHandlerFunc(func(params tapi_common.PostOperationsGetServiceInterfacePointListParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.PostOperationsGetServiceInterfacePointList has not yet been implemented")
	})
	api.TapiTopologyPostOperationsGetTopologyDetailsHandler = tapi_topology.PostOperationsGetTopologyDetailsHandlerFunc(func(params tapi_topology.PostOperationsGetTopologyDetailsParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_topology.PostOperationsGetTopologyDetails has not yet been implemented")
	})
	api.TapiTopologyPostOperationsGetTopologyListHandler = tapi_topology.PostOperationsGetTopologyListHandlerFunc(func(params tapi_topology.PostOperationsGetTopologyListParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_topology.PostOperationsGetTopologyList has not yet been implemented")
	})
	api.TapiPathComputationPostOperationsOptimizeP2PPathHandler = tapi_path_computation.PostOperationsOptimizeP2PPathHandlerFunc(func(params tapi_path_computation.PostOperationsOptimizeP2PPathParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_path_computation.PostOperationsOptimizeP2PPath has not yet been implemented")
	})
	api.TapiConnectivityPostOperationsUpdateConnectivityServiceHandler = tapi_connectivity.PostOperationsUpdateConnectivityServiceHandlerFunc(func(params tapi_connectivity.PostOperationsUpdateConnectivityServiceParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_connectivity.PostOperationsUpdateConnectivityService has not yet been implemented")
	})
	api.TapiCommonPostOperationsUpdateServiceInterfacePointHandler = tapi_common.PostOperationsUpdateServiceInterfacePointHandlerFunc(func(params tapi_common.PostOperationsUpdateServiceInterfacePointParams) middleware.Responder {
		return middleware.NotImplemented("operation tapi_common.PostOperationsUpdateServiceInterfacePoint has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
