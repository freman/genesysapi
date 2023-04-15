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

// GetOutboundSchedulesSequencesReader is a Reader for the GetOutboundSchedulesSequences structure.
type GetOutboundSchedulesSequencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundSchedulesSequencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundSchedulesSequencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundSchedulesSequencesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundSchedulesSequencesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundSchedulesSequencesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundSchedulesSequencesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundSchedulesSequencesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundSchedulesSequencesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundSchedulesSequencesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundSchedulesSequencesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundSchedulesSequencesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundSchedulesSequencesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundSchedulesSequencesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundSchedulesSequencesOK creates a GetOutboundSchedulesSequencesOK with default headers values
func NewGetOutboundSchedulesSequencesOK() *GetOutboundSchedulesSequencesOK {
	return &GetOutboundSchedulesSequencesOK{}
}

/*
GetOutboundSchedulesSequencesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOutboundSchedulesSequencesOK struct {
	Payload []*models.SequenceSchedule
}

// IsSuccess returns true when this get outbound schedules sequences o k response has a 2xx status code
func (o *GetOutboundSchedulesSequencesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get outbound schedules sequences o k response has a 3xx status code
func (o *GetOutboundSchedulesSequencesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequences o k response has a 4xx status code
func (o *GetOutboundSchedulesSequencesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound schedules sequences o k response has a 5xx status code
func (o *GetOutboundSchedulesSequencesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules sequences o k response a status code equal to that given
func (o *GetOutboundSchedulesSequencesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOutboundSchedulesSequencesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesOK  %+v", 200, o.Payload)
}

func (o *GetOutboundSchedulesSequencesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesOK  %+v", 200, o.Payload)
}

func (o *GetOutboundSchedulesSequencesOK) GetPayload() []*models.SequenceSchedule {
	return o.Payload
}

func (o *GetOutboundSchedulesSequencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequencesBadRequest creates a GetOutboundSchedulesSequencesBadRequest with default headers values
func NewGetOutboundSchedulesSequencesBadRequest() *GetOutboundSchedulesSequencesBadRequest {
	return &GetOutboundSchedulesSequencesBadRequest{}
}

/*
GetOutboundSchedulesSequencesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundSchedulesSequencesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequences bad request response has a 2xx status code
func (o *GetOutboundSchedulesSequencesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequences bad request response has a 3xx status code
func (o *GetOutboundSchedulesSequencesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequences bad request response has a 4xx status code
func (o *GetOutboundSchedulesSequencesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules sequences bad request response has a 5xx status code
func (o *GetOutboundSchedulesSequencesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules sequences bad request response a status code equal to that given
func (o *GetOutboundSchedulesSequencesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOutboundSchedulesSequencesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundSchedulesSequencesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundSchedulesSequencesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequencesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequencesUnauthorized creates a GetOutboundSchedulesSequencesUnauthorized with default headers values
func NewGetOutboundSchedulesSequencesUnauthorized() *GetOutboundSchedulesSequencesUnauthorized {
	return &GetOutboundSchedulesSequencesUnauthorized{}
}

/*
GetOutboundSchedulesSequencesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundSchedulesSequencesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequences unauthorized response has a 2xx status code
func (o *GetOutboundSchedulesSequencesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequences unauthorized response has a 3xx status code
func (o *GetOutboundSchedulesSequencesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequences unauthorized response has a 4xx status code
func (o *GetOutboundSchedulesSequencesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules sequences unauthorized response has a 5xx status code
func (o *GetOutboundSchedulesSequencesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules sequences unauthorized response a status code equal to that given
func (o *GetOutboundSchedulesSequencesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOutboundSchedulesSequencesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundSchedulesSequencesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundSchedulesSequencesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequencesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequencesForbidden creates a GetOutboundSchedulesSequencesForbidden with default headers values
func NewGetOutboundSchedulesSequencesForbidden() *GetOutboundSchedulesSequencesForbidden {
	return &GetOutboundSchedulesSequencesForbidden{}
}

/*
GetOutboundSchedulesSequencesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundSchedulesSequencesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequences forbidden response has a 2xx status code
func (o *GetOutboundSchedulesSequencesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequences forbidden response has a 3xx status code
func (o *GetOutboundSchedulesSequencesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequences forbidden response has a 4xx status code
func (o *GetOutboundSchedulesSequencesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules sequences forbidden response has a 5xx status code
func (o *GetOutboundSchedulesSequencesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules sequences forbidden response a status code equal to that given
func (o *GetOutboundSchedulesSequencesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOutboundSchedulesSequencesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundSchedulesSequencesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundSchedulesSequencesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequencesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequencesNotFound creates a GetOutboundSchedulesSequencesNotFound with default headers values
func NewGetOutboundSchedulesSequencesNotFound() *GetOutboundSchedulesSequencesNotFound {
	return &GetOutboundSchedulesSequencesNotFound{}
}

/*
GetOutboundSchedulesSequencesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOutboundSchedulesSequencesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequences not found response has a 2xx status code
func (o *GetOutboundSchedulesSequencesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequences not found response has a 3xx status code
func (o *GetOutboundSchedulesSequencesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequences not found response has a 4xx status code
func (o *GetOutboundSchedulesSequencesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules sequences not found response has a 5xx status code
func (o *GetOutboundSchedulesSequencesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules sequences not found response a status code equal to that given
func (o *GetOutboundSchedulesSequencesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOutboundSchedulesSequencesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundSchedulesSequencesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundSchedulesSequencesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequencesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequencesRequestTimeout creates a GetOutboundSchedulesSequencesRequestTimeout with default headers values
func NewGetOutboundSchedulesSequencesRequestTimeout() *GetOutboundSchedulesSequencesRequestTimeout {
	return &GetOutboundSchedulesSequencesRequestTimeout{}
}

/*
GetOutboundSchedulesSequencesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundSchedulesSequencesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequences request timeout response has a 2xx status code
func (o *GetOutboundSchedulesSequencesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequences request timeout response has a 3xx status code
func (o *GetOutboundSchedulesSequencesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequences request timeout response has a 4xx status code
func (o *GetOutboundSchedulesSequencesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules sequences request timeout response has a 5xx status code
func (o *GetOutboundSchedulesSequencesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules sequences request timeout response a status code equal to that given
func (o *GetOutboundSchedulesSequencesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOutboundSchedulesSequencesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundSchedulesSequencesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundSchedulesSequencesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequencesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequencesRequestEntityTooLarge creates a GetOutboundSchedulesSequencesRequestEntityTooLarge with default headers values
func NewGetOutboundSchedulesSequencesRequestEntityTooLarge() *GetOutboundSchedulesSequencesRequestEntityTooLarge {
	return &GetOutboundSchedulesSequencesRequestEntityTooLarge{}
}

/*
GetOutboundSchedulesSequencesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetOutboundSchedulesSequencesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequences request entity too large response has a 2xx status code
func (o *GetOutboundSchedulesSequencesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequences request entity too large response has a 3xx status code
func (o *GetOutboundSchedulesSequencesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequences request entity too large response has a 4xx status code
func (o *GetOutboundSchedulesSequencesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules sequences request entity too large response has a 5xx status code
func (o *GetOutboundSchedulesSequencesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules sequences request entity too large response a status code equal to that given
func (o *GetOutboundSchedulesSequencesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOutboundSchedulesSequencesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundSchedulesSequencesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundSchedulesSequencesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequencesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequencesUnsupportedMediaType creates a GetOutboundSchedulesSequencesUnsupportedMediaType with default headers values
func NewGetOutboundSchedulesSequencesUnsupportedMediaType() *GetOutboundSchedulesSequencesUnsupportedMediaType {
	return &GetOutboundSchedulesSequencesUnsupportedMediaType{}
}

/*
GetOutboundSchedulesSequencesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundSchedulesSequencesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequences unsupported media type response has a 2xx status code
func (o *GetOutboundSchedulesSequencesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequences unsupported media type response has a 3xx status code
func (o *GetOutboundSchedulesSequencesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequences unsupported media type response has a 4xx status code
func (o *GetOutboundSchedulesSequencesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules sequences unsupported media type response has a 5xx status code
func (o *GetOutboundSchedulesSequencesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules sequences unsupported media type response a status code equal to that given
func (o *GetOutboundSchedulesSequencesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOutboundSchedulesSequencesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundSchedulesSequencesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundSchedulesSequencesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequencesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequencesTooManyRequests creates a GetOutboundSchedulesSequencesTooManyRequests with default headers values
func NewGetOutboundSchedulesSequencesTooManyRequests() *GetOutboundSchedulesSequencesTooManyRequests {
	return &GetOutboundSchedulesSequencesTooManyRequests{}
}

/*
GetOutboundSchedulesSequencesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundSchedulesSequencesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequences too many requests response has a 2xx status code
func (o *GetOutboundSchedulesSequencesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequences too many requests response has a 3xx status code
func (o *GetOutboundSchedulesSequencesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequences too many requests response has a 4xx status code
func (o *GetOutboundSchedulesSequencesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules sequences too many requests response has a 5xx status code
func (o *GetOutboundSchedulesSequencesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules sequences too many requests response a status code equal to that given
func (o *GetOutboundSchedulesSequencesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOutboundSchedulesSequencesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundSchedulesSequencesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundSchedulesSequencesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequencesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequencesInternalServerError creates a GetOutboundSchedulesSequencesInternalServerError with default headers values
func NewGetOutboundSchedulesSequencesInternalServerError() *GetOutboundSchedulesSequencesInternalServerError {
	return &GetOutboundSchedulesSequencesInternalServerError{}
}

/*
GetOutboundSchedulesSequencesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundSchedulesSequencesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequences internal server error response has a 2xx status code
func (o *GetOutboundSchedulesSequencesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequences internal server error response has a 3xx status code
func (o *GetOutboundSchedulesSequencesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequences internal server error response has a 4xx status code
func (o *GetOutboundSchedulesSequencesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound schedules sequences internal server error response has a 5xx status code
func (o *GetOutboundSchedulesSequencesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound schedules sequences internal server error response a status code equal to that given
func (o *GetOutboundSchedulesSequencesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOutboundSchedulesSequencesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundSchedulesSequencesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundSchedulesSequencesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequencesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequencesServiceUnavailable creates a GetOutboundSchedulesSequencesServiceUnavailable with default headers values
func NewGetOutboundSchedulesSequencesServiceUnavailable() *GetOutboundSchedulesSequencesServiceUnavailable {
	return &GetOutboundSchedulesSequencesServiceUnavailable{}
}

/*
GetOutboundSchedulesSequencesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundSchedulesSequencesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequences service unavailable response has a 2xx status code
func (o *GetOutboundSchedulesSequencesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequences service unavailable response has a 3xx status code
func (o *GetOutboundSchedulesSequencesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequences service unavailable response has a 4xx status code
func (o *GetOutboundSchedulesSequencesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound schedules sequences service unavailable response has a 5xx status code
func (o *GetOutboundSchedulesSequencesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound schedules sequences service unavailable response a status code equal to that given
func (o *GetOutboundSchedulesSequencesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOutboundSchedulesSequencesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundSchedulesSequencesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundSchedulesSequencesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequencesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesSequencesGatewayTimeout creates a GetOutboundSchedulesSequencesGatewayTimeout with default headers values
func NewGetOutboundSchedulesSequencesGatewayTimeout() *GetOutboundSchedulesSequencesGatewayTimeout {
	return &GetOutboundSchedulesSequencesGatewayTimeout{}
}

/*
GetOutboundSchedulesSequencesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOutboundSchedulesSequencesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules sequences gateway timeout response has a 2xx status code
func (o *GetOutboundSchedulesSequencesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules sequences gateway timeout response has a 3xx status code
func (o *GetOutboundSchedulesSequencesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules sequences gateway timeout response has a 4xx status code
func (o *GetOutboundSchedulesSequencesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound schedules sequences gateway timeout response has a 5xx status code
func (o *GetOutboundSchedulesSequencesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound schedules sequences gateway timeout response a status code equal to that given
func (o *GetOutboundSchedulesSequencesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOutboundSchedulesSequencesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundSchedulesSequencesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/sequences][%d] getOutboundSchedulesSequencesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundSchedulesSequencesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesSequencesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
