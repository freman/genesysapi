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

// GetAlertingInteractionstatsAlertsReader is a Reader for the GetAlertingInteractionstatsAlerts structure.
type GetAlertingInteractionstatsAlertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertingInteractionstatsAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertingInteractionstatsAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAlertingInteractionstatsAlertsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAlertingInteractionstatsAlertsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAlertingInteractionstatsAlertsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAlertingInteractionstatsAlertsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAlertingInteractionstatsAlertsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAlertingInteractionstatsAlertsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAlertingInteractionstatsAlertsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAlertingInteractionstatsAlertsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAlertingInteractionstatsAlertsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAlertingInteractionstatsAlertsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAlertingInteractionstatsAlertsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAlertingInteractionstatsAlertsOK creates a GetAlertingInteractionstatsAlertsOK with default headers values
func NewGetAlertingInteractionstatsAlertsOK() *GetAlertingInteractionstatsAlertsOK {
	return &GetAlertingInteractionstatsAlertsOK{}
}

/*
GetAlertingInteractionstatsAlertsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAlertingInteractionstatsAlertsOK struct {
	Payload *models.InteractionStatsAlertContainer
}

// IsSuccess returns true when this get alerting interactionstats alerts o k response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alerting interactionstats alerts o k response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts o k response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alerting interactionstats alerts o k response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats alerts o k response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAlertingInteractionstatsAlertsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsOK  %+v", 200, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsOK  %+v", 200, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsOK) GetPayload() *models.InteractionStatsAlertContainer {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InteractionStatsAlertContainer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsBadRequest creates a GetAlertingInteractionstatsAlertsBadRequest with default headers values
func NewGetAlertingInteractionstatsAlertsBadRequest() *GetAlertingInteractionstatsAlertsBadRequest {
	return &GetAlertingInteractionstatsAlertsBadRequest{}
}

/*
GetAlertingInteractionstatsAlertsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAlertingInteractionstatsAlertsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts bad request response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts bad request response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts bad request response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats alerts bad request response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats alerts bad request response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAlertingInteractionstatsAlertsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnauthorized creates a GetAlertingInteractionstatsAlertsUnauthorized with default headers values
func NewGetAlertingInteractionstatsAlertsUnauthorized() *GetAlertingInteractionstatsAlertsUnauthorized {
	return &GetAlertingInteractionstatsAlertsUnauthorized{}
}

/*
GetAlertingInteractionstatsAlertsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAlertingInteractionstatsAlertsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts unauthorized response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts unauthorized response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts unauthorized response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats alerts unauthorized response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats alerts unauthorized response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAlertingInteractionstatsAlertsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsForbidden creates a GetAlertingInteractionstatsAlertsForbidden with default headers values
func NewGetAlertingInteractionstatsAlertsForbidden() *GetAlertingInteractionstatsAlertsForbidden {
	return &GetAlertingInteractionstatsAlertsForbidden{}
}

/*
GetAlertingInteractionstatsAlertsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetAlertingInteractionstatsAlertsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts forbidden response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts forbidden response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts forbidden response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats alerts forbidden response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats alerts forbidden response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAlertingInteractionstatsAlertsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsForbidden  %+v", 403, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsForbidden  %+v", 403, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsNotFound creates a GetAlertingInteractionstatsAlertsNotFound with default headers values
func NewGetAlertingInteractionstatsAlertsNotFound() *GetAlertingInteractionstatsAlertsNotFound {
	return &GetAlertingInteractionstatsAlertsNotFound{}
}

/*
GetAlertingInteractionstatsAlertsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetAlertingInteractionstatsAlertsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts not found response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts not found response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts not found response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats alerts not found response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats alerts not found response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAlertingInteractionstatsAlertsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsNotFound  %+v", 404, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsNotFound  %+v", 404, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsRequestTimeout creates a GetAlertingInteractionstatsAlertsRequestTimeout with default headers values
func NewGetAlertingInteractionstatsAlertsRequestTimeout() *GetAlertingInteractionstatsAlertsRequestTimeout {
	return &GetAlertingInteractionstatsAlertsRequestTimeout{}
}

/*
GetAlertingInteractionstatsAlertsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAlertingInteractionstatsAlertsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts request timeout response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts request timeout response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts request timeout response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats alerts request timeout response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats alerts request timeout response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetAlertingInteractionstatsAlertsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsRequestEntityTooLarge creates a GetAlertingInteractionstatsAlertsRequestEntityTooLarge with default headers values
func NewGetAlertingInteractionstatsAlertsRequestEntityTooLarge() *GetAlertingInteractionstatsAlertsRequestEntityTooLarge {
	return &GetAlertingInteractionstatsAlertsRequestEntityTooLarge{}
}

/*
GetAlertingInteractionstatsAlertsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetAlertingInteractionstatsAlertsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts request entity too large response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts request entity too large response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts request entity too large response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats alerts request entity too large response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats alerts request entity too large response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetAlertingInteractionstatsAlertsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsUnsupportedMediaType creates a GetAlertingInteractionstatsAlertsUnsupportedMediaType with default headers values
func NewGetAlertingInteractionstatsAlertsUnsupportedMediaType() *GetAlertingInteractionstatsAlertsUnsupportedMediaType {
	return &GetAlertingInteractionstatsAlertsUnsupportedMediaType{}
}

/*
GetAlertingInteractionstatsAlertsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAlertingInteractionstatsAlertsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts unsupported media type response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts unsupported media type response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts unsupported media type response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats alerts unsupported media type response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats alerts unsupported media type response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetAlertingInteractionstatsAlertsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsTooManyRequests creates a GetAlertingInteractionstatsAlertsTooManyRequests with default headers values
func NewGetAlertingInteractionstatsAlertsTooManyRequests() *GetAlertingInteractionstatsAlertsTooManyRequests {
	return &GetAlertingInteractionstatsAlertsTooManyRequests{}
}

/*
GetAlertingInteractionstatsAlertsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAlertingInteractionstatsAlertsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts too many requests response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts too many requests response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts too many requests response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats alerts too many requests response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats alerts too many requests response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetAlertingInteractionstatsAlertsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsInternalServerError creates a GetAlertingInteractionstatsAlertsInternalServerError with default headers values
func NewGetAlertingInteractionstatsAlertsInternalServerError() *GetAlertingInteractionstatsAlertsInternalServerError {
	return &GetAlertingInteractionstatsAlertsInternalServerError{}
}

/*
GetAlertingInteractionstatsAlertsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAlertingInteractionstatsAlertsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts internal server error response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts internal server error response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts internal server error response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alerting interactionstats alerts internal server error response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get alerting interactionstats alerts internal server error response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAlertingInteractionstatsAlertsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsServiceUnavailable creates a GetAlertingInteractionstatsAlertsServiceUnavailable with default headers values
func NewGetAlertingInteractionstatsAlertsServiceUnavailable() *GetAlertingInteractionstatsAlertsServiceUnavailable {
	return &GetAlertingInteractionstatsAlertsServiceUnavailable{}
}

/*
GetAlertingInteractionstatsAlertsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAlertingInteractionstatsAlertsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts service unavailable response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts service unavailable response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts service unavailable response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alerting interactionstats alerts service unavailable response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get alerting interactionstats alerts service unavailable response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetAlertingInteractionstatsAlertsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsAlertsGatewayTimeout creates a GetAlertingInteractionstatsAlertsGatewayTimeout with default headers values
func NewGetAlertingInteractionstatsAlertsGatewayTimeout() *GetAlertingInteractionstatsAlertsGatewayTimeout {
	return &GetAlertingInteractionstatsAlertsGatewayTimeout{}
}

/*
GetAlertingInteractionstatsAlertsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetAlertingInteractionstatsAlertsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats alerts gateway timeout response has a 2xx status code
func (o *GetAlertingInteractionstatsAlertsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats alerts gateway timeout response has a 3xx status code
func (o *GetAlertingInteractionstatsAlertsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats alerts gateway timeout response has a 4xx status code
func (o *GetAlertingInteractionstatsAlertsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alerting interactionstats alerts gateway timeout response has a 5xx status code
func (o *GetAlertingInteractionstatsAlertsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get alerting interactionstats alerts gateway timeout response a status code equal to that given
func (o *GetAlertingInteractionstatsAlertsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetAlertingInteractionstatsAlertsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/alerts][%d] getAlertingInteractionstatsAlertsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAlertingInteractionstatsAlertsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsAlertsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
