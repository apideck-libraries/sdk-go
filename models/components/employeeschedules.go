// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type EmployeeSchedules struct {
	Employee  *Employee  `json:"employee,omitempty"`
	Schedules []Schedule `json:"schedules,omitempty"`
}

func (o *EmployeeSchedules) GetEmployee() *Employee {
	if o == nil {
		return nil
	}
	return o.Employee
}

func (o *EmployeeSchedules) GetSchedules() []Schedule {
	if o == nil {
		return nil
	}
	return o.Schedules
}
