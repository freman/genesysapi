// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutIntegrationsCredentialReader is a Reader for the PutIntegrationsCredential structure.
type PutIntegrationsCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutIntegrationsCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutIntegrationsCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutIntegrationsCredentialBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutIntegrationsCredentialUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutIntegrationsCredentialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutIntegrationsCredentialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutIntegrationsCredentialRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutIntegrationsCredentialRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutIntegrationsCredentialUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutIntegrationsCredentialTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutIntegrationsCredentialInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutIntegrationsCredentialServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutIntegrationsCredentialGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutIntegrationsCredentialOK creates a PutIntegrationsCredentialOK with default headers values
func NewPutIntegrationsCredentialOK() *PutIntegrationsCredentialOK {
	return &PutIntegrationsCredentialOK{}
}

/*
PutIntegrationsCredentialOK describes a response with status code 200, with default header values.

successful operation
*/
type PutIntegrationsCredentialOK struct {
	Payload *models.CredentialInfo
}

// IsSuccess returns true when this put integrations credential o k response has a 2xx status code
func (o *PutIntegrationsCredentialOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put integrations credential o k response has a 3xx status code
func (o *PutIntegrationsCredentialOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integrations credential o k response has a 4xx status code
func (o *PutIntegrationsCredentialOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put integrations credential o k response has a 5xx status code
func (o *PutIntegrationsCredentialOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put integrations credential o k response a status code equal to that given
func (o *PutIntegrationsCredentialOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutIntegrationsCredentialOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialOK  %+v", 200, o.Payload)
}

func (o *PutIntegrationsCredentialOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialOK  %+v", 200, o.Payload)
}

func (o *PutIntegrationsCredentialOK) GetPayload() *models.CredentialInfo {
	return o.Payload
}

func (o *PutIntegrationsCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationsCredentialBadRequest creates a PutIntegrationsCredentialBadRequest with default headers values
func NewPutIntegrationsCredentialBadRequest() *PutIntegrationsCredentialBadRequest {
	return &PutIntegrationsCredentialBadRequest{}
}

/*
PutIntegrationsCredentialBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutIntegrationsCredentialBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integrations credential bad request response has a 2xx status code
func (o *PutIntegrationsCredentialBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integrations credential bad request response has a 3xx status code
func (o *PutIntegrationsCredentialBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integrations credential bad request response has a 4xx status code
func (o *PutIntegrationsCredentialBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put integrations credential bad request response has a 5xx status code
func (o *PutIntegrationsCredentialBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put integrations credential bad request response a status code equal to that given
func (o *PutIntegrationsCredentialBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutIntegrationsCredentialBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialBadRequest  %+v", 400, o.Payload)
}

func (o *PutIntegrationsCredentialBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialBadRequest  %+v", 400, o.Payload)
}

func (o *PutIntegrationsCredentialBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationsCredentialBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationsCredentialUnauthorized creates a PutIntegrationsCredentialUnauthorized with default headers values
func NewPutIntegrationsCredentialUnauthorized() *PutIntegrationsCredentialUnauthorized {
	return &PutIntegrationsCredentialUnauthorized{}
}

/*
PutIntegrationsCredentialUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutIntegrationsCredentialUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integrations credential unauthorized response has a 2xx status code
func (o *PutIntegrationsCredentialUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integrations credential unauthorized response has a 3xx status code
func (o *PutIntegrationsCredentialUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integrations credential unauthorized response has a 4xx status code
func (o *PutIntegrationsCredentialUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put integrations credential unauthorized response has a 5xx status code
func (o *PutIntegrationsCredentialUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put integrations credential unauthorized response a status code equal to that given
func (o *PutIntegrationsCredentialUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutIntegrationsCredentialUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialUnauthorized  %+v", 401, o.Payload)
}

func (o *PutIntegrationsCredentialUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialUnauthorized  %+v", 401, o.Payload)
}

func (o *PutIntegrationsCredentialUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationsCredentialUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationsCredentialForbidden creates a PutIntegrationsCredentialForbidden with default headers values
func NewPutIntegrationsCredentialForbidden() *PutIntegrationsCredentialForbidden {
	return &PutIntegrationsCredentialForbidden{}
}

/*
PutIntegrationsCredentialForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutIntegrationsCredentialForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integrations credential forbidden response has a 2xx status code
func (o *PutIntegrationsCredentialForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integrations credential forbidden response has a 3xx status code
func (o *PutIntegrationsCredentialForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integrations credential forbidden response has a 4xx status code
func (o *PutIntegrationsCredentialForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put integrations credential forbidden response has a 5xx status code
func (o *PutIntegrationsCredentialForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put integrations credential forbidden response a status code equal to that given
func (o *PutIntegrationsCredentialForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutIntegrationsCredentialForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialForbidden  %+v", 403, o.Payload)
}

func (o *PutIntegrationsCredentialForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialForbidden  %+v", 403, o.Payload)
}

func (o *PutIntegrationsCredentialForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationsCredentialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationsCredentialNotFound creates a PutIntegrationsCredentialNotFound with default headers values
func NewPutIntegrationsCredentialNotFound() *PutIntegrationsCredentialNotFound {
	return &PutIntegrationsCredentialNotFound{}
}

/*
PutIntegrationsCredentialNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutIntegrationsCredentialNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integrations credential not found response has a 2xx status code
func (o *PutIntegrationsCredentialNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integrations credential not found response has a 3xx status code
func (o *PutIntegrationsCredentialNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integrations credential not found response has a 4xx status code
func (o *PutIntegrationsCredentialNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put integrations credential not found response has a 5xx status code
func (o *PutIntegrationsCredentialNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put integrations credential not found response a status code equal to that given
func (o *PutIntegrationsCredentialNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutIntegrationsCredentialNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialNotFound  %+v", 404, o.Payload)
}

func (o *PutIntegrationsCredentialNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialNotFound  %+v", 404, o.Payload)
}

func (o *PutIntegrationsCredentialNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationsCredentialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationsCredentialRequestTimeout creates a PutIntegrationsCredentialRequestTimeout with default headers values
func NewPutIntegrationsCredentialRequestTimeout() *PutIntegrationsCredentialRequestTimeout {
	return &PutIntegrationsCredentialRequestTimeout{}
}

/*
PutIntegrationsCredentialRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutIntegrationsCredentialRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integrations credential request timeout response has a 2xx status code
func (o *PutIntegrationsCredentialRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integrations credential request timeout response has a 3xx status code
func (o *PutIntegrationsCredentialRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integrations credential request timeout response has a 4xx status code
func (o *PutIntegrationsCredentialRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put integrations credential request timeout response has a 5xx status code
func (o *PutIntegrationsCredentialRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put integrations credential request timeout response a status code equal to that given
func (o *PutIntegrationsCredentialRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutIntegrationsCredentialRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutIntegrationsCredentialRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutIntegrationsCredentialRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationsCredentialRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationsCredentialRequestEntityTooLarge creates a PutIntegrationsCredentialRequestEntityTooLarge with default headers values
func NewPutIntegrationsCredentialRequestEntityTooLarge() *PutIntegrationsCredentialRequestEntityTooLarge {
	return &PutIntegrationsCredentialRequestEntityTooLarge{}
}

/*
PutIntegrationsCredentialRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutIntegrationsCredentialRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integrations credential request entity too large response has a 2xx status code
func (o *PutIntegrationsCredentialRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integrations credential request entity too large response has a 3xx status code
func (o *PutIntegrationsCredentialRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integrations credential request entity too large response has a 4xx status code
func (o *PutIntegrationsCredentialRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put integrations credential request entity too large response has a 5xx status code
func (o *PutIntegrationsCredentialRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put integrations credential request entity too large response a status code equal to that given
func (o *PutIntegrationsCredentialRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutIntegrationsCredentialRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutIntegrationsCredentialRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutIntegrationsCredentialRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationsCredentialRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationsCredentialUnsupportedMediaType creates a PutIntegrationsCredentialUnsupportedMediaType with default headers values
func NewPutIntegrationsCredentialUnsupportedMediaType() *PutIntegrationsCredentialUnsupportedMediaType {
	return &PutIntegrationsCredentialUnsupportedMediaType{}
}

/*
PutIntegrationsCredentialUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutIntegrationsCredentialUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integrations credential unsupported media type response has a 2xx status code
func (o *PutIntegrationsCredentialUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integrations credential unsupported media type response has a 3xx status code
func (o *PutIntegrationsCredentialUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integrations credential unsupported media type response has a 4xx status code
func (o *PutIntegrationsCredentialUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put integrations credential unsupported media type response has a 5xx status code
func (o *PutIntegrationsCredentialUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put integrations credential unsupported media type response a status code equal to that given
func (o *PutIntegrationsCredentialUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutIntegrationsCredentialUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutIntegrationsCredentialUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutIntegrationsCredentialUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationsCredentialUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationsCredentialTooManyRequests creates a PutIntegrationsCredentialTooManyRequests with default headers values
func NewPutIntegrationsCredentialTooManyRequests() *PutIntegrationsCredentialTooManyRequests {
	return &PutIntegrationsCredentialTooManyRequests{}
}

/*
PutIntegrationsCredentialTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutIntegrationsCredentialTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integrations credential too many requests response has a 2xx status code
func (o *PutIntegrationsCredentialTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integrations credential too many requests response has a 3xx status code
func (o *PutIntegrationsCredentialTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integrations credential too many requests response has a 4xx status code
func (o *PutIntegrationsCredentialTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put integrations credential too many requests response has a 5xx status code
func (o *PutIntegrationsCredentialTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put integrations credential too many requests response a status code equal to that given
func (o *PutIntegrationsCredentialTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutIntegrationsCredentialTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutIntegrationsCredentialTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutIntegrationsCredentialTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationsCredentialTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationsCredentialInternalServerError creates a PutIntegrationsCredentialInternalServerError with default headers values
func NewPutIntegrationsCredentialInternalServerError() *PutIntegrationsCredentialInternalServerError {
	return &PutIntegrationsCredentialInternalServerError{}
}

/*
PutIntegrationsCredentialInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutIntegrationsCredentialInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integrations credential internal server error response has a 2xx status code
func (o *PutIntegrationsCredentialInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integrations credential internal server error response has a 3xx status code
func (o *PutIntegrationsCredentialInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integrations credential internal server error response has a 4xx status code
func (o *PutIntegrationsCredentialInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put integrations credential internal server error response has a 5xx status code
func (o *PutIntegrationsCredentialInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put integrations credential internal server error response a status code equal to that given
func (o *PutIntegrationsCredentialInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutIntegrationsCredentialInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialInternalServerError  %+v", 500, o.Payload)
}

func (o *PutIntegrationsCredentialInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialInternalServerError  %+v", 500, o.Payload)
}

func (o *PutIntegrationsCredentialInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationsCredentialInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationsCredentialServiceUnavailable creates a PutIntegrationsCredentialServiceUnavailable with default headers values
func NewPutIntegrationsCredentialServiceUnavailable() *PutIntegrationsCredentialServiceUnavailable {
	return &PutIntegrationsCredentialServiceUnavailable{}
}

/*
PutIntegrationsCredentialServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutIntegrationsCredentialServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integrations credential service unavailable response has a 2xx status code
func (o *PutIntegrationsCredentialServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integrations credential service unavailable response has a 3xx status code
func (o *PutIntegrationsCredentialServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integrations credential service unavailable response has a 4xx status code
func (o *PutIntegrationsCredentialServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put integrations credential service unavailable response has a 5xx status code
func (o *PutIntegrationsCredentialServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put integrations credential service unavailable response a status code equal to that given
func (o *PutIntegrationsCredentialServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutIntegrationsCredentialServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutIntegrationsCredentialServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutIntegrationsCredentialServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationsCredentialServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationsCredentialGatewayTimeout creates a PutIntegrationsCredentialGatewayTimeout with default headers values
func NewPutIntegrationsCredentialGatewayTimeout() *PutIntegrationsCredentialGatewayTimeout {
	return &PutIntegrationsCredentialGatewayTimeout{}
}

/*
PutIntegrationsCredentialGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutIntegrationsCredentialGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integrations credential gateway timeout response has a 2xx status code
func (o *PutIntegrationsCredentialGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integrations credential gateway timeout response has a 3xx status code
func (o *PutIntegrationsCredentialGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integrations credential gateway timeout response has a 4xx status code
func (o *PutIntegrationsCredentialGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put integrations credential gateway timeout response has a 5xx status code
func (o *PutIntegrationsCredentialGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put integrations credential gateway timeout response a status code equal to that given
func (o *PutIntegrationsCredentialGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutIntegrationsCredentialGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutIntegrationsCredentialGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutIntegrationsCredentialGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationsCredentialGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
