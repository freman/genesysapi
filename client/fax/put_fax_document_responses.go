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

// PutFaxDocumentReader is a Reader for the PutFaxDocument structure.
type PutFaxDocumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutFaxDocumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutFaxDocumentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutFaxDocumentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutFaxDocumentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutFaxDocumentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutFaxDocumentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutFaxDocumentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutFaxDocumentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutFaxDocumentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutFaxDocumentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutFaxDocumentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutFaxDocumentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutFaxDocumentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutFaxDocumentOK creates a PutFaxDocumentOK with default headers values
func NewPutFaxDocumentOK() *PutFaxDocumentOK {
	return &PutFaxDocumentOK{}
}

/*
PutFaxDocumentOK describes a response with status code 200, with default header values.

successful operation
*/
type PutFaxDocumentOK struct {
	Payload *models.FaxDocument
}

// IsSuccess returns true when this put fax document o k response has a 2xx status code
func (o *PutFaxDocumentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put fax document o k response has a 3xx status code
func (o *PutFaxDocumentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fax document o k response has a 4xx status code
func (o *PutFaxDocumentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put fax document o k response has a 5xx status code
func (o *PutFaxDocumentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put fax document o k response a status code equal to that given
func (o *PutFaxDocumentOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutFaxDocumentOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentOK  %+v", 200, o.Payload)
}

func (o *PutFaxDocumentOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentOK  %+v", 200, o.Payload)
}

func (o *PutFaxDocumentOK) GetPayload() *models.FaxDocument {
	return o.Payload
}

func (o *PutFaxDocumentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FaxDocument)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFaxDocumentBadRequest creates a PutFaxDocumentBadRequest with default headers values
func NewPutFaxDocumentBadRequest() *PutFaxDocumentBadRequest {
	return &PutFaxDocumentBadRequest{}
}

/*
PutFaxDocumentBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutFaxDocumentBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put fax document bad request response has a 2xx status code
func (o *PutFaxDocumentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fax document bad request response has a 3xx status code
func (o *PutFaxDocumentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fax document bad request response has a 4xx status code
func (o *PutFaxDocumentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put fax document bad request response has a 5xx status code
func (o *PutFaxDocumentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put fax document bad request response a status code equal to that given
func (o *PutFaxDocumentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutFaxDocumentBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentBadRequest  %+v", 400, o.Payload)
}

func (o *PutFaxDocumentBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentBadRequest  %+v", 400, o.Payload)
}

func (o *PutFaxDocumentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFaxDocumentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFaxDocumentUnauthorized creates a PutFaxDocumentUnauthorized with default headers values
func NewPutFaxDocumentUnauthorized() *PutFaxDocumentUnauthorized {
	return &PutFaxDocumentUnauthorized{}
}

/*
PutFaxDocumentUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutFaxDocumentUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put fax document unauthorized response has a 2xx status code
func (o *PutFaxDocumentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fax document unauthorized response has a 3xx status code
func (o *PutFaxDocumentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fax document unauthorized response has a 4xx status code
func (o *PutFaxDocumentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put fax document unauthorized response has a 5xx status code
func (o *PutFaxDocumentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put fax document unauthorized response a status code equal to that given
func (o *PutFaxDocumentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutFaxDocumentUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentUnauthorized  %+v", 401, o.Payload)
}

func (o *PutFaxDocumentUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentUnauthorized  %+v", 401, o.Payload)
}

func (o *PutFaxDocumentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFaxDocumentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFaxDocumentForbidden creates a PutFaxDocumentForbidden with default headers values
func NewPutFaxDocumentForbidden() *PutFaxDocumentForbidden {
	return &PutFaxDocumentForbidden{}
}

/*
PutFaxDocumentForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutFaxDocumentForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put fax document forbidden response has a 2xx status code
func (o *PutFaxDocumentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fax document forbidden response has a 3xx status code
func (o *PutFaxDocumentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fax document forbidden response has a 4xx status code
func (o *PutFaxDocumentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put fax document forbidden response has a 5xx status code
func (o *PutFaxDocumentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put fax document forbidden response a status code equal to that given
func (o *PutFaxDocumentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutFaxDocumentForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentForbidden  %+v", 403, o.Payload)
}

func (o *PutFaxDocumentForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentForbidden  %+v", 403, o.Payload)
}

func (o *PutFaxDocumentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFaxDocumentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFaxDocumentNotFound creates a PutFaxDocumentNotFound with default headers values
func NewPutFaxDocumentNotFound() *PutFaxDocumentNotFound {
	return &PutFaxDocumentNotFound{}
}

/*
PutFaxDocumentNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutFaxDocumentNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put fax document not found response has a 2xx status code
func (o *PutFaxDocumentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fax document not found response has a 3xx status code
func (o *PutFaxDocumentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fax document not found response has a 4xx status code
func (o *PutFaxDocumentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put fax document not found response has a 5xx status code
func (o *PutFaxDocumentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put fax document not found response a status code equal to that given
func (o *PutFaxDocumentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutFaxDocumentNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentNotFound  %+v", 404, o.Payload)
}

func (o *PutFaxDocumentNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentNotFound  %+v", 404, o.Payload)
}

func (o *PutFaxDocumentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFaxDocumentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFaxDocumentRequestTimeout creates a PutFaxDocumentRequestTimeout with default headers values
func NewPutFaxDocumentRequestTimeout() *PutFaxDocumentRequestTimeout {
	return &PutFaxDocumentRequestTimeout{}
}

/*
PutFaxDocumentRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutFaxDocumentRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put fax document request timeout response has a 2xx status code
func (o *PutFaxDocumentRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fax document request timeout response has a 3xx status code
func (o *PutFaxDocumentRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fax document request timeout response has a 4xx status code
func (o *PutFaxDocumentRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put fax document request timeout response has a 5xx status code
func (o *PutFaxDocumentRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put fax document request timeout response a status code equal to that given
func (o *PutFaxDocumentRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutFaxDocumentRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutFaxDocumentRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutFaxDocumentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFaxDocumentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFaxDocumentRequestEntityTooLarge creates a PutFaxDocumentRequestEntityTooLarge with default headers values
func NewPutFaxDocumentRequestEntityTooLarge() *PutFaxDocumentRequestEntityTooLarge {
	return &PutFaxDocumentRequestEntityTooLarge{}
}

/*
PutFaxDocumentRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutFaxDocumentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put fax document request entity too large response has a 2xx status code
func (o *PutFaxDocumentRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fax document request entity too large response has a 3xx status code
func (o *PutFaxDocumentRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fax document request entity too large response has a 4xx status code
func (o *PutFaxDocumentRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put fax document request entity too large response has a 5xx status code
func (o *PutFaxDocumentRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put fax document request entity too large response a status code equal to that given
func (o *PutFaxDocumentRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutFaxDocumentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutFaxDocumentRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutFaxDocumentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFaxDocumentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFaxDocumentUnsupportedMediaType creates a PutFaxDocumentUnsupportedMediaType with default headers values
func NewPutFaxDocumentUnsupportedMediaType() *PutFaxDocumentUnsupportedMediaType {
	return &PutFaxDocumentUnsupportedMediaType{}
}

/*
PutFaxDocumentUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutFaxDocumentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put fax document unsupported media type response has a 2xx status code
func (o *PutFaxDocumentUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fax document unsupported media type response has a 3xx status code
func (o *PutFaxDocumentUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fax document unsupported media type response has a 4xx status code
func (o *PutFaxDocumentUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put fax document unsupported media type response has a 5xx status code
func (o *PutFaxDocumentUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put fax document unsupported media type response a status code equal to that given
func (o *PutFaxDocumentUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutFaxDocumentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutFaxDocumentUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutFaxDocumentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFaxDocumentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFaxDocumentTooManyRequests creates a PutFaxDocumentTooManyRequests with default headers values
func NewPutFaxDocumentTooManyRequests() *PutFaxDocumentTooManyRequests {
	return &PutFaxDocumentTooManyRequests{}
}

/*
PutFaxDocumentTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutFaxDocumentTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put fax document too many requests response has a 2xx status code
func (o *PutFaxDocumentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fax document too many requests response has a 3xx status code
func (o *PutFaxDocumentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fax document too many requests response has a 4xx status code
func (o *PutFaxDocumentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put fax document too many requests response has a 5xx status code
func (o *PutFaxDocumentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put fax document too many requests response a status code equal to that given
func (o *PutFaxDocumentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutFaxDocumentTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutFaxDocumentTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutFaxDocumentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFaxDocumentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFaxDocumentInternalServerError creates a PutFaxDocumentInternalServerError with default headers values
func NewPutFaxDocumentInternalServerError() *PutFaxDocumentInternalServerError {
	return &PutFaxDocumentInternalServerError{}
}

/*
PutFaxDocumentInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutFaxDocumentInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put fax document internal server error response has a 2xx status code
func (o *PutFaxDocumentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fax document internal server error response has a 3xx status code
func (o *PutFaxDocumentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fax document internal server error response has a 4xx status code
func (o *PutFaxDocumentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put fax document internal server error response has a 5xx status code
func (o *PutFaxDocumentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put fax document internal server error response a status code equal to that given
func (o *PutFaxDocumentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutFaxDocumentInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *PutFaxDocumentInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *PutFaxDocumentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFaxDocumentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFaxDocumentServiceUnavailable creates a PutFaxDocumentServiceUnavailable with default headers values
func NewPutFaxDocumentServiceUnavailable() *PutFaxDocumentServiceUnavailable {
	return &PutFaxDocumentServiceUnavailable{}
}

/*
PutFaxDocumentServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutFaxDocumentServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put fax document service unavailable response has a 2xx status code
func (o *PutFaxDocumentServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fax document service unavailable response has a 3xx status code
func (o *PutFaxDocumentServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fax document service unavailable response has a 4xx status code
func (o *PutFaxDocumentServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put fax document service unavailable response has a 5xx status code
func (o *PutFaxDocumentServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put fax document service unavailable response a status code equal to that given
func (o *PutFaxDocumentServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutFaxDocumentServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutFaxDocumentServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutFaxDocumentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFaxDocumentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFaxDocumentGatewayTimeout creates a PutFaxDocumentGatewayTimeout with default headers values
func NewPutFaxDocumentGatewayTimeout() *PutFaxDocumentGatewayTimeout {
	return &PutFaxDocumentGatewayTimeout{}
}

/*
PutFaxDocumentGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutFaxDocumentGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put fax document gateway timeout response has a 2xx status code
func (o *PutFaxDocumentGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fax document gateway timeout response has a 3xx status code
func (o *PutFaxDocumentGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fax document gateway timeout response has a 4xx status code
func (o *PutFaxDocumentGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put fax document gateway timeout response has a 5xx status code
func (o *PutFaxDocumentGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put fax document gateway timeout response a status code equal to that given
func (o *PutFaxDocumentGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutFaxDocumentGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutFaxDocumentGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/fax/documents/{documentId}][%d] putFaxDocumentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutFaxDocumentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFaxDocumentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
