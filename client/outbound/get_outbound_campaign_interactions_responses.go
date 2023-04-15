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

// GetOutboundCampaignInteractionsReader is a Reader for the GetOutboundCampaignInteractions structure.
type GetOutboundCampaignInteractionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundCampaignInteractionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundCampaignInteractionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundCampaignInteractionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundCampaignInteractionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundCampaignInteractionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundCampaignInteractionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundCampaignInteractionsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundCampaignInteractionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundCampaignInteractionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundCampaignInteractionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundCampaignInteractionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundCampaignInteractionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundCampaignInteractionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundCampaignInteractionsOK creates a GetOutboundCampaignInteractionsOK with default headers values
func NewGetOutboundCampaignInteractionsOK() *GetOutboundCampaignInteractionsOK {
	return &GetOutboundCampaignInteractionsOK{}
}

/*
GetOutboundCampaignInteractionsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOutboundCampaignInteractionsOK struct {
	Payload *models.CampaignInteractions
}

// IsSuccess returns true when this get outbound campaign interactions o k response has a 2xx status code
func (o *GetOutboundCampaignInteractionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get outbound campaign interactions o k response has a 3xx status code
func (o *GetOutboundCampaignInteractionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign interactions o k response has a 4xx status code
func (o *GetOutboundCampaignInteractionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound campaign interactions o k response has a 5xx status code
func (o *GetOutboundCampaignInteractionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaign interactions o k response a status code equal to that given
func (o *GetOutboundCampaignInteractionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOutboundCampaignInteractionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCampaignInteractionsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCampaignInteractionsOK) GetPayload() *models.CampaignInteractions {
	return o.Payload
}

func (o *GetOutboundCampaignInteractionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CampaignInteractions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignInteractionsBadRequest creates a GetOutboundCampaignInteractionsBadRequest with default headers values
func NewGetOutboundCampaignInteractionsBadRequest() *GetOutboundCampaignInteractionsBadRequest {
	return &GetOutboundCampaignInteractionsBadRequest{}
}

/*
GetOutboundCampaignInteractionsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundCampaignInteractionsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign interactions bad request response has a 2xx status code
func (o *GetOutboundCampaignInteractionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign interactions bad request response has a 3xx status code
func (o *GetOutboundCampaignInteractionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign interactions bad request response has a 4xx status code
func (o *GetOutboundCampaignInteractionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaign interactions bad request response has a 5xx status code
func (o *GetOutboundCampaignInteractionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaign interactions bad request response a status code equal to that given
func (o *GetOutboundCampaignInteractionsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOutboundCampaignInteractionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCampaignInteractionsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCampaignInteractionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignInteractionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignInteractionsUnauthorized creates a GetOutboundCampaignInteractionsUnauthorized with default headers values
func NewGetOutboundCampaignInteractionsUnauthorized() *GetOutboundCampaignInteractionsUnauthorized {
	return &GetOutboundCampaignInteractionsUnauthorized{}
}

/*
GetOutboundCampaignInteractionsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundCampaignInteractionsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign interactions unauthorized response has a 2xx status code
func (o *GetOutboundCampaignInteractionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign interactions unauthorized response has a 3xx status code
func (o *GetOutboundCampaignInteractionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign interactions unauthorized response has a 4xx status code
func (o *GetOutboundCampaignInteractionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaign interactions unauthorized response has a 5xx status code
func (o *GetOutboundCampaignInteractionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaign interactions unauthorized response a status code equal to that given
func (o *GetOutboundCampaignInteractionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOutboundCampaignInteractionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCampaignInteractionsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCampaignInteractionsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignInteractionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignInteractionsForbidden creates a GetOutboundCampaignInteractionsForbidden with default headers values
func NewGetOutboundCampaignInteractionsForbidden() *GetOutboundCampaignInteractionsForbidden {
	return &GetOutboundCampaignInteractionsForbidden{}
}

/*
GetOutboundCampaignInteractionsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundCampaignInteractionsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign interactions forbidden response has a 2xx status code
func (o *GetOutboundCampaignInteractionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign interactions forbidden response has a 3xx status code
func (o *GetOutboundCampaignInteractionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign interactions forbidden response has a 4xx status code
func (o *GetOutboundCampaignInteractionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaign interactions forbidden response has a 5xx status code
func (o *GetOutboundCampaignInteractionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaign interactions forbidden response a status code equal to that given
func (o *GetOutboundCampaignInteractionsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOutboundCampaignInteractionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCampaignInteractionsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCampaignInteractionsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignInteractionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignInteractionsNotFound creates a GetOutboundCampaignInteractionsNotFound with default headers values
func NewGetOutboundCampaignInteractionsNotFound() *GetOutboundCampaignInteractionsNotFound {
	return &GetOutboundCampaignInteractionsNotFound{}
}

/*
GetOutboundCampaignInteractionsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOutboundCampaignInteractionsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign interactions not found response has a 2xx status code
func (o *GetOutboundCampaignInteractionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign interactions not found response has a 3xx status code
func (o *GetOutboundCampaignInteractionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign interactions not found response has a 4xx status code
func (o *GetOutboundCampaignInteractionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaign interactions not found response has a 5xx status code
func (o *GetOutboundCampaignInteractionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaign interactions not found response a status code equal to that given
func (o *GetOutboundCampaignInteractionsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOutboundCampaignInteractionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCampaignInteractionsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCampaignInteractionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignInteractionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignInteractionsRequestTimeout creates a GetOutboundCampaignInteractionsRequestTimeout with default headers values
func NewGetOutboundCampaignInteractionsRequestTimeout() *GetOutboundCampaignInteractionsRequestTimeout {
	return &GetOutboundCampaignInteractionsRequestTimeout{}
}

/*
GetOutboundCampaignInteractionsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundCampaignInteractionsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign interactions request timeout response has a 2xx status code
func (o *GetOutboundCampaignInteractionsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign interactions request timeout response has a 3xx status code
func (o *GetOutboundCampaignInteractionsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign interactions request timeout response has a 4xx status code
func (o *GetOutboundCampaignInteractionsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaign interactions request timeout response has a 5xx status code
func (o *GetOutboundCampaignInteractionsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaign interactions request timeout response a status code equal to that given
func (o *GetOutboundCampaignInteractionsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOutboundCampaignInteractionsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundCampaignInteractionsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundCampaignInteractionsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignInteractionsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignInteractionsRequestEntityTooLarge creates a GetOutboundCampaignInteractionsRequestEntityTooLarge with default headers values
func NewGetOutboundCampaignInteractionsRequestEntityTooLarge() *GetOutboundCampaignInteractionsRequestEntityTooLarge {
	return &GetOutboundCampaignInteractionsRequestEntityTooLarge{}
}

/*
GetOutboundCampaignInteractionsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetOutboundCampaignInteractionsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign interactions request entity too large response has a 2xx status code
func (o *GetOutboundCampaignInteractionsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign interactions request entity too large response has a 3xx status code
func (o *GetOutboundCampaignInteractionsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign interactions request entity too large response has a 4xx status code
func (o *GetOutboundCampaignInteractionsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaign interactions request entity too large response has a 5xx status code
func (o *GetOutboundCampaignInteractionsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaign interactions request entity too large response a status code equal to that given
func (o *GetOutboundCampaignInteractionsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOutboundCampaignInteractionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCampaignInteractionsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCampaignInteractionsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignInteractionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignInteractionsUnsupportedMediaType creates a GetOutboundCampaignInteractionsUnsupportedMediaType with default headers values
func NewGetOutboundCampaignInteractionsUnsupportedMediaType() *GetOutboundCampaignInteractionsUnsupportedMediaType {
	return &GetOutboundCampaignInteractionsUnsupportedMediaType{}
}

/*
GetOutboundCampaignInteractionsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundCampaignInteractionsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign interactions unsupported media type response has a 2xx status code
func (o *GetOutboundCampaignInteractionsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign interactions unsupported media type response has a 3xx status code
func (o *GetOutboundCampaignInteractionsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign interactions unsupported media type response has a 4xx status code
func (o *GetOutboundCampaignInteractionsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaign interactions unsupported media type response has a 5xx status code
func (o *GetOutboundCampaignInteractionsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaign interactions unsupported media type response a status code equal to that given
func (o *GetOutboundCampaignInteractionsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOutboundCampaignInteractionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCampaignInteractionsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCampaignInteractionsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignInteractionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignInteractionsTooManyRequests creates a GetOutboundCampaignInteractionsTooManyRequests with default headers values
func NewGetOutboundCampaignInteractionsTooManyRequests() *GetOutboundCampaignInteractionsTooManyRequests {
	return &GetOutboundCampaignInteractionsTooManyRequests{}
}

/*
GetOutboundCampaignInteractionsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundCampaignInteractionsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign interactions too many requests response has a 2xx status code
func (o *GetOutboundCampaignInteractionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign interactions too many requests response has a 3xx status code
func (o *GetOutboundCampaignInteractionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign interactions too many requests response has a 4xx status code
func (o *GetOutboundCampaignInteractionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaign interactions too many requests response has a 5xx status code
func (o *GetOutboundCampaignInteractionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaign interactions too many requests response a status code equal to that given
func (o *GetOutboundCampaignInteractionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOutboundCampaignInteractionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCampaignInteractionsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCampaignInteractionsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignInteractionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignInteractionsInternalServerError creates a GetOutboundCampaignInteractionsInternalServerError with default headers values
func NewGetOutboundCampaignInteractionsInternalServerError() *GetOutboundCampaignInteractionsInternalServerError {
	return &GetOutboundCampaignInteractionsInternalServerError{}
}

/*
GetOutboundCampaignInteractionsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundCampaignInteractionsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign interactions internal server error response has a 2xx status code
func (o *GetOutboundCampaignInteractionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign interactions internal server error response has a 3xx status code
func (o *GetOutboundCampaignInteractionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign interactions internal server error response has a 4xx status code
func (o *GetOutboundCampaignInteractionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound campaign interactions internal server error response has a 5xx status code
func (o *GetOutboundCampaignInteractionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound campaign interactions internal server error response a status code equal to that given
func (o *GetOutboundCampaignInteractionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOutboundCampaignInteractionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCampaignInteractionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCampaignInteractionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignInteractionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignInteractionsServiceUnavailable creates a GetOutboundCampaignInteractionsServiceUnavailable with default headers values
func NewGetOutboundCampaignInteractionsServiceUnavailable() *GetOutboundCampaignInteractionsServiceUnavailable {
	return &GetOutboundCampaignInteractionsServiceUnavailable{}
}

/*
GetOutboundCampaignInteractionsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundCampaignInteractionsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign interactions service unavailable response has a 2xx status code
func (o *GetOutboundCampaignInteractionsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign interactions service unavailable response has a 3xx status code
func (o *GetOutboundCampaignInteractionsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign interactions service unavailable response has a 4xx status code
func (o *GetOutboundCampaignInteractionsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound campaign interactions service unavailable response has a 5xx status code
func (o *GetOutboundCampaignInteractionsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound campaign interactions service unavailable response a status code equal to that given
func (o *GetOutboundCampaignInteractionsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOutboundCampaignInteractionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCampaignInteractionsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCampaignInteractionsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignInteractionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignInteractionsGatewayTimeout creates a GetOutboundCampaignInteractionsGatewayTimeout with default headers values
func NewGetOutboundCampaignInteractionsGatewayTimeout() *GetOutboundCampaignInteractionsGatewayTimeout {
	return &GetOutboundCampaignInteractionsGatewayTimeout{}
}

/*
GetOutboundCampaignInteractionsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOutboundCampaignInteractionsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign interactions gateway timeout response has a 2xx status code
func (o *GetOutboundCampaignInteractionsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign interactions gateway timeout response has a 3xx status code
func (o *GetOutboundCampaignInteractionsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign interactions gateway timeout response has a 4xx status code
func (o *GetOutboundCampaignInteractionsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound campaign interactions gateway timeout response has a 5xx status code
func (o *GetOutboundCampaignInteractionsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound campaign interactions gateway timeout response a status code equal to that given
func (o *GetOutboundCampaignInteractionsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOutboundCampaignInteractionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCampaignInteractionsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}/interactions][%d] getOutboundCampaignInteractionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCampaignInteractionsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignInteractionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
