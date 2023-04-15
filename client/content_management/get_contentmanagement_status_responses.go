// Code generated by go-swagger; DO NOT EDIT.

package content_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetContentmanagementStatusReader is a Reader for the GetContentmanagementStatus structure.
type GetContentmanagementStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContentmanagementStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContentmanagementStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetContentmanagementStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetContentmanagementStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetContentmanagementStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetContentmanagementStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetContentmanagementStatusRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetContentmanagementStatusRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetContentmanagementStatusUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetContentmanagementStatusTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetContentmanagementStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetContentmanagementStatusServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetContentmanagementStatusGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetContentmanagementStatusOK creates a GetContentmanagementStatusOK with default headers values
func NewGetContentmanagementStatusOK() *GetContentmanagementStatusOK {
	return &GetContentmanagementStatusOK{}
}

/*
GetContentmanagementStatusOK describes a response with status code 200, with default header values.

successful operation
*/
type GetContentmanagementStatusOK struct {
	Payload *models.CommandStatusEntityListing
}

// IsSuccess returns true when this get contentmanagement status o k response has a 2xx status code
func (o *GetContentmanagementStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get contentmanagement status o k response has a 3xx status code
func (o *GetContentmanagementStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement status o k response has a 4xx status code
func (o *GetContentmanagementStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement status o k response has a 5xx status code
func (o *GetContentmanagementStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement status o k response a status code equal to that given
func (o *GetContentmanagementStatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetContentmanagementStatusOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusOK  %+v", 200, o.Payload)
}

func (o *GetContentmanagementStatusOK) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusOK  %+v", 200, o.Payload)
}

func (o *GetContentmanagementStatusOK) GetPayload() *models.CommandStatusEntityListing {
	return o.Payload
}

func (o *GetContentmanagementStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommandStatusEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementStatusBadRequest creates a GetContentmanagementStatusBadRequest with default headers values
func NewGetContentmanagementStatusBadRequest() *GetContentmanagementStatusBadRequest {
	return &GetContentmanagementStatusBadRequest{}
}

/*
GetContentmanagementStatusBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetContentmanagementStatusBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement status bad request response has a 2xx status code
func (o *GetContentmanagementStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement status bad request response has a 3xx status code
func (o *GetContentmanagementStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement status bad request response has a 4xx status code
func (o *GetContentmanagementStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement status bad request response has a 5xx status code
func (o *GetContentmanagementStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement status bad request response a status code equal to that given
func (o *GetContentmanagementStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetContentmanagementStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusBadRequest  %+v", 400, o.Payload)
}

func (o *GetContentmanagementStatusBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusBadRequest  %+v", 400, o.Payload)
}

func (o *GetContentmanagementStatusBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementStatusUnauthorized creates a GetContentmanagementStatusUnauthorized with default headers values
func NewGetContentmanagementStatusUnauthorized() *GetContentmanagementStatusUnauthorized {
	return &GetContentmanagementStatusUnauthorized{}
}

/*
GetContentmanagementStatusUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetContentmanagementStatusUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement status unauthorized response has a 2xx status code
func (o *GetContentmanagementStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement status unauthorized response has a 3xx status code
func (o *GetContentmanagementStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement status unauthorized response has a 4xx status code
func (o *GetContentmanagementStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement status unauthorized response has a 5xx status code
func (o *GetContentmanagementStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement status unauthorized response a status code equal to that given
func (o *GetContentmanagementStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetContentmanagementStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *GetContentmanagementStatusUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *GetContentmanagementStatusUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementStatusForbidden creates a GetContentmanagementStatusForbidden with default headers values
func NewGetContentmanagementStatusForbidden() *GetContentmanagementStatusForbidden {
	return &GetContentmanagementStatusForbidden{}
}

/*
GetContentmanagementStatusForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetContentmanagementStatusForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement status forbidden response has a 2xx status code
func (o *GetContentmanagementStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement status forbidden response has a 3xx status code
func (o *GetContentmanagementStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement status forbidden response has a 4xx status code
func (o *GetContentmanagementStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement status forbidden response has a 5xx status code
func (o *GetContentmanagementStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement status forbidden response a status code equal to that given
func (o *GetContentmanagementStatusForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetContentmanagementStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusForbidden  %+v", 403, o.Payload)
}

func (o *GetContentmanagementStatusForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusForbidden  %+v", 403, o.Payload)
}

func (o *GetContentmanagementStatusForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementStatusNotFound creates a GetContentmanagementStatusNotFound with default headers values
func NewGetContentmanagementStatusNotFound() *GetContentmanagementStatusNotFound {
	return &GetContentmanagementStatusNotFound{}
}

/*
GetContentmanagementStatusNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetContentmanagementStatusNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement status not found response has a 2xx status code
func (o *GetContentmanagementStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement status not found response has a 3xx status code
func (o *GetContentmanagementStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement status not found response has a 4xx status code
func (o *GetContentmanagementStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement status not found response has a 5xx status code
func (o *GetContentmanagementStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement status not found response a status code equal to that given
func (o *GetContentmanagementStatusNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetContentmanagementStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusNotFound  %+v", 404, o.Payload)
}

func (o *GetContentmanagementStatusNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusNotFound  %+v", 404, o.Payload)
}

func (o *GetContentmanagementStatusNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementStatusRequestTimeout creates a GetContentmanagementStatusRequestTimeout with default headers values
func NewGetContentmanagementStatusRequestTimeout() *GetContentmanagementStatusRequestTimeout {
	return &GetContentmanagementStatusRequestTimeout{}
}

/*
GetContentmanagementStatusRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetContentmanagementStatusRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement status request timeout response has a 2xx status code
func (o *GetContentmanagementStatusRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement status request timeout response has a 3xx status code
func (o *GetContentmanagementStatusRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement status request timeout response has a 4xx status code
func (o *GetContentmanagementStatusRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement status request timeout response has a 5xx status code
func (o *GetContentmanagementStatusRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement status request timeout response a status code equal to that given
func (o *GetContentmanagementStatusRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetContentmanagementStatusRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetContentmanagementStatusRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetContentmanagementStatusRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementStatusRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementStatusRequestEntityTooLarge creates a GetContentmanagementStatusRequestEntityTooLarge with default headers values
func NewGetContentmanagementStatusRequestEntityTooLarge() *GetContentmanagementStatusRequestEntityTooLarge {
	return &GetContentmanagementStatusRequestEntityTooLarge{}
}

/*
GetContentmanagementStatusRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetContentmanagementStatusRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement status request entity too large response has a 2xx status code
func (o *GetContentmanagementStatusRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement status request entity too large response has a 3xx status code
func (o *GetContentmanagementStatusRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement status request entity too large response has a 4xx status code
func (o *GetContentmanagementStatusRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement status request entity too large response has a 5xx status code
func (o *GetContentmanagementStatusRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement status request entity too large response a status code equal to that given
func (o *GetContentmanagementStatusRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetContentmanagementStatusRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetContentmanagementStatusRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetContentmanagementStatusRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementStatusRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementStatusUnsupportedMediaType creates a GetContentmanagementStatusUnsupportedMediaType with default headers values
func NewGetContentmanagementStatusUnsupportedMediaType() *GetContentmanagementStatusUnsupportedMediaType {
	return &GetContentmanagementStatusUnsupportedMediaType{}
}

/*
GetContentmanagementStatusUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetContentmanagementStatusUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement status unsupported media type response has a 2xx status code
func (o *GetContentmanagementStatusUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement status unsupported media type response has a 3xx status code
func (o *GetContentmanagementStatusUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement status unsupported media type response has a 4xx status code
func (o *GetContentmanagementStatusUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement status unsupported media type response has a 5xx status code
func (o *GetContentmanagementStatusUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement status unsupported media type response a status code equal to that given
func (o *GetContentmanagementStatusUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetContentmanagementStatusUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetContentmanagementStatusUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetContentmanagementStatusUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementStatusUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementStatusTooManyRequests creates a GetContentmanagementStatusTooManyRequests with default headers values
func NewGetContentmanagementStatusTooManyRequests() *GetContentmanagementStatusTooManyRequests {
	return &GetContentmanagementStatusTooManyRequests{}
}

/*
GetContentmanagementStatusTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetContentmanagementStatusTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement status too many requests response has a 2xx status code
func (o *GetContentmanagementStatusTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement status too many requests response has a 3xx status code
func (o *GetContentmanagementStatusTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement status too many requests response has a 4xx status code
func (o *GetContentmanagementStatusTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement status too many requests response has a 5xx status code
func (o *GetContentmanagementStatusTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement status too many requests response a status code equal to that given
func (o *GetContentmanagementStatusTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetContentmanagementStatusTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetContentmanagementStatusTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetContentmanagementStatusTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementStatusTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementStatusInternalServerError creates a GetContentmanagementStatusInternalServerError with default headers values
func NewGetContentmanagementStatusInternalServerError() *GetContentmanagementStatusInternalServerError {
	return &GetContentmanagementStatusInternalServerError{}
}

/*
GetContentmanagementStatusInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetContentmanagementStatusInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement status internal server error response has a 2xx status code
func (o *GetContentmanagementStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement status internal server error response has a 3xx status code
func (o *GetContentmanagementStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement status internal server error response has a 4xx status code
func (o *GetContentmanagementStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement status internal server error response has a 5xx status code
func (o *GetContentmanagementStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get contentmanagement status internal server error response a status code equal to that given
func (o *GetContentmanagementStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetContentmanagementStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *GetContentmanagementStatusInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *GetContentmanagementStatusInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementStatusServiceUnavailable creates a GetContentmanagementStatusServiceUnavailable with default headers values
func NewGetContentmanagementStatusServiceUnavailable() *GetContentmanagementStatusServiceUnavailable {
	return &GetContentmanagementStatusServiceUnavailable{}
}

/*
GetContentmanagementStatusServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetContentmanagementStatusServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement status service unavailable response has a 2xx status code
func (o *GetContentmanagementStatusServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement status service unavailable response has a 3xx status code
func (o *GetContentmanagementStatusServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement status service unavailable response has a 4xx status code
func (o *GetContentmanagementStatusServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement status service unavailable response has a 5xx status code
func (o *GetContentmanagementStatusServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get contentmanagement status service unavailable response a status code equal to that given
func (o *GetContentmanagementStatusServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetContentmanagementStatusServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetContentmanagementStatusServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetContentmanagementStatusServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementStatusServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementStatusGatewayTimeout creates a GetContentmanagementStatusGatewayTimeout with default headers values
func NewGetContentmanagementStatusGatewayTimeout() *GetContentmanagementStatusGatewayTimeout {
	return &GetContentmanagementStatusGatewayTimeout{}
}

/*
GetContentmanagementStatusGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetContentmanagementStatusGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement status gateway timeout response has a 2xx status code
func (o *GetContentmanagementStatusGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement status gateway timeout response has a 3xx status code
func (o *GetContentmanagementStatusGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement status gateway timeout response has a 4xx status code
func (o *GetContentmanagementStatusGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement status gateway timeout response has a 5xx status code
func (o *GetContentmanagementStatusGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get contentmanagement status gateway timeout response a status code equal to that given
func (o *GetContentmanagementStatusGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetContentmanagementStatusGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetContentmanagementStatusGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status][%d] getContentmanagementStatusGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetContentmanagementStatusGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementStatusGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
