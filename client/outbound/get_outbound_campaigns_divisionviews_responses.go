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

// GetOutboundCampaignsDivisionviewsReader is a Reader for the GetOutboundCampaignsDivisionviews structure.
type GetOutboundCampaignsDivisionviewsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundCampaignsDivisionviewsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundCampaignsDivisionviewsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundCampaignsDivisionviewsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundCampaignsDivisionviewsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundCampaignsDivisionviewsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundCampaignsDivisionviewsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundCampaignsDivisionviewsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundCampaignsDivisionviewsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundCampaignsDivisionviewsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundCampaignsDivisionviewsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundCampaignsDivisionviewsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundCampaignsDivisionviewsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundCampaignsDivisionviewsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundCampaignsDivisionviewsOK creates a GetOutboundCampaignsDivisionviewsOK with default headers values
func NewGetOutboundCampaignsDivisionviewsOK() *GetOutboundCampaignsDivisionviewsOK {
	return &GetOutboundCampaignsDivisionviewsOK{}
}

/*
GetOutboundCampaignsDivisionviewsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOutboundCampaignsDivisionviewsOK struct {
	Payload *models.CampaignDivisionViewListing
}

// IsSuccess returns true when this get outbound campaigns divisionviews o k response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get outbound campaigns divisionviews o k response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionviews o k response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound campaigns divisionviews o k response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaigns divisionviews o k response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOutboundCampaignsDivisionviewsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsOK) GetPayload() *models.CampaignDivisionViewListing {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CampaignDivisionViewListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewsBadRequest creates a GetOutboundCampaignsDivisionviewsBadRequest with default headers values
func NewGetOutboundCampaignsDivisionviewsBadRequest() *GetOutboundCampaignsDivisionviewsBadRequest {
	return &GetOutboundCampaignsDivisionviewsBadRequest{}
}

/*
GetOutboundCampaignsDivisionviewsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundCampaignsDivisionviewsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionviews bad request response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionviews bad request response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionviews bad request response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaigns divisionviews bad request response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaigns divisionviews bad request response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOutboundCampaignsDivisionviewsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewsUnauthorized creates a GetOutboundCampaignsDivisionviewsUnauthorized with default headers values
func NewGetOutboundCampaignsDivisionviewsUnauthorized() *GetOutboundCampaignsDivisionviewsUnauthorized {
	return &GetOutboundCampaignsDivisionviewsUnauthorized{}
}

/*
GetOutboundCampaignsDivisionviewsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundCampaignsDivisionviewsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionviews unauthorized response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionviews unauthorized response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionviews unauthorized response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaigns divisionviews unauthorized response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaigns divisionviews unauthorized response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOutboundCampaignsDivisionviewsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewsForbidden creates a GetOutboundCampaignsDivisionviewsForbidden with default headers values
func NewGetOutboundCampaignsDivisionviewsForbidden() *GetOutboundCampaignsDivisionviewsForbidden {
	return &GetOutboundCampaignsDivisionviewsForbidden{}
}

/*
GetOutboundCampaignsDivisionviewsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundCampaignsDivisionviewsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionviews forbidden response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionviews forbidden response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionviews forbidden response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaigns divisionviews forbidden response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaigns divisionviews forbidden response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOutboundCampaignsDivisionviewsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewsNotFound creates a GetOutboundCampaignsDivisionviewsNotFound with default headers values
func NewGetOutboundCampaignsDivisionviewsNotFound() *GetOutboundCampaignsDivisionviewsNotFound {
	return &GetOutboundCampaignsDivisionviewsNotFound{}
}

/*
GetOutboundCampaignsDivisionviewsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOutboundCampaignsDivisionviewsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionviews not found response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionviews not found response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionviews not found response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaigns divisionviews not found response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaigns divisionviews not found response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOutboundCampaignsDivisionviewsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewsRequestTimeout creates a GetOutboundCampaignsDivisionviewsRequestTimeout with default headers values
func NewGetOutboundCampaignsDivisionviewsRequestTimeout() *GetOutboundCampaignsDivisionviewsRequestTimeout {
	return &GetOutboundCampaignsDivisionviewsRequestTimeout{}
}

/*
GetOutboundCampaignsDivisionviewsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundCampaignsDivisionviewsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionviews request timeout response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionviews request timeout response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionviews request timeout response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaigns divisionviews request timeout response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaigns divisionviews request timeout response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOutboundCampaignsDivisionviewsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewsRequestEntityTooLarge creates a GetOutboundCampaignsDivisionviewsRequestEntityTooLarge with default headers values
func NewGetOutboundCampaignsDivisionviewsRequestEntityTooLarge() *GetOutboundCampaignsDivisionviewsRequestEntityTooLarge {
	return &GetOutboundCampaignsDivisionviewsRequestEntityTooLarge{}
}

/*
GetOutboundCampaignsDivisionviewsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetOutboundCampaignsDivisionviewsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionviews request entity too large response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionviews request entity too large response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionviews request entity too large response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaigns divisionviews request entity too large response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaigns divisionviews request entity too large response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOutboundCampaignsDivisionviewsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewsUnsupportedMediaType creates a GetOutboundCampaignsDivisionviewsUnsupportedMediaType with default headers values
func NewGetOutboundCampaignsDivisionviewsUnsupportedMediaType() *GetOutboundCampaignsDivisionviewsUnsupportedMediaType {
	return &GetOutboundCampaignsDivisionviewsUnsupportedMediaType{}
}

/*
GetOutboundCampaignsDivisionviewsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundCampaignsDivisionviewsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionviews unsupported media type response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionviews unsupported media type response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionviews unsupported media type response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaigns divisionviews unsupported media type response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaigns divisionviews unsupported media type response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOutboundCampaignsDivisionviewsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewsTooManyRequests creates a GetOutboundCampaignsDivisionviewsTooManyRequests with default headers values
func NewGetOutboundCampaignsDivisionviewsTooManyRequests() *GetOutboundCampaignsDivisionviewsTooManyRequests {
	return &GetOutboundCampaignsDivisionviewsTooManyRequests{}
}

/*
GetOutboundCampaignsDivisionviewsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundCampaignsDivisionviewsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionviews too many requests response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionviews too many requests response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionviews too many requests response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaigns divisionviews too many requests response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaigns divisionviews too many requests response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOutboundCampaignsDivisionviewsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewsInternalServerError creates a GetOutboundCampaignsDivisionviewsInternalServerError with default headers values
func NewGetOutboundCampaignsDivisionviewsInternalServerError() *GetOutboundCampaignsDivisionviewsInternalServerError {
	return &GetOutboundCampaignsDivisionviewsInternalServerError{}
}

/*
GetOutboundCampaignsDivisionviewsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundCampaignsDivisionviewsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionviews internal server error response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionviews internal server error response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionviews internal server error response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound campaigns divisionviews internal server error response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound campaigns divisionviews internal server error response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOutboundCampaignsDivisionviewsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewsServiceUnavailable creates a GetOutboundCampaignsDivisionviewsServiceUnavailable with default headers values
func NewGetOutboundCampaignsDivisionviewsServiceUnavailable() *GetOutboundCampaignsDivisionviewsServiceUnavailable {
	return &GetOutboundCampaignsDivisionviewsServiceUnavailable{}
}

/*
GetOutboundCampaignsDivisionviewsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundCampaignsDivisionviewsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionviews service unavailable response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionviews service unavailable response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionviews service unavailable response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound campaigns divisionviews service unavailable response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound campaigns divisionviews service unavailable response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOutboundCampaignsDivisionviewsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignsDivisionviewsGatewayTimeout creates a GetOutboundCampaignsDivisionviewsGatewayTimeout with default headers values
func NewGetOutboundCampaignsDivisionviewsGatewayTimeout() *GetOutboundCampaignsDivisionviewsGatewayTimeout {
	return &GetOutboundCampaignsDivisionviewsGatewayTimeout{}
}

/*
GetOutboundCampaignsDivisionviewsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOutboundCampaignsDivisionviewsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaigns divisionviews gateway timeout response has a 2xx status code
func (o *GetOutboundCampaignsDivisionviewsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaigns divisionviews gateway timeout response has a 3xx status code
func (o *GetOutboundCampaignsDivisionviewsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaigns divisionviews gateway timeout response has a 4xx status code
func (o *GetOutboundCampaignsDivisionviewsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound campaigns divisionviews gateway timeout response has a 5xx status code
func (o *GetOutboundCampaignsDivisionviewsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound campaigns divisionviews gateway timeout response a status code equal to that given
func (o *GetOutboundCampaignsDivisionviewsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOutboundCampaignsDivisionviewsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/divisionviews][%d] getOutboundCampaignsDivisionviewsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCampaignsDivisionviewsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignsDivisionviewsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
