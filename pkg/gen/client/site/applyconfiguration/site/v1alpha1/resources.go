/*
Licensed under the GNU AFFERO GENERAL PUBLIC LICENSE Version 3
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// ResourcesApplyConfiguration represents an declarative configuration of the Resources type for use
// with apply.
type ResourcesApplyConfiguration struct {
	ResourceRequests *ResourceDescriptionApplyConfiguration `json:"requests,omitempty"`
	ResourceLimits   *ResourceDescriptionApplyConfiguration `json:"limits,omitempty"`
}

// ResourcesApplyConfiguration constructs an declarative configuration of the Resources type for use with
// apply.
func Resources() *ResourcesApplyConfiguration {
	return &ResourcesApplyConfiguration{}
}

// WithResourceRequests sets the ResourceRequests field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceRequests field is set to the value of the last call.
func (b *ResourcesApplyConfiguration) WithResourceRequests(value *ResourceDescriptionApplyConfiguration) *ResourcesApplyConfiguration {
	b.ResourceRequests = value
	return b
}

// WithResourceLimits sets the ResourceLimits field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceLimits field is set to the value of the last call.
func (b *ResourcesApplyConfiguration) WithResourceLimits(value *ResourceDescriptionApplyConfiguration) *ResourcesApplyConfiguration {
	b.ResourceLimits = value
	return b
}
