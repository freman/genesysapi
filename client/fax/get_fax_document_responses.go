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

// GetFaxDocumentReader is a Reader for the GetFaxDocument structure.
type GetFaxDocumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFaxDocumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFaxDocumentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFaxDocumentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFaxDocumentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFaxDocumentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFaxDocumentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetFaxDocumentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFaxDocumentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFaxDocumentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFaxDocumentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFaxDocumentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFaxDocumentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFaxDocumentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFaxDocumentOK creates a GetFaxDocumentOK with default headers values
func NewGetFaxDocumentOK() *GetFaxDocumentOK {
	return &GetFaxDocumentOK{}
}

/*
GetFaxDocumentOK describes a response with status code 200, with default header values.

successful operation
*/
type GetFaxDocumentOK struct {
	Payload *models.FaxDocument
}

// IsSuccess returns true when this get fax document o k response has a 2xx status code
func (o *GetFaxDocumentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get fax document o k response has a 3xx status code
func (o *GetFaxDocumentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax document o k response has a 4xx status code
func (o *GetFaxDocumentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fax document o k response has a 5xx status code
func (o *GetFaxDocumentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get fax document o k response a status code equal to that given
func (o *GetFaxDocumentOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetFaxDocumentOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentOK  %+v", 200, o.Payload)
}

func (o *GetFaxDocumentOK) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentOK  %+v", 200, o.Payload)
}

func (o *GetFaxDocumentOK) GetPayload() *models.FaxDocument {
	return o.Payload
}

func (o *GetFaxDocumentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FaxDocument)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentBadRequest creates a GetFaxDocumentBadRequest with default headers values
func NewGetFaxDocumentBadRequest() *GetFaxDocumentBadRequest {
	return &GetFaxDocumentBadRequest{}
}

/*
GetFaxDocumentBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFaxDocumentBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax document bad request response has a 2xx status code
func (o *GetFaxDocumentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax document bad request response has a 3xx status code
func (o *GetFaxDocumentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax document bad request response has a 4xx status code
func (o *GetFaxDocumentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fax document bad request response has a 5xx status code
func (o *GetFaxDocumentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get fax document bad request response a status code equal to that given
func (o *GetFaxDocumentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetFaxDocumentBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentBadRequest  %+v", 400, o.Payload)
}

func (o *GetFaxDocumentBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentBadRequest  %+v", 400, o.Payload)
}

func (o *GetFaxDocumentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentUnauthorized creates a GetFaxDocumentUnauthorized with default headers values
func NewGetFaxDocumentUnauthorized() *GetFaxDocumentUnauthorized {
	return &GetFaxDocumentUnauthorized{}
}

/*
GetFaxDocumentUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFaxDocumentUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax document unauthorized response has a 2xx status code
func (o *GetFaxDocumentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax document unauthorized response has a 3xx status code
func (o *GetFaxDocumentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax document unauthorized response has a 4xx status code
func (o *GetFaxDocumentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fax document unauthorized response has a 5xx status code
func (o *GetFaxDocumentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get fax document unauthorized response a status code equal to that given
func (o *GetFaxDocumentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetFaxDocumentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFaxDocumentUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFaxDocumentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentForbidden creates a GetFaxDocumentForbidden with default headers values
func NewGetFaxDocumentForbidden() *GetFaxDocumentForbidden {
	return &GetFaxDocumentForbidden{}
}

/*
GetFaxDocumentForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetFaxDocumentForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax document forbidden response has a 2xx status code
func (o *GetFaxDocumentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax document forbidden response has a 3xx status code
func (o *GetFaxDocumentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax document forbidden response has a 4xx status code
func (o *GetFaxDocumentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fax document forbidden response has a 5xx status code
func (o *GetFaxDocumentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get fax document forbidden response a status code equal to that given
func (o *GetFaxDocumentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetFaxDocumentForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentForbidden  %+v", 403, o.Payload)
}

func (o *GetFaxDocumentForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentForbidden  %+v", 403, o.Payload)
}

func (o *GetFaxDocumentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentNotFound creates a GetFaxDocumentNotFound with default headers values
func NewGetFaxDocumentNotFound() *GetFaxDocumentNotFound {
	return &GetFaxDocumentNotFound{}
}

/*
GetFaxDocumentNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetFaxDocumentNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax document not found response has a 2xx status code
func (o *GetFaxDocumentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax document not found response has a 3xx status code
func (o *GetFaxDocumentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax document not found response has a 4xx status code
func (o *GetFaxDocumentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fax document not found response has a 5xx status code
func (o *GetFaxDocumentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get fax document not found response a status code equal to that given
func (o *GetFaxDocumentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetFaxDocumentNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentNotFound  %+v", 404, o.Payload)
}

func (o *GetFaxDocumentNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentNotFound  %+v", 404, o.Payload)
}

func (o *GetFaxDocumentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentRequestTimeout creates a GetFaxDocumentRequestTimeout with default headers values
func NewGetFaxDocumentRequestTimeout() *GetFaxDocumentRequestTimeout {
	return &GetFaxDocumentRequestTimeout{}
}

/*
GetFaxDocumentRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetFaxDocumentRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax document request timeout response has a 2xx status code
func (o *GetFaxDocumentRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax document request timeout response has a 3xx status code
func (o *GetFaxDocumentRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax document request timeout response has a 4xx status code
func (o *GetFaxDocumentRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fax document request timeout response has a 5xx status code
func (o *GetFaxDocumentRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get fax document request timeout response a status code equal to that given
func (o *GetFaxDocumentRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetFaxDocumentRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFaxDocumentRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFaxDocumentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentRequestEntityTooLarge creates a GetFaxDocumentRequestEntityTooLarge with default headers values
func NewGetFaxDocumentRequestEntityTooLarge() *GetFaxDocumentRequestEntityTooLarge {
	return &GetFaxDocumentRequestEntityTooLarge{}
}

/*
GetFaxDocumentRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetFaxDocumentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax document request entity too large response has a 2xx status code
func (o *GetFaxDocumentRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax document request entity too large response has a 3xx status code
func (o *GetFaxDocumentRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax document request entity too large response has a 4xx status code
func (o *GetFaxDocumentRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fax document request entity too large response has a 5xx status code
func (o *GetFaxDocumentRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get fax document request entity too large response a status code equal to that given
func (o *GetFaxDocumentRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetFaxDocumentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFaxDocumentRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFaxDocumentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentUnsupportedMediaType creates a GetFaxDocumentUnsupportedMediaType with default headers values
func NewGetFaxDocumentUnsupportedMediaType() *GetFaxDocumentUnsupportedMediaType {
	return &GetFaxDocumentUnsupportedMediaType{}
}

/*
GetFaxDocumentUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFaxDocumentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax document unsupported media type response has a 2xx status code
func (o *GetFaxDocumentUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax document unsupported media type response has a 3xx status code
func (o *GetFaxDocumentUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax document unsupported media type response has a 4xx status code
func (o *GetFaxDocumentUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fax document unsupported media type response has a 5xx status code
func (o *GetFaxDocumentUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get fax document unsupported media type response a status code equal to that given
func (o *GetFaxDocumentUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetFaxDocumentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFaxDocumentUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFaxDocumentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentTooManyRequests creates a GetFaxDocumentTooManyRequests with default headers values
func NewGetFaxDocumentTooManyRequests() *GetFaxDocumentTooManyRequests {
	return &GetFaxDocumentTooManyRequests{}
}

/*
GetFaxDocumentTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetFaxDocumentTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax document too many requests response has a 2xx status code
func (o *GetFaxDocumentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax document too many requests response has a 3xx status code
func (o *GetFaxDocumentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax document too many requests response has a 4xx status code
func (o *GetFaxDocumentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fax document too many requests response has a 5xx status code
func (o *GetFaxDocumentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get fax document too many requests response a status code equal to that given
func (o *GetFaxDocumentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetFaxDocumentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFaxDocumentTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFaxDocumentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentInternalServerError creates a GetFaxDocumentInternalServerError with default headers values
func NewGetFaxDocumentInternalServerError() *GetFaxDocumentInternalServerError {
	return &GetFaxDocumentInternalServerError{}
}

/*
GetFaxDocumentInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFaxDocumentInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax document internal server error response has a 2xx status code
func (o *GetFaxDocumentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax document internal server error response has a 3xx status code
func (o *GetFaxDocumentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax document internal server error response has a 4xx status code
func (o *GetFaxDocumentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fax document internal server error response has a 5xx status code
func (o *GetFaxDocumentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get fax document internal server error response a status code equal to that given
func (o *GetFaxDocumentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetFaxDocumentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFaxDocumentInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFaxDocumentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentServiceUnavailable creates a GetFaxDocumentServiceUnavailable with default headers values
func NewGetFaxDocumentServiceUnavailable() *GetFaxDocumentServiceUnavailable {
	return &GetFaxDocumentServiceUnavailable{}
}

/*
GetFaxDocumentServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFaxDocumentServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax document service unavailable response has a 2xx status code
func (o *GetFaxDocumentServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax document service unavailable response has a 3xx status code
func (o *GetFaxDocumentServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax document service unavailable response has a 4xx status code
func (o *GetFaxDocumentServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fax document service unavailable response has a 5xx status code
func (o *GetFaxDocumentServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get fax document service unavailable response a status code equal to that given
func (o *GetFaxDocumentServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetFaxDocumentServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFaxDocumentServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFaxDocumentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentGatewayTimeout creates a GetFaxDocumentGatewayTimeout with default headers values
func NewGetFaxDocumentGatewayTimeout() *GetFaxDocumentGatewayTimeout {
	return &GetFaxDocumentGatewayTimeout{}
}

/*
GetFaxDocumentGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetFaxDocumentGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get fax document gateway timeout response has a 2xx status code
func (o *GetFaxDocumentGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fax document gateway timeout response has a 3xx status code
func (o *GetFaxDocumentGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fax document gateway timeout response has a 4xx status code
func (o *GetFaxDocumentGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fax document gateway timeout response has a 5xx status code
func (o *GetFaxDocumentGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get fax document gateway timeout response a status code equal to that given
func (o *GetFaxDocumentGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetFaxDocumentGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFaxDocumentGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFaxDocumentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
