// Code generated by go-swagger; DO NOT EDIT.

package identity_provider

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutIdentityprovidersPureengageReader is a Reader for the PutIdentityprovidersPureengage structure.
type PutIdentityprovidersPureengageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutIdentityprovidersPureengageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutIdentityprovidersPureengageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutIdentityprovidersPureengageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutIdentityprovidersPureengageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutIdentityprovidersPureengageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutIdentityprovidersPureengageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutIdentityprovidersPureengageRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutIdentityprovidersPureengageRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutIdentityprovidersPureengageUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutIdentityprovidersPureengageTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutIdentityprovidersPureengageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutIdentityprovidersPureengageServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutIdentityprovidersPureengageGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutIdentityprovidersPureengageOK creates a PutIdentityprovidersPureengageOK with default headers values
func NewPutIdentityprovidersPureengageOK() *PutIdentityprovidersPureengageOK {
	return &PutIdentityprovidersPureengageOK{}
}

/*
PutIdentityprovidersPureengageOK describes a response with status code 200, with default header values.

successful operation
*/
type PutIdentityprovidersPureengageOK struct {
	Payload *models.OAuthProvider
}

// IsSuccess returns true when this put identityproviders pureengage o k response has a 2xx status code
func (o *PutIdentityprovidersPureengageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put identityproviders pureengage o k response has a 3xx status code
func (o *PutIdentityprovidersPureengageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders pureengage o k response has a 4xx status code
func (o *PutIdentityprovidersPureengageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put identityproviders pureengage o k response has a 5xx status code
func (o *PutIdentityprovidersPureengageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders pureengage o k response a status code equal to that given
func (o *PutIdentityprovidersPureengageOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutIdentityprovidersPureengageOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageOK  %+v", 200, o.Payload)
}

func (o *PutIdentityprovidersPureengageOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageOK  %+v", 200, o.Payload)
}

func (o *PutIdentityprovidersPureengageOK) GetPayload() *models.OAuthProvider {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuthProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPureengageBadRequest creates a PutIdentityprovidersPureengageBadRequest with default headers values
func NewPutIdentityprovidersPureengageBadRequest() *PutIdentityprovidersPureengageBadRequest {
	return &PutIdentityprovidersPureengageBadRequest{}
}

/*
PutIdentityprovidersPureengageBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutIdentityprovidersPureengageBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders pureengage bad request response has a 2xx status code
func (o *PutIdentityprovidersPureengageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders pureengage bad request response has a 3xx status code
func (o *PutIdentityprovidersPureengageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders pureengage bad request response has a 4xx status code
func (o *PutIdentityprovidersPureengageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders pureengage bad request response has a 5xx status code
func (o *PutIdentityprovidersPureengageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders pureengage bad request response a status code equal to that given
func (o *PutIdentityprovidersPureengageBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutIdentityprovidersPureengageBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageBadRequest  %+v", 400, o.Payload)
}

func (o *PutIdentityprovidersPureengageBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageBadRequest  %+v", 400, o.Payload)
}

func (o *PutIdentityprovidersPureengageBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPureengageUnauthorized creates a PutIdentityprovidersPureengageUnauthorized with default headers values
func NewPutIdentityprovidersPureengageUnauthorized() *PutIdentityprovidersPureengageUnauthorized {
	return &PutIdentityprovidersPureengageUnauthorized{}
}

/*
PutIdentityprovidersPureengageUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutIdentityprovidersPureengageUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders pureengage unauthorized response has a 2xx status code
func (o *PutIdentityprovidersPureengageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders pureengage unauthorized response has a 3xx status code
func (o *PutIdentityprovidersPureengageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders pureengage unauthorized response has a 4xx status code
func (o *PutIdentityprovidersPureengageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders pureengage unauthorized response has a 5xx status code
func (o *PutIdentityprovidersPureengageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders pureengage unauthorized response a status code equal to that given
func (o *PutIdentityprovidersPureengageUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutIdentityprovidersPureengageUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageUnauthorized  %+v", 401, o.Payload)
}

func (o *PutIdentityprovidersPureengageUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageUnauthorized  %+v", 401, o.Payload)
}

func (o *PutIdentityprovidersPureengageUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPureengageForbidden creates a PutIdentityprovidersPureengageForbidden with default headers values
func NewPutIdentityprovidersPureengageForbidden() *PutIdentityprovidersPureengageForbidden {
	return &PutIdentityprovidersPureengageForbidden{}
}

/*
PutIdentityprovidersPureengageForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutIdentityprovidersPureengageForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders pureengage forbidden response has a 2xx status code
func (o *PutIdentityprovidersPureengageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders pureengage forbidden response has a 3xx status code
func (o *PutIdentityprovidersPureengageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders pureengage forbidden response has a 4xx status code
func (o *PutIdentityprovidersPureengageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders pureengage forbidden response has a 5xx status code
func (o *PutIdentityprovidersPureengageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders pureengage forbidden response a status code equal to that given
func (o *PutIdentityprovidersPureengageForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutIdentityprovidersPureengageForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageForbidden  %+v", 403, o.Payload)
}

func (o *PutIdentityprovidersPureengageForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageForbidden  %+v", 403, o.Payload)
}

func (o *PutIdentityprovidersPureengageForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPureengageNotFound creates a PutIdentityprovidersPureengageNotFound with default headers values
func NewPutIdentityprovidersPureengageNotFound() *PutIdentityprovidersPureengageNotFound {
	return &PutIdentityprovidersPureengageNotFound{}
}

/*
PutIdentityprovidersPureengageNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutIdentityprovidersPureengageNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders pureengage not found response has a 2xx status code
func (o *PutIdentityprovidersPureengageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders pureengage not found response has a 3xx status code
func (o *PutIdentityprovidersPureengageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders pureengage not found response has a 4xx status code
func (o *PutIdentityprovidersPureengageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders pureengage not found response has a 5xx status code
func (o *PutIdentityprovidersPureengageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders pureengage not found response a status code equal to that given
func (o *PutIdentityprovidersPureengageNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutIdentityprovidersPureengageNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageNotFound  %+v", 404, o.Payload)
}

func (o *PutIdentityprovidersPureengageNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageNotFound  %+v", 404, o.Payload)
}

func (o *PutIdentityprovidersPureengageNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPureengageRequestTimeout creates a PutIdentityprovidersPureengageRequestTimeout with default headers values
func NewPutIdentityprovidersPureengageRequestTimeout() *PutIdentityprovidersPureengageRequestTimeout {
	return &PutIdentityprovidersPureengageRequestTimeout{}
}

/*
PutIdentityprovidersPureengageRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutIdentityprovidersPureengageRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders pureengage request timeout response has a 2xx status code
func (o *PutIdentityprovidersPureengageRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders pureengage request timeout response has a 3xx status code
func (o *PutIdentityprovidersPureengageRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders pureengage request timeout response has a 4xx status code
func (o *PutIdentityprovidersPureengageRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders pureengage request timeout response has a 5xx status code
func (o *PutIdentityprovidersPureengageRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders pureengage request timeout response a status code equal to that given
func (o *PutIdentityprovidersPureengageRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutIdentityprovidersPureengageRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutIdentityprovidersPureengageRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutIdentityprovidersPureengageRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPureengageRequestEntityTooLarge creates a PutIdentityprovidersPureengageRequestEntityTooLarge with default headers values
func NewPutIdentityprovidersPureengageRequestEntityTooLarge() *PutIdentityprovidersPureengageRequestEntityTooLarge {
	return &PutIdentityprovidersPureengageRequestEntityTooLarge{}
}

/*
PutIdentityprovidersPureengageRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutIdentityprovidersPureengageRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders pureengage request entity too large response has a 2xx status code
func (o *PutIdentityprovidersPureengageRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders pureengage request entity too large response has a 3xx status code
func (o *PutIdentityprovidersPureengageRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders pureengage request entity too large response has a 4xx status code
func (o *PutIdentityprovidersPureengageRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders pureengage request entity too large response has a 5xx status code
func (o *PutIdentityprovidersPureengageRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders pureengage request entity too large response a status code equal to that given
func (o *PutIdentityprovidersPureengageRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutIdentityprovidersPureengageRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutIdentityprovidersPureengageRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutIdentityprovidersPureengageRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPureengageUnsupportedMediaType creates a PutIdentityprovidersPureengageUnsupportedMediaType with default headers values
func NewPutIdentityprovidersPureengageUnsupportedMediaType() *PutIdentityprovidersPureengageUnsupportedMediaType {
	return &PutIdentityprovidersPureengageUnsupportedMediaType{}
}

/*
PutIdentityprovidersPureengageUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutIdentityprovidersPureengageUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders pureengage unsupported media type response has a 2xx status code
func (o *PutIdentityprovidersPureengageUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders pureengage unsupported media type response has a 3xx status code
func (o *PutIdentityprovidersPureengageUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders pureengage unsupported media type response has a 4xx status code
func (o *PutIdentityprovidersPureengageUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders pureengage unsupported media type response has a 5xx status code
func (o *PutIdentityprovidersPureengageUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders pureengage unsupported media type response a status code equal to that given
func (o *PutIdentityprovidersPureengageUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutIdentityprovidersPureengageUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutIdentityprovidersPureengageUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutIdentityprovidersPureengageUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPureengageTooManyRequests creates a PutIdentityprovidersPureengageTooManyRequests with default headers values
func NewPutIdentityprovidersPureengageTooManyRequests() *PutIdentityprovidersPureengageTooManyRequests {
	return &PutIdentityprovidersPureengageTooManyRequests{}
}

/*
PutIdentityprovidersPureengageTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutIdentityprovidersPureengageTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders pureengage too many requests response has a 2xx status code
func (o *PutIdentityprovidersPureengageTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders pureengage too many requests response has a 3xx status code
func (o *PutIdentityprovidersPureengageTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders pureengage too many requests response has a 4xx status code
func (o *PutIdentityprovidersPureengageTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders pureengage too many requests response has a 5xx status code
func (o *PutIdentityprovidersPureengageTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders pureengage too many requests response a status code equal to that given
func (o *PutIdentityprovidersPureengageTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutIdentityprovidersPureengageTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutIdentityprovidersPureengageTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutIdentityprovidersPureengageTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPureengageInternalServerError creates a PutIdentityprovidersPureengageInternalServerError with default headers values
func NewPutIdentityprovidersPureengageInternalServerError() *PutIdentityprovidersPureengageInternalServerError {
	return &PutIdentityprovidersPureengageInternalServerError{}
}

/*
PutIdentityprovidersPureengageInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutIdentityprovidersPureengageInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders pureengage internal server error response has a 2xx status code
func (o *PutIdentityprovidersPureengageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders pureengage internal server error response has a 3xx status code
func (o *PutIdentityprovidersPureengageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders pureengage internal server error response has a 4xx status code
func (o *PutIdentityprovidersPureengageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put identityproviders pureengage internal server error response has a 5xx status code
func (o *PutIdentityprovidersPureengageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put identityproviders pureengage internal server error response a status code equal to that given
func (o *PutIdentityprovidersPureengageInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutIdentityprovidersPureengageInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageInternalServerError  %+v", 500, o.Payload)
}

func (o *PutIdentityprovidersPureengageInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageInternalServerError  %+v", 500, o.Payload)
}

func (o *PutIdentityprovidersPureengageInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPureengageServiceUnavailable creates a PutIdentityprovidersPureengageServiceUnavailable with default headers values
func NewPutIdentityprovidersPureengageServiceUnavailable() *PutIdentityprovidersPureengageServiceUnavailable {
	return &PutIdentityprovidersPureengageServiceUnavailable{}
}

/*
PutIdentityprovidersPureengageServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutIdentityprovidersPureengageServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders pureengage service unavailable response has a 2xx status code
func (o *PutIdentityprovidersPureengageServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders pureengage service unavailable response has a 3xx status code
func (o *PutIdentityprovidersPureengageServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders pureengage service unavailable response has a 4xx status code
func (o *PutIdentityprovidersPureengageServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put identityproviders pureengage service unavailable response has a 5xx status code
func (o *PutIdentityprovidersPureengageServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put identityproviders pureengage service unavailable response a status code equal to that given
func (o *PutIdentityprovidersPureengageServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutIdentityprovidersPureengageServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutIdentityprovidersPureengageServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutIdentityprovidersPureengageServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPureengageGatewayTimeout creates a PutIdentityprovidersPureengageGatewayTimeout with default headers values
func NewPutIdentityprovidersPureengageGatewayTimeout() *PutIdentityprovidersPureengageGatewayTimeout {
	return &PutIdentityprovidersPureengageGatewayTimeout{}
}

/*
PutIdentityprovidersPureengageGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutIdentityprovidersPureengageGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders pureengage gateway timeout response has a 2xx status code
func (o *PutIdentityprovidersPureengageGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders pureengage gateway timeout response has a 3xx status code
func (o *PutIdentityprovidersPureengageGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders pureengage gateway timeout response has a 4xx status code
func (o *PutIdentityprovidersPureengageGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put identityproviders pureengage gateway timeout response has a 5xx status code
func (o *PutIdentityprovidersPureengageGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put identityproviders pureengage gateway timeout response a status code equal to that given
func (o *PutIdentityprovidersPureengageGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutIdentityprovidersPureengageGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutIdentityprovidersPureengageGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutIdentityprovidersPureengageGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
