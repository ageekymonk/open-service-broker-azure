package fake

import "github.com/Azure/azure-service-broker/pkg/service"

// ProvisioningParameters represents parameters specific to provisioning a
// service using the fake service module. Note that, ordinarily, service module-
// specific types such as this do not need to be exported. An exception is made
// here because the fake service module is used to facilitate testing of the
// broker framework itself.
type ProvisioningParameters struct {
	SomeParameter string `json:"someParameter"`
}

// ProvisioningContext represents context collected and modified over the course
// of the fake service module's provisioning and deprovisioning processes. Note
// that, ordinarily, service module-specific types such as this do not need to
// be exported. An exception is made here because the fake service module is
// used to facilitate testing of the broker framework itself.
type ProvisioningContext struct {
	ResourceGroupName string `json:"resourceGroup"`
}

// UpdatingParameters represents parameters specific to binding to a service
// instance using the fake service module. Note that, ordinarily, service
// module-specific types such as this do not need to be exported. An exception
// is made here because the fake service module is used to facilitate testing of
// the broker framework itself.
type UpdatingParameters struct {
	SomeParameter string `json:"someParameter"`
}

// BindingParameters represents parameters specific to binding to a service
// instance using the fake service module. Note that, ordinarily, service
// module-specific types such as this do not need to be exported. An exception
// is made here because the fake service module is used to facilitate testing of
// the broker framework itself.
type BindingParameters struct {
	SomeParameter string `json:"someParameter"`
}

// BindingContext represents context collected and modified over the course
// of the fake service module's binding and unbinding processes. Note that,
// ordinarily, service module-specific types such as this do not need to be
// exported. An exception is made here because the fake service module is used
// to facilitate testing of the broker framework itself.
type BindingContext struct {
}

// Credentials generally represent credentials AND/OR ANY OTHER DETAILS (e.g.
// URLs, port numbers, etc.) that will be conveyed back to the client upon
// successful completion of a bind. In the specific case of the fake service
// module, which doesn't do much of anything (other than generate a messageID
// and write some logs), there are no important details to convey back to the
// client. The messageID is included just to provide an example of HOW details
// such as these can be conveyed to the client. Note that, ordinarily, service
// module-specific types such as this do not need to be exported. An exception
// is made here because the fake service module is used to facilitate testing of
// the broker framework itself.
type Credentials struct {
}

// GetEmptyProvisioningParameters returns an empty instance of module-specific
// provisioningParameters
func (
	m *Module,
) GetEmptyProvisioningParameters() service.ProvisioningParameters {
	return &ProvisioningParameters{}
}

// SetResourceGroup sets the name of the resource group into which service
// instances will be deployed
func (p *ProvisioningParameters) SetResourceGroup(resourceGroup string) {}

// GetEmptyProvisioningContext returns an empty instance of a module-specific
// provisioningContext
func (m *Module) GetEmptyProvisioningContext() service.ProvisioningContext {
	return &ProvisioningContext{}
}

// GetEmptyUpdatingParameters returns an empty instance of module-specific
// updatingParameters
func (
	m *Module,
) GetEmptyUpdatingParameters() service.UpdatingParameters {
	return &UpdatingParameters{}
}

// GetEmptyBindingParameters returns an empty instance of module-specific
// bindingParameters
func (m *Module) GetEmptyBindingParameters() service.BindingParameters {
	return &BindingParameters{}
}

// GetEmptyBindingContext returns an empty instance of a module-specific
// bindingContext
func (m *Module) GetEmptyBindingContext() service.BindingContext {
	return &BindingContext{}
}

// GetEmptyCredentials returns an empty instance of module-specific
// credentials
func (m *Module) GetEmptyCredentials() service.Credentials {
	return &Credentials{}
}
