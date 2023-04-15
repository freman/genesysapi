// Code generated by go-swagger; DO NOT EDIT.

package alerting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetAlertingAlertsActiveReader is a Reader for the GetAlertingAlertsActive structure.
type GetAlertingAlertsActiveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertingAlertsActiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertingAlertsActiveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAlertingAlertsActiveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAlertingAlertsActiveUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAlertingAlertsActiveForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAlertingAlertsActiveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAlertingAlertsActiveRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAlertingAlertsActiveRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAlertingAlertsActiveUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAlertingAlertsActiveTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAlertingAlertsActiveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAlertingAlertsActiveServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAlertingAlertsActiveGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAlertingAlertsActiveOK creates a GetAlertingAlertsActiveOK with default headers values
func NewGetAlertingAlertsActiveOK() *GetAlertingAlertsActiveOK {
	return &GetAlertingAlertsActiveOK{}
}

/*
GetAlertingAlertsActiveOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAlertingAlertsActiveOK struct {
	Payload *models.ActiveAlertCount
}

// IsSuccess returns true when this get alerting alerts active o k response has a 2xx status code
func (o *GetAlertingAlertsActiveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alerting alerts active o k response has a 3xx status code
func (o *GetAlertingAlertsActiveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting alerts active o k response has a 4xx status code
func (o *GetAlertingAlertsActiveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alerting alerts active o k response has a 5xx status code
func (o *GetAlertingAlertsActiveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting alerts active o k response a status code equal to that given
func (o *GetAlertingAlertsActiveOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAlertingAlertsActiveOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveOK  %+v", 200, o.Payload)
}

func (o *GetAlertingAlertsActiveOK) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveOK  %+v", 200, o.Payload)
}

func (o *GetAlertingAlertsActiveOK) GetPayload() *models.ActiveAlertCount {
	return o.Payload
}

func (o *GetAlertingAlertsActiveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActiveAlertCount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingAlertsActiveBadRequest creates a GetAlertingAlertsActiveBadRequest with default headers values
func NewGetAlertingAlertsActiveBadRequest() *GetAlertingAlertsActiveBadRequest {
	return &GetAlertingAlertsActiveBadRequest{}
}

/*
GetAlertingAlertsActiveBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAlertingAlertsActiveBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting alerts active bad request response has a 2xx status code
func (o *GetAlertingAlertsActiveBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting alerts active bad request response has a 3xx status code
func (o *GetAlertingAlertsActiveBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting alerts active bad request response has a 4xx status code
func (o *GetAlertingAlertsActiveBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting alerts active bad request response has a 5xx status code
func (o *GetAlertingAlertsActiveBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting alerts active bad request response a status code equal to that given
func (o *GetAlertingAlertsActiveBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAlertingAlertsActiveBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveBadRequest  %+v", 400, o.Payload)
}

func (o *GetAlertingAlertsActiveBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveBadRequest  %+v", 400, o.Payload)
}

func (o *GetAlertingAlertsActiveBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingAlertsActiveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingAlertsActiveUnauthorized creates a GetAlertingAlertsActiveUnauthorized with default headers values
func NewGetAlertingAlertsActiveUnauthorized() *GetAlertingAlertsActiveUnauthorized {
	return &GetAlertingAlertsActiveUnauthorized{}
}

/*
GetAlertingAlertsActiveUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAlertingAlertsActiveUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting alerts active unauthorized response has a 2xx status code
func (o *GetAlertingAlertsActiveUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting alerts active unauthorized response has a 3xx status code
func (o *GetAlertingAlertsActiveUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting alerts active unauthorized response has a 4xx status code
func (o *GetAlertingAlertsActiveUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting alerts active unauthorized response has a 5xx status code
func (o *GetAlertingAlertsActiveUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting alerts active unauthorized response a status code equal to that given
func (o *GetAlertingAlertsActiveUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAlertingAlertsActiveUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAlertingAlertsActiveUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAlertingAlertsActiveUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingAlertsActiveUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingAlertsActiveForbidden creates a GetAlertingAlertsActiveForbidden with default headers values
func NewGetAlertingAlertsActiveForbidden() *GetAlertingAlertsActiveForbidden {
	return &GetAlertingAlertsActiveForbidden{}
}

/*
GetAlertingAlertsActiveForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetAlertingAlertsActiveForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting alerts active forbidden response has a 2xx status code
func (o *GetAlertingAlertsActiveForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting alerts active forbidden response has a 3xx status code
func (o *GetAlertingAlertsActiveForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting alerts active forbidden response has a 4xx status code
func (o *GetAlertingAlertsActiveForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting alerts active forbidden response has a 5xx status code
func (o *GetAlertingAlertsActiveForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting alerts active forbidden response a status code equal to that given
func (o *GetAlertingAlertsActiveForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAlertingAlertsActiveForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveForbidden  %+v", 403, o.Payload)
}

func (o *GetAlertingAlertsActiveForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveForbidden  %+v", 403, o.Payload)
}

func (o *GetAlertingAlertsActiveForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingAlertsActiveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingAlertsActiveNotFound creates a GetAlertingAlertsActiveNotFound with default headers values
func NewGetAlertingAlertsActiveNotFound() *GetAlertingAlertsActiveNotFound {
	return &GetAlertingAlertsActiveNotFound{}
}

/*
GetAlertingAlertsActiveNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetAlertingAlertsActiveNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting alerts active not found response has a 2xx status code
func (o *GetAlertingAlertsActiveNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting alerts active not found response has a 3xx status code
func (o *GetAlertingAlertsActiveNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting alerts active not found response has a 4xx status code
func (o *GetAlertingAlertsActiveNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting alerts active not found response has a 5xx status code
func (o *GetAlertingAlertsActiveNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting alerts active not found response a status code equal to that given
func (o *GetAlertingAlertsActiveNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAlertingAlertsActiveNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveNotFound  %+v", 404, o.Payload)
}

func (o *GetAlertingAlertsActiveNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveNotFound  %+v", 404, o.Payload)
}

func (o *GetAlertingAlertsActiveNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingAlertsActiveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingAlertsActiveRequestTimeout creates a GetAlertingAlertsActiveRequestTimeout with default headers values
func NewGetAlertingAlertsActiveRequestTimeout() *GetAlertingAlertsActiveRequestTimeout {
	return &GetAlertingAlertsActiveRequestTimeout{}
}

/*
GetAlertingAlertsActiveRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAlertingAlertsActiveRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting alerts active request timeout response has a 2xx status code
func (o *GetAlertingAlertsActiveRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting alerts active request timeout response has a 3xx status code
func (o *GetAlertingAlertsActiveRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting alerts active request timeout response has a 4xx status code
func (o *GetAlertingAlertsActiveRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting alerts active request timeout response has a 5xx status code
func (o *GetAlertingAlertsActiveRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting alerts active request timeout response a status code equal to that given
func (o *GetAlertingAlertsActiveRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetAlertingAlertsActiveRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAlertingAlertsActiveRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAlertingAlertsActiveRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingAlertsActiveRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingAlertsActiveRequestEntityTooLarge creates a GetAlertingAlertsActiveRequestEntityTooLarge with default headers values
func NewGetAlertingAlertsActiveRequestEntityTooLarge() *GetAlertingAlertsActiveRequestEntityTooLarge {
	return &GetAlertingAlertsActiveRequestEntityTooLarge{}
}

/*
GetAlertingAlertsActiveRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetAlertingAlertsActiveRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting alerts active request entity too large response has a 2xx status code
func (o *GetAlertingAlertsActiveRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting alerts active request entity too large response has a 3xx status code
func (o *GetAlertingAlertsActiveRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting alerts active request entity too large response has a 4xx status code
func (o *GetAlertingAlertsActiveRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting alerts active request entity too large response has a 5xx status code
func (o *GetAlertingAlertsActiveRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting alerts active request entity too large response a status code equal to that given
func (o *GetAlertingAlertsActiveRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetAlertingAlertsActiveRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAlertingAlertsActiveRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAlertingAlertsActiveRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingAlertsActiveRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingAlertsActiveUnsupportedMediaType creates a GetAlertingAlertsActiveUnsupportedMediaType with default headers values
func NewGetAlertingAlertsActiveUnsupportedMediaType() *GetAlertingAlertsActiveUnsupportedMediaType {
	return &GetAlertingAlertsActiveUnsupportedMediaType{}
}

/*
GetAlertingAlertsActiveUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAlertingAlertsActiveUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting alerts active unsupported media type response has a 2xx status code
func (o *GetAlertingAlertsActiveUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting alerts active unsupported media type response has a 3xx status code
func (o *GetAlertingAlertsActiveUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting alerts active unsupported media type response has a 4xx status code
func (o *GetAlertingAlertsActiveUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting alerts active unsupported media type response has a 5xx status code
func (o *GetAlertingAlertsActiveUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting alerts active unsupported media type response a status code equal to that given
func (o *GetAlertingAlertsActiveUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetAlertingAlertsActiveUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAlertingAlertsActiveUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAlertingAlertsActiveUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingAlertsActiveUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingAlertsActiveTooManyRequests creates a GetAlertingAlertsActiveTooManyRequests with default headers values
func NewGetAlertingAlertsActiveTooManyRequests() *GetAlertingAlertsActiveTooManyRequests {
	return &GetAlertingAlertsActiveTooManyRequests{}
}

/*
GetAlertingAlertsActiveTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAlertingAlertsActiveTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting alerts active too many requests response has a 2xx status code
func (o *GetAlertingAlertsActiveTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting alerts active too many requests response has a 3xx status code
func (o *GetAlertingAlertsActiveTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting alerts active too many requests response has a 4xx status code
func (o *GetAlertingAlertsActiveTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting alerts active too many requests response has a 5xx status code
func (o *GetAlertingAlertsActiveTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting alerts active too many requests response a status code equal to that given
func (o *GetAlertingAlertsActiveTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetAlertingAlertsActiveTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAlertingAlertsActiveTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAlertingAlertsActiveTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingAlertsActiveTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingAlertsActiveInternalServerError creates a GetAlertingAlertsActiveInternalServerError with default headers values
func NewGetAlertingAlertsActiveInternalServerError() *GetAlertingAlertsActiveInternalServerError {
	return &GetAlertingAlertsActiveInternalServerError{}
}

/*
GetAlertingAlertsActiveInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAlertingAlertsActiveInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting alerts active internal server error response has a 2xx status code
func (o *GetAlertingAlertsActiveInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting alerts active internal server error response has a 3xx status code
func (o *GetAlertingAlertsActiveInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting alerts active internal server error response has a 4xx status code
func (o *GetAlertingAlertsActiveInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alerting alerts active internal server error response has a 5xx status code
func (o *GetAlertingAlertsActiveInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get alerting alerts active internal server error response a status code equal to that given
func (o *GetAlertingAlertsActiveInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAlertingAlertsActiveInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlertingAlertsActiveInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlertingAlertsActiveInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingAlertsActiveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingAlertsActiveServiceUnavailable creates a GetAlertingAlertsActiveServiceUnavailable with default headers values
func NewGetAlertingAlertsActiveServiceUnavailable() *GetAlertingAlertsActiveServiceUnavailable {
	return &GetAlertingAlertsActiveServiceUnavailable{}
}

/*
GetAlertingAlertsActiveServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAlertingAlertsActiveServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting alerts active service unavailable response has a 2xx status code
func (o *GetAlertingAlertsActiveServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting alerts active service unavailable response has a 3xx status code
func (o *GetAlertingAlertsActiveServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting alerts active service unavailable response has a 4xx status code
func (o *GetAlertingAlertsActiveServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alerting alerts active service unavailable response has a 5xx status code
func (o *GetAlertingAlertsActiveServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get alerting alerts active service unavailable response a status code equal to that given
func (o *GetAlertingAlertsActiveServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetAlertingAlertsActiveServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAlertingAlertsActiveServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAlertingAlertsActiveServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingAlertsActiveServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingAlertsActiveGatewayTimeout creates a GetAlertingAlertsActiveGatewayTimeout with default headers values
func NewGetAlertingAlertsActiveGatewayTimeout() *GetAlertingAlertsActiveGatewayTimeout {
	return &GetAlertingAlertsActiveGatewayTimeout{}
}

/*
GetAlertingAlertsActiveGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetAlertingAlertsActiveGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting alerts active gateway timeout response has a 2xx status code
func (o *GetAlertingAlertsActiveGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting alerts active gateway timeout response has a 3xx status code
func (o *GetAlertingAlertsActiveGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting alerts active gateway timeout response has a 4xx status code
func (o *GetAlertingAlertsActiveGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alerting alerts active gateway timeout response has a 5xx status code
func (o *GetAlertingAlertsActiveGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get alerting alerts active gateway timeout response a status code equal to that given
func (o *GetAlertingAlertsActiveGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetAlertingAlertsActiveGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAlertingAlertsActiveGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/alerts/active][%d] getAlertingAlertsActiveGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAlertingAlertsActiveGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingAlertsActiveGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
