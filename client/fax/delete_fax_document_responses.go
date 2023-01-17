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

// DeleteFaxDocumentReader is a Reader for the DeleteFaxDocument structure.
type DeleteFaxDocumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFaxDocumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteFaxDocumentAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteFaxDocumentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteFaxDocumentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteFaxDocumentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteFaxDocumentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteFaxDocumentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteFaxDocumentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteFaxDocumentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteFaxDocumentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteFaxDocumentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteFaxDocumentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteFaxDocumentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteFaxDocumentAccepted creates a DeleteFaxDocumentAccepted with default headers values
func NewDeleteFaxDocumentAccepted() *DeleteFaxDocumentAccepted {
	return &DeleteFaxDocumentAccepted{}
}

/*
DeleteFaxDocumentAccepted describes a response with status code 202, with default header values.

Accepted - Processing Delete
*/
type DeleteFaxDocumentAccepted struct {
}

// IsSuccess returns true when this delete fax document accepted response has a 2xx status code
func (o *DeleteFaxDocumentAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete fax document accepted response has a 3xx status code
func (o *DeleteFaxDocumentAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fax document accepted response has a 4xx status code
func (o *DeleteFaxDocumentAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete fax document accepted response has a 5xx status code
func (o *DeleteFaxDocumentAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fax document accepted response a status code equal to that given
func (o *DeleteFaxDocumentAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *DeleteFaxDocumentAccepted) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentAccepted ", 202)
}

func (o *DeleteFaxDocumentAccepted) String() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentAccepted ", 202)
}

func (o *DeleteFaxDocumentAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFaxDocumentBadRequest creates a DeleteFaxDocumentBadRequest with default headers values
func NewDeleteFaxDocumentBadRequest() *DeleteFaxDocumentBadRequest {
	return &DeleteFaxDocumentBadRequest{}
}

/*
DeleteFaxDocumentBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteFaxDocumentBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete fax document bad request response has a 2xx status code
func (o *DeleteFaxDocumentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fax document bad request response has a 3xx status code
func (o *DeleteFaxDocumentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fax document bad request response has a 4xx status code
func (o *DeleteFaxDocumentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete fax document bad request response has a 5xx status code
func (o *DeleteFaxDocumentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fax document bad request response a status code equal to that given
func (o *DeleteFaxDocumentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteFaxDocumentBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFaxDocumentBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFaxDocumentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFaxDocumentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFaxDocumentUnauthorized creates a DeleteFaxDocumentUnauthorized with default headers values
func NewDeleteFaxDocumentUnauthorized() *DeleteFaxDocumentUnauthorized {
	return &DeleteFaxDocumentUnauthorized{}
}

/*
DeleteFaxDocumentUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteFaxDocumentUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete fax document unauthorized response has a 2xx status code
func (o *DeleteFaxDocumentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fax document unauthorized response has a 3xx status code
func (o *DeleteFaxDocumentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fax document unauthorized response has a 4xx status code
func (o *DeleteFaxDocumentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete fax document unauthorized response has a 5xx status code
func (o *DeleteFaxDocumentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fax document unauthorized response a status code equal to that given
func (o *DeleteFaxDocumentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteFaxDocumentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteFaxDocumentUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteFaxDocumentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFaxDocumentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFaxDocumentForbidden creates a DeleteFaxDocumentForbidden with default headers values
func NewDeleteFaxDocumentForbidden() *DeleteFaxDocumentForbidden {
	return &DeleteFaxDocumentForbidden{}
}

/*
DeleteFaxDocumentForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteFaxDocumentForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete fax document forbidden response has a 2xx status code
func (o *DeleteFaxDocumentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fax document forbidden response has a 3xx status code
func (o *DeleteFaxDocumentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fax document forbidden response has a 4xx status code
func (o *DeleteFaxDocumentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete fax document forbidden response has a 5xx status code
func (o *DeleteFaxDocumentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fax document forbidden response a status code equal to that given
func (o *DeleteFaxDocumentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteFaxDocumentForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFaxDocumentForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFaxDocumentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFaxDocumentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFaxDocumentNotFound creates a DeleteFaxDocumentNotFound with default headers values
func NewDeleteFaxDocumentNotFound() *DeleteFaxDocumentNotFound {
	return &DeleteFaxDocumentNotFound{}
}

/*
DeleteFaxDocumentNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteFaxDocumentNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete fax document not found response has a 2xx status code
func (o *DeleteFaxDocumentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fax document not found response has a 3xx status code
func (o *DeleteFaxDocumentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fax document not found response has a 4xx status code
func (o *DeleteFaxDocumentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete fax document not found response has a 5xx status code
func (o *DeleteFaxDocumentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fax document not found response a status code equal to that given
func (o *DeleteFaxDocumentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteFaxDocumentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFaxDocumentNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFaxDocumentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFaxDocumentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFaxDocumentRequestTimeout creates a DeleteFaxDocumentRequestTimeout with default headers values
func NewDeleteFaxDocumentRequestTimeout() *DeleteFaxDocumentRequestTimeout {
	return &DeleteFaxDocumentRequestTimeout{}
}

/*
DeleteFaxDocumentRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteFaxDocumentRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete fax document request timeout response has a 2xx status code
func (o *DeleteFaxDocumentRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fax document request timeout response has a 3xx status code
func (o *DeleteFaxDocumentRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fax document request timeout response has a 4xx status code
func (o *DeleteFaxDocumentRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete fax document request timeout response has a 5xx status code
func (o *DeleteFaxDocumentRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fax document request timeout response a status code equal to that given
func (o *DeleteFaxDocumentRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteFaxDocumentRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteFaxDocumentRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteFaxDocumentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFaxDocumentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFaxDocumentRequestEntityTooLarge creates a DeleteFaxDocumentRequestEntityTooLarge with default headers values
func NewDeleteFaxDocumentRequestEntityTooLarge() *DeleteFaxDocumentRequestEntityTooLarge {
	return &DeleteFaxDocumentRequestEntityTooLarge{}
}

/*
DeleteFaxDocumentRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteFaxDocumentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete fax document request entity too large response has a 2xx status code
func (o *DeleteFaxDocumentRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fax document request entity too large response has a 3xx status code
func (o *DeleteFaxDocumentRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fax document request entity too large response has a 4xx status code
func (o *DeleteFaxDocumentRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete fax document request entity too large response has a 5xx status code
func (o *DeleteFaxDocumentRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fax document request entity too large response a status code equal to that given
func (o *DeleteFaxDocumentRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteFaxDocumentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteFaxDocumentRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteFaxDocumentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFaxDocumentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFaxDocumentUnsupportedMediaType creates a DeleteFaxDocumentUnsupportedMediaType with default headers values
func NewDeleteFaxDocumentUnsupportedMediaType() *DeleteFaxDocumentUnsupportedMediaType {
	return &DeleteFaxDocumentUnsupportedMediaType{}
}

/*
DeleteFaxDocumentUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteFaxDocumentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete fax document unsupported media type response has a 2xx status code
func (o *DeleteFaxDocumentUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fax document unsupported media type response has a 3xx status code
func (o *DeleteFaxDocumentUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fax document unsupported media type response has a 4xx status code
func (o *DeleteFaxDocumentUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete fax document unsupported media type response has a 5xx status code
func (o *DeleteFaxDocumentUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fax document unsupported media type response a status code equal to that given
func (o *DeleteFaxDocumentUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteFaxDocumentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteFaxDocumentUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteFaxDocumentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFaxDocumentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFaxDocumentTooManyRequests creates a DeleteFaxDocumentTooManyRequests with default headers values
func NewDeleteFaxDocumentTooManyRequests() *DeleteFaxDocumentTooManyRequests {
	return &DeleteFaxDocumentTooManyRequests{}
}

/*
DeleteFaxDocumentTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteFaxDocumentTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete fax document too many requests response has a 2xx status code
func (o *DeleteFaxDocumentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fax document too many requests response has a 3xx status code
func (o *DeleteFaxDocumentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fax document too many requests response has a 4xx status code
func (o *DeleteFaxDocumentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete fax document too many requests response has a 5xx status code
func (o *DeleteFaxDocumentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fax document too many requests response a status code equal to that given
func (o *DeleteFaxDocumentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteFaxDocumentTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteFaxDocumentTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteFaxDocumentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFaxDocumentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFaxDocumentInternalServerError creates a DeleteFaxDocumentInternalServerError with default headers values
func NewDeleteFaxDocumentInternalServerError() *DeleteFaxDocumentInternalServerError {
	return &DeleteFaxDocumentInternalServerError{}
}

/*
DeleteFaxDocumentInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteFaxDocumentInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete fax document internal server error response has a 2xx status code
func (o *DeleteFaxDocumentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fax document internal server error response has a 3xx status code
func (o *DeleteFaxDocumentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fax document internal server error response has a 4xx status code
func (o *DeleteFaxDocumentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete fax document internal server error response has a 5xx status code
func (o *DeleteFaxDocumentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete fax document internal server error response a status code equal to that given
func (o *DeleteFaxDocumentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteFaxDocumentInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteFaxDocumentInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteFaxDocumentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFaxDocumentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFaxDocumentServiceUnavailable creates a DeleteFaxDocumentServiceUnavailable with default headers values
func NewDeleteFaxDocumentServiceUnavailable() *DeleteFaxDocumentServiceUnavailable {
	return &DeleteFaxDocumentServiceUnavailable{}
}

/*
DeleteFaxDocumentServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteFaxDocumentServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete fax document service unavailable response has a 2xx status code
func (o *DeleteFaxDocumentServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fax document service unavailable response has a 3xx status code
func (o *DeleteFaxDocumentServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fax document service unavailable response has a 4xx status code
func (o *DeleteFaxDocumentServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete fax document service unavailable response has a 5xx status code
func (o *DeleteFaxDocumentServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete fax document service unavailable response a status code equal to that given
func (o *DeleteFaxDocumentServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteFaxDocumentServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteFaxDocumentServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteFaxDocumentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFaxDocumentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFaxDocumentGatewayTimeout creates a DeleteFaxDocumentGatewayTimeout with default headers values
func NewDeleteFaxDocumentGatewayTimeout() *DeleteFaxDocumentGatewayTimeout {
	return &DeleteFaxDocumentGatewayTimeout{}
}

/*
DeleteFaxDocumentGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteFaxDocumentGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete fax document gateway timeout response has a 2xx status code
func (o *DeleteFaxDocumentGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fax document gateway timeout response has a 3xx status code
func (o *DeleteFaxDocumentGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fax document gateway timeout response has a 4xx status code
func (o *DeleteFaxDocumentGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete fax document gateway timeout response has a 5xx status code
func (o *DeleteFaxDocumentGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete fax document gateway timeout response a status code equal to that given
func (o *DeleteFaxDocumentGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteFaxDocumentGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteFaxDocumentGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/fax/documents/{documentId}][%d] deleteFaxDocumentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteFaxDocumentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFaxDocumentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
