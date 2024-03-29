// Code generated by go-swagger; DO NOT EDIT.

package general_data_protection_regulation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostGdprRequestsReader is a Reader for the PostGdprRequests structure.
type PostGdprRequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostGdprRequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostGdprRequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostGdprRequestsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostGdprRequestsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostGdprRequestsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostGdprRequestsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostGdprRequestsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostGdprRequestsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostGdprRequestsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostGdprRequestsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostGdprRequestsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostGdprRequestsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostGdprRequestsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostGdprRequestsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostGdprRequestsOK creates a PostGdprRequestsOK with default headers values
func NewPostGdprRequestsOK() *PostGdprRequestsOK {
	return &PostGdprRequestsOK{}
}

/*
PostGdprRequestsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostGdprRequestsOK struct {
	Payload *models.GDPRRequest
}

// IsSuccess returns true when this post gdpr requests o k response has a 2xx status code
func (o *PostGdprRequestsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post gdpr requests o k response has a 3xx status code
func (o *PostGdprRequestsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gdpr requests o k response has a 4xx status code
func (o *PostGdprRequestsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post gdpr requests o k response has a 5xx status code
func (o *PostGdprRequestsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post gdpr requests o k response a status code equal to that given
func (o *PostGdprRequestsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostGdprRequestsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsOK  %+v", 200, o.Payload)
}

func (o *PostGdprRequestsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsOK  %+v", 200, o.Payload)
}

func (o *PostGdprRequestsOK) GetPayload() *models.GDPRRequest {
	return o.Payload
}

func (o *PostGdprRequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GDPRRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGdprRequestsAccepted creates a PostGdprRequestsAccepted with default headers values
func NewPostGdprRequestsAccepted() *PostGdprRequestsAccepted {
	return &PostGdprRequestsAccepted{}
}

/*
PostGdprRequestsAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PostGdprRequestsAccepted struct {
	Payload *models.GDPRRequest
}

// IsSuccess returns true when this post gdpr requests accepted response has a 2xx status code
func (o *PostGdprRequestsAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post gdpr requests accepted response has a 3xx status code
func (o *PostGdprRequestsAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gdpr requests accepted response has a 4xx status code
func (o *PostGdprRequestsAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this post gdpr requests accepted response has a 5xx status code
func (o *PostGdprRequestsAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this post gdpr requests accepted response a status code equal to that given
func (o *PostGdprRequestsAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PostGdprRequestsAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsAccepted  %+v", 202, o.Payload)
}

func (o *PostGdprRequestsAccepted) String() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsAccepted  %+v", 202, o.Payload)
}

func (o *PostGdprRequestsAccepted) GetPayload() *models.GDPRRequest {
	return o.Payload
}

func (o *PostGdprRequestsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GDPRRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGdprRequestsBadRequest creates a PostGdprRequestsBadRequest with default headers values
func NewPostGdprRequestsBadRequest() *PostGdprRequestsBadRequest {
	return &PostGdprRequestsBadRequest{}
}

/*
PostGdprRequestsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostGdprRequestsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gdpr requests bad request response has a 2xx status code
func (o *PostGdprRequestsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gdpr requests bad request response has a 3xx status code
func (o *PostGdprRequestsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gdpr requests bad request response has a 4xx status code
func (o *PostGdprRequestsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gdpr requests bad request response has a 5xx status code
func (o *PostGdprRequestsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post gdpr requests bad request response a status code equal to that given
func (o *PostGdprRequestsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostGdprRequestsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsBadRequest  %+v", 400, o.Payload)
}

func (o *PostGdprRequestsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsBadRequest  %+v", 400, o.Payload)
}

func (o *PostGdprRequestsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGdprRequestsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGdprRequestsUnauthorized creates a PostGdprRequestsUnauthorized with default headers values
func NewPostGdprRequestsUnauthorized() *PostGdprRequestsUnauthorized {
	return &PostGdprRequestsUnauthorized{}
}

/*
PostGdprRequestsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostGdprRequestsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gdpr requests unauthorized response has a 2xx status code
func (o *PostGdprRequestsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gdpr requests unauthorized response has a 3xx status code
func (o *PostGdprRequestsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gdpr requests unauthorized response has a 4xx status code
func (o *PostGdprRequestsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gdpr requests unauthorized response has a 5xx status code
func (o *PostGdprRequestsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post gdpr requests unauthorized response a status code equal to that given
func (o *PostGdprRequestsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostGdprRequestsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostGdprRequestsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostGdprRequestsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGdprRequestsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGdprRequestsForbidden creates a PostGdprRequestsForbidden with default headers values
func NewPostGdprRequestsForbidden() *PostGdprRequestsForbidden {
	return &PostGdprRequestsForbidden{}
}

/*
PostGdprRequestsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostGdprRequestsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gdpr requests forbidden response has a 2xx status code
func (o *PostGdprRequestsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gdpr requests forbidden response has a 3xx status code
func (o *PostGdprRequestsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gdpr requests forbidden response has a 4xx status code
func (o *PostGdprRequestsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gdpr requests forbidden response has a 5xx status code
func (o *PostGdprRequestsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post gdpr requests forbidden response a status code equal to that given
func (o *PostGdprRequestsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostGdprRequestsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsForbidden  %+v", 403, o.Payload)
}

func (o *PostGdprRequestsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsForbidden  %+v", 403, o.Payload)
}

func (o *PostGdprRequestsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGdprRequestsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGdprRequestsNotFound creates a PostGdprRequestsNotFound with default headers values
func NewPostGdprRequestsNotFound() *PostGdprRequestsNotFound {
	return &PostGdprRequestsNotFound{}
}

/*
PostGdprRequestsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostGdprRequestsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gdpr requests not found response has a 2xx status code
func (o *PostGdprRequestsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gdpr requests not found response has a 3xx status code
func (o *PostGdprRequestsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gdpr requests not found response has a 4xx status code
func (o *PostGdprRequestsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gdpr requests not found response has a 5xx status code
func (o *PostGdprRequestsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post gdpr requests not found response a status code equal to that given
func (o *PostGdprRequestsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostGdprRequestsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsNotFound  %+v", 404, o.Payload)
}

func (o *PostGdprRequestsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsNotFound  %+v", 404, o.Payload)
}

func (o *PostGdprRequestsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGdprRequestsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGdprRequestsRequestTimeout creates a PostGdprRequestsRequestTimeout with default headers values
func NewPostGdprRequestsRequestTimeout() *PostGdprRequestsRequestTimeout {
	return &PostGdprRequestsRequestTimeout{}
}

/*
PostGdprRequestsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostGdprRequestsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gdpr requests request timeout response has a 2xx status code
func (o *PostGdprRequestsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gdpr requests request timeout response has a 3xx status code
func (o *PostGdprRequestsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gdpr requests request timeout response has a 4xx status code
func (o *PostGdprRequestsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gdpr requests request timeout response has a 5xx status code
func (o *PostGdprRequestsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post gdpr requests request timeout response a status code equal to that given
func (o *PostGdprRequestsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostGdprRequestsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostGdprRequestsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostGdprRequestsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGdprRequestsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGdprRequestsRequestEntityTooLarge creates a PostGdprRequestsRequestEntityTooLarge with default headers values
func NewPostGdprRequestsRequestEntityTooLarge() *PostGdprRequestsRequestEntityTooLarge {
	return &PostGdprRequestsRequestEntityTooLarge{}
}

/*
PostGdprRequestsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostGdprRequestsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gdpr requests request entity too large response has a 2xx status code
func (o *PostGdprRequestsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gdpr requests request entity too large response has a 3xx status code
func (o *PostGdprRequestsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gdpr requests request entity too large response has a 4xx status code
func (o *PostGdprRequestsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gdpr requests request entity too large response has a 5xx status code
func (o *PostGdprRequestsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post gdpr requests request entity too large response a status code equal to that given
func (o *PostGdprRequestsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostGdprRequestsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostGdprRequestsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostGdprRequestsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGdprRequestsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGdprRequestsUnsupportedMediaType creates a PostGdprRequestsUnsupportedMediaType with default headers values
func NewPostGdprRequestsUnsupportedMediaType() *PostGdprRequestsUnsupportedMediaType {
	return &PostGdprRequestsUnsupportedMediaType{}
}

/*
PostGdprRequestsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostGdprRequestsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gdpr requests unsupported media type response has a 2xx status code
func (o *PostGdprRequestsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gdpr requests unsupported media type response has a 3xx status code
func (o *PostGdprRequestsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gdpr requests unsupported media type response has a 4xx status code
func (o *PostGdprRequestsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gdpr requests unsupported media type response has a 5xx status code
func (o *PostGdprRequestsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post gdpr requests unsupported media type response a status code equal to that given
func (o *PostGdprRequestsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostGdprRequestsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostGdprRequestsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostGdprRequestsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGdprRequestsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGdprRequestsTooManyRequests creates a PostGdprRequestsTooManyRequests with default headers values
func NewPostGdprRequestsTooManyRequests() *PostGdprRequestsTooManyRequests {
	return &PostGdprRequestsTooManyRequests{}
}

/*
PostGdprRequestsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostGdprRequestsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gdpr requests too many requests response has a 2xx status code
func (o *PostGdprRequestsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gdpr requests too many requests response has a 3xx status code
func (o *PostGdprRequestsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gdpr requests too many requests response has a 4xx status code
func (o *PostGdprRequestsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gdpr requests too many requests response has a 5xx status code
func (o *PostGdprRequestsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post gdpr requests too many requests response a status code equal to that given
func (o *PostGdprRequestsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostGdprRequestsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostGdprRequestsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostGdprRequestsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGdprRequestsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGdprRequestsInternalServerError creates a PostGdprRequestsInternalServerError with default headers values
func NewPostGdprRequestsInternalServerError() *PostGdprRequestsInternalServerError {
	return &PostGdprRequestsInternalServerError{}
}

/*
PostGdprRequestsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostGdprRequestsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gdpr requests internal server error response has a 2xx status code
func (o *PostGdprRequestsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gdpr requests internal server error response has a 3xx status code
func (o *PostGdprRequestsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gdpr requests internal server error response has a 4xx status code
func (o *PostGdprRequestsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post gdpr requests internal server error response has a 5xx status code
func (o *PostGdprRequestsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post gdpr requests internal server error response a status code equal to that given
func (o *PostGdprRequestsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostGdprRequestsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostGdprRequestsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostGdprRequestsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGdprRequestsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGdprRequestsServiceUnavailable creates a PostGdprRequestsServiceUnavailable with default headers values
func NewPostGdprRequestsServiceUnavailable() *PostGdprRequestsServiceUnavailable {
	return &PostGdprRequestsServiceUnavailable{}
}

/*
PostGdprRequestsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostGdprRequestsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gdpr requests service unavailable response has a 2xx status code
func (o *PostGdprRequestsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gdpr requests service unavailable response has a 3xx status code
func (o *PostGdprRequestsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gdpr requests service unavailable response has a 4xx status code
func (o *PostGdprRequestsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post gdpr requests service unavailable response has a 5xx status code
func (o *PostGdprRequestsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post gdpr requests service unavailable response a status code equal to that given
func (o *PostGdprRequestsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostGdprRequestsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostGdprRequestsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostGdprRequestsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGdprRequestsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGdprRequestsGatewayTimeout creates a PostGdprRequestsGatewayTimeout with default headers values
func NewPostGdprRequestsGatewayTimeout() *PostGdprRequestsGatewayTimeout {
	return &PostGdprRequestsGatewayTimeout{}
}

/*
PostGdprRequestsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostGdprRequestsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gdpr requests gateway timeout response has a 2xx status code
func (o *PostGdprRequestsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gdpr requests gateway timeout response has a 3xx status code
func (o *PostGdprRequestsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gdpr requests gateway timeout response has a 4xx status code
func (o *PostGdprRequestsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post gdpr requests gateway timeout response has a 5xx status code
func (o *PostGdprRequestsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post gdpr requests gateway timeout response a status code equal to that given
func (o *PostGdprRequestsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostGdprRequestsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostGdprRequestsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/gdpr/requests][%d] postGdprRequestsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostGdprRequestsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGdprRequestsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
