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

// GetOutboundAttemptlimitsReader is a Reader for the GetOutboundAttemptlimits structure.
type GetOutboundAttemptlimitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundAttemptlimitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundAttemptlimitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundAttemptlimitsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundAttemptlimitsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundAttemptlimitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundAttemptlimitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundAttemptlimitsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundAttemptlimitsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundAttemptlimitsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundAttemptlimitsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundAttemptlimitsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundAttemptlimitsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundAttemptlimitsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundAttemptlimitsOK creates a GetOutboundAttemptlimitsOK with default headers values
func NewGetOutboundAttemptlimitsOK() *GetOutboundAttemptlimitsOK {
	return &GetOutboundAttemptlimitsOK{}
}

/*
GetOutboundAttemptlimitsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOutboundAttemptlimitsOK struct {
	Payload *models.AttemptLimitsEntityListing
}

// IsSuccess returns true when this get outbound attemptlimits o k response has a 2xx status code
func (o *GetOutboundAttemptlimitsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get outbound attemptlimits o k response has a 3xx status code
func (o *GetOutboundAttemptlimitsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimits o k response has a 4xx status code
func (o *GetOutboundAttemptlimitsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound attemptlimits o k response has a 5xx status code
func (o *GetOutboundAttemptlimitsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound attemptlimits o k response a status code equal to that given
func (o *GetOutboundAttemptlimitsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOutboundAttemptlimitsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundAttemptlimitsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundAttemptlimitsOK) GetPayload() *models.AttemptLimitsEntityListing {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AttemptLimitsEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsBadRequest creates a GetOutboundAttemptlimitsBadRequest with default headers values
func NewGetOutboundAttemptlimitsBadRequest() *GetOutboundAttemptlimitsBadRequest {
	return &GetOutboundAttemptlimitsBadRequest{}
}

/*
GetOutboundAttemptlimitsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundAttemptlimitsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimits bad request response has a 2xx status code
func (o *GetOutboundAttemptlimitsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimits bad request response has a 3xx status code
func (o *GetOutboundAttemptlimitsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimits bad request response has a 4xx status code
func (o *GetOutboundAttemptlimitsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound attemptlimits bad request response has a 5xx status code
func (o *GetOutboundAttemptlimitsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound attemptlimits bad request response a status code equal to that given
func (o *GetOutboundAttemptlimitsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOutboundAttemptlimitsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundAttemptlimitsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundAttemptlimitsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsUnauthorized creates a GetOutboundAttemptlimitsUnauthorized with default headers values
func NewGetOutboundAttemptlimitsUnauthorized() *GetOutboundAttemptlimitsUnauthorized {
	return &GetOutboundAttemptlimitsUnauthorized{}
}

/*
GetOutboundAttemptlimitsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundAttemptlimitsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimits unauthorized response has a 2xx status code
func (o *GetOutboundAttemptlimitsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimits unauthorized response has a 3xx status code
func (o *GetOutboundAttemptlimitsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimits unauthorized response has a 4xx status code
func (o *GetOutboundAttemptlimitsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound attemptlimits unauthorized response has a 5xx status code
func (o *GetOutboundAttemptlimitsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound attemptlimits unauthorized response a status code equal to that given
func (o *GetOutboundAttemptlimitsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOutboundAttemptlimitsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundAttemptlimitsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundAttemptlimitsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsForbidden creates a GetOutboundAttemptlimitsForbidden with default headers values
func NewGetOutboundAttemptlimitsForbidden() *GetOutboundAttemptlimitsForbidden {
	return &GetOutboundAttemptlimitsForbidden{}
}

/*
GetOutboundAttemptlimitsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundAttemptlimitsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimits forbidden response has a 2xx status code
func (o *GetOutboundAttemptlimitsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimits forbidden response has a 3xx status code
func (o *GetOutboundAttemptlimitsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimits forbidden response has a 4xx status code
func (o *GetOutboundAttemptlimitsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound attemptlimits forbidden response has a 5xx status code
func (o *GetOutboundAttemptlimitsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound attemptlimits forbidden response a status code equal to that given
func (o *GetOutboundAttemptlimitsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOutboundAttemptlimitsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundAttemptlimitsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundAttemptlimitsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsNotFound creates a GetOutboundAttemptlimitsNotFound with default headers values
func NewGetOutboundAttemptlimitsNotFound() *GetOutboundAttemptlimitsNotFound {
	return &GetOutboundAttemptlimitsNotFound{}
}

/*
GetOutboundAttemptlimitsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOutboundAttemptlimitsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimits not found response has a 2xx status code
func (o *GetOutboundAttemptlimitsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimits not found response has a 3xx status code
func (o *GetOutboundAttemptlimitsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimits not found response has a 4xx status code
func (o *GetOutboundAttemptlimitsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound attemptlimits not found response has a 5xx status code
func (o *GetOutboundAttemptlimitsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound attemptlimits not found response a status code equal to that given
func (o *GetOutboundAttemptlimitsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOutboundAttemptlimitsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundAttemptlimitsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundAttemptlimitsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsRequestTimeout creates a GetOutboundAttemptlimitsRequestTimeout with default headers values
func NewGetOutboundAttemptlimitsRequestTimeout() *GetOutboundAttemptlimitsRequestTimeout {
	return &GetOutboundAttemptlimitsRequestTimeout{}
}

/*
GetOutboundAttemptlimitsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundAttemptlimitsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimits request timeout response has a 2xx status code
func (o *GetOutboundAttemptlimitsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimits request timeout response has a 3xx status code
func (o *GetOutboundAttemptlimitsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimits request timeout response has a 4xx status code
func (o *GetOutboundAttemptlimitsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound attemptlimits request timeout response has a 5xx status code
func (o *GetOutboundAttemptlimitsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound attemptlimits request timeout response a status code equal to that given
func (o *GetOutboundAttemptlimitsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOutboundAttemptlimitsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundAttemptlimitsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundAttemptlimitsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsRequestEntityTooLarge creates a GetOutboundAttemptlimitsRequestEntityTooLarge with default headers values
func NewGetOutboundAttemptlimitsRequestEntityTooLarge() *GetOutboundAttemptlimitsRequestEntityTooLarge {
	return &GetOutboundAttemptlimitsRequestEntityTooLarge{}
}

/*
GetOutboundAttemptlimitsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetOutboundAttemptlimitsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimits request entity too large response has a 2xx status code
func (o *GetOutboundAttemptlimitsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimits request entity too large response has a 3xx status code
func (o *GetOutboundAttemptlimitsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimits request entity too large response has a 4xx status code
func (o *GetOutboundAttemptlimitsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound attemptlimits request entity too large response has a 5xx status code
func (o *GetOutboundAttemptlimitsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound attemptlimits request entity too large response a status code equal to that given
func (o *GetOutboundAttemptlimitsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOutboundAttemptlimitsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundAttemptlimitsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundAttemptlimitsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsUnsupportedMediaType creates a GetOutboundAttemptlimitsUnsupportedMediaType with default headers values
func NewGetOutboundAttemptlimitsUnsupportedMediaType() *GetOutboundAttemptlimitsUnsupportedMediaType {
	return &GetOutboundAttemptlimitsUnsupportedMediaType{}
}

/*
GetOutboundAttemptlimitsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundAttemptlimitsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimits unsupported media type response has a 2xx status code
func (o *GetOutboundAttemptlimitsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimits unsupported media type response has a 3xx status code
func (o *GetOutboundAttemptlimitsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimits unsupported media type response has a 4xx status code
func (o *GetOutboundAttemptlimitsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound attemptlimits unsupported media type response has a 5xx status code
func (o *GetOutboundAttemptlimitsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound attemptlimits unsupported media type response a status code equal to that given
func (o *GetOutboundAttemptlimitsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOutboundAttemptlimitsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundAttemptlimitsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundAttemptlimitsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsTooManyRequests creates a GetOutboundAttemptlimitsTooManyRequests with default headers values
func NewGetOutboundAttemptlimitsTooManyRequests() *GetOutboundAttemptlimitsTooManyRequests {
	return &GetOutboundAttemptlimitsTooManyRequests{}
}

/*
GetOutboundAttemptlimitsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundAttemptlimitsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimits too many requests response has a 2xx status code
func (o *GetOutboundAttemptlimitsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimits too many requests response has a 3xx status code
func (o *GetOutboundAttemptlimitsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimits too many requests response has a 4xx status code
func (o *GetOutboundAttemptlimitsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound attemptlimits too many requests response has a 5xx status code
func (o *GetOutboundAttemptlimitsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound attemptlimits too many requests response a status code equal to that given
func (o *GetOutboundAttemptlimitsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOutboundAttemptlimitsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundAttemptlimitsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundAttemptlimitsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsInternalServerError creates a GetOutboundAttemptlimitsInternalServerError with default headers values
func NewGetOutboundAttemptlimitsInternalServerError() *GetOutboundAttemptlimitsInternalServerError {
	return &GetOutboundAttemptlimitsInternalServerError{}
}

/*
GetOutboundAttemptlimitsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundAttemptlimitsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimits internal server error response has a 2xx status code
func (o *GetOutboundAttemptlimitsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimits internal server error response has a 3xx status code
func (o *GetOutboundAttemptlimitsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimits internal server error response has a 4xx status code
func (o *GetOutboundAttemptlimitsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound attemptlimits internal server error response has a 5xx status code
func (o *GetOutboundAttemptlimitsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound attemptlimits internal server error response a status code equal to that given
func (o *GetOutboundAttemptlimitsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOutboundAttemptlimitsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundAttemptlimitsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundAttemptlimitsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsServiceUnavailable creates a GetOutboundAttemptlimitsServiceUnavailable with default headers values
func NewGetOutboundAttemptlimitsServiceUnavailable() *GetOutboundAttemptlimitsServiceUnavailable {
	return &GetOutboundAttemptlimitsServiceUnavailable{}
}

/*
GetOutboundAttemptlimitsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundAttemptlimitsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimits service unavailable response has a 2xx status code
func (o *GetOutboundAttemptlimitsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimits service unavailable response has a 3xx status code
func (o *GetOutboundAttemptlimitsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimits service unavailable response has a 4xx status code
func (o *GetOutboundAttemptlimitsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound attemptlimits service unavailable response has a 5xx status code
func (o *GetOutboundAttemptlimitsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound attemptlimits service unavailable response a status code equal to that given
func (o *GetOutboundAttemptlimitsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOutboundAttemptlimitsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundAttemptlimitsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundAttemptlimitsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsGatewayTimeout creates a GetOutboundAttemptlimitsGatewayTimeout with default headers values
func NewGetOutboundAttemptlimitsGatewayTimeout() *GetOutboundAttemptlimitsGatewayTimeout {
	return &GetOutboundAttemptlimitsGatewayTimeout{}
}

/*
GetOutboundAttemptlimitsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOutboundAttemptlimitsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound attemptlimits gateway timeout response has a 2xx status code
func (o *GetOutboundAttemptlimitsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound attemptlimits gateway timeout response has a 3xx status code
func (o *GetOutboundAttemptlimitsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound attemptlimits gateway timeout response has a 4xx status code
func (o *GetOutboundAttemptlimitsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound attemptlimits gateway timeout response has a 5xx status code
func (o *GetOutboundAttemptlimitsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound attemptlimits gateway timeout response a status code equal to that given
func (o *GetOutboundAttemptlimitsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOutboundAttemptlimitsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundAttemptlimitsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundAttemptlimitsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
