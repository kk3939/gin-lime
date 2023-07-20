package controllers

import "fmt"

func FmtErrMsg(err error) string {
	return fmt.Sprintf("%v", err)
}
