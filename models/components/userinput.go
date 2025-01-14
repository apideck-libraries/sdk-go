// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UserInput struct {
	// The parent user id
	ParentID *string `json:"parent_id,omitempty"`
	// The username of the user
	Username *string `json:"username,omitempty"`
	// The first name of the person.
	FirstName *string `json:"first_name,omitempty"`
	// The last name of the person.
	LastName *string `json:"last_name,omitempty"`
	// The job title of the person.
	Title *string `json:"title,omitempty"`
	// The division the person is currently in. Usually a collection of departments or teams or regions.
	Division *string `json:"division,omitempty"`
	// The department the person is currently in. [Deprecated](https://developers.apideck.com/changelog) in favor of the dedicated department_id and department_name field.
	//
	// Deprecated: This will be removed in a future release, please migrate away from it as soon as possible.
	Department *string `json:"department,omitempty"`
	// The name of the company.
	CompanyName *string `json:"company_name,omitempty"`
	// An Employee Number, Employee ID or Employee Code, is a unique number that has been assigned to each individual staff member within a company.
	EmployeeNumber *string `json:"employee_number,omitempty"`
	// A description of the object.
	Description *string `json:"description,omitempty"`
	// The URL of the user's avatar
	Image *string `json:"image,omitempty"`
	// language code according to ISO 639-1. For the United States - EN
	Language *string `json:"language,omitempty"`
	// The status of the user
	Status *string `json:"status,omitempty"`
	// The password of the user
	Password     *string       `json:"password,omitempty"`
	Addresses    []Address     `json:"addresses,omitempty"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers,omitempty"`
	Emails       []Email       `json:"emails"`
	// The pass_through property allows passing service-specific, custom data or structured modifications in request body when creating or updating resources.
	PassThrough []PassThroughBody `json:"pass_through,omitempty"`
}

func (o *UserInput) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *UserInput) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *UserInput) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *UserInput) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *UserInput) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *UserInput) GetDivision() *string {
	if o == nil {
		return nil
	}
	return o.Division
}

func (o *UserInput) GetDepartment() *string {
	if o == nil {
		return nil
	}
	return o.Department
}

func (o *UserInput) GetCompanyName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyName
}

func (o *UserInput) GetEmployeeNumber() *string {
	if o == nil {
		return nil
	}
	return o.EmployeeNumber
}

func (o *UserInput) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UserInput) GetImage() *string {
	if o == nil {
		return nil
	}
	return o.Image
}

func (o *UserInput) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *UserInput) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *UserInput) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *UserInput) GetAddresses() []Address {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *UserInput) GetPhoneNumbers() []PhoneNumber {
	if o == nil {
		return nil
	}
	return o.PhoneNumbers
}

func (o *UserInput) GetEmails() []Email {
	if o == nil {
		return []Email{}
	}
	return o.Emails
}

func (o *UserInput) GetPassThrough() []PassThroughBody {
	if o == nil {
		return nil
	}
	return o.PassThrough
}
