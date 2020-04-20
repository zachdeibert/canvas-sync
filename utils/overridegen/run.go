package overridegen

import (
	"io/ioutil"
	"reflect"

	"github.com/zachdeibert/canvas-sync/utils/apisync"
)

// Run the program
func Run() {
	ctx := createContext()
	ctx.addSlice("models", reflect.TypeOf(&[]*apisync.Model{}))
	ctx.addSlice("methods", reflect.TypeOf(&[]apisync.MethodAPIPair{}))
	if err := ioutil.WriteFile("apisync/overrides/helpers.go", []byte(ctx.writeCode()), 0644); err != nil {
		panic(err)
	}
}
