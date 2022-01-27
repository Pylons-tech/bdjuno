module github.com/pylons-tech/bdjuno

go 1.16

require (
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/cosmos/btcutil v1.0.4 // indirect
	github.com/cosmos/cosmos-sdk v0.44.1
	github.com/cosmos/gaia/v6 v6.0.0-rc1
	github.com/cosmos/iavl v0.17.3 // indirect
	github.com/forbole/juno/v2 v2.0.0-20211215134503-2c58dc73913b
	github.com/gin-gonic/gin v1.7.0 // indirect
	github.com/go-co-op/gocron v1.11.0
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/google/cel-go v0.9.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jmoiron/sqlx v1.2.1-0.20200324155115-ee514944af4b
	github.com/lib/pq v1.10.4
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/opencontainers/runc v1.0.3 // indirect
	github.com/pelletier/go-toml v1.9.4
	github.com/proullon/ramsql v0.0.0-20181213202341-817cee58a244
	github.com/rogpeppe/go-internal v1.6.2
	github.com/rs/zerolog v1.26.0
	github.com/spf13/cobra v1.3.0
	github.com/spf13/viper v1.10.1 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.14
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa
	google.golang.org/grpc v1.43.0
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/tendermint/tendermint => github.com/forbole/tendermint v0.34.13-0.20210820072129-a2a4af55563d
