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

// GetOutboundCampaignsDivisionviewReader is a Reader for the GetOutboundCampaignsDivisionview structure.
type GetOutboundCampaignsDivisionviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundCampaignsDivisionviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundCampaignsDivisionviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundCampaignsDivisionviewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundCampaignsDivisionviewUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundCampaignsDivisionviewForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundCampaignsDivisionviewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundCampaignsDivisionviewRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundCampaignsDivisionviewRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundCampaignsDivisionviewUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundCampaignsDivisionviewTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundCampaignsDivisionviewInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundCampaignsDivisionviewServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundCampaignsDivisionviewGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundCampaignsDivisionviewOK creates a GetOutboundCampaignsDivisionviewOK with default headers values
func NewGetOutboundCampaignsDivisionviewOK() *GetOutboundCampaignsDivisionviewOK {
	return &GetOutboundCampaignsDivisionviewOK{}
}

/*
GetOutboundCampaignsDivisionviewOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOutboundCampaignsDivisionviewOK struct {
	Payload *models.CampaignDivisionView
}

// IsSuccess returns true when this get outbound campaigns divisionview o k response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get outbound campaigns divisionview o k response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionview o k response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound campaigns divisionview o k response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaigns divisionview o k response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOutboundCampaignsDivisionviewOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewOK) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewOK) GetPayload() *models.CampaignDivisionView {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CampaignDivisionView)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewBadRequest creates a GetOutboundCampaignsDivisionviewBadRequest with default headers values
func NewGetOutboundCampaignsDivisionviewBadRequest() *GetOutboundCampaignsDivisionviewBadRequest {
	return &GetOutboundCampaignsDivisionviewBadRequest{}
}

/*
GetOutboundCampaignsDivisionviewBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundCampaignsDivisionviewBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionview bad request response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionview bad request response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionview bad request response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaigns divisionview bad request response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaigns divisionview bad request response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOutboundCampaignsDivisionviewBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewUnauthorized creates a GetOutboundCampaignsDivisionviewUnauthorized with default headers values
func NewGetOutboundCampaignsDivisionviewUnauthorized() *GetOutboundCampaignsDivisionviewUnauthorized {
	return &GetOutboundCampaignsDivisionviewUnauthorized{}
}

/*
GetOutboundCampaignsDivisionviewUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundCampaignsDivisionviewUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionview unauthorized response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionview unauthorized response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionview unauthorized response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaigns divisionview unauthorized response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaigns divisionview unauthorized response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOutboundCampaignsDivisionviewUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewForbidden creates a GetOutboundCampaignsDivisionviewForbidden with default headers values
func NewGetOutboundCampaignsDivisionviewForbidden() *GetOutboundCampaignsDivisionviewForbidden {
	return &GetOutboundCampaignsDivisionviewForbidden{}
}

/*
GetOutboundCampaignsDivisionviewForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundCampaignsDivisionviewForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionview forbidden response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionview forbidden response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionview forbidden response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaigns divisionview forbidden response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaigns divisionview forbidden response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOutboundCampaignsDivisionviewForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewNotFound creates a GetOutboundCampaignsDivisionviewNotFound with default headers values
func NewGetOutboundCampaignsDivisionviewNotFound() *GetOutboundCampaignsDivisionviewNotFound {
	return &GetOutboundCampaignsDivisionviewNotFound{}
}

/*
GetOutboundCampaignsDivisionviewNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOutboundCampaignsDivisionviewNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionview not found response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionview not found response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionview not found response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaigns divisionview not found response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaigns divisionview not found response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOutboundCampaignsDivisionviewNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewRequestTimeout creates a GetOutboundCampaignsDivisionviewRequestTimeout with default headers values
func NewGetOutboundCampaignsDivisionviewRequestTimeout() *GetOutboundCampaignsDivisionviewRequestTimeout {
	return &GetOutboundCampaignsDivisionviewRequestTimeout{}
}

/*
GetOutboundCampaignsDivisionviewRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundCampaignsDivisionviewRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionview request timeout response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionview request timeout response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionview request timeout response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaigns divisionview request timeout response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaigns divisionview request timeout response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOutboundCampaignsDivisionviewRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewRequestEntityTooLarge creates a GetOutboundCampaignsDivisionviewRequestEntityTooLarge with default headers values
func NewGetOutboundCampaignsDivisionviewRequestEntityTooLarge() *GetOutboundCampaignsDivisionviewRequestEntityTooLarge {
	return &GetOutboundCampaignsDivisionviewRequestEntityTooLarge{}
}

/*
GetOutboundCampaignsDivisionviewRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetOutboundCampaignsDivisionviewRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionview request entity too large response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionview request entity too large response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionview request entity too large response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaigns divisionview request entity too large response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaigns divisionview request entity too large response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOutboundCampaignsDivisionviewRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewUnsupportedMediaType creates a GetOutboundCampaignsDivisionviewUnsupportedMediaType with default headers values
func NewGetOutboundCampaignsDivisionviewUnsupportedMediaType() *GetOutboundCampaignsDivisionviewUnsupportedMediaType {
	return &GetOutboundCampaignsDivisionviewUnsupportedMediaType{}
}

/*
GetOutboundCampaignsDivisionviewUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundCampaignsDivisionviewUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionview unsupported media type response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionview unsupported media type response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionview unsupported media type response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaigns divisionview unsupported media type response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaigns divisionview unsupported media type response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOutboundCampaignsDivisionviewUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewTooManyRequests creates a GetOutboundCampaignsDivisionviewTooManyRequests with default headers values
func NewGetOutboundCampaignsDivisionviewTooManyRequests() *GetOutboundCampaignsDivisionviewTooManyRequests {
	return &GetOutboundCampaignsDivisionviewTooManyRequests{}
}

/*
GetOutboundCampaignsDivisionviewTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundCampaignsDivisionviewTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionview too many requests response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionview too many requests response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionview too many requests response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaigns divisionview too many requests response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaigns divisionview too many requests response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOutboundCampaignsDivisionviewTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewInternalServerError creates a GetOutboundCampaignsDivisionviewInternalServerError with default headers values
func NewGetOutboundCampaignsDivisionviewInternalServerError() *GetOutboundCampaignsDivisionviewInternalServerError {
	return &GetOutboundCampaignsDivisionviewInternalServerError{}
}

/*
GetOutboundCampaignsDivisionviewInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundCampaignsDivisionviewInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionview internal server error response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionview internal server error response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionview internal server error response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound campaigns divisionview internal server error response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound campaigns divisionview internal server error response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOutboundCampaignsDivisionviewInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewServiceUnavailable creates a GetOutboundCampaignsDivisionviewServiceUnavailable with default headers values
func NewGetOutboundCampaignsDivisionviewServiceUnavailable() *GetOutboundCampaignsDivisionviewServiceUnavailable {
	return &GetOutboundCampaignsDivisionviewServiceUnavailable{}
}

/*
GetOutboundCampaignsDivisionviewServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundCampaignsDivisionviewServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionview service unavailable response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionview service unavailable response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionview service unavailable response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound campaigns divisionview service unavailable response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound campaigns divisionview service unavailable response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOutboundCampaignsDivisionviewServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewGatewayTimeout creates a GetOutboundCampaignsDivisionviewGatewayTimeout with default headers values
func NewGetOutboundCampaignsDivisionviewGatewayTimeout() *GetOutboundCampaignsDivisionviewGatewayTimeout {
	return &GetOutboundCampaignsDivisionviewGatewayTimeout{}
}

/*
GetOutboundCampaignsDivisionviewGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOutboundCampaignsDivisionviewGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionview gateway timeout response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionview gateway timeout response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionview gateway timeout response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound campaigns divisionview gateway timeout response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound campaigns divisionview gateway timeout response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOutboundCampaignsDivisionviewGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews/{campaignId}][%d] getOutboundCampaignsDivisionviewGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
