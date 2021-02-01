package v1alpha1

import (
	"context"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// SaveStatus will trigger an object update to save the current status
// conditions
func (httpso *HTTPScaledObject) SaveStatus(
	logger logr.Logger,
	cl client.Client,
) {
	ctx := context.TODO()
	logger.Info("Updating status on object", "scaledobject", *httpso)
	tmpHTTPSo := &HTTPScaledObject{}
	if err := cl.Get(ctx, types.NamespacedName{
		Namespace: httpso.Namespace,
		Name:      httpso.Name,
	}, tmpHTTPSo); err != nil {
		logger.Error(err, "HTTPScaledObject not found on test check")
	} else {
		logger.Info("Found HTTPScaledObject on test check", "scaled object", *tmpHTTPSo)
	}

	tmpHTTPSo.Status = httpso.Status
	err := cl.Status().Update(ctx, tmpHTTPSo)
	httpso = tmpHTTPSo

	if err != nil {
		logger.Error(err, "failed to update status on HTTPScaledObject", "httpso", httpso)
	} else {
		logger.Info("Updated status on HTTPScaledObject", "HTTPScaledObject", *httpso)
	}
}

// AddCondition adds a new condition to the HTTPScaledObject
func (httpso *HTTPScaledObject) AddCondition(condition HTTPScaledObjectCondition) *HTTPScaledObject {
	httpso.Status.Conditions = append(httpso.Status.Conditions, condition)
	return httpso
}

// CreateCondition initializes a new status condition
func CreateCondition (
	condType HTTPScaledObjectCreationStatus,
	status metav1.ConditionStatus,
	reason HTTPScaledObjectConditionReason,
) *HTTPScaledObjectCondition {
	cond := HTTPScaledObjectCondition{
		Type: condType,
		Status: status,
		Reason: reason,
	}
	return &cond
}

// SetMessage sets the optional reason for the condition
func (c *HTTPScaledObjectCondition) SetMessage (message string) *HTTPScaledObjectCondition {
	c.Message = message
	return c
}
