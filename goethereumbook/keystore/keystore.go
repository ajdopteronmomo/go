package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	//createKs()
	importKs()
}

//这是一个完整的生成新的keystore账户的示例
func createKs() {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex())
}

//现在要导入您的keystore，您基本上像往常一样再次调用NewKeyStore，然后调用Import方法，该方法接收keystore的JSON数据作为字节。第二个参数是用于加密私钥的口令。
//第三个参数是指定一个新的加密口令，但我们在示例中使用一样的口令。导入账户将允许您按期访问该账户，但它将生成新keystore文件！有两个相同的事物是没有意义的，所以我们将删除旧的。
//这是一个导入keystore和访问账户的示例。
func importKs() {
	file := "./wallets/UTC--2020-04-24T08-19-03.284122700Z--778c589d4d5a8981f1983de4ba17666af3b933a4"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)

	fmt.Println(account.Address.Hex())

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}
