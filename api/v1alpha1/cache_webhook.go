package v1alpha1

import (
	gingersnapapi "github.com/gingersnap-project/operator/gen/gingersnap-api/config/cache/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (c *Cache) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(c).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-gingersnap-project-io-v1alpha1-cache,mutating=true,failurePolicy=fail,sideEffects=None,groups=gingersnap-project.io,resources=caches,verbs=create;update,versions=v1alpha1,name=mcache.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Cache{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (c *Cache) Default() {
	spec := &c.Spec
	if spec.Redis == nil && spec.Infinispan == nil {
		spec.Infinispan = &gingersnapapi.InfinispanSpec{}
	}
}

//+kubebuilder:webhook:path=/validate-gingersnap-project-io-v1alpha1-cache,mutating=false,failurePolicy=fail,sideEffects=None,groups=gingersnap-project.io,resources=caches,verbs=create;update,versions=v1alpha1,name=vcache.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Cache{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (c *Cache) ValidateCreate() error {
	var allErrs field.ErrorList

	if c.Spec.Redis != nil && c.Spec.Infinispan != nil {
		allErrs = append(allErrs, field.Required(field.NewPath("spec"), "At most one of ['spec.infinispan', 'spec.redis'] must be configured"))
	}
	return c.StatusError(allErrs)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (c *Cache) ValidateUpdate(old runtime.Object) error {
	if err := c.ValidateCreate(); err != nil {
		return err
	}

	oldSpec := old.(*Cache).Spec
	var allErrs field.ErrorList

	if c.Spec.Infinispan == nil && oldSpec.Infinispan != nil {
		allErrs = append(allErrs, field.Forbidden(field.NewPath("spec").Child("infinispan"), "The Infinispan spec is immutable and cannot be unset after initial Cache creation"))
	}

	if c.Spec.Redis == nil && oldSpec.Redis != nil {
		allErrs = append(allErrs, field.Forbidden(field.NewPath("spec").Child("redis"), "The Redis spec is immutable and cannot be unset after initial Cache creation"))
	}

	return c.StatusError(allErrs)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (c *Cache) ValidateDelete() error {
	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

func (c *Cache) StatusError(allErrs field.ErrorList) error {
	if len(allErrs) != 0 {
		return apierrors.NewInvalid(
			schema.GroupKind{Group: GroupVersion.Group, Kind: KindCache},
			c.Name, allErrs)
	}
	return nil
}
