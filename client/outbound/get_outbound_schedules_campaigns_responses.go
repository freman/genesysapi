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

// GetOutboundSchedulesCampaignsReader is a Reader for the GetOutboundSchedulesCampaigns structure.
type GetOutboundSchedulesCampaignsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundSchedulesCampaignsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundSchedulesCampaignsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundSchedulesCampaignsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundSchedulesCampaignsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundSchedulesCampaignsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundSchedulesCampaignsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundSchedulesCampaignsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundSchedulesCampaignsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundSchedulesCampaignsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundSchedulesCampaignsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundSchedulesCampaignsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundSchedulesCampaignsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundSchedulesCampaignsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundSchedulesCampaignsOK creates a GetOutboundSchedulesCampaignsOK with default headers values
func NewGetOutboundSchedulesCampaignsOK() *GetOutboundSchedulesCampaignsOK {
	return &GetOutboundSchedulesCampaignsOK{}
}

/*
GetOutboundSchedulesCampaignsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOutboundSchedulesCampaignsOK struct {
	Payload []*models.CampaignSchedule
}

// IsSuccess returns true when this get outbound schedules campaigns o k response has a 2xx status code
func (o *GetOutboundSchedulesCampaignsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get outbound schedules campaigns o k response has a 3xx status code
func (o *GetOutboundSchedulesCampaignsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules campaigns o k response has a 4xx status code
func (o *GetOutboundSchedulesCampaignsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound schedules campaigns o k response has a 5xx status code
func (o *GetOutboundSchedulesCampaignsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules campaigns o k response a status code equal to that given
func (o *GetOutboundSchedulesCampaignsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOutboundSchedulesCampaignsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsOK) GetPayload() []*models.CampaignSchedule {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesCampaignsBadRequest creates a GetOutboundSchedulesCampaignsBadRequest with default headers values
func NewGetOutboundSchedulesCampaignsBadRequest() *GetOutboundSchedulesCampaignsBadRequest {
	return &GetOutboundSchedulesCampaignsBadRequest{}
}

/*
GetOutboundSchedulesCampaignsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundSchedulesCampaignsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules campaigns bad request response has a 2xx status code
func (o *GetOutboundSchedulesCampaignsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules campaigns bad request response has a 3xx status code
func (o *GetOutboundSchedulesCampaignsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules campaigns bad request response has a 4xx status code
func (o *GetOutboundSchedulesCampaignsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules campaigns bad request response has a 5xx status code
func (o *GetOutboundSchedulesCampaignsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules campaigns bad request response a status code equal to that given
func (o *GetOutboundSchedulesCampaignsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOutboundSchedulesCampaignsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesCampaignsUnauthorized creates a GetOutboundSchedulesCampaignsUnauthorized with default headers values
func NewGetOutboundSchedulesCampaignsUnauthorized() *GetOutboundSchedulesCampaignsUnauthorized {
	return &GetOutboundSchedulesCampaignsUnauthorized{}
}

/*
GetOutboundSchedulesCampaignsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundSchedulesCampaignsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules campaigns unauthorized response has a 2xx status code
func (o *GetOutboundSchedulesCampaignsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules campaigns unauthorized response has a 3xx status code
func (o *GetOutboundSchedulesCampaignsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules campaigns unauthorized response has a 4xx status code
func (o *GetOutboundSchedulesCampaignsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules campaigns unauthorized response has a 5xx status code
func (o *GetOutboundSchedulesCampaignsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules campaigns unauthorized response a status code equal to that given
func (o *GetOutboundSchedulesCampaignsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOutboundSchedulesCampaignsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesCampaignsForbidden creates a GetOutboundSchedulesCampaignsForbidden with default headers values
func NewGetOutboundSchedulesCampaignsForbidden() *GetOutboundSchedulesCampaignsForbidden {
	return &GetOutboundSchedulesCampaignsForbidden{}
}

/*
GetOutboundSchedulesCampaignsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundSchedulesCampaignsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules campaigns forbidden response has a 2xx status code
func (o *GetOutboundSchedulesCampaignsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules campaigns forbidden response has a 3xx status code
func (o *GetOutboundSchedulesCampaignsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules campaigns forbidden response has a 4xx status code
func (o *GetOutboundSchedulesCampaignsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules campaigns forbidden response has a 5xx status code
func (o *GetOutboundSchedulesCampaignsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules campaigns forbidden response a status code equal to that given
func (o *GetOutboundSchedulesCampaignsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOutboundSchedulesCampaignsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesCampaignsNotFound creates a GetOutboundSchedulesCampaignsNotFound with default headers values
func NewGetOutboundSchedulesCampaignsNotFound() *GetOutboundSchedulesCampaignsNotFound {
	return &GetOutboundSchedulesCampaignsNotFound{}
}

/*
GetOutboundSchedulesCampaignsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOutboundSchedulesCampaignsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules campaigns not found response has a 2xx status code
func (o *GetOutboundSchedulesCampaignsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules campaigns not found response has a 3xx status code
func (o *GetOutboundSchedulesCampaignsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules campaigns not found response has a 4xx status code
func (o *GetOutboundSchedulesCampaignsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules campaigns not found response has a 5xx status code
func (o *GetOutboundSchedulesCampaignsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules campaigns not found response a status code equal to that given
func (o *GetOutboundSchedulesCampaignsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOutboundSchedulesCampaignsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesCampaignsRequestTimeout creates a GetOutboundSchedulesCampaignsRequestTimeout with default headers values
func NewGetOutboundSchedulesCampaignsRequestTimeout() *GetOutboundSchedulesCampaignsRequestTimeout {
	return &GetOutboundSchedulesCampaignsRequestTimeout{}
}

/*
GetOutboundSchedulesCampaignsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundSchedulesCampaignsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules campaigns request timeout response has a 2xx status code
func (o *GetOutboundSchedulesCampaignsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules campaigns request timeout response has a 3xx status code
func (o *GetOutboundSchedulesCampaignsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules campaigns request timeout response has a 4xx status code
func (o *GetOutboundSchedulesCampaignsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules campaigns request timeout response has a 5xx status code
func (o *GetOutboundSchedulesCampaignsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules campaigns request timeout response a status code equal to that given
func (o *GetOutboundSchedulesCampaignsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOutboundSchedulesCampaignsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesCampaignsRequestEntityTooLarge creates a GetOutboundSchedulesCampaignsRequestEntityTooLarge with default headers values
func NewGetOutboundSchedulesCampaignsRequestEntityTooLarge() *GetOutboundSchedulesCampaignsRequestEntityTooLarge {
	return &GetOutboundSchedulesCampaignsRequestEntityTooLarge{}
}

/*
GetOutboundSchedulesCampaignsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetOutboundSchedulesCampaignsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules campaigns request entity too large response has a 2xx status code
func (o *GetOutboundSchedulesCampaignsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules campaigns request entity too large response has a 3xx status code
func (o *GetOutboundSchedulesCampaignsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules campaigns request entity too large response has a 4xx status code
func (o *GetOutboundSchedulesCampaignsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules campaigns request entity too large response has a 5xx status code
func (o *GetOutboundSchedulesCampaignsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules campaigns request entity too large response a status code equal to that given
func (o *GetOutboundSchedulesCampaignsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOutboundSchedulesCampaignsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesCampaignsUnsupportedMediaType creates a GetOutboundSchedulesCampaignsUnsupportedMediaType with default headers values
func NewGetOutboundSchedulesCampaignsUnsupportedMediaType() *GetOutboundSchedulesCampaignsUnsupportedMediaType {
	return &GetOutboundSchedulesCampaignsUnsupportedMediaType{}
}

/*
GetOutboundSchedulesCampaignsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundSchedulesCampaignsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules campaigns unsupported media type response has a 2xx status code
func (o *GetOutboundSchedulesCampaignsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules campaigns unsupported media type response has a 3xx status code
func (o *GetOutboundSchedulesCampaignsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules campaigns unsupported media type response has a 4xx status code
func (o *GetOutboundSchedulesCampaignsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules campaigns unsupported media type response has a 5xx status code
func (o *GetOutboundSchedulesCampaignsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules campaigns unsupported media type response a status code equal to that given
func (o *GetOutboundSchedulesCampaignsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOutboundSchedulesCampaignsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesCampaignsTooManyRequests creates a GetOutboundSchedulesCampaignsTooManyRequests with default headers values
func NewGetOutboundSchedulesCampaignsTooManyRequests() *GetOutboundSchedulesCampaignsTooManyRequests {
	return &GetOutboundSchedulesCampaignsTooManyRequests{}
}

/*
GetOutboundSchedulesCampaignsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundSchedulesCampaignsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules campaigns too many requests response has a 2xx status code
func (o *GetOutboundSchedulesCampaignsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules campaigns too many requests response has a 3xx status code
func (o *GetOutboundSchedulesCampaignsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules campaigns too many requests response has a 4xx status code
func (o *GetOutboundSchedulesCampaignsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound schedules campaigns too many requests response has a 5xx status code
func (o *GetOutboundSchedulesCampaignsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound schedules campaigns too many requests response a status code equal to that given
func (o *GetOutboundSchedulesCampaignsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOutboundSchedulesCampaignsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesCampaignsInternalServerError creates a GetOutboundSchedulesCampaignsInternalServerError with default headers values
func NewGetOutboundSchedulesCampaignsInternalServerError() *GetOutboundSchedulesCampaignsInternalServerError {
	return &GetOutboundSchedulesCampaignsInternalServerError{}
}

/*
GetOutboundSchedulesCampaignsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundSchedulesCampaignsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules campaigns internal server error response has a 2xx status code
func (o *GetOutboundSchedulesCampaignsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules campaigns internal server error response has a 3xx status code
func (o *GetOutboundSchedulesCampaignsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules campaigns internal server error response has a 4xx status code
func (o *GetOutboundSchedulesCampaignsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound schedules campaigns internal server error response has a 5xx status code
func (o *GetOutboundSchedulesCampaignsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound schedules campaigns internal server error response a status code equal to that given
func (o *GetOutboundSchedulesCampaignsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOutboundSchedulesCampaignsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesCampaignsServiceUnavailable creates a GetOutboundSchedulesCampaignsServiceUnavailable with default headers values
func NewGetOutboundSchedulesCampaignsServiceUnavailable() *GetOutboundSchedulesCampaignsServiceUnavailable {
	return &GetOutboundSchedulesCampaignsServiceUnavailable{}
}

/*
GetOutboundSchedulesCampaignsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundSchedulesCampaignsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules campaigns service unavailable response has a 2xx status code
func (o *GetOutboundSchedulesCampaignsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules campaigns service unavailable response has a 3xx status code
func (o *GetOutboundSchedulesCampaignsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules campaigns service unavailable response has a 4xx status code
func (o *GetOutboundSchedulesCampaignsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound schedules campaigns service unavailable response has a 5xx status code
func (o *GetOutboundSchedulesCampaignsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound schedules campaigns service unavailable response a status code equal to that given
func (o *GetOutboundSchedulesCampaignsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOutboundSchedulesCampaignsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesCampaignsGatewayTimeout creates a GetOutboundSchedulesCampaignsGatewayTimeout with default headers values
func NewGetOutboundSchedulesCampaignsGatewayTimeout() *GetOutboundSchedulesCampaignsGatewayTimeout {
	return &GetOutboundSchedulesCampaignsGatewayTimeout{}
}

/*
GetOutboundSchedulesCampaignsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOutboundSchedulesCampaignsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound schedules campaigns gateway timeout response has a 2xx status code
func (o *GetOutboundSchedulesCampaignsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound schedules campaigns gateway timeout response has a 3xx status code
func (o *GetOutboundSchedulesCampaignsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound schedules campaigns gateway timeout response has a 4xx status code
func (o *GetOutboundSchedulesCampaignsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound schedules campaigns gateway timeout response has a 5xx status code
func (o *GetOutboundSchedulesCampaignsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound schedules campaigns gateway timeout response a status code equal to that given
func (o *GetOutboundSchedulesCampaignsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOutboundSchedulesCampaignsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
