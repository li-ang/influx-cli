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

// AuthorizationPostRequestAllOf struct for AuthorizationPostRequestAllOf
type AuthorizationPostRequestAllOf struct {
	// ID of org that authorization is scoped to.
	OrgID *string `json:"orgID,omitempty" yaml:"orgID,omitempty"`
	// ID of user that authorization is scoped to.
	UserID *string `json:"userID,omitempty" yaml:"userID,omitempty"`
	// List of permissions for an auth.  An auth must have at least one Permission.
	Permissions *[]Permission `json:"permissions,omitempty" yaml:"permissions,omitempty"`
}

// NewAuthorizationPostRequestAllOf instantiates a new AuthorizationPostRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationPostRequestAllOf() *AuthorizationPostRequestAllOf {
	this := AuthorizationPostRequestAllOf{}
	return &this
}

// NewAuthorizationPostRequestAllOfWithDefaults instantiates a new AuthorizationPostRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationPostRequestAllOfWithDefaults() *AuthorizationPostRequestAllOf {
	this := AuthorizationPostRequestAllOf{}
	return &this
}

// GetOrgID returns the OrgID field value if set, zero value otherwise.
func (o *AuthorizationPostRequestAllOf) GetOrgID() string {
	if o == nil || o.OrgID == nil {
		var ret string
		return ret
	}
	return *o.OrgID
}

// GetOrgIDOk returns a tuple with the OrgID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationPostRequestAllOf) GetOrgIDOk() (*string, bool) {
	if o == nil || o.OrgID == nil {
		return nil, false
	}
	return o.OrgID, true
}

// HasOrgID returns a boolean if a field has been set.
func (o *AuthorizationPostRequestAllOf) HasOrgID() bool {
	if o != nil && o.OrgID != nil {
		return true
	}

	return false
}

// SetOrgID gets a reference to the given string and assigns it to the OrgID field.
func (o *AuthorizationPostRequestAllOf) SetOrgID(v string) {
	o.OrgID = &v
}

// GetUserID returns the UserID field value if set, zero value otherwise.
func (o *AuthorizationPostRequestAllOf) GetUserID() string {
	if o == nil || o.UserID == nil {
		var ret string
		return ret
	}
	return *o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationPostRequestAllOf) GetUserIDOk() (*string, bool) {
	if o == nil || o.UserID == nil {
		return nil, false
	}
	return o.UserID, true
}

// HasUserID returns a boolean if a field has been set.
func (o *AuthorizationPostRequestAllOf) HasUserID() bool {
	if o != nil && o.UserID != nil {
		return true
	}

	return false
}

// SetUserID gets a reference to the given string and assigns it to the UserID field.
func (o *AuthorizationPostRequestAllOf) SetUserID(v string) {
	o.UserID = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *AuthorizationPostRequestAllOf) GetPermissions() []Permission {
	if o == nil || o.Permissions == nil {
		var ret []Permission
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationPostRequestAllOf) GetPermissionsOk() (*[]Permission, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *AuthorizationPostRequestAllOf) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []Permission and assigns it to the Permissions field.
func (o *AuthorizationPostRequestAllOf) SetPermissions(v []Permission) {
	o.Permissions = &v
}

func (o AuthorizationPostRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OrgID != nil {
		toSerialize["orgID"] = o.OrgID
	}
	if o.UserID != nil {
		toSerialize["userID"] = o.UserID
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizationPostRequestAllOf struct {
	value *AuthorizationPostRequestAllOf
	isSet bool
}

func (v NullableAuthorizationPostRequestAllOf) Get() *AuthorizationPostRequestAllOf {
	return v.value
}

func (v *NullableAuthorizationPostRequestAllOf) Set(val *AuthorizationPostRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationPostRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationPostRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationPostRequestAllOf(val *AuthorizationPostRequestAllOf) *NullableAuthorizationPostRequestAllOf {
	return &NullableAuthorizationPostRequestAllOf{value: val, isSet: true}
}

func (v NullableAuthorizationPostRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationPostRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
