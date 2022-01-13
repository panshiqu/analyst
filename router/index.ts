import { ethers } from 'ethers';
import { Token, TradeType, CurrencyAmount } from '@uniswap/sdk-core';
import { AlphaRouter, V3QuoteProvider, UniswapMulticallProvider, ChainId, routeAmountsToString } from '@uniswap/smart-order-router';

const web3Provider = new ethers.providers.JsonRpcProvider('https://polygon-rpc.com');

const multicall2Provider = new UniswapMulticallProvider(ChainId.POLYGON, web3Provider, 375_000);

const v3QuoteProvider = new V3QuoteProvider(
  ChainId.POLYGON,
  web3Provider,
  multicall2Provider,
  {
    retries: 2,
    minTimeout: 100,
    maxTimeout: 1000,
  },
  {
    multicallChunk: 105,
    gasLimitPerCall: 352_500,
    quoteMinSuccessRate: 0.15,
  },
  {
    gasLimitOverride: 1_000_000,
    multicallChunk: 35,
  }
);

const router = new AlphaRouter({
  chainId: ChainId.POLYGON,
  provider: web3Provider,
  v3QuoteProvider: v3QuoteProvider,
});

const USDC = new Token(ChainId.POLYGON, '0x2791bca1f2de4661ed88a30c99a7a9449aa84174', 6, 'USDC', 'USD Coin');

require('log-timestamp');

async function asyncHandle(req, res) {
  res.writeHead(200, { 'Content-Type': 'application/json' });

  try {
    console.log('handle', req.url);

    const params = new URLSearchParams(req.url.replace(/^(\/)/, ''));

    const token = new Token(ChainId.POLYGON, params.get('address'), Number(params.get('decimals')));

    const currencyAmount = CurrencyAmount.fromRawAmount(token, params.get('amount'));

    let swap = await router.route(currencyAmount, USDC, TradeType.EXACT_INPUT);
    if (!swap) {
      throw Error('could not find route.');
    }

    let response = {
      quote: swap.quote.toFixed(4),
      quoteGasAdjusted: swap.quoteGasAdjusted.toFixed(4),
      estimatedGasUsed: swap.estimatedGasUsed.toString(),
      estimatedGasUsedQuoteToken: swap.estimatedGasUsedQuoteToken.toFixed(4),
      estimatedGasUsedUSD: swap.estimatedGasUsedUSD.toFixed(4),
      gasPriceWei: swap.gasPriceWei.toString(),
      blockNumber: swap.blockNumber.toString(),
      routeString: routeAmountsToString(swap.route),
    };

    console.log('response', response);

    res.end(JSON.stringify(response));
  } catch (err) {
    console.log('catch', err);

    res.end(JSON.stringify({
      message: err.message,
    }));
  }
}

require('http').createServer((req, res) => {
  asyncHandle(req, res);
}).listen(8000, '127.0.0.1');
