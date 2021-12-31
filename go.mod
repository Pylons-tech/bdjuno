module github.com/pylons-tech/bdjuno

go 1.15

require (
	filippo.io/edwards25519 v1.0.0-rc.1 // indirect
	github.com/ChainSafe/go-schnorrkel v1.0.0 // indirect
	github.com/Workiva/go-datastructures v1.0.53 // indirect
	github.com/cosmos/cosmos-sdk v0.44.5
	github.com/cosmos/ibc-go v1.2.5
	github.com/danieljoos/wincred v1.1.2 // indirect
	github.com/desmos-labs/juno v0.0.0-20211005090705-514187767199
	github.com/dvsekhvalnov/jose2go v1.5.0 // indirect
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/go-co-op/gocron v1.11.0
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/gogo/protobuf v1.3.3
	github.com/hdevalence/ed25519consensus v0.0.0-20210430192048-0962ce16b305 // indirect
	github.com/jmoiron/sqlx v1.3.4
	github.com/keybase/go-keychain v0.0.0-20211119201326-e02f34051621 // indirect
	github.com/lib/pq v1.10.4
	github.com/mimoo/StrobeGo v0.0.0-20210601165009-122bf33a46e0 // indirect
	github.com/pelletier/go-toml v1.9.4
	github.com/petermattis/goid v0.0.0-20211229010228-4d14c490ee36 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/proullon/ramsql v0.0.0-20211120092837-c8d0a408b939
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/rs/cors v1.8.2 // indirect
	github.com/rs/zerolog v1.26.1
	github.com/sasha-s/go-deadlock v0.3.1 // indirect
	github.com/spf13/afero v1.7.1 // indirect
	github.com/spf13/cobra v1.3.0
	github.com/spf13/viper v1.10.1 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.14
	github.com/tendermint/tm-db v0.6.6 // indirect
	github.com/ziutek/mymysql v1.5.4 // indirect
	golang.org/x/net v0.0.0-20211216030914-fe4d6282115f // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	google.golang.org/genproto v0.0.0-20211223182754-3ac035c7e7cb // indirect
	google.golang.org/grpc v1.43.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/tendermint/tendermint => github.com/forbole/tendermint v0.34.13-0.20211005074050-33591287eca5
