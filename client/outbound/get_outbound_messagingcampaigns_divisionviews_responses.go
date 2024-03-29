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

// GetOutboundMessagingcampaignsDivisionviewsReader is a Reader for the GetOutboundMessagingcampaignsDivisionviews structure.
type GetOutboundMessagingcampaignsDivisionviewsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundMessagingcampaignsDivisionviewsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundMessagingcampaignsDivisionviewsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundMessagingcampaignsDivisionviewsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundMessagingcampaignsDivisionviewsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundMessagingcampaignsDivisionviewsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundMessagingcampaignsDivisionviewsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundMessagingcampaignsDivisionviewsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundMessagingcampaignsDivisionviewsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundMessagingcampaignsDivisionviewsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundMessagingcampaignsDivisionviewsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundMessagingcampaignsDivisionviewsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundMessagingcampaignsDivisionviewsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundMessagingcampaignsDivisionviewsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundMessagingcampaignsDivisionviewsOK creates a GetOutboundMessagingcampaignsDivisionviewsOK with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewsOK() *GetOutboundMessagingcampaignsDivisionviewsOK {
	return &GetOutboundMessagingcampaignsDivisionviewsOK{}
}

/*
GetOutboundMessagingcampaignsDivisionviewsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOutboundMessagingcampaignsDivisionviewsOK struct {
	Payload *models.MessagingCampaignDivisionViewEntityListing
}

// IsSuccess returns true when this get outbound messagingcampaigns divisionviews o k response has a 2xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get outbound messagingcampaigns divisionviews o k response has a 3xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound messagingcampaigns divisionviews o k response has a 4xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound messagingcampaigns divisionviews o k response has a 5xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound messagingcampaigns divisionviews o k response a status code equal to that given
func (o *GetOutboundMessagingcampaignsDivisionviewsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOutboundMessagingcampaignsDivisionviewsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsOK) GetPayload() *models.MessagingCampaignDivisionViewEntityListing {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessagingCampaignDivisionViewEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewsBadRequest creates a GetOutboundMessagingcampaignsDivisionviewsBadRequest with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewsBadRequest() *GetOutboundMessagingcampaignsDivisionviewsBadRequest {
	return &GetOutboundMessagingcampaignsDivisionviewsBadRequest{}
}

/*
GetOutboundMessagingcampaignsDivisionviewsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundMessagingcampaignsDivisionviewsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound messagingcampaigns divisionviews bad request response has a 2xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound messagingcampaigns divisionviews bad request response has a 3xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound messagingcampaigns divisionviews bad request response has a 4xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound messagingcampaigns divisionviews bad request response has a 5xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound messagingcampaigns divisionviews bad request response a status code equal to that given
func (o *GetOutboundMessagingcampaignsDivisionviewsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOutboundMessagingcampaignsDivisionviewsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewsUnauthorized creates a GetOutboundMessagingcampaignsDivisionviewsUnauthorized with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewsUnauthorized() *GetOutboundMessagingcampaignsDivisionviewsUnauthorized {
	return &GetOutboundMessagingcampaignsDivisionviewsUnauthorized{}
}

/*
GetOutboundMessagingcampaignsDivisionviewsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundMessagingcampaignsDivisionviewsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound messagingcampaigns divisionviews unauthorized response has a 2xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound messagingcampaigns divisionviews unauthorized response has a 3xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound messagingcampaigns divisionviews unauthorized response has a 4xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound messagingcampaigns divisionviews unauthorized response has a 5xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound messagingcampaigns divisionviews unauthorized response a status code equal to that given
func (o *GetOutboundMessagingcampaignsDivisionviewsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOutboundMessagingcampaignsDivisionviewsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewsForbidden creates a GetOutboundMessagingcampaignsDivisionviewsForbidden with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewsForbidden() *GetOutboundMessagingcampaignsDivisionviewsForbidden {
	return &GetOutboundMessagingcampaignsDivisionviewsForbidden{}
}

/*
GetOutboundMessagingcampaignsDivisionviewsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundMessagingcampaignsDivisionviewsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound messagingcampaigns divisionviews forbidden response has a 2xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound messagingcampaigns divisionviews forbidden response has a 3xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound messagingcampaigns divisionviews forbidden response has a 4xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound messagingcampaigns divisionviews forbidden response has a 5xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound messagingcampaigns divisionviews forbidden response a status code equal to that given
func (o *GetOutboundMessagingcampaignsDivisionviewsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOutboundMessagingcampaignsDivisionviewsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewsNotFound creates a GetOutboundMessagingcampaignsDivisionviewsNotFound with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewsNotFound() *GetOutboundMessagingcampaignsDivisionviewsNotFound {
	return &GetOutboundMessagingcampaignsDivisionviewsNotFound{}
}

/*
GetOutboundMessagingcampaignsDivisionviewsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOutboundMessagingcampaignsDivisionviewsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound messagingcampaigns divisionviews not found response has a 2xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound messagingcampaigns divisionviews not found response has a 3xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound messagingcampaigns divisionviews not found response has a 4xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound messagingcampaigns divisionviews not found response has a 5xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound messagingcampaigns divisionviews not found response a status code equal to that given
func (o *GetOutboundMessagingcampaignsDivisionviewsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOutboundMessagingcampaignsDivisionviewsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewsRequestTimeout creates a GetOutboundMessagingcampaignsDivisionviewsRequestTimeout with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewsRequestTimeout() *GetOutboundMessagingcampaignsDivisionviewsRequestTimeout {
	return &GetOutboundMessagingcampaignsDivisionviewsRequestTimeout{}
}

/*
GetOutboundMessagingcampaignsDivisionviewsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundMessagingcampaignsDivisionviewsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound messagingcampaigns divisionviews request timeout response has a 2xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound messagingcampaigns divisionviews request timeout response has a 3xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound messagingcampaigns divisionviews request timeout response has a 4xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound messagingcampaigns divisionviews request timeout response has a 5xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound messagingcampaigns divisionviews request timeout response a status code equal to that given
func (o *GetOutboundMessagingcampaignsDivisionviewsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOutboundMessagingcampaignsDivisionviewsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewsRequestEntityTooLarge creates a GetOutboundMessagingcampaignsDivisionviewsRequestEntityTooLarge with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewsRequestEntityTooLarge() *GetOutboundMessagingcampaignsDivisionviewsRequestEntityTooLarge {
	return &GetOutboundMessagingcampaignsDivisionviewsRequestEntityTooLarge{}
}

/*
GetOutboundMessagingcampaignsDivisionviewsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetOutboundMessagingcampaignsDivisionviewsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound messagingcampaigns divisionviews request entity too large response has a 2xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound messagingcampaigns divisionviews request entity too large response has a 3xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound messagingcampaigns divisionviews request entity too large response has a 4xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound messagingcampaigns divisionviews request entity too large response has a 5xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound messagingcampaigns divisionviews request entity too large response a status code equal to that given
func (o *GetOutboundMessagingcampaignsDivisionviewsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOutboundMessagingcampaignsDivisionviewsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewsUnsupportedMediaType creates a GetOutboundMessagingcampaignsDivisionviewsUnsupportedMediaType with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewsUnsupportedMediaType() *GetOutboundMessagingcampaignsDivisionviewsUnsupportedMediaType {
	return &GetOutboundMessagingcampaignsDivisionviewsUnsupportedMediaType{}
}

/*
GetOutboundMessagingcampaignsDivisionviewsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundMessagingcampaignsDivisionviewsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound messagingcampaigns divisionviews unsupported media type response has a 2xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound messagingcampaigns divisionviews unsupported media type response has a 3xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound messagingcampaigns divisionviews unsupported media type response has a 4xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound messagingcampaigns divisionviews unsupported media type response has a 5xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound messagingcampaigns divisionviews unsupported media type response a status code equal to that given
func (o *GetOutboundMessagingcampaignsDivisionviewsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOutboundMessagingcampaignsDivisionviewsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewsTooManyRequests creates a GetOutboundMessagingcampaignsDivisionviewsTooManyRequests with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewsTooManyRequests() *GetOutboundMessagingcampaignsDivisionviewsTooManyRequests {
	return &GetOutboundMessagingcampaignsDivisionviewsTooManyRequests{}
}

/*
GetOutboundMessagingcampaignsDivisionviewsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundMessagingcampaignsDivisionviewsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound messagingcampaigns divisionviews too many requests response has a 2xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound messagingcampaigns divisionviews too many requests response has a 3xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound messagingcampaigns divisionviews too many requests response has a 4xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound messagingcampaigns divisionviews too many requests response has a 5xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound messagingcampaigns divisionviews too many requests response a status code equal to that given
func (o *GetOutboundMessagingcampaignsDivisionviewsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOutboundMessagingcampaignsDivisionviewsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewsInternalServerError creates a GetOutboundMessagingcampaignsDivisionviewsInternalServerError with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewsInternalServerError() *GetOutboundMessagingcampaignsDivisionviewsInternalServerError {
	return &GetOutboundMessagingcampaignsDivisionviewsInternalServerError{}
}

/*
GetOutboundMessagingcampaignsDivisionviewsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundMessagingcampaignsDivisionviewsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound messagingcampaigns divisionviews internal server error response has a 2xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound messagingcampaigns divisionviews internal server error response has a 3xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound messagingcampaigns divisionviews internal server error response has a 4xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound messagingcampaigns divisionviews internal server error response has a 5xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound messagingcampaigns divisionviews internal server error response a status code equal to that given
func (o *GetOutboundMessagingcampaignsDivisionviewsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOutboundMessagingcampaignsDivisionviewsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewsServiceUnavailable creates a GetOutboundMessagingcampaignsDivisionviewsServiceUnavailable with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewsServiceUnavailable() *GetOutboundMessagingcampaignsDivisionviewsServiceUnavailable {
	return &GetOutboundMessagingcampaignsDivisionviewsServiceUnavailable{}
}

/*
GetOutboundMessagingcampaignsDivisionviewsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundMessagingcampaignsDivisionviewsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound messagingcampaigns divisionviews service unavailable response has a 2xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound messagingcampaigns divisionviews service unavailable response has a 3xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound messagingcampaigns divisionviews service unavailable response has a 4xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound messagingcampaigns divisionviews service unavailable response has a 5xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound messagingcampaigns divisionviews service unavailable response a status code equal to that given
func (o *GetOutboundMessagingcampaignsDivisionviewsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOutboundMessagingcampaignsDivisionviewsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundMessagingcampaignsDivisionviewsGatewayTimeout creates a GetOutboundMessagingcampaignsDivisionviewsGatewayTimeout with default headers values
func NewGetOutboundMessagingcampaignsDivisionviewsGatewayTimeout() *GetOutboundMessagingcampaignsDivisionviewsGatewayTimeout {
	return &GetOutboundMessagingcampaignsDivisionviewsGatewayTimeout{}
}

/*
GetOutboundMessagingcampaignsDivisionviewsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOutboundMessagingcampaignsDivisionviewsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound messagingcampaigns divisionviews gateway timeout response has a 2xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound messagingcampaigns divisionviews gateway timeout response has a 3xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound messagingcampaigns divisionviews gateway timeout response has a 4xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound messagingcampaigns divisionviews gateway timeout response has a 5xx status code
func (o *GetOutboundMessagingcampaignsDivisionviewsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound messagingcampaigns divisionviews gateway timeout response a status code equal to that given
func (o *GetOutboundMessagingcampaignsDivisionviewsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOutboundMessagingcampaignsDivisionviewsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/messagingcampaigns/divisionviews][%d] getOutboundMessagingcampaignsDivisionviewsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundMessagingcampaignsDivisionviewsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundMessagingcampaignsDivisionviewsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
