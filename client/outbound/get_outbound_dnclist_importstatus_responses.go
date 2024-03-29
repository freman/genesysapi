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

// GetOutboundDnclistImportstatusReader is a Reader for the GetOutboundDnclistImportstatus structure.
type GetOutboundDnclistImportstatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundDnclistImportstatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundDnclistImportstatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundDnclistImportstatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundDnclistImportstatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundDnclistImportstatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundDnclistImportstatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundDnclistImportstatusRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundDnclistImportstatusRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundDnclistImportstatusUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundDnclistImportstatusTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundDnclistImportstatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundDnclistImportstatusServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundDnclistImportstatusGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundDnclistImportstatusOK creates a GetOutboundDnclistImportstatusOK with default headers values
func NewGetOutboundDnclistImportstatusOK() *GetOutboundDnclistImportstatusOK {
	return &GetOutboundDnclistImportstatusOK{}
}

/*
GetOutboundDnclistImportstatusOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOutboundDnclistImportstatusOK struct {
	Payload *models.ImportStatus
}

// IsSuccess returns true when this get outbound dnclist importstatus o k response has a 2xx status code
func (o *GetOutboundDnclistImportstatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get outbound dnclist importstatus o k response has a 3xx status code
func (o *GetOutboundDnclistImportstatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound dnclist importstatus o k response has a 4xx status code
func (o *GetOutboundDnclistImportstatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound dnclist importstatus o k response has a 5xx status code
func (o *GetOutboundDnclistImportstatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound dnclist importstatus o k response a status code equal to that given
func (o *GetOutboundDnclistImportstatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOutboundDnclistImportstatusOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusOK  %+v", 200, o.Payload)
}

func (o *GetOutboundDnclistImportstatusOK) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusOK  %+v", 200, o.Payload)
}

func (o *GetOutboundDnclistImportstatusOK) GetPayload() *models.ImportStatus {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImportStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistImportstatusBadRequest creates a GetOutboundDnclistImportstatusBadRequest with default headers values
func NewGetOutboundDnclistImportstatusBadRequest() *GetOutboundDnclistImportstatusBadRequest {
	return &GetOutboundDnclistImportstatusBadRequest{}
}

/*
GetOutboundDnclistImportstatusBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundDnclistImportstatusBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound dnclist importstatus bad request response has a 2xx status code
func (o *GetOutboundDnclistImportstatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound dnclist importstatus bad request response has a 3xx status code
func (o *GetOutboundDnclistImportstatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound dnclist importstatus bad request response has a 4xx status code
func (o *GetOutboundDnclistImportstatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound dnclist importstatus bad request response has a 5xx status code
func (o *GetOutboundDnclistImportstatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound dnclist importstatus bad request response a status code equal to that given
func (o *GetOutboundDnclistImportstatusBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOutboundDnclistImportstatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundDnclistImportstatusBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundDnclistImportstatusBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistImportstatusUnauthorized creates a GetOutboundDnclistImportstatusUnauthorized with default headers values
func NewGetOutboundDnclistImportstatusUnauthorized() *GetOutboundDnclistImportstatusUnauthorized {
	return &GetOutboundDnclistImportstatusUnauthorized{}
}

/*
GetOutboundDnclistImportstatusUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundDnclistImportstatusUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound dnclist importstatus unauthorized response has a 2xx status code
func (o *GetOutboundDnclistImportstatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound dnclist importstatus unauthorized response has a 3xx status code
func (o *GetOutboundDnclistImportstatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound dnclist importstatus unauthorized response has a 4xx status code
func (o *GetOutboundDnclistImportstatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound dnclist importstatus unauthorized response has a 5xx status code
func (o *GetOutboundDnclistImportstatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound dnclist importstatus unauthorized response a status code equal to that given
func (o *GetOutboundDnclistImportstatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOutboundDnclistImportstatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundDnclistImportstatusUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundDnclistImportstatusUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistImportstatusForbidden creates a GetOutboundDnclistImportstatusForbidden with default headers values
func NewGetOutboundDnclistImportstatusForbidden() *GetOutboundDnclistImportstatusForbidden {
	return &GetOutboundDnclistImportstatusForbidden{}
}

/*
GetOutboundDnclistImportstatusForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundDnclistImportstatusForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound dnclist importstatus forbidden response has a 2xx status code
func (o *GetOutboundDnclistImportstatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound dnclist importstatus forbidden response has a 3xx status code
func (o *GetOutboundDnclistImportstatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound dnclist importstatus forbidden response has a 4xx status code
func (o *GetOutboundDnclistImportstatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound dnclist importstatus forbidden response has a 5xx status code
func (o *GetOutboundDnclistImportstatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound dnclist importstatus forbidden response a status code equal to that given
func (o *GetOutboundDnclistImportstatusForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOutboundDnclistImportstatusForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundDnclistImportstatusForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundDnclistImportstatusForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistImportstatusNotFound creates a GetOutboundDnclistImportstatusNotFound with default headers values
func NewGetOutboundDnclistImportstatusNotFound() *GetOutboundDnclistImportstatusNotFound {
	return &GetOutboundDnclistImportstatusNotFound{}
}

/*
GetOutboundDnclistImportstatusNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOutboundDnclistImportstatusNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound dnclist importstatus not found response has a 2xx status code
func (o *GetOutboundDnclistImportstatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound dnclist importstatus not found response has a 3xx status code
func (o *GetOutboundDnclistImportstatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound dnclist importstatus not found response has a 4xx status code
func (o *GetOutboundDnclistImportstatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound dnclist importstatus not found response has a 5xx status code
func (o *GetOutboundDnclistImportstatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound dnclist importstatus not found response a status code equal to that given
func (o *GetOutboundDnclistImportstatusNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOutboundDnclistImportstatusNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundDnclistImportstatusNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundDnclistImportstatusNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistImportstatusRequestTimeout creates a GetOutboundDnclistImportstatusRequestTimeout with default headers values
func NewGetOutboundDnclistImportstatusRequestTimeout() *GetOutboundDnclistImportstatusRequestTimeout {
	return &GetOutboundDnclistImportstatusRequestTimeout{}
}

/*
GetOutboundDnclistImportstatusRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundDnclistImportstatusRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound dnclist importstatus request timeout response has a 2xx status code
func (o *GetOutboundDnclistImportstatusRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound dnclist importstatus request timeout response has a 3xx status code
func (o *GetOutboundDnclistImportstatusRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound dnclist importstatus request timeout response has a 4xx status code
func (o *GetOutboundDnclistImportstatusRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound dnclist importstatus request timeout response has a 5xx status code
func (o *GetOutboundDnclistImportstatusRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound dnclist importstatus request timeout response a status code equal to that given
func (o *GetOutboundDnclistImportstatusRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOutboundDnclistImportstatusRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundDnclistImportstatusRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundDnclistImportstatusRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistImportstatusRequestEntityTooLarge creates a GetOutboundDnclistImportstatusRequestEntityTooLarge with default headers values
func NewGetOutboundDnclistImportstatusRequestEntityTooLarge() *GetOutboundDnclistImportstatusRequestEntityTooLarge {
	return &GetOutboundDnclistImportstatusRequestEntityTooLarge{}
}

/*
GetOutboundDnclistImportstatusRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetOutboundDnclistImportstatusRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound dnclist importstatus request entity too large response has a 2xx status code
func (o *GetOutboundDnclistImportstatusRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound dnclist importstatus request entity too large response has a 3xx status code
func (o *GetOutboundDnclistImportstatusRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound dnclist importstatus request entity too large response has a 4xx status code
func (o *GetOutboundDnclistImportstatusRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound dnclist importstatus request entity too large response has a 5xx status code
func (o *GetOutboundDnclistImportstatusRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound dnclist importstatus request entity too large response a status code equal to that given
func (o *GetOutboundDnclistImportstatusRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOutboundDnclistImportstatusRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundDnclistImportstatusRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundDnclistImportstatusRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistImportstatusUnsupportedMediaType creates a GetOutboundDnclistImportstatusUnsupportedMediaType with default headers values
func NewGetOutboundDnclistImportstatusUnsupportedMediaType() *GetOutboundDnclistImportstatusUnsupportedMediaType {
	return &GetOutboundDnclistImportstatusUnsupportedMediaType{}
}

/*
GetOutboundDnclistImportstatusUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundDnclistImportstatusUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound dnclist importstatus unsupported media type response has a 2xx status code
func (o *GetOutboundDnclistImportstatusUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound dnclist importstatus unsupported media type response has a 3xx status code
func (o *GetOutboundDnclistImportstatusUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound dnclist importstatus unsupported media type response has a 4xx status code
func (o *GetOutboundDnclistImportstatusUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound dnclist importstatus unsupported media type response has a 5xx status code
func (o *GetOutboundDnclistImportstatusUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound dnclist importstatus unsupported media type response a status code equal to that given
func (o *GetOutboundDnclistImportstatusUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOutboundDnclistImportstatusUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundDnclistImportstatusUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundDnclistImportstatusUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistImportstatusTooManyRequests creates a GetOutboundDnclistImportstatusTooManyRequests with default headers values
func NewGetOutboundDnclistImportstatusTooManyRequests() *GetOutboundDnclistImportstatusTooManyRequests {
	return &GetOutboundDnclistImportstatusTooManyRequests{}
}

/*
GetOutboundDnclistImportstatusTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundDnclistImportstatusTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound dnclist importstatus too many requests response has a 2xx status code
func (o *GetOutboundDnclistImportstatusTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound dnclist importstatus too many requests response has a 3xx status code
func (o *GetOutboundDnclistImportstatusTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound dnclist importstatus too many requests response has a 4xx status code
func (o *GetOutboundDnclistImportstatusTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound dnclist importstatus too many requests response has a 5xx status code
func (o *GetOutboundDnclistImportstatusTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound dnclist importstatus too many requests response a status code equal to that given
func (o *GetOutboundDnclistImportstatusTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOutboundDnclistImportstatusTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundDnclistImportstatusTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundDnclistImportstatusTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistImportstatusInternalServerError creates a GetOutboundDnclistImportstatusInternalServerError with default headers values
func NewGetOutboundDnclistImportstatusInternalServerError() *GetOutboundDnclistImportstatusInternalServerError {
	return &GetOutboundDnclistImportstatusInternalServerError{}
}

/*
GetOutboundDnclistImportstatusInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundDnclistImportstatusInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound dnclist importstatus internal server error response has a 2xx status code
func (o *GetOutboundDnclistImportstatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound dnclist importstatus internal server error response has a 3xx status code
func (o *GetOutboundDnclistImportstatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound dnclist importstatus internal server error response has a 4xx status code
func (o *GetOutboundDnclistImportstatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound dnclist importstatus internal server error response has a 5xx status code
func (o *GetOutboundDnclistImportstatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound dnclist importstatus internal server error response a status code equal to that given
func (o *GetOutboundDnclistImportstatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOutboundDnclistImportstatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundDnclistImportstatusInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundDnclistImportstatusInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistImportstatusServiceUnavailable creates a GetOutboundDnclistImportstatusServiceUnavailable with default headers values
func NewGetOutboundDnclistImportstatusServiceUnavailable() *GetOutboundDnclistImportstatusServiceUnavailable {
	return &GetOutboundDnclistImportstatusServiceUnavailable{}
}

/*
GetOutboundDnclistImportstatusServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundDnclistImportstatusServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound dnclist importstatus service unavailable response has a 2xx status code
func (o *GetOutboundDnclistImportstatusServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound dnclist importstatus service unavailable response has a 3xx status code
func (o *GetOutboundDnclistImportstatusServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound dnclist importstatus service unavailable response has a 4xx status code
func (o *GetOutboundDnclistImportstatusServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound dnclist importstatus service unavailable response has a 5xx status code
func (o *GetOutboundDnclistImportstatusServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound dnclist importstatus service unavailable response a status code equal to that given
func (o *GetOutboundDnclistImportstatusServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOutboundDnclistImportstatusServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundDnclistImportstatusServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundDnclistImportstatusServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistImportstatusGatewayTimeout creates a GetOutboundDnclistImportstatusGatewayTimeout with default headers values
func NewGetOutboundDnclistImportstatusGatewayTimeout() *GetOutboundDnclistImportstatusGatewayTimeout {
	return &GetOutboundDnclistImportstatusGatewayTimeout{}
}

/*
GetOutboundDnclistImportstatusGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOutboundDnclistImportstatusGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound dnclist importstatus gateway timeout response has a 2xx status code
func (o *GetOutboundDnclistImportstatusGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound dnclist importstatus gateway timeout response has a 3xx status code
func (o *GetOutboundDnclistImportstatusGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound dnclist importstatus gateway timeout response has a 4xx status code
func (o *GetOutboundDnclistImportstatusGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound dnclist importstatus gateway timeout response has a 5xx status code
func (o *GetOutboundDnclistImportstatusGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound dnclist importstatus gateway timeout response a status code equal to that given
func (o *GetOutboundDnclistImportstatusGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOutboundDnclistImportstatusGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundDnclistImportstatusGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundDnclistImportstatusGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
