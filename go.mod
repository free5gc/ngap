module github.com/free5gc/ngap

go 1.21

require (
	github.com/free5gc/aper v1.0.6-0.20240503143507-2c4c4780b98f
	github.com/free5gc/openapi v1.0.9-0.20240730084323-449098e08462
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.8.4
	github.com/tim-ywliu/nested-logrus-formatter v1.3.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/free5gc/openapi => github.com/yccodr/openapi v1.0.9-0.20240801032959-f8c907cee3a4
