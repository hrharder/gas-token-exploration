package watcher

import (
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/hrharder/gas-token-exploration/contracts"
	"github.com/hrharder/go-gas"
)

type WatcherConfig struct {
	FromAddress common.Address       // address signing/sending transactions (required)
	Signer      bind.SignerFn        // function to use to sign transaction (required)
	Backend     bind.ContractBackend // ethereum rpc backend for interacting with contracts (required)

	GasToken       common.Address // address of the gas token contract for the current network (required)
	UniswapFactory common.Address // address of the uniswap factory for the current network (required)
	SwapDeadline   time.Duration  // duration to use for deadline on uniswap swaps (required)

	GasPrice         *big.Int        // gas price to use for transactions (nil = fast estimate)
	GasPriceTimeout  time.Duration   // maximum age of a gas price estimate (if not using pre-defined price)
	GasPricePriority gas.GasPriority // gas price priority to use for price estimates (not necessary if setting gas price)
}

type Watcher struct {
	sync.Mutex

	// ok to change these by direct property access; all other require .Configure
	GasPrice *big.Int
	From     common.Address

	signer  bind.SignerFn
	backend bind.ContractBackend

	gasTokenContract *contracts.GasToken2
	gasTokenExchange *contracts.UniswapExchange
	swapDeadline     time.Duration

	gasPriceSuggester gas.GasPriceSuggester
	gasPricePriority  gas.GasPriority
}

func New(cfg WatcherConfig) (*Watcher, error) {
	watcher := new(Watcher)
	return watcher, watcher.Configure(cfg)
}

// Configure configures or re-configures the watcher with new options
func (w *Watcher) Configure(cfg WatcherConfig) error {
	w.Lock()
	defer w.Unlock()

	factory, err := contracts.NewUniswapFactoryCaller(cfg.UniswapFactory, cfg.Backend)
	if err != nil {
		return err
	}

	exchangeAddress, err := factory.GetExchange(nil, cfg.GasToken)
	if err != nil {
		return err
	}

	exchange, err := contracts.NewUniswapExchange(exchangeAddress, cfg.Backend)
	if err != nil {
		return err
	}

	token, err := contracts.NewGasToken2(cfg.GasToken, cfg.Backend)
	if err != nil {
		return err
	}

	suggester, err := gas.NewGasPriceSuggester(cfg.GasPriceTimeout)
	if err != nil {
		return err
	}

	w.backend = cfg.Backend
	w.gasTokenContract = token
	w.gasTokenExchange = exchange
	w.gasPriceSuggester = suggester
	w.gasPricePriority = cfg.GasPricePriority
	w.swapDeadline = cfg.SwapDeadline
	return nil
}

// TransactOpts generates new transaction options for a transaction
func (w *Watcher) TransactOpts() (*bind.TransactOpts, error) {
	gasPrice, err := w.SuggestGasPrice()
	if err != nil {
		return nil, err
	}

	return &bind.TransactOpts{
		From:     w.From,
		Signer:   w.signer,
		GasPrice: gasPrice,
	}, nil
}

// MintGasTokens calls `Mint` on the gas token contract to create GST amount of tokens
func (w *Watcher) MintGasTokens(amount *big.Int) (*types.Transaction, error) {
	opts, err := w.TransactOpts()
	if err != nil {
		return nil, err
	}

	return w.gasTokenContract.Mint(opts, amount)
}

// BuyGasTokens buys amount of GST from uniswap for ETH at the current rate
func (w *Watcher) BuyGasTokens(amount *big.Int) (*types.Transaction, error) {
	opts, err := w.TransactOpts()
	if err != nil {
		return nil, err
	}

	estimate, err := w.gasTokenExchange.GetEthToTokenOutputPrice(nil, amount)
	if err != nil {
		return nil, err
	}

	deadline := big.NewInt(time.Now().Add(w.swapDeadline).Unix())
	opts.Value = new(big.Int).Set(estimate)
	return w.gasTokenExchange.EthToTokenSwapOutput(opts, amount, deadline)
}

// SellGasTokens sells amount of GST on uniswap for ETH at the current rate
func (w *Watcher) SellGasTokens(amount *big.Int) (*types.Transaction, error) {
	opts, err := w.TransactOpts()
	if err != nil {
		return nil, err
	}

	estimate, err := w.gasTokenExchange.GetTokenToEthInputPrice(nil, amount)
	if err != nil {
		return nil, err
	}

	deadline := big.NewInt(time.Now().Add(w.swapDeadline).Unix())
	return w.gasTokenExchange.TokenToEthSwapInput(opts, amount, estimate, deadline)
}

// SuggestGasPrice fetches a gas price estimate for a transacion either from a user-defined price or an estimate
// from ETH Gas Station with a user-defined priority
func (w *Watcher) SuggestGasPrice() (*big.Int, error) {
	if w.GasPrice != nil {
		return w.GasPrice, nil
	}
	return w.gasPriceSuggester(w.gasPricePriority)
}
