/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package javascript

import (
	"github.com/robertkrimen/otto"
)

func (p *JavascriptPlugin) storageUse(call otto.FunctionCall) otto.Value {
	table, _ := call.Argument(0).ToString()

	err := p.c.Storage.Use(table)
	if err != nil {
		return otto.FalseValue()
	}

	return otto.TrueValue()
}

func (p *JavascriptPlugin) storagePut(call otto.FunctionCall) otto.Value {
	key, _ := call.Argument(0).ToString()
	val, _ := call.Argument(0).ToString()

	err := p.c.Storage.Put(key, val)
	if err != nil {
		return otto.FalseValue()
	}

	return otto.TrueValue()
}

func (p *JavascriptPlugin) storageGet(call otto.FunctionCall) otto.Value {
	key, _ := call.Argument(0).ToString()

	val, err := p.c.Storage.Get(key)
	if err != nil {
		return otto.NullValue()
	}

	value, err := otto.ToValue(val)
	if err != nil {
		return otto.NullValue()
	}

	return value
}

func (p *JavascriptPlugin) storageRemove(call otto.FunctionCall) otto.Value {
	key, _ := call.Argument(0).ToString()

	err := p.c.Storage.Remove(key)
	if err != nil {
		return otto.FalseValue()
	}

	return otto.TrueValue()
}
