// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// EmploymentStatus - The employment status of the employee, indicating whether they are currently employed, inactive, terminated, or in another status.
type EmploymentStatus string

const (
	EmploymentStatusActive     EmploymentStatus = "active"
	EmploymentStatusInactive   EmploymentStatus = "inactive"
	EmploymentStatusTerminated EmploymentStatus = "terminated"
	EmploymentStatusOther      EmploymentStatus = "other"
)

func (e EmploymentStatus) ToPointer() *EmploymentStatus {
	return &e
}
func (e *EmploymentStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "active":
		fallthrough
	case "inactive":
		fallthrough
	case "terminated":
		fallthrough
	case "other":
		*e = EmploymentStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EmploymentStatus: %v", v)
	}
}
