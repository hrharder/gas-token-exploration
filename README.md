# Gas Token Exploration

This repo provides a simple script to view the cost of buying/selling a single unit of [`GasToken2`](https://etherscan.io/address/gst2.gastokenio.eth) on Uniswap, and the cost to mint the same amount at current market gas prices.

It simply logs the results to the console (see below).

**This is an early exploration/demonstration that will be improved upon. Not meant for usage outside of experimentation**

## Usage

Set the required environment variables:
- `ETHEREUM_JSONRPC_URL` - URL of an Ethereum JSONRPC provider
- `GAS_TOKEN_ADDRESS` - Address of the gas token to use (defaults to mainnet address of GST2)
- `UNISWAP_FACTORY_ADDRESS` - Address of a Uniswap factory contract (defaults to mainnet Uniswap factory)
- `TEST_AMOUNT` - Number of base units of gas token to test with (default is `100` which is 1 gas token)

Run the script:
```
go run main.go
```

View the results:
```
UNISWAP PRICES for 100 token(s)
- price to buy 1 gas tokens on uniswap (in wei):        3341219422754341
- price to sell 1 gas tokens on uniswap (in wei):       3316452665934509

MINT PRICES for 100 token(s) direct from contract
- buy gas token by minting (in wei @ 1 gwei):           3702884000000000
- buy gas token by minting (in wei @ standard):         3702884000000000
- buy gas token by minting (in wei @ fast):             29623072000000000
```