// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/types"
	"time"
)

// Visibility - The visibility of the job
type Visibility string

const (
	VisibilityDraft    Visibility = "draft"
	VisibilityPublic   Visibility = "public"
	VisibilityInternal Visibility = "internal"
)

func (e Visibility) ToPointer() *Visibility {
	return &e
}
func (e *Visibility) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "draft":
		fallthrough
	case "public":
		fallthrough
	case "internal":
		*e = Visibility(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Visibility: %v", v)
	}
}

type EmploymentTerms string

const (
	EmploymentTermsFullTime   EmploymentTerms = "full-time"
	EmploymentTermsPartTime   EmploymentTerms = "part-time"
	EmploymentTermsInternship EmploymentTerms = "internship"
	EmploymentTermsContractor EmploymentTerms = "contractor"
	EmploymentTermsEmployee   EmploymentTerms = "employee"
	EmploymentTermsFreelance  EmploymentTerms = "freelance"
	EmploymentTermsTemp       EmploymentTerms = "temp"
	EmploymentTermsSeasonal   EmploymentTerms = "seasonal"
	EmploymentTermsVolunteer  EmploymentTerms = "volunteer"
	EmploymentTermsOther      EmploymentTerms = "other"
)

func (e EmploymentTerms) ToPointer() *EmploymentTerms {
	return &e
}
func (e *EmploymentTerms) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "full-time":
		fallthrough
	case "part-time":
		fallthrough
	case "internship":
		fallthrough
	case "contractor":
		fallthrough
	case "employee":
		fallthrough
	case "freelance":
		fallthrough
	case "temp":
		fallthrough
	case "seasonal":
		fallthrough
	case "volunteer":
		fallthrough
	case "other":
		*e = EmploymentTerms(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EmploymentTerms: %v", v)
	}
}

// Branch - Details of the branch for which the job is created.
type Branch struct {
	// A unique identifier for an object.
	ID *string `json:"id,omitempty"`
	// Name of the branch.
	Name *string `json:"name,omitempty"`
}

func (o *Branch) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Branch) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type Blocks struct {
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
}

func (o *Blocks) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *Blocks) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

type Salary struct {
	// Minimum salary payable for the job role.
	Min *int64 `json:"min,omitempty"`
	// Maximum salary payable for the job role.
	Max *int64 `json:"max,omitempty"`
	// Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Currency *Currency `json:"currency,omitempty"`
	Interval *string   `json:"interval,omitempty"`
}

func (o *Salary) GetMin() *int64 {
	if o == nil {
		return nil
	}
	return o.Min
}

func (o *Salary) GetMax() *int64 {
	if o == nil {
		return nil
	}
	return o.Max
}

func (o *Salary) GetCurrency() *Currency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *Salary) GetInterval() *string {
	if o == nil {
		return nil
	}
	return o.Interval
}

type JobType string

const (
	JobTypeJobPortal      JobType = "job_portal"
	JobTypeJobDescription JobType = "job_description"
)

func (e JobType) ToPointer() *JobType {
	return &e
}
func (e *JobType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "job_portal":
		fallthrough
	case "job_description":
		*e = JobType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for JobType: %v", v)
	}
}

type JobLinks struct {
	Type *JobType `json:"type,omitempty"`
	URL  *string  `json:"url,omitempty"`
}

func (o *JobLinks) GetType() *JobType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *JobLinks) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

