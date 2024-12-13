module github.com/deviavir/bitcoind-healthcheck

go 1.21

// `go get github.com/deviavir/btcd@fffd8b2def3be6238c9bdfcf9bebdde1c41de312`
// to get and replace using a specific commit version
replace github.com/btcsuite/btcd => github.com/deviavir/btcd v0.0.0-20241213141921-f9d94bc301ba

require (
	github.com/btcsuite/btcd v0.24.2
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/btcsuite/btcd/btcec/v2 v2.3.4 // indirect
	github.com/btcsuite/btcd/btcutil v1.1.5 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.1.0 // indirect
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f // indirect
	github.com/btcsuite/go-socks v0.0.0-20170105172521-4720035b7bfd // indirect
	github.com/btcsuite/websocket v0.0.0-20150119174127-31079b680792 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/decred/dcrd/crypto/blake256 v1.0.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
