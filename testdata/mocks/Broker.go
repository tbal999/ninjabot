// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	model "github.com/tbal999/ninjabot/model"
	mock "github.com/stretchr/testify/mock"
)

// Broker is an autogenerated mock type for the Broker type
type Broker struct {
	mock.Mock
}

type Broker_Expecter struct {
	mock *mock.Mock
}

func (_m *Broker) EXPECT() *Broker_Expecter {
	return &Broker_Expecter{mock: &_m.Mock}
}

// Account provides a mock function with given fields:
func (_m *Broker) Account() (model.Account, error) {
	ret := _m.Called()

	var r0 model.Account
	if rf, ok := ret.Get(0).(func() model.Account); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Broker_Account_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Account'
type Broker_Account_Call struct {
	*mock.Call
}

// Account is a helper method to define mock.On call
func (_e *Broker_Expecter) Account() *Broker_Account_Call {
	return &Broker_Account_Call{Call: _e.mock.On("Account")}
}

func (_c *Broker_Account_Call) Run(run func()) *Broker_Account_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Broker_Account_Call) Return(_a0 model.Account, _a1 error) *Broker_Account_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Cancel provides a mock function with given fields: _a0
func (_m *Broker) Cancel(_a0 model.Order) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Order) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Broker_Cancel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Cancel'
type Broker_Cancel_Call struct {
	*mock.Call
}

// Cancel is a helper method to define mock.On call
//   - _a0 model.Order
func (_e *Broker_Expecter) Cancel(_a0 interface{}) *Broker_Cancel_Call {
	return &Broker_Cancel_Call{Call: _e.mock.On("Cancel", _a0)}
}

func (_c *Broker_Cancel_Call) Run(run func(_a0 model.Order)) *Broker_Cancel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.Order))
	})
	return _c
}

func (_c *Broker_Cancel_Call) Return(_a0 error) *Broker_Cancel_Call {
	_c.Call.Return(_a0)
	return _c
}

// CreateOrderLimit provides a mock function with given fields: side, pair, size, limit
func (_m *Broker) CreateOrderLimit(side model.SideType, pair string, size float64, limit float64) (model.Order, error) {
	ret := _m.Called(side, pair, size, limit)

	var r0 model.Order
	if rf, ok := ret.Get(0).(func(model.SideType, string, float64, float64) model.Order); ok {
		r0 = rf(side, pair, size, limit)
	} else {
		r0 = ret.Get(0).(model.Order)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.SideType, string, float64, float64) error); ok {
		r1 = rf(side, pair, size, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Broker_CreateOrderLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrderLimit'
type Broker_CreateOrderLimit_Call struct {
	*mock.Call
}

// CreateOrderLimit is a helper method to define mock.On call
//   - side model.SideType
//   - pair string
//   - size float64
//   - limit float64
func (_e *Broker_Expecter) CreateOrderLimit(side interface{}, pair interface{}, size interface{}, limit interface{}) *Broker_CreateOrderLimit_Call {
	return &Broker_CreateOrderLimit_Call{Call: _e.mock.On("CreateOrderLimit", side, pair, size, limit)}
}

func (_c *Broker_CreateOrderLimit_Call) Run(run func(side model.SideType, pair string, size float64, limit float64)) *Broker_CreateOrderLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.SideType), args[1].(string), args[2].(float64), args[3].(float64))
	})
	return _c
}

