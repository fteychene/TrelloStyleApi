package rest

import (
	"regexp"
	"strings"
	"log"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type Handler func(parameters map[string][]string) (int, interface{})


// TODO : Improve management of route by creating interface equals(method, route) 
// and getHandler to use polymorphism Route -> [SimpleRoute, ParametrizedRoute]
type Route struct {
	method          string
	path            string
	handler         Handler
	isParameterized bool
	parametersNames []string
}

type RouteMatcher struct {
	routes []*Route
}

/*
 * Register a Route in the RouteMatche and compite all the inforamtion needed
 */
func (matcher *RouteMatcher) AddRoute(method string, path string, handler Handler) {
	urlParamPattern, _ := regexp.Compile("/(:[^/]+)/?")
	route := new(Route)
	route.method = method
	route.path = path
	route.handler = handler
	if urlParamPattern.MatchString(path) {
		// Case of a parametrized path
		route.isParameterized = true
		// Computed parameters name
		parametersName := urlParamPattern.FindAllStringSubmatch(route.path, -1)
		route.parametersNames = make([]string, 0)
		for i := 0; i < len(parametersName); i++ {
			route.parametersNames = append(route.parametersNames, parametersName[i][1])
		}
		// Compute path as pattern
		route.path = "^"+urlParamPattern.ReplaceAllStringFunc(route.path, func(matched string) string {
			if strings.HasSuffix(matched, "/") {
				return "/([^/]+)/"
			} else {
				return "/([^/]+)"
			}
		})+"$"
		log.Printf("Adding parametrized route -> %s %s %s\n", route.method, route.path, route.parametersNames)
	} else {
		// Case of simple path
		route.isParameterized = false
		log.Printf("Adding route -> %s %s \n", route.method, route.path)
	}
	matcher.routes = append(matcher.routes, route)
	
}

/*
 * Check if a Route with url parameters is matching a given path and construct the url datas if true
 */
func searchWithParametrizedRoute(path string, route *Route) (bool, map[string][]string) {

	routePattern, _ := regexp.Compile(route.path)
	if !routePattern.MatchString(path) {
		return false, make(map[string][]string)
	}
	
	paramResult := make(map[string][]string)
	parameters := routePattern.FindAllStringSubmatch(path, -1)
	for i := 0; i < len(route.parametersNames); i++ {
		
		paramResult[strings.TrimPrefix(route.parametersNames[i], ":")] = []string{parameters[0][i+1]}
	}
	return true, paramResult
}

/*
 * Check if a Route is matching a given method and path
 * TODO : Should be in a method of Route to refactor
 */
func routeEquals(method string, path string, route *Route) (bool, map[string][]string) {
	if route.method != method {
		return false,  make(map[string][]string)
	}
	if !route.isParameterized {
		return path == route.path, make(map[string][]string)
	}
	return searchWithParametrizedRoute(path, route)
}

/*
 * Function to search in the configured handlers if one exist matching the pethod and the path received
 */
func (matcher *RouteMatcher) getHandlerAndUrlDatas(method string, path string) (*Route, map[string][]string) {

	for _, route := range matcher.routes {
		match, computedDatas := routeEquals(method, path, route)
		if match {
			return route, computedDatas
		}
	}
	return nil, make(map[string][]string)
}

/*
 * Constructor RouteMatcher
 */
func NewRouteMatcher() *RouteMatcher {
	return &RouteMatcher{routes: make([]*Route, 0)}
}
