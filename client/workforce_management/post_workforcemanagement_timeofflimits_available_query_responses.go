// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostWorkforcemanagementTimeofflimitsAvailableQueryReader is a Reader for the PostWorkforcemanagementTimeofflimitsAvailableQuery structure.
type PostWorkforcemanagementTimeofflimitsAvailableQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWorkforcemanagementTimeofflimitsAvailableQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkforcemanagementTimeofflimitsAvailableQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementTimeofflimitsAvailableQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementTimeofflimitsAvailableQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementTimeofflimitsAvailableQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostWorkforcemanagementTimeofflimitsAvailableQueryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementTimeofflimitsAvailableQueryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementTimeofflimitsAvailableQueryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementTimeofflimitsAvailableQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementTimeofflimitsAvailableQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementTimeofflimitsAvailableQueryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementTimeofflimitsAvailableQueryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkforcemanagementTimeofflimitsAvailableQueryOK creates a PostWorkforcemanagementTimeofflimitsAvailableQueryOK with default headers values
func NewPostWorkforcemanagementTimeofflimitsAvailableQueryOK() *PostWorkforcemanagementTimeofflimitsAvailableQueryOK {
	return &PostWorkforcemanagementTimeofflimitsAvailableQueryOK{}
}

/*
PostWorkforcemanagementTimeofflimitsAvailableQueryOK describes a response with status code 200, with default header values.

successful operation
*/
type PostWorkforcemanagementTimeofflimitsAvailableQueryOK struct {
	Payload *models.AvailableTimeOffResponse
}

