/*
 * Subset of Influx API covered by Influx CLI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TemplateApplyActionProperties struct for TemplateApplyActionProperties
type TemplateApplyActionProperties struct {
	Kind                 string  `json:"kind" yaml:"kind"`
	ResourceTemplateName *string `json:"resourceTemplateName,omitempty" yaml:"resourceTemplateName,omitempty"`
}

// NewTemplateApplyActionProperties instantiates a new TemplateApplyActionProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateApplyActionProperties(kind string) *TemplateApplyActionProperties {
	this := TemplateApplyActionProperties{}
	this.Kind = kind
	return &this
}

// NewTemplateApplyActionPropertiesWithDefaults instantiates a new TemplateApplyActionProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateApplyActionPropertiesWithDefaults() *TemplateApplyActionProperties {
	this := TemplateApplyActionProperties{}
	return &this
}

// GetKind returns the Kind field value
func (o *TemplateApplyActionProperties) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *TemplateApplyActionProperties) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *TemplateApplyActionProperties) SetKind(v string) {
	o.Kind = v
}

// GetResourceTemplateName returns the ResourceTemplateName field value if set, zero value otherwise.
func (o *TemplateApplyActionProperties) GetResourceTemplateName() string {
	if o == nil || o.ResourceTemplateName == nil {
		var ret string
		return ret
	}
	return *o.ResourceTemplateName
}

// GetResourceTemplateNameOk returns a tuple with the ResourceTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateApplyActionProperties) GetResourceTemplateNameOk() (*string, bool) {
	if o == nil || o.ResourceTemplateName == nil {
		return nil, false
	}
	return o.ResourceTemplateName, true
}

// HasResourceTemplateName returns a boolean if a field has been set.
func (o *TemplateApplyActionProperties) HasResourceTemplateName() bool {
	if o != nil && o.ResourceTemplateName != nil {
		return true
	}

	return false
}

// SetResourceTemplateName gets a reference to the given string and assigns it to the ResourceTemplateName field.
func (o *TemplateApplyActionProperties) SetResourceTemplateName(v string) {
	o.ResourceTemplateName = &v
}

func (o TemplateApplyActionProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if o.ResourceTemplateName != nil {
		toSerialize["resourceTemplateName"] = o.ResourceTemplateName
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateApplyActionProperties struct {
	value *TemplateApplyActionProperties
	isSet bool
}

func (v NullableTemplateApplyActionProperties) Get() *TemplateApplyActionProperties {
	return v.value
}

func (v *NullableTemplateApplyActionProperties) Set(val *TemplateApplyActionProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateApplyActionProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateApplyActionProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateApplyActionProperties(val *TemplateApplyActionProperties) *NullableTemplateApplyActionProperties {
	return &NullableTemplateApplyActionProperties{value: val, isSet: true}
}

func (v NullableTemplateApplyActionProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateApplyActionProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}