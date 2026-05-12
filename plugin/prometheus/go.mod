module github.com/ycxwi/micro/plugin/prometheus/v3

go 1.18

require (
	github.com/ycxwi/micro/v3 v3.2.1
	github.com/prometheus/client_golang v1.12.1
	github.com/stretchr/testify v1.7.0
)

replace github.com/ycxwi/micro/v3 => ../..
