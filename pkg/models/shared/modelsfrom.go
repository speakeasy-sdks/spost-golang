// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ModelsFrom struct {
	Email *string
	Name  *string
}

func (o *ModelsFrom) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *ModelsFrom) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}