package main

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"AttributeList",
		"GET",
		"/attributes",
		AttributeList,
	},
	Route{
		"AttributeDetail",
		"GET",
		"/attributes/{attributeId}",
		AttributeDetail,
	},
}
