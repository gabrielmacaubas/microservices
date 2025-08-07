module github.com/gabrielmacaubas/microservices/order

go 1.24.3

require github.com/gabrielmacaubas/microservices-proto/golang/order v0.0.0-00010101000000-000000000000
replace github.com/gabrielmacaubas/microservices-proto/golang/order => ../../microservices-proto/golang/order
require github.com/gabrielmacaubas/microservices-proto/golang/payment v0.0.0-00010101000000-000000000000
replace github.com/gabrielmacaubas/microservices-proto/golang/payment => ../../microservices-proto/golang/payment