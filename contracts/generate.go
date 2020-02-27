//go:generate abigen --abi ./abis/uniswap_exchange.json --type UniswapExchange --pkg contracts --out uniswap_exchange.go
//go:generate abigen --abi ./abis/uniswap_factory.json --type UniswapFactory --pkg contracts --out uniswap_factory.go
//go:generate abigen --abi ./abis/gas_token_2.json --type GasToken2 --pkg contracts --out gas_token_2.go
package contracts
