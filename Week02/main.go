package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

// 1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
// 需要抛给上层，因为需要将堆栈信息返回给上层，方便定位问题。

func main() {
	fmt.Printf("%+v\n", service())
}

func dao() error {
	// 假如出现查询错误
	return errors.Wrap(sql.ErrNoRows, "query failed")
}

func biz() error {
	return dao()
}

func service() error {
	return biz()
}
