/*
 * Simple MEC Discovery API
 *
 * # Find your nearest MEC platform --- Network operators will typically have multiple MEC sites in a given territory. Connecting your application to a server on the closest MEC platform means the lowest latency - however, the physical location of a user is not an accurate match to the closest MEC site, due to the way operator networks are configured. This API returns the MEC platforms with the _shortest network path_ to the client making the request, and hence the lowest propagation delay. * If you have a server instance deployed there, connect to it to gain the lowest latency * Or if not, you may wish to deploy an instance there using the APIs of the cloud provider supporting that zone.    This API is intended to be called by a client application hosted on a UE attached to the operator network. _Note that the API parameters have been listed in this 'simple API' to align with the full API, but are optional and may not be supported by the API server_ ---
 *
 * API version: 0.8.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}


var routes = Routes{
	Route{
		"endpoint for ClientTest",
		strings.ToUpper("Get"),
		"/clienttest1",
		Test1,
	},

	Route{
		"endpoint for ClientTest",
		strings.ToUpper("Get"),
		"/clienttest2",
		Test2,
	},


	Route{
		"endpoint for ClientTest",
		strings.ToUpper("Get"),
		"/testrequest/{value}",
		TestRequest,
	},
}
