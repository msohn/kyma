// Code generated by mockery v1.0.0
package automock

import chart "k8s.io/helm/pkg/proto/hapi/chart"
import helm "k8s.io/helm/pkg/helm"

import mock "github.com/stretchr/testify/mock"
import services "k8s.io/helm/pkg/proto/hapi/services"

// HelmDeleteInstaller is an autogenerated mock type for the HelmDeleteInstaller type
type HelmDeleteInstaller struct {
	mock.Mock
}

// DeleteRelease provides a mock function with given fields: rlsName, opts
func (_m *HelmDeleteInstaller) DeleteRelease(rlsName string, opts ...helm.DeleteOption) (*services.UninstallReleaseResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, rlsName)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *services.UninstallReleaseResponse
	if rf, ok := ret.Get(0).(func(string, ...helm.DeleteOption) *services.UninstallReleaseResponse); ok {
		r0 = rf(rlsName, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*services.UninstallReleaseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...helm.DeleteOption) error); ok {
		r1 = rf(rlsName, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InstallReleaseFromChart provides a mock function with given fields: _a0, ns, opts
func (_m *HelmDeleteInstaller) InstallReleaseFromChart(_a0 *chart.Chart, ns string, opts ...helm.InstallOption) (*services.InstallReleaseResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, ns)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *services.InstallReleaseResponse
	if rf, ok := ret.Get(0).(func(*chart.Chart, string, ...helm.InstallOption) *services.InstallReleaseResponse); ok {
		r0 = rf(_a0, ns, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*services.InstallReleaseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*chart.Chart, string, ...helm.InstallOption) error); ok {
		r1 = rf(_a0, ns, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}