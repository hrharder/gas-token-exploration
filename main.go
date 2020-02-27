package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/hrharder/go-gas"

	"github.com/caarlos0/env/v6"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hrharder/gas-token-exploration/contracts"
)

var NULL_ADDRESS = common.Address{}

type config struct {
	EthereumURL           string `env:"ETHEREUM_JSONRPC_URL" envDefault:"https://mainnet.infura.io"`
	GasTokenAddress       string `env:"GAS_TOKEN_ADDRESS" envDefault:"0x0000000000b3F879cb30FE243b4Dfee438691c04"`
	UniswapFactoryAddress string `env:"UNISWAP_FACTORY_ADDRESS" envDefault:"0xc0a47dFe034B400B47bDaD5FecDa2621de6c4d95"`
	TestAmount            int64  `env:"TEST_AMOUNT", envDefault:"100"`

	gasToken       common.Address
	uniswapFactory common.Address
}

func (cfg *config) parse() error {
	if err := env.Parse(cfg); err != nil {
		return err
	}

	cfg.gasToken = common.HexToAddress(cfg.GasTokenAddress)
	cfg.uniswapFactory = common.HexToAddress(cfg.UniswapFactoryAddress)

	return nil
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	cfg := new(config)
	check(cfg.parse())

	client, err := ethclient.Dial(cfg.EthereumURL)
	check(err)

	factory, err := contracts.NewUniswapFactoryCaller(cfg.uniswapFactory, client)
	check(err)

	exchangeAddress, err := factory.GetExchange(nil, cfg.gasToken)
	check(err)

	if exchangeAddress == NULL_ADDRESS {
		log.Fatal("no exchange for token address")
	}

	// fetch buy and sell prices for GST2 on uniswap
	exchange, err := contracts.NewUniswapExchangeCaller(exchangeAddress, client)
	check(err)
	tradeAmount := big.NewInt(cfg.TestAmount)
	buyPrice, err := exchange.GetEthToTokenOutputPrice(nil, tradeAmount)
	check(err)
	sellPrice, err := exchange.GetTokenToEthInputPrice(nil, tradeAmount)
	check(err)
	fmt.Printf("UNISWAP PRICES for %v token(s)\n", cfg.TestAmount)
	fmt.Printf("- price to buy gas tokens on uniswap (in wei):\t\t%s\n", buyPrice.String())
	fmt.Printf("- price to sell gas tokens on uniswap (in wei):\t\t%s\n", sellPrice.String())
	fmt.Println()

	// estimate gas cost of minting 1 gas token (100 base units)
	gasTokenABI, err := abi.JSON(strings.NewReader(contracts.GasToken2ABI))
	check(err)
	callData, err := gasTokenABI.Pack("mint", big.NewInt(cfg.TestAmount))
	check(err)
	gasCost, err := client.EstimateGas(context.Background(), ethereum.CallMsg{From: NULL_ADDRESS, To: &cfg.gasToken, Data: callData})
	check(err)

	// fetch gas prices to generate real-world ETH costs of minting GST2
	gasCostBig := big.NewInt(int64(gasCost))
	gasPriceFast, err := gas.SuggestGasPrice(gas.GasPriorityFast)
	check(err)
	gasPriceStandard, err := gas.SuggestGasPrice(gas.GasPriorityAverage)
	check(err)
	fmt.Printf("MINT PRICES for %v token(s) direct from contract\n", cfg.TestAmount)
	fmt.Printf("- buy gas token by minting (in wei @ 1 gwei):\t\t%v\n", new(big.Int).Mul(gasCostBig, big.NewInt(1000000000)))
	fmt.Printf("- buy gas token by minting (in wei @ standard):\t\t%v\n", new(big.Int).Mul(gasCostBig, gasPriceStandard))
	fmt.Printf("- buy gas token by minting (in wei @ fast):\t\t%v\n", new(big.Int).Mul(gasCostBig, gasPriceFast))
}
