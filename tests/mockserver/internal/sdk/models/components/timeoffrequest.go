// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"mockserver/internal/sdk/utils"
	"time"
)

// TimeOffRequestStatus - The status of the time off request.
type TimeOffRequestStatus string

const (
	TimeOffRequestStatusRequested TimeOffRequestStatus = "requested"
	TimeOffRequestStatusApproved  TimeOffRequestStatus = "approved"
	TimeOffRequestStatusDeclined  TimeOffRequestStatus = "declined"
	TimeOffRequestStatusCancelled TimeOffRequestStatus = "cancelled"
	TimeOffRequestStatusDeleted   TimeOffRequestStatus = "deleted"
	TimeOffRequestStatusOther     TimeOffRequestStatus = "other"
)

func (e TimeOffRequestStatus) ToPointer() *TimeOffRequestStatus {
	return &e
}
func (e *TimeOffRequestStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "requested":
		fallthrough
	case "approved":
		fallthrough
	case "declined":
		fallthrough
	case "cancelled":
		fallthrough
	case "deleted":
		fallthrough
	case "other":
		*e = TimeOffRequestStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TimeOffRequestStatus: %v", v)
	}
}

// RequestType - The type of request
type RequestType string

const (
	RequestTypeVacation    RequestType = "vacation"
	RequestTypeSick        RequestType = "sick"
	RequestTypePersonal    RequestType = "personal"
	RequestTypeJuryDuty    RequestType = "jury_duty"
	RequestTypeVolunteer   RequestType = "volunteer"
	RequestTypeBereavement RequestType = "bereavement"
	RequestTypeOther       RequestType = "other"
)

func (e RequestType) ToPointer() *RequestType {
	return &e
}
func (e *RequestType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "vacation":
		fallthrough
	case "sick":
		fallthrough
	case "personal":
		fallthrough
	case "jury_duty":
		fallthrough
	case "volunteer":
		fallthrough
	case "bereavement":
		fallthrough
	case "other":
		*e = RequestType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestType: %v", v)
	}
}

// Units - The unit of time off requested. Possible values include: `hours`, `days`, or `other`.
type Units string

const (
	UnitsDays  Units = "days"
	UnitsHours Units = "hours"
	UnitsOther Units = "other"
)

func (e Units) ToPointer() *Units {
	return &e
}
func (e *Units) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "days":
		fallthrough
	case "hours":
		fallthrough
	case "other":
		*e = Units(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Units: %v", v)
	}
}

type Notes struct {
	Employee *string `json:"employee,omitempty"`
	Manager  *string `json:"manager,omitempty"`
}

func (o *Notes) GetEmployee() *string {
	if o == nil {
		return nil
	}
	return o.Employee
}

func (o *Notes) GetManager() *string {
	if o == nil {
		return nil
	}
	return o.Manager
}

type TimeOffRequest struct {
	// A unique identifier for an object.
	ID *string `json:"id,omitempty"`
	// ID of the employee
	EmployeeID *string `json:"employee_id,omitempty"`
	// ID of the policy
	PolicyID *string `json:"policy_id,omitempty"`
	// The status of the time off request.
	Status *TimeOffRequestStatus `json:"status,omitempty"`
	// Description of the time off request.
	Description *string `json:"description,omitempty"`
	// The start date of the time off request.
	StartDate *string `json:"start_date,omitempty"`
	// The end date of the time off request.
	EndDate *string `json:"end_date,omitempty"`
	// The date the request was made.
	RequestDate *string `json:"request_date,omitempty"`
	// The type of request
	RequestType *RequestType `json:"request_type,omitempty"`
	// The date the request was approved
	ApprovalDate *string `json:"approval_date,omitempty"`
	// The unit of time off requested. Possible values include: `hours`, `days`, or `other`.
	Units *Units `json:"units,omitempty"`
	// The amount of time off requested.
	Amount *float64 `json:"amount,omitempty"`
	// The day part of the time off request.
	DayPart *string `json:"day_part,omitempty"`
	Notes   *Notes  `json:"notes,omitempty"`
	// When custom mappings are configured on the resource, the result is included here.
	CustomMappings map[string]any `json:"custom_mappings,omitempty"`
	// The user who last updated the object.
	UpdatedBy *string `json:"updated_by,omitempty"`
	// The user who created the object.
	CreatedBy *string `json:"created_by,omitempty"`
	// The date and time when the object was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The date and time when the object was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The pass_through property allows passing service-specific, custom data or structured modifications in request body when creating or updating resources.
	PassThrough []PassThroughBody `json:"pass_through,omitempty"`
	// The policy type of the time off request
	PolicyType *string `json:"policy_type,omitempty"`
}

func (t TimeOffRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TimeOffRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TimeOffRequest) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TimeOffRequest) GetEmployeeID() *string {
	if o == nil {
		return nil
	}
	return o.EmployeeID
}

