module github.com/pendulum-labs/market

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.44.5
	github.com/cosmos/ibc-go/v2 v2.0.2
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/regen-network/cosmos-proto v0.3.1 // indirect
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.2.1
	github.com/stretchr/testify v1.8.1
	github.com/tendermint/starport v0.19.2
	github.com/tendermint/tendermint v0.34.14
	github.com/tendermint/tm-db v0.6.4
	golang.org/x/net v0.10.0 // indirect
	google.golang.org/genproto v0.0.0-20230526203410-71b5a4ffd15e // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230530153820-e85fd2cbaebc
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/grpc v1.55.0
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
