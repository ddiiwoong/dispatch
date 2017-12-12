///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package secret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/secret-store/gen/models"
)

// AddSecretCreatedCode is the HTTP code returned for type AddSecretCreated
const AddSecretCreatedCode int = 201

/*AddSecretCreated The created secret.

swagger:response addSecretCreated
*/
type AddSecretCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Secret `json:"body,omitempty"`
}

// NewAddSecretCreated creates AddSecretCreated with default headers values
func NewAddSecretCreated() *AddSecretCreated {
	return &AddSecretCreated{}
}

// WithPayload adds the payload to the add secret created response
func (o *AddSecretCreated) WithPayload(payload *models.Secret) *AddSecretCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add secret created response
func (o *AddSecretCreated) SetPayload(payload *models.Secret) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddSecretCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AddSecretDefault Standard error

swagger:response addSecretDefault
*/
type AddSecretDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddSecretDefault creates AddSecretDefault with default headers values
func NewAddSecretDefault(code int) *AddSecretDefault {
	if code <= 0 {
		code = 500
	}

	return &AddSecretDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add secret default response
func (o *AddSecretDefault) WithStatusCode(code int) *AddSecretDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add secret default response
func (o *AddSecretDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add secret default response
func (o *AddSecretDefault) WithPayload(payload *models.Error) *AddSecretDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add secret default response
func (o *AddSecretDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddSecretDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}