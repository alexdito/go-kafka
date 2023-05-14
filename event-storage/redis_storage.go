package eventStorage

import (
	"KafkaProject/model"
	"context"
	"encoding/json"
	"fmt"
)

var ctx = context.Background()

func Set(key string, value []byte) (model.TransactionList, model.Balance) {
	rdb := GetClient()

	err := rdb.SAdd(ctx, key, value).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.SMembers(ctx, key).Result()

	if err != nil {
		panic(err)
	}

	var transactionList model.TransactionList
	var balance model.Balance

	for _, v := range val {
		var transaction model.Transaction
		err = json.Unmarshal([]byte(v), &transaction)
		transactionList = append(transactionList, transaction)
		balance.User = transaction.User

		switch transaction.Operation {
		case "withdrawal":
			balance.Amount -= transaction.Amount
			break
		case "additional":
			balance.Amount += transaction.Amount
			break
		}

		fmt.Println(transaction)
	}

	fmt.Println(balance)
	fmt.Println(transactionList)
	return transactionList, balance
}
