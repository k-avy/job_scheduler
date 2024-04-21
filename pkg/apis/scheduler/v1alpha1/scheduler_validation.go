package v1alpha1

import (
	"context"

	"knative.dev/pkg/apis"
)

// Validate implements apis.Validatable
func (as *Scheduler) Validate(ctx context.Context) *apis.FieldError {
	return as.Spec.Validate(ctx).ViaField("spec")
}

// Validate implements apis.Validatable
func (ass *SchedulerSpec) Validate(ctx context.Context) *apis.FieldError {
	if ass.ServiceName == "" {
		return apis.ErrMissingField("serviceName")
	}
	return nil
}