// IsSuccess returns true when this post workforcemanagement timeofflimits available query o k response has a 2xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post workforcemanagement timeofflimits available query o k response has a 3xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement timeofflimits available query o k response has a 4xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement timeofflimits available query o k response has a 5xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement timeofflimits available query o k response a status code equal to that given
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryOK) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryOK) GetPayload() *models.AvailableTimeOffResponse {
	return o.Payload
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AvailableTimeOffResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementTimeofflimitsAvailableQueryBadRequest creates a PostWorkforcemanagementTimeofflimitsAvailableQueryBadRequest with default headers values
func NewPostWorkforcemanagementTimeofflimitsAvailableQueryBadRequest() *PostWorkforcemanagementTimeofflimitsAvailableQueryBadRequest {
	return &PostWorkforcemanagementTimeofflimitsAvailableQueryBadRequest{}
}

/*
PostWorkforcemanagementTimeofflimitsAvailableQueryBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementTimeofflimitsAvailableQueryBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement timeofflimits available query bad request response has a 2xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement timeofflimits available query bad request response has a 3xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement timeofflimits available query bad request response has a 4xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement timeofflimits available query bad request response has a 5xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement timeofflimits available query bad request response a status code equal to that given
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementTimeofflimitsAvailableQueryUnauthorized creates a PostWorkforcemanagementTimeofflimitsAvailableQueryUnauthorized with default headers values
func NewPostWorkforcemanagementTimeofflimitsAvailableQueryUnauthorized() *PostWorkforcemanagementTimeofflimitsAvailableQueryUnauthorized {
	return &PostWorkforcemanagementTimeofflimitsAvailableQueryUnauthorized{}
}

/*
PostWorkforcemanagementTimeofflimitsAvailableQueryUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementTimeofflimitsAvailableQueryUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement timeofflimits available query unauthorized response has a 2xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement timeofflimits available query unauthorized response has a 3xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement timeofflimits available query unauthorized response has a 4xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement timeofflimits available query unauthorized response has a 5xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement timeofflimits available query unauthorized response a status code equal to that given
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementTimeofflimitsAvailableQueryForbidden creates a PostWorkforcemanagementTimeofflimitsAvailableQueryForbidden with default headers values
func NewPostWorkforcemanagementTimeofflimitsAvailableQueryForbidden() *PostWorkforcemanagementTimeofflimitsAvailableQueryForbidden {
	return &PostWorkforcemanagementTimeofflimitsAvailableQueryForbidden{}
}

/*
PostWorkforcemanagementTimeofflimitsAvailableQueryForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementTimeofflimitsAvailableQueryForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement timeofflimits available query forbidden response has a 2xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement timeofflimits available query forbidden response has a 3xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement timeofflimits available query forbidden response has a 4xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement timeofflimits available query forbidden response has a 5xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement timeofflimits available query forbidden response a status code equal to that given
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementTimeofflimitsAvailableQueryNotFound creates a PostWorkforcemanagementTimeofflimitsAvailableQueryNotFound with default headers values
func NewPostWorkforcemanagementTimeofflimitsAvailableQueryNotFound() *PostWorkforcemanagementTimeofflimitsAvailableQueryNotFound {
	return &PostWorkforcemanagementTimeofflimitsAvailableQueryNotFound{}
}

/*
PostWorkforcemanagementTimeofflimitsAvailableQueryNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementTimeofflimitsAvailableQueryNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement timeofflimits available query not found response has a 2xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement timeofflimits available query not found response has a 3xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement timeofflimits available query not found response has a 4xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement timeofflimits available query not found response has a 5xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement timeofflimits available query not found response a status code equal to that given
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementTimeofflimitsAvailableQueryRequestTimeout creates a PostWorkforcemanagementTimeofflimitsAvailableQueryRequestTimeout with default headers values
func NewPostWorkforcemanagementTimeofflimitsAvailableQueryRequestTimeout() *PostWorkforcemanagementTimeofflimitsAvailableQueryRequestTimeout {
	return &PostWorkforcemanagementTimeofflimitsAvailableQueryRequestTimeout{}
}

/*
PostWorkforcemanagementTimeofflimitsAvailableQueryRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostWorkforcemanagementTimeofflimitsAvailableQueryRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement timeofflimits available query request timeout response has a 2xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement timeofflimits available query request timeout response has a 3xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement timeofflimits available query request timeout response has a 4xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement timeofflimits available query request timeout response has a 5xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement timeofflimits available query request timeout response a status code equal to that given
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementTimeofflimitsAvailableQueryRequestEntityTooLarge creates a PostWorkforcemanagementTimeofflimitsAvailableQueryRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementTimeofflimitsAvailableQueryRequestEntityTooLarge() *PostWorkforcemanagementTimeofflimitsAvailableQueryRequestEntityTooLarge {
	return &PostWorkforcemanagementTimeofflimitsAvailableQueryRequestEntityTooLarge{}
}

/*
PostWorkforcemanagementTimeofflimitsAvailableQueryRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostWorkforcemanagementTimeofflimitsAvailableQueryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement timeofflimits available query request entity too large response has a 2xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement timeofflimits available query request entity too large response has a 3xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement timeofflimits available query request entity too large response has a 4xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement timeofflimits available query request entity too large response has a 5xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement timeofflimits available query request entity too large response a status code equal to that given
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementTimeofflimitsAvailableQueryUnsupportedMediaType creates a PostWorkforcemanagementTimeofflimitsAvailableQueryUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementTimeofflimitsAvailableQueryUnsupportedMediaType() *PostWorkforcemanagementTimeofflimitsAvailableQueryUnsupportedMediaType {
	return &PostWorkforcemanagementTimeofflimitsAvailableQueryUnsupportedMediaType{}
}

/*
PostWorkforcemanagementTimeofflimitsAvailableQueryUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementTimeofflimitsAvailableQueryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement timeofflimits available query unsupported media type response has a 2xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement timeofflimits available query unsupported media type response has a 3xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement timeofflimits available query unsupported media type response has a 4xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement timeofflimits available query unsupported media type response has a 5xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement timeofflimits available query unsupported media type response a status code equal to that given
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementTimeofflimitsAvailableQueryTooManyRequests creates a PostWorkforcemanagementTimeofflimitsAvailableQueryTooManyRequests with default headers values
func NewPostWorkforcemanagementTimeofflimitsAvailableQueryTooManyRequests() *PostWorkforcemanagementTimeofflimitsAvailableQueryTooManyRequests {
	return &PostWorkforcemanagementTimeofflimitsAvailableQueryTooManyRequests{}
}

/*
PostWorkforcemanagementTimeofflimitsAvailableQueryTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostWorkforcemanagementTimeofflimitsAvailableQueryTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement timeofflimits available query too many requests response has a 2xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement timeofflimits available query too many requests response has a 3xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement timeofflimits available query too many requests response has a 4xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement timeofflimits available query too many requests response has a 5xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement timeofflimits available query too many requests response a status code equal to that given
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementTimeofflimitsAvailableQueryInternalServerError creates a PostWorkforcemanagementTimeofflimitsAvailableQueryInternalServerError with default headers values
func NewPostWorkforcemanagementTimeofflimitsAvailableQueryInternalServerError() *PostWorkforcemanagementTimeofflimitsAvailableQueryInternalServerError {
	return &PostWorkforcemanagementTimeofflimitsAvailableQueryInternalServerError{}
}

/*
PostWorkforcemanagementTimeofflimitsAvailableQueryInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementTimeofflimitsAvailableQueryInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement timeofflimits available query internal server error response has a 2xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement timeofflimits available query internal server error response has a 3xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement timeofflimits available query internal server error response has a 4xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement timeofflimits available query internal server error response has a 5xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post workforcemanagement timeofflimits available query internal server error response a status code equal to that given
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementTimeofflimitsAvailableQueryServiceUnavailable creates a PostWorkforcemanagementTimeofflimitsAvailableQueryServiceUnavailable with default headers values
func NewPostWorkforcemanagementTimeofflimitsAvailableQueryServiceUnavailable() *PostWorkforcemanagementTimeofflimitsAvailableQueryServiceUnavailable {
	return &PostWorkforcemanagementTimeofflimitsAvailableQueryServiceUnavailable{}
}

/*
PostWorkforcemanagementTimeofflimitsAvailableQueryServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementTimeofflimitsAvailableQueryServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement timeofflimits available query service unavailable response has a 2xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement timeofflimits available query service unavailable response has a 3xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement timeofflimits available query service unavailable response has a 4xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement timeofflimits available query service unavailable response has a 5xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post workforcemanagement timeofflimits available query service unavailable response a status code equal to that given
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementTimeofflimitsAvailableQueryGatewayTimeout creates a PostWorkforcemanagementTimeofflimitsAvailableQueryGatewayTimeout with default headers values
func NewPostWorkforcemanagementTimeofflimitsAvailableQueryGatewayTimeout() *PostWorkforcemanagementTimeofflimitsAvailableQueryGatewayTimeout {
	return &PostWorkforcemanagementTimeofflimitsAvailableQueryGatewayTimeout{}
}

/*
PostWorkforcemanagementTimeofflimitsAvailableQueryGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostWorkforcemanagementTimeofflimitsAvailableQueryGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement timeofflimits available query gateway timeout response has a 2xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement timeofflimits available query gateway timeout response has a 3xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement timeofflimits available query gateway timeout response has a 4xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement timeofflimits available query gateway timeout response has a 5xx status code
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post workforcemanagement timeofflimits available query gateway timeout response a status code equal to that given
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/timeofflimits/available/query][%d] postWorkforcemanagementTimeofflimitsAvailableQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
