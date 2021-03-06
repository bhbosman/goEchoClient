package components

import (
	"context"
	"github.com/bhbosman/gocomms/intf"
	"github.com/bhbosman/gologging"
)

type connectionReactorFactory struct {
	name string
}

func (self *connectionReactorFactory) Values(inputValues map[string]interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	return result, nil
}

func (self *connectionReactorFactory) Create(
	name string,
	cancelCtx context.Context,
	cancelFunc context.CancelFunc,
	logger *gologging.SubSystemLogger,
	userContext interface{}) intf.IConnectionReactor {
	result := newConnectionReactor(logger, cancelCtx, cancelFunc, name, userContext)
	return result
}

func (self *connectionReactorFactory) Name() string {
	return self.name
}
