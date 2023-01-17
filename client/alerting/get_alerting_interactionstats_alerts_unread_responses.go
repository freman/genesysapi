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

// GetAlertingInteractionstatsAlertsUnreadReader is a Reader for the GetAlertingInteractionstatsAlertsUnread structure.
type GetAlertingInteractionstatsAlertsUnreadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertingInteractionstatsAlertsUnreadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertingInteractionstatsAlertsUnreadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAlertingInteractionstatsAlertsUnreadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAlertingInteractionstatsAlertsUnreadUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAlertingInteractionstatsAlertsUnreadForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAlertingInteractionstatsAlertsUnreadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAlertingInteractionstatsAlertsUnreadRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAlertingInteractionstatsAlertsUnreadTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAlertingInteractionstatsAlertsUnreadInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAlertingInteractionstatsAlertsUnreadServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAlertingInteractionstatsAlertsUnreadGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAlertingInteractionstatsAlertsUnreadOK creates a GetAlertingInteractionstatsAlertsUnreadOK with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadOK() *GetAlertingInteractionstatsAlertsUnreadOK {
	return &GetAlertingInteractionstatsAlertsUnreadOK{}
}

/*
GetAlertingInteractionstatsAlertsUnreadOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAlertingInteractionstatsAlertsUnreadOK struct {
	Payload *models.UnreadMetric
}

// IsSuccess returns true when this get alerting interactionstats alerts unread o k response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alerting interactionstats alerts unread o k response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts unread o k response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alerting interactionstats alerts unread o k response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats alerts unread o k response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsUnreadOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAlertingInteractionstatsAlertsUnreadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadOK  %+v", 200, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadOK) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadOK  %+v", 200, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadOK) GetPayload() *models.UnreadMetric {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnreadMetric)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnreadBadRequest creates a GetAlertingInteractionstatsAlertsUnreadBadRequest with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadBadRequest() *GetAlertingInteractionstatsAlertsUnreadBadRequest {
	return &GetAlertingInteractionstatsAlertsUnreadBadRequest{}
}

/*
GetAlertingInteractionstatsAlertsUnreadBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAlertingInteractionstatsAlertsUnreadBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts unread bad request response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts unread bad request response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts unread bad request response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats alerts unread bad request response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats alerts unread bad request response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsUnreadBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAlertingInteractionstatsAlertsUnreadBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadBadRequest  %+v", 400, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadBadRequest  %+v", 400, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnreadUnauthorized creates a GetAlertingInteractionstatsAlertsUnreadUnauthorized with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadUnauthorized() *GetAlertingInteractionstatsAlertsUnreadUnauthorized {
	return &GetAlertingInteractionstatsAlertsUnreadUnauthorized{}
}

/*
GetAlertingInteractionstatsAlertsUnreadUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAlertingInteractionstatsAlertsUnreadUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts unread unauthorized response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts unread unauthorized response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts unread unauthorized response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats alerts unread unauthorized response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats alerts unread unauthorized response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsUnreadUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAlertingInteractionstatsAlertsUnreadUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnreadForbidden creates a GetAlertingInteractionstatsAlertsUnreadForbidden with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadForbidden() *GetAlertingInteractionstatsAlertsUnreadForbidden {
	return &GetAlertingInteractionstatsAlertsUnreadForbidden{}
}

/*
GetAlertingInteractionstatsAlertsUnreadForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetAlertingInteractionstatsAlertsUnreadForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts unread forbidden response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts unread forbidden response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts unread forbidden response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats alerts unread forbidden response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats alerts unread forbidden response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsUnreadForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAlertingInteractionstatsAlertsUnreadForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadForbidden  %+v", 403, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadForbidden  %+v", 403, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnreadNotFound creates a GetAlertingInteractionstatsAlertsUnreadNotFound with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadNotFound() *GetAlertingInteractionstatsAlertsUnreadNotFound {
	return &GetAlertingInteractionstatsAlertsUnreadNotFound{}
}

/*
GetAlertingInteractionstatsAlertsUnreadNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetAlertingInteractionstatsAlertsUnreadNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts unread not found response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts unread not found response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts unread not found response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats alerts unread not found response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats alerts unread not found response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsUnreadNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAlertingInteractionstatsAlertsUnreadNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadNotFound  %+v", 404, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadNotFound  %+v", 404, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnreadRequestTimeout creates a GetAlertingInteractionstatsAlertsUnreadRequestTimeout with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadRequestTimeout() *GetAlertingInteractionstatsAlertsUnreadRequestTimeout {
	return &GetAlertingInteractionstatsAlertsUnreadRequestTimeout{}
}

/*
GetAlertingInteractionstatsAlertsUnreadRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAlertingInteractionstatsAlertsUnreadRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts unread request timeout response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts unread request timeout response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts unread request timeout response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats alerts unread request timeout response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats alerts unread request timeout response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsUnreadRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetAlertingInteractionstatsAlertsUnreadRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge creates a GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge() *GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge {
	return &GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge{}
}

/*
GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts unread request entity too large response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts unread request entity too large response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts unread request entity too large response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats alerts unread request entity too large response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats alerts unread request entity too large response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType creates a GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType() *GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType {
	return &GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType{}
}

/*
GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts unread unsupported media type response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts unread unsupported media type response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts unread unsupported media type response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats alerts unread unsupported media type response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats alerts unread unsupported media type response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnreadTooManyRequests creates a GetAlertingInteractionstatsAlertsUnreadTooManyRequests with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadTooManyRequests() *GetAlertingInteractionstatsAlertsUnreadTooManyRequests {
	return &GetAlertingInteractionstatsAlertsUnreadTooManyRequests{}
}

/*
GetAlertingInteractionstatsAlertsUnreadTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAlertingInteractionstatsAlertsUnreadTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts unread too many requests response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts unread too many requests response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts unread too many requests response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats alerts unread too many requests response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats alerts unread too many requests response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsUnreadTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetAlertingInteractionstatsAlertsUnreadTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnreadInternalServerError creates a GetAlertingInteractionstatsAlertsUnreadInternalServerError with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadInternalServerError() *GetAlertingInteractionstatsAlertsUnreadInternalServerError {
	return &GetAlertingInteractionstatsAlertsUnreadInternalServerError{}
}

/*
GetAlertingInteractionstatsAlertsUnreadInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAlertingInteractionstatsAlertsUnreadInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts unread internal server error response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts unread internal server error response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts unread internal server error response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alerting interactionstats alerts unread internal server error response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get alerting interactionstats alerts unread internal server error response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsUnreadInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAlertingInteractionstatsAlertsUnreadInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnreadServiceUnavailable creates a GetAlertingInteractionstatsAlertsUnreadServiceUnavailable with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadServiceUnavailable() *GetAlertingInteractionstatsAlertsUnreadServiceUnavailable {
	return &GetAlertingInteractionstatsAlertsUnreadServiceUnavailable{}
}

/*
GetAlertingInteractionstatsAlertsUnreadServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAlertingInteractionstatsAlertsUnreadServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts unread service unavailable response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts unread service unavailable response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts unread service unavailable response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alerting interactionstats alerts unread service unavailable response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get alerting interactionstats alerts unread service unavailable response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsUnreadServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetAlertingInteractionstatsAlertsUnreadServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnreadGatewayTimeout creates a GetAlertingInteractionstatsAlertsUnreadGatewayTimeout with default headers values
func NewGetAlertingInteractionstatsAlertsUnreadGatewayTimeout() *GetAlertingInteractionstatsAlertsUnreadGatewayTimeout {
	return &GetAlertingInteractionstatsAlertsUnreadGatewayTimeout{}
}

/*
GetAlertingInteractionstatsAlertsUnreadGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetAlertingInteractionstatsAlertsUnreadGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts unread gateway timeout response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts unread gateway timeout response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts unread gateway timeout response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alerting interactionstats alerts unread gateway timeout response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsUnreadGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get alerting interactionstats alerts unread gateway timeout response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsUnreadGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetAlertingInteractionstatsAlertsUnreadGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts/unread][%d] getAlertingInteractionstatsAlertsUnreadGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnreadGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnreadGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
