// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type AccountingLocationsFilter struct {
	// Id of the subsidiary to search for
	Subsidiary *string `queryParam:"name=subsidiary"`
}

func (o *AccountingLocationsFilter) GetSubsidiary() *string {
	if o == nil {
		return nil
	}
	return o.Subsidiary
}
