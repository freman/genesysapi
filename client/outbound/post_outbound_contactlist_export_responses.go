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

// PostOutboundContactlistExportReader is a Reader for the PostOutboundContactlistExport structure.
type PostOutboundContactlistExportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOutboundContactlistExportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOutboundContactlistExportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOutboundContactlistExportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOutboundContactlistExportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOutboundContactlistExportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOutboundContactlistExportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostOutboundContactlistExportRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOutboundContactlistExportRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOutboundContactlistExportUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOutboundContactlistExportTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOutboundContactlistExportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOutboundContactlistExportServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOutboundContactlistExportGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOutboundContactlistExportOK creates a PostOutboundContactlistExportOK with default headers values
func NewPostOutboundContactlistExportOK() *PostOutboundContactlistExportOK {
	return &PostOutboundContactlistExportOK{}
}

/*
PostOutboundContactlistExportOK describes a response with status code 200, with default header values.

successful operation
*/
type PostOutboundContactlistExportOK struct {
	Payload *models.DomainEntityRef
}

// IsSuccess returns true when this post outbound contactlist export o k response has a 2xx status code
func (o *PostOutboundContactlistExportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post outbound contactlist export o k response has a 3xx status code
func (o *PostOutboundContactlistExportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist export o k response has a 4xx status code
func (o *PostOutboundContactlistExportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound contactlist export o k response has a 5xx status code
func (o *PostOutboundContactlistExportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound contactlist export o k response a status code equal to that given
func (o *PostOutboundContactlistExportOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostOutboundContactlistExportOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportOK  %+v", 200, o.Payload)
}

func (o *PostOutboundContactlistExportOK) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportOK  %+v", 200, o.Payload)
}

func (o *PostOutboundContactlistExportOK) GetPayload() *models.DomainEntityRef {
	return o.Payload
}

func (o *PostOutboundContactlistExportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainEntityRef)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistExportBadRequest creates a PostOutboundContactlistExportBadRequest with default headers values
func NewPostOutboundContactlistExportBadRequest() *PostOutboundContactlistExportBadRequest {
	return &PostOutboundContactlistExportBadRequest{}
}

/*
PostOutboundContactlistExportBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOutboundContactlistExportBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist export bad request response has a 2xx status code
func (o *PostOutboundContactlistExportBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist export bad request response has a 3xx status code
func (o *PostOutboundContactlistExportBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist export bad request response has a 4xx status code
func (o *PostOutboundContactlistExportBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound contactlist export bad request response has a 5xx status code
func (o *PostOutboundContactlistExportBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound contactlist export bad request response a status code equal to that given
func (o *PostOutboundContactlistExportBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostOutboundContactlistExportBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportBadRequest  %+v", 400, o.Payload)
}

func (o *PostOutboundContactlistExportBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportBadRequest  %+v", 400, o.Payload)
}

func (o *PostOutboundContactlistExportBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistExportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistExportUnauthorized creates a PostOutboundContactlistExportUnauthorized with default headers values
func NewPostOutboundContactlistExportUnauthorized() *PostOutboundContactlistExportUnauthorized {
	return &PostOutboundContactlistExportUnauthorized{}
}

/*
PostOutboundContactlistExportUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOutboundContactlistExportUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist export unauthorized response has a 2xx status code
func (o *PostOutboundContactlistExportUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist export unauthorized response has a 3xx status code
func (o *PostOutboundContactlistExportUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist export unauthorized response has a 4xx status code
func (o *PostOutboundContactlistExportUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound contactlist export unauthorized response has a 5xx status code
func (o *PostOutboundContactlistExportUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound contactlist export unauthorized response a status code equal to that given
func (o *PostOutboundContactlistExportUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostOutboundContactlistExportUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOutboundContactlistExportUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOutboundContactlistExportUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistExportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistExportForbidden creates a PostOutboundContactlistExportForbidden with default headers values
func NewPostOutboundContactlistExportForbidden() *PostOutboundContactlistExportForbidden {
	return &PostOutboundContactlistExportForbidden{}
}

/*
PostOutboundContactlistExportForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostOutboundContactlistExportForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist export forbidden response has a 2xx status code
func (o *PostOutboundContactlistExportForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist export forbidden response has a 3xx status code
func (o *PostOutboundContactlistExportForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist export forbidden response has a 4xx status code
func (o *PostOutboundContactlistExportForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound contactlist export forbidden response has a 5xx status code
func (o *PostOutboundContactlistExportForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound contactlist export forbidden response a status code equal to that given
func (o *PostOutboundContactlistExportForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostOutboundContactlistExportForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportForbidden  %+v", 403, o.Payload)
}

func (o *PostOutboundContactlistExportForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportForbidden  %+v", 403, o.Payload)
}

func (o *PostOutboundContactlistExportForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistExportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistExportNotFound creates a PostOutboundContactlistExportNotFound with default headers values
func NewPostOutboundContactlistExportNotFound() *PostOutboundContactlistExportNotFound {
	return &PostOutboundContactlistExportNotFound{}
}

/*
PostOutboundContactlistExportNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostOutboundContactlistExportNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist export not found response has a 2xx status code
func (o *PostOutboundContactlistExportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist export not found response has a 3xx status code
func (o *PostOutboundContactlistExportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist export not found response has a 4xx status code
func (o *PostOutboundContactlistExportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound contactlist export not found response has a 5xx status code
func (o *PostOutboundContactlistExportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound contactlist export not found response a status code equal to that given
func (o *PostOutboundContactlistExportNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostOutboundContactlistExportNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportNotFound  %+v", 404, o.Payload)
}

func (o *PostOutboundContactlistExportNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportNotFound  %+v", 404, o.Payload)
}

func (o *PostOutboundContactlistExportNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistExportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistExportRequestTimeout creates a PostOutboundContactlistExportRequestTimeout with default headers values
func NewPostOutboundContactlistExportRequestTimeout() *PostOutboundContactlistExportRequestTimeout {
	return &PostOutboundContactlistExportRequestTimeout{}
}

/*
PostOutboundContactlistExportRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostOutboundContactlistExportRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist export request timeout response has a 2xx status code
func (o *PostOutboundContactlistExportRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist export request timeout response has a 3xx status code
func (o *PostOutboundContactlistExportRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist export request timeout response has a 4xx status code
func (o *PostOutboundContactlistExportRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound contactlist export request timeout response has a 5xx status code
func (o *PostOutboundContactlistExportRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound contactlist export request timeout response a status code equal to that given
func (o *PostOutboundContactlistExportRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostOutboundContactlistExportRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOutboundContactlistExportRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOutboundContactlistExportRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistExportRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistExportRequestEntityTooLarge creates a PostOutboundContactlistExportRequestEntityTooLarge with default headers values
func NewPostOutboundContactlistExportRequestEntityTooLarge() *PostOutboundContactlistExportRequestEntityTooLarge {
	return &PostOutboundContactlistExportRequestEntityTooLarge{}
}

/*
PostOutboundContactlistExportRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostOutboundContactlistExportRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist export request entity too large response has a 2xx status code
func (o *PostOutboundContactlistExportRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist export request entity too large response has a 3xx status code
func (o *PostOutboundContactlistExportRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist export request entity too large response has a 4xx status code
func (o *PostOutboundContactlistExportRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound contactlist export request entity too large response has a 5xx status code
func (o *PostOutboundContactlistExportRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound contactlist export request entity too large response a status code equal to that given
func (o *PostOutboundContactlistExportRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostOutboundContactlistExportRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOutboundContactlistExportRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOutboundContactlistExportRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistExportRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistExportUnsupportedMediaType creates a PostOutboundContactlistExportUnsupportedMediaType with default headers values
func NewPostOutboundContactlistExportUnsupportedMediaType() *PostOutboundContactlistExportUnsupportedMediaType {
	return &PostOutboundContactlistExportUnsupportedMediaType{}
}

/*
PostOutboundContactlistExportUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOutboundContactlistExportUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist export unsupported media type response has a 2xx status code
func (o *PostOutboundContactlistExportUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist export unsupported media type response has a 3xx status code
func (o *PostOutboundContactlistExportUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist export unsupported media type response has a 4xx status code
func (o *PostOutboundContactlistExportUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound contactlist export unsupported media type response has a 5xx status code
func (o *PostOutboundContactlistExportUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound contactlist export unsupported media type response a status code equal to that given
func (o *PostOutboundContactlistExportUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostOutboundContactlistExportUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOutboundContactlistExportUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOutboundContactlistExportUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistExportUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistExportTooManyRequests creates a PostOutboundContactlistExportTooManyRequests with default headers values
func NewPostOutboundContactlistExportTooManyRequests() *PostOutboundContactlistExportTooManyRequests {
	return &PostOutboundContactlistExportTooManyRequests{}
}

/*
PostOutboundContactlistExportTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostOutboundContactlistExportTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist export too many requests response has a 2xx status code
func (o *PostOutboundContactlistExportTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist export too many requests response has a 3xx status code
func (o *PostOutboundContactlistExportTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist export too many requests response has a 4xx status code
func (o *PostOutboundContactlistExportTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound contactlist export too many requests response has a 5xx status code
func (o *PostOutboundContactlistExportTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound contactlist export too many requests response a status code equal to that given
func (o *PostOutboundContactlistExportTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostOutboundContactlistExportTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOutboundContactlistExportTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOutboundContactlistExportTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistExportTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistExportInternalServerError creates a PostOutboundContactlistExportInternalServerError with default headers values
func NewPostOutboundContactlistExportInternalServerError() *PostOutboundContactlistExportInternalServerError {
	return &PostOutboundContactlistExportInternalServerError{}
}

/*
PostOutboundContactlistExportInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOutboundContactlistExportInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist export internal server error response has a 2xx status code
func (o *PostOutboundContactlistExportInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist export internal server error response has a 3xx status code
func (o *PostOutboundContactlistExportInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist export internal server error response has a 4xx status code
func (o *PostOutboundContactlistExportInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound contactlist export internal server error response has a 5xx status code
func (o *PostOutboundContactlistExportInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post outbound contactlist export internal server error response a status code equal to that given
func (o *PostOutboundContactlistExportInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostOutboundContactlistExportInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOutboundContactlistExportInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOutboundContactlistExportInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistExportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistExportServiceUnavailable creates a PostOutboundContactlistExportServiceUnavailable with default headers values
func NewPostOutboundContactlistExportServiceUnavailable() *PostOutboundContactlistExportServiceUnavailable {
	return &PostOutboundContactlistExportServiceUnavailable{}
}

/*
PostOutboundContactlistExportServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOutboundContactlistExportServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist export service unavailable response has a 2xx status code
func (o *PostOutboundContactlistExportServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist export service unavailable response has a 3xx status code
func (o *PostOutboundContactlistExportServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist export service unavailable response has a 4xx status code
func (o *PostOutboundContactlistExportServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound contactlist export service unavailable response has a 5xx status code
func (o *PostOutboundContactlistExportServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post outbound contactlist export service unavailable response a status code equal to that given
func (o *PostOutboundContactlistExportServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostOutboundContactlistExportServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOutboundContactlistExportServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOutboundContactlistExportServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistExportServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistExportGatewayTimeout creates a PostOutboundContactlistExportGatewayTimeout with default headers values
func NewPostOutboundContactlistExportGatewayTimeout() *PostOutboundContactlistExportGatewayTimeout {
	return &PostOutboundContactlistExportGatewayTimeout{}
}

/*
PostOutboundContactlistExportGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostOutboundContactlistExportGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist export gateway timeout response has a 2xx status code
func (o *PostOutboundContactlistExportGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist export gateway timeout response has a 3xx status code
func (o *PostOutboundContactlistExportGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist export gateway timeout response has a 4xx status code
func (o *PostOutboundContactlistExportGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound contactlist export gateway timeout response has a 5xx status code
func (o *PostOutboundContactlistExportGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post outbound contactlist export gateway timeout response a status code equal to that given
func (o *PostOutboundContactlistExportGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostOutboundContactlistExportGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOutboundContactlistExportGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/export][%d] postOutboundContactlistExportGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOutboundContactlistExportGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistExportGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
