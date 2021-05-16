package cas

import (
	"ddl-bot/pkg/cas/assets"
	"errors"
	"fmt"
	"github.com/dop251/goja"
)

func GetRSAValue(userName string, password string, ltValue string) (string, error) {
	vm := goja.New()
	_, err := vm.RunString(string(assets.DES_JS))
	if err != nil {
		panic(err)
	}
	var inputStr = userName + password + ltValue
	strEnc, valid := goja.AssertFunction(vm.Get("strEnc"))
	if !valid {
		return "", errors.New("invalid js")
	}
	value, err := strEnc(nil, vm.ToValue(inputStr), vm.ToValue("1"), vm.ToValue("2"), vm.ToValue("3"))
	if err != nil {
		panic(err)
	}
	var result = value.String()
	fmt.Println(result)
	return result, nil
}
