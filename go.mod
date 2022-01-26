module cj_converter

go 1.17

replace github.com/adelberteng/cj_converter/converter => ./converter

require (
	github.com/adelberteng/cj_converter/converter v0.0.0-00010101000000-000000000000
	github.com/adelberteng/go_logger v0.0.0-20220125054226-5ec1a7ca43ad
	gopkg.in/ini.v1 v1.66.3
)

require github.com/stretchr/testify v1.7.0 // indirect