type Job struct {
	// A unique identifier for an object.
	ID   *string `json:"id,omitempty"`
	Slug *string `json:"slug,omitempty"`
	// The job title of the person.
	Title *string `json:"title,omitempty"`
	// Sequence in relation to other jobs.
	Sequence *int64 `json:"sequence,omitempty"`
	// The visibility of the job
	Visibility *Visibility `json:"visibility,omitempty"`
	// The status of the job.
	Status *JobStatus `json:"status,omitempty"`
	// The code of the job.
	Code *string `json:"code,omitempty"`
	// language code according to ISO 639-1. For the United States - EN
	Language        *string          `json:"language,omitempty"`
	EmploymentTerms *EmploymentTerms `json:"employment_terms,omitempty"`
	// Level of experience required for the job role.
	Experience *string `json:"experience,omitempty"`
	// Specifies the location for the job posting.
	Location *string `json:"location,omitempty"`
	// Specifies whether the posting is for a remote job.
	Remote *bool `json:"remote,omitempty"`
	// A job's Requisition ID (Req ID) allows your organization to identify and track a job based on alphanumeric naming conventions unique to your company's internal processes.
	RequisitionID *string     `json:"requisition_id,omitempty"`
	Department    *Department `json:"department,omitempty"`
	// Details of the branch for which the job is created.
	Branch *Branch `json:"branch,omitempty"`
	// The recruiter is generally someone who is tasked to help the hiring manager find and screen qualified applicant
	Recruiters     []string `json:"recruiters,omitempty"`
	HiringManagers []string `json:"hiring_managers,omitempty"`
	Followers      []string `json:"followers,omitempty"`
	// A description of the object.
	Description *string `json:"description,omitempty"`
	// The job description in HTML format
	DescriptionHTML *string  `json:"description_html,omitempty"`
	Blocks          []Blocks `json:"blocks,omitempty"`
	Closing         *string  `json:"closing,omitempty"`
	// The closing section of the job description in HTML format
	ClosingHTML *string     `json:"closing_html,omitempty"`
	ClosingDate *types.Date `json:"closing_date,omitempty"`
	Salary      *Salary     `json:"salary,omitempty"`
	// URL of the job description
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	URL *string `json:"url,omitempty"`
	// URL of the job portal
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	JobPortalURL *string `json:"job_portal_url,omitempty"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	RecordURL    *string    `json:"record_url,omitempty"`
	Links        []JobLinks `json:"links,omitempty"`
	Confidential *bool      `json:"confidential,omitempty"`
	// Specifies whether an employee of the organization can apply for the job.
	AvailableToEmployees *bool         `json:"available_to_employees,omitempty"`
	Tags                 []string      `json:"tags,omitempty"`
	Addresses            []Address     `json:"addresses,omitempty"`
	CustomFields         []CustomField `json:"custom_fields,omitempty"`
	// Flag to indicate if the object is deleted.
	Deleted     *bool      `json:"deleted,omitempty"`
	OwnerID     *string    `json:"owner_id,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	// When custom mappings are configured on the resource, the result is included here.
	CustomMappings *CustomMappings `json:"custom_mappings,omitempty"`
	// The user who last updated the object.
	UpdatedBy *string `json:"updated_by,omitempty"`
	// The user who created the object.
	CreatedBy *string `json:"created_by,omitempty"`
	// The date and time when the object was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The date and time when the object was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

func (j Job) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(j, "", false)
}

func (j *Job) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &j, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Job) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Job) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *Job) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *Job) GetSequence() *int64 {
	if o == nil {
		return nil
	}
	return o.Sequence
}

func (o *Job) GetVisibility() *Visibility {
	if o == nil {
		return nil
	}
	return o.Visibility
}

func (o *Job) GetStatus() *JobStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Job) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *Job) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *Job) GetEmploymentTerms() *EmploymentTerms {
	if o == nil {
		return nil
	}
	return o.EmploymentTerms
}

func (o *Job) GetExperience() *string {
	if o == nil {
		return nil
	}
	return o.Experience
}

func (o *Job) GetLocation() *string {
	if o == nil {
		return nil
	}
	return o.Location
}

func (o *Job) GetRemote() *bool {
	if o == nil {
		return nil
	}
	return o.Remote
}

func (o *Job) GetRequisitionID() *string {
	if o == nil {
		return nil
	}
	return o.RequisitionID
}

func (o *Job) GetDepartment() *Department {
	if o == nil {
		return nil
	}
	return o.Department
}

func (o *Job) GetBranch() *Branch {
	if o == nil {
		return nil
	}
	return o.Branch
}

func (o *Job) GetRecruiters() []string {
	if o == nil {
		return nil
	}
	return o.Recruiters
}

func (o *Job) GetHiringManagers() []string {
	if o == nil {
		return nil
	}
	return o.HiringManagers
}

func (o *Job) GetFollowers() []string {
	if o == nil {
		return nil
	}
	return o.Followers
}

func (o *Job) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Job) GetDescriptionHTML() *string {
	if o == nil {
		return nil
	}
	return o.DescriptionHTML
}

func (o *Job) GetBlocks() []Blocks {
	if o == nil {
		return nil
	}
	return o.Blocks
}

func (o *Job) GetClosing() *string {
	if o == nil {
		return nil
	}
	return o.Closing
}

func (o *Job) GetClosingHTML() *string {
	if o == nil {
		return nil
	}
	return o.ClosingHTML
}

func (o *Job) GetClosingDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.ClosingDate
}

func (o *Job) GetSalary() *Salary {
	if o == nil {
		return nil
	}
	return o.Salary
}

func (o *Job) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *Job) GetJobPortalURL() *string {
	if o == nil {
		return nil
	}
	return o.JobPortalURL
}

func (o *Job) GetRecordURL() *string {
	if o == nil {
		return nil
	}
	return o.RecordURL
}

func (o *Job) GetLinks() []JobLinks {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *Job) GetConfidential() *bool {
	if o == nil {
		return nil
	}
	return o.Confidential
}

func (o *Job) GetAvailableToEmployees() *bool {
	if o == nil {
		return nil
	}
	return o.AvailableToEmployees
}

func (o *Job) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Job) GetAddresses() []Address {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *Job) GetCustomFields() []CustomField {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

func (o *Job) GetDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.Deleted
}

func (o *Job) GetOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.OwnerID
}

func (o *Job) GetPublishedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.PublishedAt
}

func (o *Job) GetCustomMappings() *CustomMappings {
	if o == nil {
		return nil
	}
	return o.CustomMappings
}

func (o *Job) GetUpdatedBy() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedBy
}

func (o *Job) GetCreatedBy() *string {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *Job) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Job) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}
