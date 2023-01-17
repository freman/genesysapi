// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutUserStationDefaultstationStationIDReader is a Reader for the PutUserStationDefaultstationStationID structure.
type PutUserStationDefaultstationStationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutUserStationDefaultstationStationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPutUserStationDefaultstationStationIDAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutUserStationDefaultstationStationIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutUserStationDefaultstationStationIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutUserStationDefaultstationStationIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutUserStationDefaultstationStationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutUserStationDefaultstationStationIDRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutUserStationDefaultstationStationIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutUserStationDefaultstationStationIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutUserStationDefaultstationStationIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutUserStationDefaultstationStationIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutUserStationDefaultstationStationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutUserStationDefaultstationStationIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutUserStationDefaultstationStationIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutUserStationDefaultstationStationIDAccepted creates a PutUserStationDefaultstationStationIDAccepted with default headers values
func NewPutUserStationDefaultstationStationIDAccepted() *PutUserStationDefaultstationStationIDAccepted {
	return &PutUserStationDefaultstationStationIDAccepted{}
}

/*
PutUserStationDefaultstationStationIDAccepted describes a response with status code 202, with default header values.

Success
*/
type PutUserStationDefaultstationStationIDAccepted struct {
}

// IsSuccess returns true when this put user station defaultstation station Id accepted response has a 2xx status code
func (o *PutUserStationDefaultstationStationIDAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put user station defaultstation station Id accepted response has a 3xx status code
func (o *PutUserStationDefaultstationStationIDAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put user station defaultstation station Id accepted response has a 4xx status code
func (o *PutUserStationDefaultstationStationIDAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this put user station defaultstation station Id accepted response has a 5xx status code
func (o *PutUserStationDefaultstationStationIDAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this put user station defaultstation station Id accepted response a status code equal to that given
func (o *PutUserStationDefaultstationStationIDAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PutUserStationDefaultstationStationIDAccepted) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdAccepted ", 202)
}

func (o *PutUserStationDefaultstationStationIDAccepted) String() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdAccepted ", 202)
}

func (o *PutUserStationDefaultstationStationIDAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutUserStationDefaultstationStationIDBadRequest creates a PutUserStationDefaultstationStationIDBadRequest with default headers values
func NewPutUserStationDefaultstationStationIDBadRequest() *PutUserStationDefaultstationStationIDBadRequest {
	return &PutUserStationDefaultstationStationIDBadRequest{}
}

/*
PutUserStationDefaultstationStationIDBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutUserStationDefaultstationStationIDBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put user station defaultstation station Id bad request response has a 2xx status code
func (o *PutUserStationDefaultstationStationIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put user station defaultstation station Id bad request response has a 3xx status code
func (o *PutUserStationDefaultstationStationIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put user station defaultstation station Id bad request response has a 4xx status code
func (o *PutUserStationDefaultstationStationIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put user station defaultstation station Id bad request response has a 5xx status code
func (o *PutUserStationDefaultstationStationIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put user station defaultstation station Id bad request response a status code equal to that given
func (o *PutUserStationDefaultstationStationIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutUserStationDefaultstationStationIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserStationDefaultstationStationIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserStationDefaultstationStationIDUnauthorized creates a PutUserStationDefaultstationStationIDUnauthorized with default headers values
func NewPutUserStationDefaultstationStationIDUnauthorized() *PutUserStationDefaultstationStationIDUnauthorized {
	return &PutUserStationDefaultstationStationIDUnauthorized{}
}

/*
PutUserStationDefaultstationStationIDUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutUserStationDefaultstationStationIDUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put user station defaultstation station Id unauthorized response has a 2xx status code
func (o *PutUserStationDefaultstationStationIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put user station defaultstation station Id unauthorized response has a 3xx status code
func (o *PutUserStationDefaultstationStationIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put user station defaultstation station Id unauthorized response has a 4xx status code
func (o *PutUserStationDefaultstationStationIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put user station defaultstation station Id unauthorized response has a 5xx status code
func (o *PutUserStationDefaultstationStationIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put user station defaultstation station Id unauthorized response a status code equal to that given
func (o *PutUserStationDefaultstationStationIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutUserStationDefaultstationStationIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserStationDefaultstationStationIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserStationDefaultstationStationIDForbidden creates a PutUserStationDefaultstationStationIDForbidden with default headers values
func NewPutUserStationDefaultstationStationIDForbidden() *PutUserStationDefaultstationStationIDForbidden {
	return &PutUserStationDefaultstationStationIDForbidden{}
}

/*
PutUserStationDefaultstationStationIDForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutUserStationDefaultstationStationIDForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put user station defaultstation station Id forbidden response has a 2xx status code
func (o *PutUserStationDefaultstationStationIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put user station defaultstation station Id forbidden response has a 3xx status code
func (o *PutUserStationDefaultstationStationIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put user station defaultstation station Id forbidden response has a 4xx status code
func (o *PutUserStationDefaultstationStationIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put user station defaultstation station Id forbidden response has a 5xx status code
func (o *PutUserStationDefaultstationStationIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put user station defaultstation station Id forbidden response a status code equal to that given
func (o *PutUserStationDefaultstationStationIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutUserStationDefaultstationStationIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdForbidden  %+v", 403, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdForbidden  %+v", 403, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserStationDefaultstationStationIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserStationDefaultstationStationIDNotFound creates a PutUserStationDefaultstationStationIDNotFound with default headers values
func NewPutUserStationDefaultstationStationIDNotFound() *PutUserStationDefaultstationStationIDNotFound {
	return &PutUserStationDefaultstationStationIDNotFound{}
}

/*
PutUserStationDefaultstationStationIDNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutUserStationDefaultstationStationIDNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put user station defaultstation station Id not found response has a 2xx status code
func (o *PutUserStationDefaultstationStationIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put user station defaultstation station Id not found response has a 3xx status code
func (o *PutUserStationDefaultstationStationIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put user station defaultstation station Id not found response has a 4xx status code
func (o *PutUserStationDefaultstationStationIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put user station defaultstation station Id not found response has a 5xx status code
func (o *PutUserStationDefaultstationStationIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put user station defaultstation station Id not found response a status code equal to that given
func (o *PutUserStationDefaultstationStationIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutUserStationDefaultstationStationIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdNotFound  %+v", 404, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdNotFound  %+v", 404, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserStationDefaultstationStationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserStationDefaultstationStationIDRequestTimeout creates a PutUserStationDefaultstationStationIDRequestTimeout with default headers values
func NewPutUserStationDefaultstationStationIDRequestTimeout() *PutUserStationDefaultstationStationIDRequestTimeout {
	return &PutUserStationDefaultstationStationIDRequestTimeout{}
}

/*
PutUserStationDefaultstationStationIDRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutUserStationDefaultstationStationIDRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put user station defaultstation station Id request timeout response has a 2xx status code
func (o *PutUserStationDefaultstationStationIDRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put user station defaultstation station Id request timeout response has a 3xx status code
func (o *PutUserStationDefaultstationStationIDRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put user station defaultstation station Id request timeout response has a 4xx status code
func (o *PutUserStationDefaultstationStationIDRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put user station defaultstation station Id request timeout response has a 5xx status code
func (o *PutUserStationDefaultstationStationIDRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put user station defaultstation station Id request timeout response a status code equal to that given
func (o *PutUserStationDefaultstationStationIDRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutUserStationDefaultstationStationIDRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserStationDefaultstationStationIDRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserStationDefaultstationStationIDConflict creates a PutUserStationDefaultstationStationIDConflict with default headers values
func NewPutUserStationDefaultstationStationIDConflict() *PutUserStationDefaultstationStationIDConflict {
	return &PutUserStationDefaultstationStationIDConflict{}
}

/*
PutUserStationDefaultstationStationIDConflict describes a response with status code 409, with default header values.

Conflict
*/
type PutUserStationDefaultstationStationIDConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put user station defaultstation station Id conflict response has a 2xx status code
func (o *PutUserStationDefaultstationStationIDConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put user station defaultstation station Id conflict response has a 3xx status code
func (o *PutUserStationDefaultstationStationIDConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put user station defaultstation station Id conflict response has a 4xx status code
func (o *PutUserStationDefaultstationStationIDConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this put user station defaultstation station Id conflict response has a 5xx status code
func (o *PutUserStationDefaultstationStationIDConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this put user station defaultstation station Id conflict response a status code equal to that given
func (o *PutUserStationDefaultstationStationIDConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PutUserStationDefaultstationStationIDConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdConflict  %+v", 409, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDConflict) String() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdConflict  %+v", 409, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserStationDefaultstationStationIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserStationDefaultstationStationIDRequestEntityTooLarge creates a PutUserStationDefaultstationStationIDRequestEntityTooLarge with default headers values
func NewPutUserStationDefaultstationStationIDRequestEntityTooLarge() *PutUserStationDefaultstationStationIDRequestEntityTooLarge {
	return &PutUserStationDefaultstationStationIDRequestEntityTooLarge{}
}

/*
PutUserStationDefaultstationStationIDRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutUserStationDefaultstationStationIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put user station defaultstation station Id request entity too large response has a 2xx status code
func (o *PutUserStationDefaultstationStationIDRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put user station defaultstation station Id request entity too large response has a 3xx status code
func (o *PutUserStationDefaultstationStationIDRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put user station defaultstation station Id request entity too large response has a 4xx status code
func (o *PutUserStationDefaultstationStationIDRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put user station defaultstation station Id request entity too large response has a 5xx status code
func (o *PutUserStationDefaultstationStationIDRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put user station defaultstation station Id request entity too large response a status code equal to that given
func (o *PutUserStationDefaultstationStationIDRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutUserStationDefaultstationStationIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserStationDefaultstationStationIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserStationDefaultstationStationIDUnsupportedMediaType creates a PutUserStationDefaultstationStationIDUnsupportedMediaType with default headers values
func NewPutUserStationDefaultstationStationIDUnsupportedMediaType() *PutUserStationDefaultstationStationIDUnsupportedMediaType {
	return &PutUserStationDefaultstationStationIDUnsupportedMediaType{}
}

/*
PutUserStationDefaultstationStationIDUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutUserStationDefaultstationStationIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put user station defaultstation station Id unsupported media type response has a 2xx status code
func (o *PutUserStationDefaultstationStationIDUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put user station defaultstation station Id unsupported media type response has a 3xx status code
func (o *PutUserStationDefaultstationStationIDUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put user station defaultstation station Id unsupported media type response has a 4xx status code
func (o *PutUserStationDefaultstationStationIDUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put user station defaultstation station Id unsupported media type response has a 5xx status code
func (o *PutUserStationDefaultstationStationIDUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put user station defaultstation station Id unsupported media type response a status code equal to that given
func (o *PutUserStationDefaultstationStationIDUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutUserStationDefaultstationStationIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserStationDefaultstationStationIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserStationDefaultstationStationIDTooManyRequests creates a PutUserStationDefaultstationStationIDTooManyRequests with default headers values
func NewPutUserStationDefaultstationStationIDTooManyRequests() *PutUserStationDefaultstationStationIDTooManyRequests {
	return &PutUserStationDefaultstationStationIDTooManyRequests{}
}

/*
PutUserStationDefaultstationStationIDTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutUserStationDefaultstationStationIDTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put user station defaultstation station Id too many requests response has a 2xx status code
func (o *PutUserStationDefaultstationStationIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put user station defaultstation station Id too many requests response has a 3xx status code
func (o *PutUserStationDefaultstationStationIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put user station defaultstation station Id too many requests response has a 4xx status code
func (o *PutUserStationDefaultstationStationIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put user station defaultstation station Id too many requests response has a 5xx status code
func (o *PutUserStationDefaultstationStationIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put user station defaultstation station Id too many requests response a status code equal to that given
func (o *PutUserStationDefaultstationStationIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutUserStationDefaultstationStationIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserStationDefaultstationStationIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserStationDefaultstationStationIDInternalServerError creates a PutUserStationDefaultstationStationIDInternalServerError with default headers values
func NewPutUserStationDefaultstationStationIDInternalServerError() *PutUserStationDefaultstationStationIDInternalServerError {
	return &PutUserStationDefaultstationStationIDInternalServerError{}
}

/*
PutUserStationDefaultstationStationIDInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutUserStationDefaultstationStationIDInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put user station defaultstation station Id internal server error response has a 2xx status code
func (o *PutUserStationDefaultstationStationIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put user station defaultstation station Id internal server error response has a 3xx status code
func (o *PutUserStationDefaultstationStationIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put user station defaultstation station Id internal server error response has a 4xx status code
func (o *PutUserStationDefaultstationStationIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put user station defaultstation station Id internal server error response has a 5xx status code
func (o *PutUserStationDefaultstationStationIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put user station defaultstation station Id internal server error response a status code equal to that given
func (o *PutUserStationDefaultstationStationIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutUserStationDefaultstationStationIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserStationDefaultstationStationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserStationDefaultstationStationIDServiceUnavailable creates a PutUserStationDefaultstationStationIDServiceUnavailable with default headers values
func NewPutUserStationDefaultstationStationIDServiceUnavailable() *PutUserStationDefaultstationStationIDServiceUnavailable {
	return &PutUserStationDefaultstationStationIDServiceUnavailable{}
}

/*
PutUserStationDefaultstationStationIDServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutUserStationDefaultstationStationIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put user station defaultstation station Id service unavailable response has a 2xx status code
func (o *PutUserStationDefaultstationStationIDServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put user station defaultstation station Id service unavailable response has a 3xx status code
func (o *PutUserStationDefaultstationStationIDServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put user station defaultstation station Id service unavailable response has a 4xx status code
func (o *PutUserStationDefaultstationStationIDServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put user station defaultstation station Id service unavailable response has a 5xx status code
func (o *PutUserStationDefaultstationStationIDServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put user station defaultstation station Id service unavailable response a status code equal to that given
func (o *PutUserStationDefaultstationStationIDServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutUserStationDefaultstationStationIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserStationDefaultstationStationIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserStationDefaultstationStationIDGatewayTimeout creates a PutUserStationDefaultstationStationIDGatewayTimeout with default headers values
func NewPutUserStationDefaultstationStationIDGatewayTimeout() *PutUserStationDefaultstationStationIDGatewayTimeout {
	return &PutUserStationDefaultstationStationIDGatewayTimeout{}
}

/*
PutUserStationDefaultstationStationIDGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutUserStationDefaultstationStationIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put user station defaultstation station Id gateway timeout response has a 2xx status code
func (o *PutUserStationDefaultstationStationIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put user station defaultstation station Id gateway timeout response has a 3xx status code
func (o *PutUserStationDefaultstationStationIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put user station defaultstation station Id gateway timeout response has a 4xx status code
func (o *PutUserStationDefaultstationStationIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put user station defaultstation station Id gateway timeout response has a 5xx status code
func (o *PutUserStationDefaultstationStationIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put user station defaultstation station Id gateway timeout response a status code equal to that given
func (o *PutUserStationDefaultstationStationIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutUserStationDefaultstationStationIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/station/defaultstation/{stationId}][%d] putUserStationDefaultstationStationIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutUserStationDefaultstationStationIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserStationDefaultstationStationIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
