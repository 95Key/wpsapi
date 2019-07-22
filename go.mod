module github.com/95key/wpsapi

go 1.12

replace (
	golang.org/x/net => github.com/golang/net v0.0.0-20190301231341-16b79f2e4e95
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190309122539-980fc434d28e

)

require (
	github.com/95key/util v0.0.0-20190722071218-84ce68825d2e
	github.com/patrickmn/go-cache v2.1.0+incompatible
)
