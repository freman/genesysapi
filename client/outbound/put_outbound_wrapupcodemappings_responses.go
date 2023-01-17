// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutOutboundWrapupcodemappingsReader is a Reader for the PutOutboundWrapupcodemappings structure.
type PutOutboundWrapupcodemappingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOutboundWrapupcodemappingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOutboundWrapupcodemappingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutOutboundWrapupcodemappingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutOutboundWrapupcodemappingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutOutboundWrapupcodemappingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutOutboundWrapupcodemappingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutOutboundWrapupcodemappingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutOutboundWrapupcodemappingsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutOutboundWrapupcodemappingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutOutboundWrapupcodemappingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutOutboundWrapupcodemappingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutOutboundWrapupcodemappingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutOutboundWrapupcodemappingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutOutboundWrapupcodemappingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutOutboundWrapupcodemappingsOK creates a PutOutboundWrapupcodemappingsOK with default headers values
func NewPutOutboundWrapupcodemappingsOK() *PutOutboundWrapupcodemappingsOK {
	return &PutOutboundWrapupcodemappingsOK{}
}

/*
PutOutboundWrapupcodemappingsOK describes a response with status code 200, with default header values.

successful operation
*/
type PutOutboundWrapupcodemappingsOK struct {
	Payload *models.WrapUpCodeMapping
}

