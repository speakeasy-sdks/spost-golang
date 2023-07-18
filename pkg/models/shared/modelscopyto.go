// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ModelsCopyToCustomFields struct {
}

type ModelsCopyTo struct {
	CustomFields *ModelsCopyToCustomFields
	Email        *string
	Name         *string
}

func (o *ModelsCopyTo) GetCustomFields() *ModelsCopyToCustomFields {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

func (o *ModelsCopyTo) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *ModelsCopyTo) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
