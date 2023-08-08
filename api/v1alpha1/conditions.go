package v1alpha1

import "github.com/redhat-appstudio/operator-goodies/conditions"

const (
	// processedConditionType is the type used to track the status of a resource processing
	processedConditionType conditions.ConditionType = "Processed"

	// validatedConditionType is the type used to track the status of a resource validation
	validatedConditionType conditions.ConditionType = "Validated"
)

const (
	// FailedReason is the reason set when a failure occurs
	FailedReason conditions.ConditionReason = "Failed"

	// ProgressingReason is the reason set when an action is progressing
	ProgressingReason conditions.ConditionReason = "Progressing"

	// SucceededReason is the reason set when an action succeeds
	SucceededReason conditions.ConditionReason = "Succeeded"
)
