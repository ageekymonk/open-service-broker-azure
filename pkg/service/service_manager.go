package service

// ServiceManager is an interface to be implemented by module components
// responsible for managing the lifecycle of services and plans thereof
type ServiceManager interface { // nolint: golint
	// GetEmptyProvisioningParameters returns an empty instance of
	// service-specific non-sensitive  provisioning parameters
	GetEmptyProvisioningParameters() ProvisioningParameters
	// GetEmptySecureProvisioningParameters returns an empty instance of
	// service-specific secured (sensitive) provisioning parameters
	GetEmptySecureProvisioningParameters() SecureProvisioningParameters
	// ValidateProvisioningParameters validates the provided
	// provisioningParameters and returns an error if there is any problem
	ValidateProvisioningParameters(
		ProvisioningParameters,
		SecureProvisioningParameters,
	) error
	// GetProvisioner returns a provisioner that defines the steps a module must
	// execute asynchronously to provision a service.
	GetProvisioner(Plan) (Provisioner, error)
	// GetEmptyInstanceDetails returns an empty instance of non-sensitive
	// service-specific instance details
	GetEmptyInstanceDetails() InstanceDetails
	// GetEmptySecureInstanceDetails returns an empty instance of sensitive
	// service-specific instance details
	GetEmptySecureInstanceDetails() SecureInstanceDetails
	// ValidateUpdatingParameters validates the provided
	// updating parameters against allowed values and current instance state
	// and returns an error if there is any problem
	ValidateUpdatingParameters(Instance) error
	// GetUpdater returns a updater that defines the steps a module must
	// execute asynchronously to update a service.
	GetUpdater(Plan) (Updater, error)
	// GetEmptyBindingParameters returns an empty instance of non-sensitive
	// service-specific binding parameters
	GetEmptyBindingParameters() BindingParameters
	// GetEmptySecureBindingParameters returns an empty instance of sensitive
	// service-specific binding parameters
	GetEmptySecureBindingParameters() SecureBindingParameters
	// ValidateBindingParameters validates the provided bindingParameters and
	// returns an error if there is any problem
	ValidateBindingParameters(BindingParameters, SecureBindingParameters) error
	// Bind synchronously binds to a service
	Bind(
		Instance,
		BindingParameters,
		SecureBindingParameters,
	) (BindingDetails, SecureBindingDetails, error)
	// GetEmptyBindingDetails returns an empty instance of service-specific
	// non-sensitive binding details
	GetEmptyBindingDetails() BindingDetails
	// GetEmptySecureBindingDetails returns an empty instance of service-specific
	// secured (sensitive) binding details
	GetEmptySecureBindingDetails() SecureBindingDetails
	// GetCredentials returns service-specific credentials populated from instance
	// and binding details
	GetCredentials(Instance, Binding) (Credentials, error)
	// Unbind synchronously unbinds from a service
	Unbind(Instance, Binding) error
	// GetDeprovisioner returns a deprovisioner that defines the steps a module
	// must execute asynchronously to deprovision a service
	GetDeprovisioner(Plan) (Deprovisioner, error)
}
