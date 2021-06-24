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

// TemplateSummaryLabelMapping struct for TemplateSummaryLabelMapping
type TemplateSummaryLabelMapping struct {
	Status                   string `json:"status"`
	ResourceTemplateMetaName string `json:"resourceTemplateMetaName"`
	ResourceName             string `json:"resourceName"`
	ResourceID               string `json:"resourceID"`
	ResourceType             string `json:"resourceType"`
	LabelTemplateMetaName    string `json:"labelTemplateMetaName"`
	LabelName                string `json:"labelName"`
	LabelID                  string `json:"labelID"`
}

// NewTemplateSummaryLabelMapping instantiates a new TemplateSummaryLabelMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryLabelMapping(status string, resourceTemplateMetaName string, resourceName string, resourceID string, resourceType string, labelTemplateMetaName string, labelName string, labelID string) *TemplateSummaryLabelMapping {
	this := TemplateSummaryLabelMapping{}
	this.Status = status
	this.ResourceTemplateMetaName = resourceTemplateMetaName
	this.ResourceName = resourceName
	this.ResourceID = resourceID
	this.ResourceType = resourceType
	this.LabelTemplateMetaName = labelTemplateMetaName
	this.LabelName = labelName
	this.LabelID = labelID
	return &this
}

// NewTemplateSummaryLabelMappingWithDefaults instantiates a new TemplateSummaryLabelMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryLabelMappingWithDefaults() *TemplateSummaryLabelMapping {
	this := TemplateSummaryLabelMapping{}
	return &this
}

// GetStatus returns the Status field value
func (o *TemplateSummaryLabelMapping) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryLabelMapping) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TemplateSummaryLabelMapping) SetStatus(v string) {
	o.Status = v
}

// GetResourceTemplateMetaName returns the ResourceTemplateMetaName field value
func (o *TemplateSummaryLabelMapping) GetResourceTemplateMetaName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceTemplateMetaName
}

// GetResourceTemplateMetaNameOk returns a tuple with the ResourceTemplateMetaName field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryLabelMapping) GetResourceTemplateMetaNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceTemplateMetaName, true
}

// SetResourceTemplateMetaName sets field value
func (o *TemplateSummaryLabelMapping) SetResourceTemplateMetaName(v string) {
	o.ResourceTemplateMetaName = v
}

// GetResourceName returns the ResourceName field value
func (o *TemplateSummaryLabelMapping) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryLabelMapping) GetResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value
func (o *TemplateSummaryLabelMapping) SetResourceName(v string) {
	o.ResourceName = v
}

// GetResourceID returns the ResourceID field value
func (o *TemplateSummaryLabelMapping) GetResourceID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceID
}

// GetResourceIDOk returns a tuple with the ResourceID field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryLabelMapping) GetResourceIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceID, true
}

// SetResourceID sets field value
func (o *TemplateSummaryLabelMapping) SetResourceID(v string) {
	o.ResourceID = v
}

// GetResourceType returns the ResourceType field value
func (o *TemplateSummaryLabelMapping) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryLabelMapping) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *TemplateSummaryLabelMapping) SetResourceType(v string) {
	o.ResourceType = v
}

// GetLabelTemplateMetaName returns the LabelTemplateMetaName field value
func (o *TemplateSummaryLabelMapping) GetLabelTemplateMetaName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LabelTemplateMetaName
}

// GetLabelTemplateMetaNameOk returns a tuple with the LabelTemplateMetaName field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryLabelMapping) GetLabelTemplateMetaNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelTemplateMetaName, true
}

// SetLabelTemplateMetaName sets field value
func (o *TemplateSummaryLabelMapping) SetLabelTemplateMetaName(v string) {
	o.LabelTemplateMetaName = v
}

// GetLabelName returns the LabelName field value
func (o *TemplateSummaryLabelMapping) GetLabelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LabelName
}

// GetLabelNameOk returns a tuple with the LabelName field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryLabelMapping) GetLabelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelName, true
}

// SetLabelName sets field value
func (o *TemplateSummaryLabelMapping) SetLabelName(v string) {
	o.LabelName = v
}

// GetLabelID returns the LabelID field value
func (o *TemplateSummaryLabelMapping) GetLabelID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LabelID
}

// GetLabelIDOk returns a tuple with the LabelID field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryLabelMapping) GetLabelIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelID, true
}

// SetLabelID sets field value
func (o *TemplateSummaryLabelMapping) SetLabelID(v string) {
	o.LabelID = v
}

func (o TemplateSummaryLabelMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["resourceTemplateMetaName"] = o.ResourceTemplateMetaName
	}
	if true {
		toSerialize["resourceName"] = o.ResourceName
	}
	if true {
		toSerialize["resourceID"] = o.ResourceID
	}
	if true {
		toSerialize["resourceType"] = o.ResourceType
	}
	if true {
		toSerialize["labelTemplateMetaName"] = o.LabelTemplateMetaName
	}
	if true {
		toSerialize["labelName"] = o.LabelName
	}
	if true {
		toSerialize["labelID"] = o.LabelID
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryLabelMapping struct {
	value *TemplateSummaryLabelMapping
	isSet bool
}

func (v NullableTemplateSummaryLabelMapping) Get() *TemplateSummaryLabelMapping {
	return v.value
}

func (v *NullableTemplateSummaryLabelMapping) Set(val *TemplateSummaryLabelMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryLabelMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryLabelMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryLabelMapping(val *TemplateSummaryLabelMapping) *NullableTemplateSummaryLabelMapping {
	return &NullableTemplateSummaryLabelMapping{value: val, isSet: true}
}

func (v NullableTemplateSummaryLabelMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryLabelMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
