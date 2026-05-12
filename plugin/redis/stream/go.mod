module github.com/ycxwi/micro/plugin/redis/stream/v3

go 1.18

require (
	github.com/go-redis/redis/v8 v8.11.4
	github.com/google/uuid v1.3.0
	github.com/ycxwi/micro/v3 v3.2.1
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
)

replace github.com/ycxwi/micro/v3 => ../../..
