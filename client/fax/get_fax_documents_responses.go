// Code generated by go-swagger; DO NOT EDIT.

package fax

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetFaxDocumentsReader is a Reader for the GetFaxDocuments structure.
type GetFaxDocumentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFaxDocumentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFaxDocumentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFaxDocumentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFaxDocumentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFaxDocumentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFaxDocumentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetFaxDocumentsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFaxDocumentsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFaxDocumentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFaxDocumentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFaxDocumentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFaxDocumentsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFaxDocumentsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFaxDocumentsOK creates a GetFaxDocumentsOK with default headers values
func NewGetFaxDocumentsOK() *GetFaxDocumentsOK {
	return &GetFaxDocumentsOK{}
}

/*
GetFaxDocumentsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetFaxDocumentsOK struct {
	Payload *models.FaxDocumentEntityListing
}

// IsSuccess returns true when this get fax documents o k response has a 2xx status code
func (o *GetFaxDocumentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get fax documents o k response has a 3xx status code
func (o *GetFaxDocumentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax documents o k response has a 4xx status code
func (o *GetFaxDocumentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fax documents o k response has a 5xx status code
func (o *GetFaxDocumentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get fax documents o k response a status code equal to that given
func (o *GetFaxDocumentsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetFaxDocumentsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsOK  %+v", 200, o.Payload)
}

func (o *GetFaxDocumentsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsOK  %+v", 200, o.Payload)
}

func (o *GetFaxDocumentsOK) GetPayload() *models.FaxDocumentEntityListing {
	return o.Payload
}

func (o *GetFaxDocumentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FaxDocumentEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentsBadRequest creates a GetFaxDocumentsBadRequest with default headers values
func NewGetFaxDocumentsBadRequest() *GetFaxDocumentsBadRequest {
	return &GetFaxDocumentsBadRequest{}
}

/*
GetFaxDocumentsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFaxDocumentsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax documents bad request response has a 2xx status code
func (o *GetFaxDocumentsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax documents bad request response has a 3xx status code
func (o *GetFaxDocumentsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax documents bad request response has a 4xx status code
func (o *GetFaxDocumentsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fax documents bad request response has a 5xx status code
func (o *GetFaxDocumentsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get fax documents bad request response a status code equal to that given
func (o *GetFaxDocumentsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetFaxDocumentsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsBadRequest  %+v", 400, o.Payload)
}

func (o *GetFaxDocumentsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsBadRequest  %+v", 400, o.Payload)
}

func (o *GetFaxDocumentsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentsUnauthorized creates a GetFaxDocumentsUnauthorized with default headers values
func NewGetFaxDocumentsUnauthorized() *GetFaxDocumentsUnauthorized {
	return &GetFaxDocumentsUnauthorized{}
}

/*
GetFaxDocumentsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFaxDocumentsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax documents unauthorized response has a 2xx status code
func (o *GetFaxDocumentsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax documents unauthorized response has a 3xx status code
func (o *GetFaxDocumentsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax documents unauthorized response has a 4xx status code
func (o *GetFaxDocumentsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fax documents unauthorized response has a 5xx status code
func (o *GetFaxDocumentsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get fax documents unauthorized response a status code equal to that given
func (o *GetFaxDocumentsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetFaxDocumentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFaxDocumentsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFaxDocumentsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentsForbidden creates a GetFaxDocumentsForbidden with default headers values
func NewGetFaxDocumentsForbidden() *GetFaxDocumentsForbidden {
	return &GetFaxDocumentsForbidden{}
}

/*
GetFaxDocumentsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetFaxDocumentsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax documents forbidden response has a 2xx status code
func (o *GetFaxDocumentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax documents forbidden response has a 3xx status code
func (o *GetFaxDocumentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax documents forbidden response has a 4xx status code
func (o *GetFaxDocumentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fax documents forbidden response has a 5xx status code
func (o *GetFaxDocumentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get fax documents forbidden response a status code equal to that given
func (o *GetFaxDocumentsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetFaxDocumentsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsForbidden  %+v", 403, o.Payload)
}

func (o *GetFaxDocumentsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsForbidden  %+v", 403, o.Payload)
}

func (o *GetFaxDocumentsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentsNotFound creates a GetFaxDocumentsNotFound with default headers values
func NewGetFaxDocumentsNotFound() *GetFaxDocumentsNotFound {
	return &GetFaxDocumentsNotFound{}
}

/*
GetFaxDocumentsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetFaxDocumentsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax documents not found response has a 2xx status code
func (o *GetFaxDocumentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax documents not found response has a 3xx status code
func (o *GetFaxDocumentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax documents not found response has a 4xx status code
func (o *GetFaxDocumentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fax documents not found response has a 5xx status code
func (o *GetFaxDocumentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get fax documents not found response a status code equal to that given
func (o *GetFaxDocumentsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetFaxDocumentsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsNotFound  %+v", 404, o.Payload)
}

func (o *GetFaxDocumentsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsNotFound  %+v", 404, o.Payload)
}

func (o *GetFaxDocumentsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentsRequestTimeout creates a GetFaxDocumentsRequestTimeout with default headers values
func NewGetFaxDocumentsRequestTimeout() *GetFaxDocumentsRequestTimeout {
	return &GetFaxDocumentsRequestTimeout{}
}

/*
GetFaxDocumentsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetFaxDocumentsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax documents request timeout response has a 2xx status code
func (o *GetFaxDocumentsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax documents request timeout response has a 3xx status code
func (o *GetFaxDocumentsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax documents request timeout response has a 4xx status code
func (o *GetFaxDocumentsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fax documents request timeout response has a 5xx status code
func (o *GetFaxDocumentsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get fax documents request timeout response a status code equal to that given
func (o *GetFaxDocumentsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetFaxDocumentsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFaxDocumentsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFaxDocumentsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentsRequestEntityTooLarge creates a GetFaxDocumentsRequestEntityTooLarge with default headers values
func NewGetFaxDocumentsRequestEntityTooLarge() *GetFaxDocumentsRequestEntityTooLarge {
	return &GetFaxDocumentsRequestEntityTooLarge{}
}

/*
GetFaxDocumentsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetFaxDocumentsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax documents request entity too large response has a 2xx status code
func (o *GetFaxDocumentsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax documents request entity too large response has a 3xx status code
func (o *GetFaxDocumentsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax documents request entity too large response has a 4xx status code
func (o *GetFaxDocumentsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fax documents request entity too large response has a 5xx status code
func (o *GetFaxDocumentsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get fax documents request entity too large response a status code equal to that given
func (o *GetFaxDocumentsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetFaxDocumentsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFaxDocumentsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFaxDocumentsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentsUnsupportedMediaType creates a GetFaxDocumentsUnsupportedMediaType with default headers values
func NewGetFaxDocumentsUnsupportedMediaType() *GetFaxDocumentsUnsupportedMediaType {
	return &GetFaxDocumentsUnsupportedMediaType{}
}

/*
GetFaxDocumentsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFaxDocumentsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax documents unsupported media type response has a 2xx status code
func (o *GetFaxDocumentsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax documents unsupported media type response has a 3xx status code
func (o *GetFaxDocumentsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax documents unsupported media type response has a 4xx status code
func (o *GetFaxDocumentsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fax documents unsupported media type response has a 5xx status code
func (o *GetFaxDocumentsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get fax documents unsupported media type response a status code equal to that given
func (o *GetFaxDocumentsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetFaxDocumentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFaxDocumentsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFaxDocumentsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentsTooManyRequests creates a GetFaxDocumentsTooManyRequests with default headers values
func NewGetFaxDocumentsTooManyRequests() *GetFaxDocumentsTooManyRequests {
	return &GetFaxDocumentsTooManyRequests{}
}

/*
GetFaxDocumentsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetFaxDocumentsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax documents too many requests response has a 2xx status code
func (o *GetFaxDocumentsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax documents too many requests response has a 3xx status code
func (o *GetFaxDocumentsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax documents too many requests response has a 4xx status code
func (o *GetFaxDocumentsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fax documents too many requests response has a 5xx status code
func (o *GetFaxDocumentsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get fax documents too many requests response a status code equal to that given
func (o *GetFaxDocumentsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetFaxDocumentsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFaxDocumentsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFaxDocumentsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentsInternalServerError creates a GetFaxDocumentsInternalServerError with default headers values
func NewGetFaxDocumentsInternalServerError() *GetFaxDocumentsInternalServerError {
	return &GetFaxDocumentsInternalServerError{}
}

/*
GetFaxDocumentsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFaxDocumentsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax documents internal server error response has a 2xx status code
func (o *GetFaxDocumentsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax documents internal server error response has a 3xx status code
func (o *GetFaxDocumentsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax documents internal server error response has a 4xx status code
func (o *GetFaxDocumentsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fax documents internal server error response has a 5xx status code
func (o *GetFaxDocumentsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get fax documents internal server error response a status code equal to that given
func (o *GetFaxDocumentsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetFaxDocumentsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFaxDocumentsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFaxDocumentsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentsServiceUnavailable creates a GetFaxDocumentsServiceUnavailable with default headers values
func NewGetFaxDocumentsServiceUnavailable() *GetFaxDocumentsServiceUnavailable {
	return &GetFaxDocumentsServiceUnavailable{}
}

/*
GetFaxDocumentsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFaxDocumentsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax documents service unavailable response has a 2xx status code
func (o *GetFaxDocumentsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax documents service unavailable response has a 3xx status code
func (o *GetFaxDocumentsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax documents service unavailable response has a 4xx status code
func (o *GetFaxDocumentsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fax documents service unavailable response has a 5xx status code
func (o *GetFaxDocumentsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get fax documents service unavailable response a status code equal to that given
func (o *GetFaxDocumentsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetFaxDocumentsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFaxDocumentsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFaxDocumentsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentsGatewayTimeout creates a GetFaxDocumentsGatewayTimeout with default headers values
func NewGetFaxDocumentsGatewayTimeout() *GetFaxDocumentsGatewayTimeout {
	return &GetFaxDocumentsGatewayTimeout{}
}

/*
GetFaxDocumentsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetFaxDocumentsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax documents gateway timeout response has a 2xx status code
func (o *GetFaxDocumentsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax documents gateway timeout response has a 3xx status code
func (o *GetFaxDocumentsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax documents gateway timeout response has a 4xx status code
func (o *GetFaxDocumentsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fax documents gateway timeout response has a 5xx status code
func (o *GetFaxDocumentsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get fax documents gateway timeout response a status code equal to that given
func (o *GetFaxDocumentsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetFaxDocumentsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFaxDocumentsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFaxDocumentsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
