// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutArchitectPromptResourceReader is a Reader for the PutArchitectPromptResource structure.
type PutArchitectPromptResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutArchitectPromptResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutArchitectPromptResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutArchitectPromptResourceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutArchitectPromptResourceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutArchitectPromptResourceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutArchitectPromptResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutArchitectPromptResourceRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutArchitectPromptResourceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutArchitectPromptResourceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutArchitectPromptResourceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutArchitectPromptResourceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutArchitectPromptResourceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutArchitectPromptResourceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutArchitectPromptResourceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutArchitectPromptResourceOK creates a PutArchitectPromptResourceOK with default headers values
func NewPutArchitectPromptResourceOK() *PutArchitectPromptResourceOK {
	return &PutArchitectPromptResourceOK{}
}

/*
PutArchitectPromptResourceOK describes a response with status code 200, with default header values.

successful operation
*/
type PutArchitectPromptResourceOK struct {
	Payload *models.PromptAsset
}

// IsSuccess returns true when this put architect prompt resource o k response has a 2xx status code
func (o *PutArchitectPromptResourceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put architect prompt resource o k response has a 3xx status code
func (o *PutArchitectPromptResourceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect prompt resource o k response has a 4xx status code
func (o *PutArchitectPromptResourceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put architect prompt resource o k response has a 5xx status code
func (o *PutArchitectPromptResourceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put architect prompt resource o k response a status code equal to that given
func (o *PutArchitectPromptResourceOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutArchitectPromptResourceOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceOK  %+v", 200, o.Payload)
}

func (o *PutArchitectPromptResourceOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceOK  %+v", 200, o.Payload)
}

func (o *PutArchitectPromptResourceOK) GetPayload() *models.PromptAsset {
	return o.Payload
}

func (o *PutArchitectPromptResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PromptAsset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectPromptResourceBadRequest creates a PutArchitectPromptResourceBadRequest with default headers values
func NewPutArchitectPromptResourceBadRequest() *PutArchitectPromptResourceBadRequest {
	return &PutArchitectPromptResourceBadRequest{}
}

/*
PutArchitectPromptResourceBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutArchitectPromptResourceBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect prompt resource bad request response has a 2xx status code
func (o *PutArchitectPromptResourceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect prompt resource bad request response has a 3xx status code
func (o *PutArchitectPromptResourceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect prompt resource bad request response has a 4xx status code
func (o *PutArchitectPromptResourceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put architect prompt resource bad request response has a 5xx status code
func (o *PutArchitectPromptResourceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put architect prompt resource bad request response a status code equal to that given
func (o *PutArchitectPromptResourceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutArchitectPromptResourceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceBadRequest  %+v", 400, o.Payload)
}

func (o *PutArchitectPromptResourceBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceBadRequest  %+v", 400, o.Payload)
}

func (o *PutArchitectPromptResourceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectPromptResourceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectPromptResourceUnauthorized creates a PutArchitectPromptResourceUnauthorized with default headers values
func NewPutArchitectPromptResourceUnauthorized() *PutArchitectPromptResourceUnauthorized {
	return &PutArchitectPromptResourceUnauthorized{}
}

/*
PutArchitectPromptResourceUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutArchitectPromptResourceUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect prompt resource unauthorized response has a 2xx status code
func (o *PutArchitectPromptResourceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect prompt resource unauthorized response has a 3xx status code
func (o *PutArchitectPromptResourceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect prompt resource unauthorized response has a 4xx status code
func (o *PutArchitectPromptResourceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put architect prompt resource unauthorized response has a 5xx status code
func (o *PutArchitectPromptResourceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put architect prompt resource unauthorized response a status code equal to that given
func (o *PutArchitectPromptResourceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutArchitectPromptResourceUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceUnauthorized  %+v", 401, o.Payload)
}

func (o *PutArchitectPromptResourceUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceUnauthorized  %+v", 401, o.Payload)
}

func (o *PutArchitectPromptResourceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectPromptResourceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectPromptResourceForbidden creates a PutArchitectPromptResourceForbidden with default headers values
func NewPutArchitectPromptResourceForbidden() *PutArchitectPromptResourceForbidden {
	return &PutArchitectPromptResourceForbidden{}
}

/*
PutArchitectPromptResourceForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutArchitectPromptResourceForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect prompt resource forbidden response has a 2xx status code
func (o *PutArchitectPromptResourceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect prompt resource forbidden response has a 3xx status code
func (o *PutArchitectPromptResourceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect prompt resource forbidden response has a 4xx status code
func (o *PutArchitectPromptResourceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put architect prompt resource forbidden response has a 5xx status code
func (o *PutArchitectPromptResourceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put architect prompt resource forbidden response a status code equal to that given
func (o *PutArchitectPromptResourceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutArchitectPromptResourceForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceForbidden  %+v", 403, o.Payload)
}

func (o *PutArchitectPromptResourceForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceForbidden  %+v", 403, o.Payload)
}

func (o *PutArchitectPromptResourceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectPromptResourceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectPromptResourceNotFound creates a PutArchitectPromptResourceNotFound with default headers values
func NewPutArchitectPromptResourceNotFound() *PutArchitectPromptResourceNotFound {
	return &PutArchitectPromptResourceNotFound{}
}

/*
PutArchitectPromptResourceNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutArchitectPromptResourceNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect prompt resource not found response has a 2xx status code
func (o *PutArchitectPromptResourceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect prompt resource not found response has a 3xx status code
func (o *PutArchitectPromptResourceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect prompt resource not found response has a 4xx status code
func (o *PutArchitectPromptResourceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put architect prompt resource not found response has a 5xx status code
func (o *PutArchitectPromptResourceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put architect prompt resource not found response a status code equal to that given
func (o *PutArchitectPromptResourceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutArchitectPromptResourceNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceNotFound  %+v", 404, o.Payload)
}

func (o *PutArchitectPromptResourceNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceNotFound  %+v", 404, o.Payload)
}

func (o *PutArchitectPromptResourceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectPromptResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectPromptResourceRequestTimeout creates a PutArchitectPromptResourceRequestTimeout with default headers values
func NewPutArchitectPromptResourceRequestTimeout() *PutArchitectPromptResourceRequestTimeout {
	return &PutArchitectPromptResourceRequestTimeout{}
}

/*
PutArchitectPromptResourceRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutArchitectPromptResourceRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect prompt resource request timeout response has a 2xx status code
func (o *PutArchitectPromptResourceRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect prompt resource request timeout response has a 3xx status code
func (o *PutArchitectPromptResourceRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect prompt resource request timeout response has a 4xx status code
func (o *PutArchitectPromptResourceRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put architect prompt resource request timeout response has a 5xx status code
func (o *PutArchitectPromptResourceRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put architect prompt resource request timeout response a status code equal to that given
func (o *PutArchitectPromptResourceRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutArchitectPromptResourceRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutArchitectPromptResourceRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutArchitectPromptResourceRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectPromptResourceRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectPromptResourceConflict creates a PutArchitectPromptResourceConflict with default headers values
func NewPutArchitectPromptResourceConflict() *PutArchitectPromptResourceConflict {
	return &PutArchitectPromptResourceConflict{}
}

/*
PutArchitectPromptResourceConflict describes a response with status code 409, with default header values.

Conflict
*/
type PutArchitectPromptResourceConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect prompt resource conflict response has a 2xx status code
func (o *PutArchitectPromptResourceConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect prompt resource conflict response has a 3xx status code
func (o *PutArchitectPromptResourceConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect prompt resource conflict response has a 4xx status code
func (o *PutArchitectPromptResourceConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this put architect prompt resource conflict response has a 5xx status code
func (o *PutArchitectPromptResourceConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this put architect prompt resource conflict response a status code equal to that given
func (o *PutArchitectPromptResourceConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PutArchitectPromptResourceConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceConflict  %+v", 409, o.Payload)
}

func (o *PutArchitectPromptResourceConflict) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceConflict  %+v", 409, o.Payload)
}

func (o *PutArchitectPromptResourceConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectPromptResourceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectPromptResourceRequestEntityTooLarge creates a PutArchitectPromptResourceRequestEntityTooLarge with default headers values
func NewPutArchitectPromptResourceRequestEntityTooLarge() *PutArchitectPromptResourceRequestEntityTooLarge {
	return &PutArchitectPromptResourceRequestEntityTooLarge{}
}

/*
PutArchitectPromptResourceRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutArchitectPromptResourceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect prompt resource request entity too large response has a 2xx status code
func (o *PutArchitectPromptResourceRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect prompt resource request entity too large response has a 3xx status code
func (o *PutArchitectPromptResourceRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect prompt resource request entity too large response has a 4xx status code
func (o *PutArchitectPromptResourceRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put architect prompt resource request entity too large response has a 5xx status code
func (o *PutArchitectPromptResourceRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put architect prompt resource request entity too large response a status code equal to that given
func (o *PutArchitectPromptResourceRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutArchitectPromptResourceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutArchitectPromptResourceRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutArchitectPromptResourceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectPromptResourceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectPromptResourceUnsupportedMediaType creates a PutArchitectPromptResourceUnsupportedMediaType with default headers values
func NewPutArchitectPromptResourceUnsupportedMediaType() *PutArchitectPromptResourceUnsupportedMediaType {
	return &PutArchitectPromptResourceUnsupportedMediaType{}
}

/*
PutArchitectPromptResourceUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutArchitectPromptResourceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect prompt resource unsupported media type response has a 2xx status code
func (o *PutArchitectPromptResourceUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect prompt resource unsupported media type response has a 3xx status code
func (o *PutArchitectPromptResourceUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect prompt resource unsupported media type response has a 4xx status code
func (o *PutArchitectPromptResourceUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put architect prompt resource unsupported media type response has a 5xx status code
func (o *PutArchitectPromptResourceUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put architect prompt resource unsupported media type response a status code equal to that given
func (o *PutArchitectPromptResourceUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutArchitectPromptResourceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutArchitectPromptResourceUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutArchitectPromptResourceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectPromptResourceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectPromptResourceTooManyRequests creates a PutArchitectPromptResourceTooManyRequests with default headers values
func NewPutArchitectPromptResourceTooManyRequests() *PutArchitectPromptResourceTooManyRequests {
	return &PutArchitectPromptResourceTooManyRequests{}
}

/*
PutArchitectPromptResourceTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutArchitectPromptResourceTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect prompt resource too many requests response has a 2xx status code
func (o *PutArchitectPromptResourceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect prompt resource too many requests response has a 3xx status code
func (o *PutArchitectPromptResourceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect prompt resource too many requests response has a 4xx status code
func (o *PutArchitectPromptResourceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put architect prompt resource too many requests response has a 5xx status code
func (o *PutArchitectPromptResourceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put architect prompt resource too many requests response a status code equal to that given
func (o *PutArchitectPromptResourceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutArchitectPromptResourceTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutArchitectPromptResourceTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutArchitectPromptResourceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectPromptResourceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectPromptResourceInternalServerError creates a PutArchitectPromptResourceInternalServerError with default headers values
func NewPutArchitectPromptResourceInternalServerError() *PutArchitectPromptResourceInternalServerError {
	return &PutArchitectPromptResourceInternalServerError{}
}

/*
PutArchitectPromptResourceInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutArchitectPromptResourceInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect prompt resource internal server error response has a 2xx status code
func (o *PutArchitectPromptResourceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect prompt resource internal server error response has a 3xx status code
func (o *PutArchitectPromptResourceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect prompt resource internal server error response has a 4xx status code
func (o *PutArchitectPromptResourceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put architect prompt resource internal server error response has a 5xx status code
func (o *PutArchitectPromptResourceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put architect prompt resource internal server error response a status code equal to that given
func (o *PutArchitectPromptResourceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutArchitectPromptResourceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceInternalServerError  %+v", 500, o.Payload)
}

func (o *PutArchitectPromptResourceInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceInternalServerError  %+v", 500, o.Payload)
}

func (o *PutArchitectPromptResourceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectPromptResourceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectPromptResourceServiceUnavailable creates a PutArchitectPromptResourceServiceUnavailable with default headers values
func NewPutArchitectPromptResourceServiceUnavailable() *PutArchitectPromptResourceServiceUnavailable {
	return &PutArchitectPromptResourceServiceUnavailable{}
}

/*
PutArchitectPromptResourceServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutArchitectPromptResourceServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect prompt resource service unavailable response has a 2xx status code
func (o *PutArchitectPromptResourceServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect prompt resource service unavailable response has a 3xx status code
func (o *PutArchitectPromptResourceServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect prompt resource service unavailable response has a 4xx status code
func (o *PutArchitectPromptResourceServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put architect prompt resource service unavailable response has a 5xx status code
func (o *PutArchitectPromptResourceServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put architect prompt resource service unavailable response a status code equal to that given
func (o *PutArchitectPromptResourceServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutArchitectPromptResourceServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutArchitectPromptResourceServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutArchitectPromptResourceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectPromptResourceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectPromptResourceGatewayTimeout creates a PutArchitectPromptResourceGatewayTimeout with default headers values
func NewPutArchitectPromptResourceGatewayTimeout() *PutArchitectPromptResourceGatewayTimeout {
	return &PutArchitectPromptResourceGatewayTimeout{}
}

/*
PutArchitectPromptResourceGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutArchitectPromptResourceGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put architect prompt resource gateway timeout response has a 2xx status code
func (o *PutArchitectPromptResourceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put architect prompt resource gateway timeout response has a 3xx status code
func (o *PutArchitectPromptResourceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put architect prompt resource gateway timeout response has a 4xx status code
func (o *PutArchitectPromptResourceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put architect prompt resource gateway timeout response has a 5xx status code
func (o *PutArchitectPromptResourceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put architect prompt resource gateway timeout response a status code equal to that given
func (o *PutArchitectPromptResourceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutArchitectPromptResourceGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutArchitectPromptResourceGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] putArchitectPromptResourceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutArchitectPromptResourceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectPromptResourceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
