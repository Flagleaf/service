//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package starter

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &application_{}
		},
	})
	applicationStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &Application{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(applicationStructDescriptor)
}

type application_ struct {
	Run_ func()
}

func (a *application_) Run() {
	a.Run_()
}

type ApplicationIOCInterface interface {
	Run()
}

var _applicationSDID string

func GetApplicationSingleton() (*Application, error) {
	if _applicationSDID == "" {
		_applicationSDID = util.GetSDIDByStructPtr(new(Application))
	}
	i, err := singleton.GetImpl(_applicationSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*Application)
	return impl, nil
}

func GetApplicationIOCInterfaceSingleton() (ApplicationIOCInterface, error) {
	if _applicationSDID == "" {
		_applicationSDID = util.GetSDIDByStructPtr(new(Application))
	}
	i, err := singleton.GetImplWithProxy(_applicationSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(ApplicationIOCInterface)
	return impl, nil
}

type ThisApplication struct {
}

func (t *ThisApplication) This() ApplicationIOCInterface {
	thisPtr, _ := GetApplicationIOCInterfaceSingleton()
	return thisPtr
}