func (_c *Broker_CreateOrderLimit_Call) Return(_a0 model.Order, _a1 error) *Broker_CreateOrderLimit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// CreateOrderMarket provides a mock function with given fields: side, pair, size
func (_m *Broker) CreateOrderMarket(side model.SideType, pair string, size float64) (model.Order, error) {
	ret := _m.Called(side, pair, size)

	var r0 model.Order
	if rf, ok := ret.Get(0).(func(model.SideType, string, float64) model.Order); ok {
		r0 = rf(side, pair, size)
	} else {
		r0 = ret.Get(0).(model.Order)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.SideType, string, float64) error); ok {
		r1 = rf(side, pair, size)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Broker_CreateOrderMarket_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrderMarket'
type Broker_CreateOrderMarket_Call struct {
	*mock.Call
}

// CreateOrderMarket is a helper method to define mock.On call
//   - side model.SideType
//   - pair string
//   - size float64
func (_e *Broker_Expecter) CreateOrderMarket(side interface{}, pair interface{}, size interface{}) *Broker_CreateOrderMarket_Call {
	return &Broker_CreateOrderMarket_Call{Call: _e.mock.On("CreateOrderMarket", side, pair, size)}
}

func (_c *Broker_CreateOrderMarket_Call) Run(run func(side model.SideType, pair string, size float64)) *Broker_CreateOrderMarket_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.SideType), args[1].(string), args[2].(float64))
	})
	return _c
}

func (_c *Broker_CreateOrderMarket_Call) Return(_a0 model.Order, _a1 error) *Broker_CreateOrderMarket_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// CreateOrderMarketQuote provides a mock function with given fields: side, pair, quote
func (_m *Broker) CreateOrderMarketQuote(side model.SideType, pair string, quote float64) (model.Order, error) {
	ret := _m.Called(side, pair, quote)

	var r0 model.Order
	if rf, ok := ret.Get(0).(func(model.SideType, string, float64) model.Order); ok {
		r0 = rf(side, pair, quote)
	} else {
		r0 = ret.Get(0).(model.Order)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.SideType, string, float64) error); ok {
		r1 = rf(side, pair, quote)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Broker_CreateOrderMarketQuote_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrderMarketQuote'
type Broker_CreateOrderMarketQuote_Call struct {
	*mock.Call
}

// CreateOrderMarketQuote is a helper method to define mock.On call
//   - side model.SideType
//   - pair string
//   - quote float64
func (_e *Broker_Expecter) CreateOrderMarketQuote(side interface{}, pair interface{}, quote interface{}) *Broker_CreateOrderMarketQuote_Call {
	return &Broker_CreateOrderMarketQuote_Call{Call: _e.mock.On("CreateOrderMarketQuote", side, pair, quote)}
}

func (_c *Broker_CreateOrderMarketQuote_Call) Run(run func(side model.SideType, pair string, quote float64)) *Broker_CreateOrderMarketQuote_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.SideType), args[1].(string), args[2].(float64))
	})
	return _c
}

func (_c *Broker_CreateOrderMarketQuote_Call) Return(_a0 model.Order, _a1 error) *Broker_CreateOrderMarketQuote_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// CreateOrderOCO provides a mock function with given fields: side, pair, size, price, stop, stopLimit
func (_m *Broker) CreateOrderOCO(side model.SideType, pair string, size float64, price float64, stop float64, stopLimit float64) ([]model.Order, error) {
	ret := _m.Called(side, pair, size, price, stop, stopLimit)

	var r0 []model.Order
	if rf, ok := ret.Get(0).(func(model.SideType, string, float64, float64, float64, float64) []model.Order); ok {
		r0 = rf(side, pair, size, price, stop, stopLimit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.SideType, string, float64, float64, float64, float64) error); ok {
		r1 = rf(side, pair, size, price, stop, stopLimit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Broker_CreateOrderOCO_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrderOCO'
type Broker_CreateOrderOCO_Call struct {
	*mock.Call
}

// CreateOrderOCO is a helper method to define mock.On call
//   - side model.SideType
//   - pair string
//   - size float64
//   - price float64
//   - stop float64
//   - stopLimit float64
func (_e *Broker_Expecter) CreateOrderOCO(side interface{}, pair interface{}, size interface{}, price interface{}, stop interface{}, stopLimit interface{}) *Broker_CreateOrderOCO_Call {
	return &Broker_CreateOrderOCO_Call{Call: _e.mock.On("CreateOrderOCO", side, pair, size, price, stop, stopLimit)}
}

func (_c *Broker_CreateOrderOCO_Call) Run(run func(side model.SideType, pair string, size float64, price float64, stop float64, stopLimit float64)) *Broker_CreateOrderOCO_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.SideType), args[1].(string), args[2].(float64), args[3].(float64), args[4].(float64), args[5].(float64))
	})
	return _c
}

