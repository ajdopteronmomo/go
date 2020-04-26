package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/ff8dd216d182458ab681895ebbdca51f")
	if err != nil {
		log.Fatal(err)
	}
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	//账户余额
	//读取一个账户的余额相当简单。调用客户端的BalanceAt方法，给它传递账户地址和可选的区块号。将区块号设置为nil将返回最新的余额。
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	//传区块号能让您读取该区块时的账户余额。区块号必须是big.Int类型。
	blockNumber := big.NewInt(5532993)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	fmt.Println(balanceAt)

	//以太坊中的数字是使用尽可能小的单位来处理的，因为它们是定点精度，在ETH中它是wei。要读取ETH值，您必须做计算wei/10^18。
	//因为我们正在处理大数，我们得导入原生的Gomath和math/big包。这是您做的转换。
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)

	//待处理的余额
	//有时您想知道待处理的账户余额是多少，例如，在提交或等待交易确认后。客户端提供了类似BalanceAt的方法，名为PendingBalanceAt，它接收账户地址作为参数。
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance)
}
