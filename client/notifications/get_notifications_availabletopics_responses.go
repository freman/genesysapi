// Code generated by go-swagger; DO NOT EDIT.

package notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetNotificationsAvailabletopicsReader is a Reader for the GetNotificationsAvailabletopics structure.
type GetNotificationsAvailabletopicsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNotificationsAvailabletopicsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNotificationsAvailabletopicsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNotificationsAvailabletopicsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetNotificationsAvailabletopicsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNotificationsAvailabletopicsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetNotificationsAvailabletopicsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetNotificationsAvailabletopicsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetNotificationsAvailabletopicsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetNotificationsAvailabletopicsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetNotificationsAvailabletopicsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNotificationsAvailabletopicsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetNotificationsAvailabletopicsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetNotificationsAvailabletopicsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNotificationsAvailabletopicsOK creates a GetNotificationsAvailabletopicsOK with default headers values
func NewGetNotificationsAvailabletopicsOK() *GetNotificationsAvailabletopicsOK {
	return &GetNotificationsAvailabletopicsOK{}
}

/*
GetNotificationsAvailabletopicsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetNotificationsAvailabletopicsOK struct {
	Payload *models.AvailableTopicEntityListing
}

// IsSuccess returns true when this get notifications availabletopics o k response has a 2xx status code
func (o *GetNotificationsAvailabletopicsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get notifications availabletopics o k response has a 3xx status code
func (o *GetNotificationsAvailabletopicsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications availabletopics o k response has a 4xx status code
func (o *GetNotificationsAvailabletopicsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get notifications availabletopics o k response has a 5xx status code
func (o *GetNotificationsAvailabletopicsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications availabletopics o k response a status code equal to that given
func (o *GetNotificationsAvailabletopicsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNotificationsAvailabletopicsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsOK  %+v", 200, o.Payload)
}

func (o *GetNotificationsAvailabletopicsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsOK  %+v", 200, o.Payload)
}

func (o *GetNotificationsAvailabletopicsOK) GetPayload() *models.AvailableTopicEntityListing {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AvailableTopicEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsAvailabletopicsBadRequest creates a GetNotificationsAvailabletopicsBadRequest with default headers values
func NewGetNotificationsAvailabletopicsBadRequest() *GetNotificationsAvailabletopicsBadRequest {
	return &GetNotificationsAvailabletopicsBadRequest{}
}

/*
GetNotificationsAvailabletopicsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetNotificationsAvailabletopicsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications availabletopics bad request response has a 2xx status code
func (o *GetNotificationsAvailabletopicsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications availabletopics bad request response has a 3xx status code
func (o *GetNotificationsAvailabletopicsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications availabletopics bad request response has a 4xx status code
func (o *GetNotificationsAvailabletopicsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get notifications availabletopics bad request response has a 5xx status code
func (o *GetNotificationsAvailabletopicsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications availabletopics bad request response a status code equal to that given
func (o *GetNotificationsAvailabletopicsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetNotificationsAvailabletopicsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsBadRequest  %+v", 400, o.Payload)
}

func (o *GetNotificationsAvailabletopicsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsBadRequest  %+v", 400, o.Payload)
}

func (o *GetNotificationsAvailabletopicsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsAvailabletopicsUnauthorized creates a GetNotificationsAvailabletopicsUnauthorized with default headers values
func NewGetNotificationsAvailabletopicsUnauthorized() *GetNotificationsAvailabletopicsUnauthorized {
	return &GetNotificationsAvailabletopicsUnauthorized{}
}

/*
GetNotificationsAvailabletopicsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetNotificationsAvailabletopicsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications availabletopics unauthorized response has a 2xx status code
func (o *GetNotificationsAvailabletopicsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications availabletopics unauthorized response has a 3xx status code
func (o *GetNotificationsAvailabletopicsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications availabletopics unauthorized response has a 4xx status code
func (o *GetNotificationsAvailabletopicsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get notifications availabletopics unauthorized response has a 5xx status code
func (o *GetNotificationsAvailabletopicsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications availabletopics unauthorized response a status code equal to that given
func (o *GetNotificationsAvailabletopicsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetNotificationsAvailabletopicsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetNotificationsAvailabletopicsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetNotificationsAvailabletopicsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsAvailabletopicsForbidden creates a GetNotificationsAvailabletopicsForbidden with default headers values
func NewGetNotificationsAvailabletopicsForbidden() *GetNotificationsAvailabletopicsForbidden {
	return &GetNotificationsAvailabletopicsForbidden{}
}

/*
GetNotificationsAvailabletopicsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetNotificationsAvailabletopicsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications availabletopics forbidden response has a 2xx status code
func (o *GetNotificationsAvailabletopicsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications availabletopics forbidden response has a 3xx status code
func (o *GetNotificationsAvailabletopicsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications availabletopics forbidden response has a 4xx status code
func (o *GetNotificationsAvailabletopicsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get notifications availabletopics forbidden response has a 5xx status code
func (o *GetNotificationsAvailabletopicsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications availabletopics forbidden response a status code equal to that given
func (o *GetNotificationsAvailabletopicsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetNotificationsAvailabletopicsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsForbidden  %+v", 403, o.Payload)
}

func (o *GetNotificationsAvailabletopicsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsForbidden  %+v", 403, o.Payload)
}

func (o *GetNotificationsAvailabletopicsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsAvailabletopicsNotFound creates a GetNotificationsAvailabletopicsNotFound with default headers values
func NewGetNotificationsAvailabletopicsNotFound() *GetNotificationsAvailabletopicsNotFound {
	return &GetNotificationsAvailabletopicsNotFound{}
}

/*
GetNotificationsAvailabletopicsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetNotificationsAvailabletopicsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications availabletopics not found response has a 2xx status code
func (o *GetNotificationsAvailabletopicsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications availabletopics not found response has a 3xx status code
func (o *GetNotificationsAvailabletopicsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications availabletopics not found response has a 4xx status code
func (o *GetNotificationsAvailabletopicsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get notifications availabletopics not found response has a 5xx status code
func (o *GetNotificationsAvailabletopicsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications availabletopics not found response a status code equal to that given
func (o *GetNotificationsAvailabletopicsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetNotificationsAvailabletopicsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsNotFound  %+v", 404, o.Payload)
}

func (o *GetNotificationsAvailabletopicsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsNotFound  %+v", 404, o.Payload)
}

func (o *GetNotificationsAvailabletopicsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsAvailabletopicsRequestTimeout creates a GetNotificationsAvailabletopicsRequestTimeout with default headers values
func NewGetNotificationsAvailabletopicsRequestTimeout() *GetNotificationsAvailabletopicsRequestTimeout {
	return &GetNotificationsAvailabletopicsRequestTimeout{}
}

/*
GetNotificationsAvailabletopicsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetNotificationsAvailabletopicsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications availabletopics request timeout response has a 2xx status code
func (o *GetNotificationsAvailabletopicsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications availabletopics request timeout response has a 3xx status code
func (o *GetNotificationsAvailabletopicsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications availabletopics request timeout response has a 4xx status code
func (o *GetNotificationsAvailabletopicsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get notifications availabletopics request timeout response has a 5xx status code
func (o *GetNotificationsAvailabletopicsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications availabletopics request timeout response a status code equal to that given
func (o *GetNotificationsAvailabletopicsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetNotificationsAvailabletopicsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetNotificationsAvailabletopicsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetNotificationsAvailabletopicsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsAvailabletopicsRequestEntityTooLarge creates a GetNotificationsAvailabletopicsRequestEntityTooLarge with default headers values
func NewGetNotificationsAvailabletopicsRequestEntityTooLarge() *GetNotificationsAvailabletopicsRequestEntityTooLarge {
	return &GetNotificationsAvailabletopicsRequestEntityTooLarge{}
}

/*
GetNotificationsAvailabletopicsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetNotificationsAvailabletopicsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications availabletopics request entity too large response has a 2xx status code
func (o *GetNotificationsAvailabletopicsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications availabletopics request entity too large response has a 3xx status code
func (o *GetNotificationsAvailabletopicsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications availabletopics request entity too large response has a 4xx status code
func (o *GetNotificationsAvailabletopicsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get notifications availabletopics request entity too large response has a 5xx status code
func (o *GetNotificationsAvailabletopicsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications availabletopics request entity too large response a status code equal to that given
func (o *GetNotificationsAvailabletopicsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetNotificationsAvailabletopicsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetNotificationsAvailabletopicsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetNotificationsAvailabletopicsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsAvailabletopicsUnsupportedMediaType creates a GetNotificationsAvailabletopicsUnsupportedMediaType with default headers values
func NewGetNotificationsAvailabletopicsUnsupportedMediaType() *GetNotificationsAvailabletopicsUnsupportedMediaType {
	return &GetNotificationsAvailabletopicsUnsupportedMediaType{}
}

/*
GetNotificationsAvailabletopicsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetNotificationsAvailabletopicsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications availabletopics unsupported media type response has a 2xx status code
func (o *GetNotificationsAvailabletopicsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications availabletopics unsupported media type response has a 3xx status code
func (o *GetNotificationsAvailabletopicsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications availabletopics unsupported media type response has a 4xx status code
func (o *GetNotificationsAvailabletopicsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get notifications availabletopics unsupported media type response has a 5xx status code
func (o *GetNotificationsAvailabletopicsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications availabletopics unsupported media type response a status code equal to that given
func (o *GetNotificationsAvailabletopicsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetNotificationsAvailabletopicsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetNotificationsAvailabletopicsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetNotificationsAvailabletopicsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsAvailabletopicsTooManyRequests creates a GetNotificationsAvailabletopicsTooManyRequests with default headers values
func NewGetNotificationsAvailabletopicsTooManyRequests() *GetNotificationsAvailabletopicsTooManyRequests {
	return &GetNotificationsAvailabletopicsTooManyRequests{}
}

/*
GetNotificationsAvailabletopicsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetNotificationsAvailabletopicsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications availabletopics too many requests response has a 2xx status code
func (o *GetNotificationsAvailabletopicsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications availabletopics too many requests response has a 3xx status code
func (o *GetNotificationsAvailabletopicsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications availabletopics too many requests response has a 4xx status code
func (o *GetNotificationsAvailabletopicsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get notifications availabletopics too many requests response has a 5xx status code
func (o *GetNotificationsAvailabletopicsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications availabletopics too many requests response a status code equal to that given
func (o *GetNotificationsAvailabletopicsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetNotificationsAvailabletopicsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetNotificationsAvailabletopicsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetNotificationsAvailabletopicsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsAvailabletopicsInternalServerError creates a GetNotificationsAvailabletopicsInternalServerError with default headers values
func NewGetNotificationsAvailabletopicsInternalServerError() *GetNotificationsAvailabletopicsInternalServerError {
	return &GetNotificationsAvailabletopicsInternalServerError{}
}

/*
GetNotificationsAvailabletopicsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetNotificationsAvailabletopicsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications availabletopics internal server error response has a 2xx status code
func (o *GetNotificationsAvailabletopicsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications availabletopics internal server error response has a 3xx status code
func (o *GetNotificationsAvailabletopicsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications availabletopics internal server error response has a 4xx status code
func (o *GetNotificationsAvailabletopicsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get notifications availabletopics internal server error response has a 5xx status code
func (o *GetNotificationsAvailabletopicsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get notifications availabletopics internal server error response a status code equal to that given
func (o *GetNotificationsAvailabletopicsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetNotificationsAvailabletopicsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNotificationsAvailabletopicsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNotificationsAvailabletopicsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsAvailabletopicsServiceUnavailable creates a GetNotificationsAvailabletopicsServiceUnavailable with default headers values
func NewGetNotificationsAvailabletopicsServiceUnavailable() *GetNotificationsAvailabletopicsServiceUnavailable {
	return &GetNotificationsAvailabletopicsServiceUnavailable{}
}

/*
GetNotificationsAvailabletopicsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetNotificationsAvailabletopicsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications availabletopics service unavailable response has a 2xx status code
func (o *GetNotificationsAvailabletopicsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications availabletopics service unavailable response has a 3xx status code
func (o *GetNotificationsAvailabletopicsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications availabletopics service unavailable response has a 4xx status code
func (o *GetNotificationsAvailabletopicsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get notifications availabletopics service unavailable response has a 5xx status code
func (o *GetNotificationsAvailabletopicsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get notifications availabletopics service unavailable response a status code equal to that given
func (o *GetNotificationsAvailabletopicsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetNotificationsAvailabletopicsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetNotificationsAvailabletopicsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetNotificationsAvailabletopicsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsAvailabletopicsGatewayTimeout creates a GetNotificationsAvailabletopicsGatewayTimeout with default headers values
func NewGetNotificationsAvailabletopicsGatewayTimeout() *GetNotificationsAvailabletopicsGatewayTimeout {
	return &GetNotificationsAvailabletopicsGatewayTimeout{}
}

/*
GetNotificationsAvailabletopicsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetNotificationsAvailabletopicsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications availabletopics gateway timeout response has a 2xx status code
func (o *GetNotificationsAvailabletopicsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications availabletopics gateway timeout response has a 3xx status code
func (o *GetNotificationsAvailabletopicsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications availabletopics gateway timeout response has a 4xx status code
func (o *GetNotificationsAvailabletopicsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get notifications availabletopics gateway timeout response has a 5xx status code
func (o *GetNotificationsAvailabletopicsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get notifications availabletopics gateway timeout response a status code equal to that given
func (o *GetNotificationsAvailabletopicsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetNotificationsAvailabletopicsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetNotificationsAvailabletopicsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetNotificationsAvailabletopicsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
