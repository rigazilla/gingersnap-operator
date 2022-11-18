// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// ResourceQuantityApplyConfiguration represents an declarative configuration of the ResourceQuantity type for use
// with apply.
type ResourceQuantityApplyConfiguration struct {
	Memory *string `json:"memory,omitempty"`
	Cpu    *string `json:"cpu,omitempty"`
}

// ResourceQuantityApplyConfiguration constructs an declarative configuration of the ResourceQuantity type for use with
// apply.
func ResourceQuantity() *ResourceQuantityApplyConfiguration {
	return &ResourceQuantityApplyConfiguration{}
}

// WithMemory sets the Memory field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Memory field is set to the value of the last call.
func (b *ResourceQuantityApplyConfiguration) WithMemory(value string) *ResourceQuantityApplyConfiguration {
	b.Memory = &value
	return b
}

// WithCpu sets the Cpu field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Cpu field is set to the value of the last call.
func (b *ResourceQuantityApplyConfiguration) WithCpu(value string) *ResourceQuantityApplyConfiguration {
	b.Cpu = &value
	return b
}
