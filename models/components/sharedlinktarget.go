// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type SharedLinkTarget struct {
	// A unique identifier for an object.
	ID string `json:"id"`
	// The name of the file
	Name *string `json:"name,omitempty"`
	// The type of resource. Could be file, folder or url
	Type *FileType `json:"type,omitempty"`
}

func (o *SharedLinkTarget) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *SharedLinkTarget) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SharedLinkTarget) GetType() *FileType {
	if o == nil {
		return nil
	}
	return o.Type
}