func (o *TimeOffRequest) GetPolicyID() *string {
	if o == nil {
		return nil
	}
	return o.PolicyID
}

func (o *TimeOffRequest) GetStatus() *TimeOffRequestStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *TimeOffRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *TimeOffRequest) GetStartDate() *string {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *TimeOffRequest) GetEndDate() *string {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *TimeOffRequest) GetRequestDate() *string {
	if o == nil {
		return nil
	}
	return o.RequestDate
}

func (o *TimeOffRequest) GetRequestType() *RequestType {
	if o == nil {
		return nil
	}
	return o.RequestType
}

func (o *TimeOffRequest) GetApprovalDate() *string {
	if o == nil {
		return nil
	}
	return o.ApprovalDate
}

func (o *TimeOffRequest) GetUnits() *Units {
	if o == nil {
		return nil
	}
	return o.Units
}

func (o *TimeOffRequest) GetAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *TimeOffRequest) GetDayPart() *string {
	if o == nil {
		return nil
	}
	return o.DayPart
}

func (o *TimeOffRequest) GetNotes() *Notes {
	if o == nil {
		return nil
	}
	return o.Notes
}

func (o *TimeOffRequest) GetCustomMappings() map[string]any {
	if o == nil {
		return nil
	}
	return o.CustomMappings
}

func (o *TimeOffRequest) GetUpdatedBy() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedBy
}

func (o *TimeOffRequest) GetCreatedBy() *string {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *TimeOffRequest) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *TimeOffRequest) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *TimeOffRequest) GetPassThrough() []PassThroughBody {
	if o == nil {
		return nil
	}
	return o.PassThrough
}

func (o *TimeOffRequest) GetPolicyType() *string {
	if o == nil {
		return nil
	}
	return o.PolicyType
}

type TimeOffRequestInput struct {
	// ID of the employee
	EmployeeID *string `json:"employee_id,omitempty"`
	// ID of the policy
	PolicyID *string `json:"policy_id,omitempty"`
	// The status of the time off request.
	Status *TimeOffRequestStatus `json:"status,omitempty"`
	// Description of the time off request.
	Description *string `json:"description,omitempty"`
	// The start date of the time off request.
	StartDate *string `json:"start_date,omitempty"`
	// The end date of the time off request.
	EndDate *string `json:"end_date,omitempty"`
	// The date the request was made.
	RequestDate *string `json:"request_date,omitempty"`
	// The type of request
	RequestType *RequestType `json:"request_type,omitempty"`
	// The date the request was approved
	ApprovalDate *string `json:"approval_date,omitempty"`
	// The unit of time off requested. Possible values include: `hours`, `days`, or `other`.
	Units *Units `json:"units,omitempty"`
	// The amount of time off requested.
	Amount *float64 `json:"amount,omitempty"`
	// The day part of the time off request.
	DayPart *string `json:"day_part,omitempty"`
	Notes   *Notes  `json:"notes,omitempty"`
	// The pass_through property allows passing service-specific, custom data or structured modifications in request body when creating or updating resources.
	PassThrough []PassThroughBody `json:"pass_through,omitempty"`
	// The policy type of the time off request
	PolicyType *string `json:"policy_type,omitempty"`
}

func (o *TimeOffRequestInput) GetEmployeeID() *string {
	if o == nil {
		return nil
	}
	return o.EmployeeID
}

func (o *TimeOffRequestInput) GetPolicyID() *string {
	if o == nil {
		return nil
	}
	return o.PolicyID
}

func (o *TimeOffRequestInput) GetStatus() *TimeOffRequestStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *TimeOffRequestInput) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *TimeOffRequestInput) GetStartDate() *string {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *TimeOffRequestInput) GetEndDate() *string {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *TimeOffRequestInput) GetRequestDate() *string {
	if o == nil {
		return nil
	}
	return o.RequestDate
}

func (o *TimeOffRequestInput) GetRequestType() *RequestType {
	if o == nil {
		return nil
	}
	return o.RequestType
}

func (o *TimeOffRequestInput) GetApprovalDate() *string {
	if o == nil {
		return nil
	}
	return o.ApprovalDate
}

func (o *TimeOffRequestInput) GetUnits() *Units {
	if o == nil {
		return nil
	}
	return o.Units
}

func (o *TimeOffRequestInput) GetAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *TimeOffRequestInput) GetDayPart() *string {
	if o == nil {
		return nil
	}
	return o.DayPart
}

func (o *TimeOffRequestInput) GetNotes() *Notes {
	if o == nil {
		return nil
	}
	return o.Notes
}

func (o *TimeOffRequestInput) GetPassThrough() []PassThroughBody {
	if o == nil {
		return nil
	}
	return o.PassThrough
}

func (o *TimeOffRequestInput) GetPolicyType() *string {
	if o == nil {
		return nil
	}
	return o.PolicyType
}
