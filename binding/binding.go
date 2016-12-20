// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import "net/http"

const (
	MIMEJSON              = "application/json"
	MIMEHTML              = "text/html"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
)

type Binding interface {
	Name() string
	Bind(*http.Request, interface{}) error
}

type StructValidator interface {
	// ValidateStruct can receive any kind of type and it should never panic, even if the configuration is not right.
	// If the received type is not a struct, any validation should be skipped and nil must be returned.
	// If the received type is a struct or pointer to a struct, the validation should be performed.
	// If the struct is not valid or the validation itself fails, a descriptive error should be returned.
	// Otherwise nil must be returned.
	ValidateStruct(interface{}) error
}

var Validator StructValidator = nil

var (
	JSON          = jsonBinding{}
	Form          = formBinding{}
	FormPost      = formPostBinding{}
	FormMultipart = formMultipartBinding{}
)

func Default(method, contentType string) Binding {
	if method == "GET" {
		return Form
	} else {
		switch contentType {
		case MIMEJSON:
			return JSON
		default: //case MIMEPOSTForm, MIMEMultipartPOSTForm:
			return Form
		}
	}
}

func validate(obj interface{}) error {
	if Validator == nil {
		return nil
	}
	return Validator.ValidateStruct(obj)
}
