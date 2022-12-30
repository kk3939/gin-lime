package controllers

import "fmt"

func fmtErrMsg(err error) string {
	return fmt.Sprintf("%v", err)
}
