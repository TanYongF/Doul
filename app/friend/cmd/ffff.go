package main

import (
	"fmt"
	"github.com/pkg/errors"
	"go_code/Doul/common/xerr"
)

func main() {

	err := t2()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	//fmt.Print(err)
	return
}

func t2() error {
	err := t1()
	if err != nil {
		return err
	}
	return nil
}

func t1() error {
	return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "fafaf")
}