func (_c *Broker_CreateOrderOCO_Call) Return(_a0 []model.Order, _a1 error) *Broker_CreateOrderOCO_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// CreateOrderStop provides a mock function with given fields: pair, quantity, limit
func (_m *Broker) CreateOrderStop(pair string, quantity float64, limit float64) (model.Order, error) {
	ret := _m.Called(pair, quantity, limit)

	var r0 model.Order
	if rf, ok := ret.Get(0).(func(string, float64, float64) model.Order); ok {
		r0 = rf(pair, quantity, limit)
	} else {
		r0 = ret.Get(0).(model.Order)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, float64, float64) error); ok {
		r1 = rf(pair, quantity, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Broker_CreateOrderStop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrderStop'
type Broker_CreateOrderStop_Call struct {
	*mock.Call
}

// CreateOrderStop is a helper method to define mock.On call
//   - pair string
//   - quantity float64
//   - limit float64
func (_e *Broker_Expecter) CreateOrderStop(pair interface{}, quantity interface{}, limit interface{}) *Broker_CreateOrderStop_Call {
	return &Broker_CreateOrderStop_Call{Call: _e.mock.On("CreateOrderStop", pair, quantity, limit)}
}

func (_c *Broker_CreateOrderStop_Call) Run(run func(pair string, quantity float64, limit float64)) *Broker_CreateOrderStop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(float64), args[2].(float64))
	})
	return _c
}

func (_c *Broker_CreateOrderStop_Call) Return(_a0 model.Order, _a1 error) *Broker_CreateOrderStop_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Order provides a mock function with given fields: pair, id
func (_m *Broker) Order(pair string, id int64) (model.Order, error) {
	ret := _m.Called(pair, id)

	var r0 model.Order
	if rf, ok := ret.Get(0).(func(string, int64) model.Order); ok {
		r0 = rf(pair, id)
	} else {
		r0 = ret.Get(0).(model.Order)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int64) error); ok {
		r1 = rf(pair, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Broker_Order_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Order'
type Broker_Order_Call struct {
	*mock.Call
}

// Order is a helper method to define mock.On call
//   - pair string
//   - id int64
func (_e *Broker_Expecter) Order(pair interface{}, id interface{}) *Broker_Order_Call {
	return &Broker_Order_Call{Call: _e.mock.On("Order", pair, id)}
}

func (_c *Broker_Order_Call) Run(run func(pair string, id int64)) *Broker_Order_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int64))
	})
	return _c
}

func (_c *Broker_Order_Call) Return(_a0 model.Order, _a1 error) *Broker_Order_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Position provides a mock function with given fields: pair
func (_m *Broker) Position(pair string) (float64, float64, error) {
	ret := _m.Called(pair)

	var r0 float64
	if rf, ok := ret.Get(0).(func(string) float64); ok {
		r0 = rf(pair)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 float64
	if rf, ok := ret.Get(1).(func(string) float64); ok {
		r1 = rf(pair)
	} else {
		r1 = ret.Get(1).(float64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(pair)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Broker_Position_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Position'
type Broker_Position_Call struct {
	*mock.Call
}

// Position is a helper method to define mock.On call
//   - pair string
func (_e *Broker_Expecter) Position(pair interface{}) *Broker_Position_Call {
	return &Broker_Position_Call{Call: _e.mock.On("Position", pair)}
}

func (_c *Broker_Position_Call) Run(run func(pair string)) *Broker_Position_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Broker_Position_Call) Return(asset float64, quote float64, err error) *Broker_Position_Call {
	_c.Call.Return(asset, quote, err)
	return _c
}

type mockConstructorTestingTNewBroker interface {
	mock.TestingT
	Cleanup(func())
}

// NewBroker creates a new instance of Broker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBroker(t mockConstructorTestingTNewBroker) *Broker {
	mock := &Broker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
