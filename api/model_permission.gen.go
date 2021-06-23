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

// Permission struct for Permission
type Permission struct {
	Action   string             `json:"action"`
	Resource PermissionResource `json:"resource"`
}

// NewPermission instantiates a new Permission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermission(action string, resource PermissionResource) *Permission {
	this := Permission{}
	this.Action = action
	this.Resource = resource
	return &this
}

// NewPermissionWithDefaults instantiates a new Permission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionWithDefaults() *Permission {
	this := Permission{}
	return &this
}

// GetAction returns the Action field value
func (o *Permission) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *Permission) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *Permission) SetAction(v string) {
	o.Action = v
}

// GetResource returns the Resource field value
func (o *Permission) GetResource() PermissionResource {
	if o == nil {
		var ret PermissionResource
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *Permission) GetResourceOk() (*PermissionResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *Permission) SetResource(v PermissionResource) {
	o.Resource = v
}

func (o Permission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullablePermission struct {
	value *Permission
	isSet bool
}

func (v NullablePermission) Get() *Permission {
	return v.value
}

func (v *NullablePermission) Set(val *Permission) {
	v.value = val
	v.isSet = true
}

func (v NullablePermission) IsSet() bool {
	return v.isSet
}

func (v *NullablePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermission(val *Permission) *NullablePermission {
	return &NullablePermission{value: val, isSet: true}
}

func (v NullablePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}