// IsSuccess returns true when this put outbound wrapupcodemappings o k response has a 2xx status code
func (o *PutOutboundWrapupcodemappingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put outbound wrapupcodemappings o k response has a 3xx status code
func (o *PutOutboundWrapupcodemappingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound wrapupcodemappings o k response has a 4xx status code
func (o *PutOutboundWrapupcodemappingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put outbound wrapupcodemappings o k response has a 5xx status code
func (o *PutOutboundWrapupcodemappingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound wrapupcodemappings o k response a status code equal to that given
func (o *PutOutboundWrapupcodemappingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutOutboundWrapupcodemappingsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsOK  %+v", 200, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsOK  %+v", 200, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsOK) GetPayload() *models.WrapUpCodeMapping {
	return o.Payload
}

func (o *PutOutboundWrapupcodemappingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WrapUpCodeMapping)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundWrapupcodemappingsBadRequest creates a PutOutboundWrapupcodemappingsBadRequest with default headers values
func NewPutOutboundWrapupcodemappingsBadRequest() *PutOutboundWrapupcodemappingsBadRequest {
	return &PutOutboundWrapupcodemappingsBadRequest{}
}

/*
PutOutboundWrapupcodemappingsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutOutboundWrapupcodemappingsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound wrapupcodemappings bad request response has a 2xx status code
func (o *PutOutboundWrapupcodemappingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound wrapupcodemappings bad request response has a 3xx status code
func (o *PutOutboundWrapupcodemappingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound wrapupcodemappings bad request response has a 4xx status code
func (o *PutOutboundWrapupcodemappingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound wrapupcodemappings bad request response has a 5xx status code
func (o *PutOutboundWrapupcodemappingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound wrapupcodemappings bad request response a status code equal to that given
func (o *PutOutboundWrapupcodemappingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutOutboundWrapupcodemappingsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsBadRequest  %+v", 400, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsBadRequest  %+v", 400, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundWrapupcodemappingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundWrapupcodemappingsUnauthorized creates a PutOutboundWrapupcodemappingsUnauthorized with default headers values
func NewPutOutboundWrapupcodemappingsUnauthorized() *PutOutboundWrapupcodemappingsUnauthorized {
	return &PutOutboundWrapupcodemappingsUnauthorized{}
}

/*
PutOutboundWrapupcodemappingsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutOutboundWrapupcodemappingsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound wrapupcodemappings unauthorized response has a 2xx status code
func (o *PutOutboundWrapupcodemappingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound wrapupcodemappings unauthorized response has a 3xx status code
func (o *PutOutboundWrapupcodemappingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound wrapupcodemappings unauthorized response has a 4xx status code
func (o *PutOutboundWrapupcodemappingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound wrapupcodemappings unauthorized response has a 5xx status code
func (o *PutOutboundWrapupcodemappingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound wrapupcodemappings unauthorized response a status code equal to that given
func (o *PutOutboundWrapupcodemappingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutOutboundWrapupcodemappingsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundWrapupcodemappingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundWrapupcodemappingsForbidden creates a PutOutboundWrapupcodemappingsForbidden with default headers values
func NewPutOutboundWrapupcodemappingsForbidden() *PutOutboundWrapupcodemappingsForbidden {
	return &PutOutboundWrapupcodemappingsForbidden{}
}

/*
PutOutboundWrapupcodemappingsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutOutboundWrapupcodemappingsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound wrapupcodemappings forbidden response has a 2xx status code
func (o *PutOutboundWrapupcodemappingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound wrapupcodemappings forbidden response has a 3xx status code
func (o *PutOutboundWrapupcodemappingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound wrapupcodemappings forbidden response has a 4xx status code
func (o *PutOutboundWrapupcodemappingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound wrapupcodemappings forbidden response has a 5xx status code
func (o *PutOutboundWrapupcodemappingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound wrapupcodemappings forbidden response a status code equal to that given
func (o *PutOutboundWrapupcodemappingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutOutboundWrapupcodemappingsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsForbidden  %+v", 403, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsForbidden  %+v", 403, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundWrapupcodemappingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundWrapupcodemappingsNotFound creates a PutOutboundWrapupcodemappingsNotFound with default headers values
func NewPutOutboundWrapupcodemappingsNotFound() *PutOutboundWrapupcodemappingsNotFound {
	return &PutOutboundWrapupcodemappingsNotFound{}
}

/*
PutOutboundWrapupcodemappingsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutOutboundWrapupcodemappingsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound wrapupcodemappings not found response has a 2xx status code
func (o *PutOutboundWrapupcodemappingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound wrapupcodemappings not found response has a 3xx status code
func (o *PutOutboundWrapupcodemappingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound wrapupcodemappings not found response has a 4xx status code
func (o *PutOutboundWrapupcodemappingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound wrapupcodemappings not found response has a 5xx status code
func (o *PutOutboundWrapupcodemappingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound wrapupcodemappings not found response a status code equal to that given
func (o *PutOutboundWrapupcodemappingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutOutboundWrapupcodemappingsNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsNotFound  %+v", 404, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsNotFound  %+v", 404, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundWrapupcodemappingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundWrapupcodemappingsRequestTimeout creates a PutOutboundWrapupcodemappingsRequestTimeout with default headers values
func NewPutOutboundWrapupcodemappingsRequestTimeout() *PutOutboundWrapupcodemappingsRequestTimeout {
	return &PutOutboundWrapupcodemappingsRequestTimeout{}
}

/*
PutOutboundWrapupcodemappingsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutOutboundWrapupcodemappingsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound wrapupcodemappings request timeout response has a 2xx status code
func (o *PutOutboundWrapupcodemappingsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound wrapupcodemappings request timeout response has a 3xx status code
func (o *PutOutboundWrapupcodemappingsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound wrapupcodemappings request timeout response has a 4xx status code
func (o *PutOutboundWrapupcodemappingsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound wrapupcodemappings request timeout response has a 5xx status code
func (o *PutOutboundWrapupcodemappingsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound wrapupcodemappings request timeout response a status code equal to that given
func (o *PutOutboundWrapupcodemappingsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutOutboundWrapupcodemappingsRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundWrapupcodemappingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundWrapupcodemappingsConflict creates a PutOutboundWrapupcodemappingsConflict with default headers values
func NewPutOutboundWrapupcodemappingsConflict() *PutOutboundWrapupcodemappingsConflict {
	return &PutOutboundWrapupcodemappingsConflict{}
}

/*
PutOutboundWrapupcodemappingsConflict describes a response with status code 409, with default header values.

Conflict
*/
type PutOutboundWrapupcodemappingsConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound wrapupcodemappings conflict response has a 2xx status code
func (o *PutOutboundWrapupcodemappingsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound wrapupcodemappings conflict response has a 3xx status code
func (o *PutOutboundWrapupcodemappingsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound wrapupcodemappings conflict response has a 4xx status code
func (o *PutOutboundWrapupcodemappingsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound wrapupcodemappings conflict response has a 5xx status code
func (o *PutOutboundWrapupcodemappingsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound wrapupcodemappings conflict response a status code equal to that given
func (o *PutOutboundWrapupcodemappingsConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PutOutboundWrapupcodemappingsConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsConflict  %+v", 409, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsConflict) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsConflict  %+v", 409, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundWrapupcodemappingsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundWrapupcodemappingsRequestEntityTooLarge creates a PutOutboundWrapupcodemappingsRequestEntityTooLarge with default headers values
func NewPutOutboundWrapupcodemappingsRequestEntityTooLarge() *PutOutboundWrapupcodemappingsRequestEntityTooLarge {
	return &PutOutboundWrapupcodemappingsRequestEntityTooLarge{}
}

/*
PutOutboundWrapupcodemappingsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutOutboundWrapupcodemappingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound wrapupcodemappings request entity too large response has a 2xx status code
func (o *PutOutboundWrapupcodemappingsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound wrapupcodemappings request entity too large response has a 3xx status code
func (o *PutOutboundWrapupcodemappingsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound wrapupcodemappings request entity too large response has a 4xx status code
func (o *PutOutboundWrapupcodemappingsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound wrapupcodemappings request entity too large response has a 5xx status code
func (o *PutOutboundWrapupcodemappingsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound wrapupcodemappings request entity too large response a status code equal to that given
func (o *PutOutboundWrapupcodemappingsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutOutboundWrapupcodemappingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundWrapupcodemappingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundWrapupcodemappingsUnsupportedMediaType creates a PutOutboundWrapupcodemappingsUnsupportedMediaType with default headers values
func NewPutOutboundWrapupcodemappingsUnsupportedMediaType() *PutOutboundWrapupcodemappingsUnsupportedMediaType {
	return &PutOutboundWrapupcodemappingsUnsupportedMediaType{}
}

/*
PutOutboundWrapupcodemappingsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutOutboundWrapupcodemappingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound wrapupcodemappings unsupported media type response has a 2xx status code
func (o *PutOutboundWrapupcodemappingsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound wrapupcodemappings unsupported media type response has a 3xx status code
func (o *PutOutboundWrapupcodemappingsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound wrapupcodemappings unsupported media type response has a 4xx status code
func (o *PutOutboundWrapupcodemappingsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound wrapupcodemappings unsupported media type response has a 5xx status code
func (o *PutOutboundWrapupcodemappingsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound wrapupcodemappings unsupported media type response a status code equal to that given
func (o *PutOutboundWrapupcodemappingsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutOutboundWrapupcodemappingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundWrapupcodemappingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundWrapupcodemappingsTooManyRequests creates a PutOutboundWrapupcodemappingsTooManyRequests with default headers values
func NewPutOutboundWrapupcodemappingsTooManyRequests() *PutOutboundWrapupcodemappingsTooManyRequests {
	return &PutOutboundWrapupcodemappingsTooManyRequests{}
}

/*
PutOutboundWrapupcodemappingsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutOutboundWrapupcodemappingsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound wrapupcodemappings too many requests response has a 2xx status code
func (o *PutOutboundWrapupcodemappingsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound wrapupcodemappings too many requests response has a 3xx status code
func (o *PutOutboundWrapupcodemappingsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound wrapupcodemappings too many requests response has a 4xx status code
func (o *PutOutboundWrapupcodemappingsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put outbound wrapupcodemappings too many requests response has a 5xx status code
func (o *PutOutboundWrapupcodemappingsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put outbound wrapupcodemappings too many requests response a status code equal to that given
func (o *PutOutboundWrapupcodemappingsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutOutboundWrapupcodemappingsTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundWrapupcodemappingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundWrapupcodemappingsInternalServerError creates a PutOutboundWrapupcodemappingsInternalServerError with default headers values
func NewPutOutboundWrapupcodemappingsInternalServerError() *PutOutboundWrapupcodemappingsInternalServerError {
	return &PutOutboundWrapupcodemappingsInternalServerError{}
}

/*
PutOutboundWrapupcodemappingsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutOutboundWrapupcodemappingsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound wrapupcodemappings internal server error response has a 2xx status code
func (o *PutOutboundWrapupcodemappingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound wrapupcodemappings internal server error response has a 3xx status code
func (o *PutOutboundWrapupcodemappingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound wrapupcodemappings internal server error response has a 4xx status code
func (o *PutOutboundWrapupcodemappingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put outbound wrapupcodemappings internal server error response has a 5xx status code
func (o *PutOutboundWrapupcodemappingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put outbound wrapupcodemappings internal server error response a status code equal to that given
func (o *PutOutboundWrapupcodemappingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutOutboundWrapupcodemappingsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundWrapupcodemappingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundWrapupcodemappingsServiceUnavailable creates a PutOutboundWrapupcodemappingsServiceUnavailable with default headers values
func NewPutOutboundWrapupcodemappingsServiceUnavailable() *PutOutboundWrapupcodemappingsServiceUnavailable {
	return &PutOutboundWrapupcodemappingsServiceUnavailable{}
}

/*
PutOutboundWrapupcodemappingsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutOutboundWrapupcodemappingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound wrapupcodemappings service unavailable response has a 2xx status code
func (o *PutOutboundWrapupcodemappingsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound wrapupcodemappings service unavailable response has a 3xx status code
func (o *PutOutboundWrapupcodemappingsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound wrapupcodemappings service unavailable response has a 4xx status code
func (o *PutOutboundWrapupcodemappingsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put outbound wrapupcodemappings service unavailable response has a 5xx status code
func (o *PutOutboundWrapupcodemappingsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put outbound wrapupcodemappings service unavailable response a status code equal to that given
func (o *PutOutboundWrapupcodemappingsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutOutboundWrapupcodemappingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundWrapupcodemappingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundWrapupcodemappingsGatewayTimeout creates a PutOutboundWrapupcodemappingsGatewayTimeout with default headers values
func NewPutOutboundWrapupcodemappingsGatewayTimeout() *PutOutboundWrapupcodemappingsGatewayTimeout {
	return &PutOutboundWrapupcodemappingsGatewayTimeout{}
}

/*
PutOutboundWrapupcodemappingsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutOutboundWrapupcodemappingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put outbound wrapupcodemappings gateway timeout response has a 2xx status code
func (o *PutOutboundWrapupcodemappingsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put outbound wrapupcodemappings gateway timeout response has a 3xx status code
func (o *PutOutboundWrapupcodemappingsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put outbound wrapupcodemappings gateway timeout response has a 4xx status code
func (o *PutOutboundWrapupcodemappingsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put outbound wrapupcodemappings gateway timeout response has a 5xx status code
func (o *PutOutboundWrapupcodemappingsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put outbound wrapupcodemappings gateway timeout response a status code equal to that given
func (o *PutOutboundWrapupcodemappingsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutOutboundWrapupcodemappingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/wrapupcodemappings][%d] putOutboundWrapupcodemappingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutOutboundWrapupcodemappingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundWrapupcodemappingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
