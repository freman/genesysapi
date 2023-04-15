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

// GetOutboundCampaignReader is a Reader for the GetOutboundCampaign structure.
type GetOutboundCampaignReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundCampaignReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundCampaignOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundCampaignBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundCampaignUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundCampaignForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundCampaignNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundCampaignRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundCampaignRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundCampaignUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundCampaignTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundCampaignInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundCampaignServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundCampaignGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundCampaignOK creates a GetOutboundCampaignOK with default headers values
func NewGetOutboundCampaignOK() *GetOutboundCampaignOK {
	return &GetOutboundCampaignOK{}
}

/*
GetOutboundCampaignOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOutboundCampaignOK struct {
	Payload *models.Campaign
}

// IsSuccess returns true when this get outbound campaign o k response has a 2xx status code
func (o *GetOutboundCampaignOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get outbound campaign o k response has a 3xx status code
func (o *GetOutboundCampaignOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign o k response has a 4xx status code
func (o *GetOutboundCampaignOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound campaign o k response has a 5xx status code
func (o *GetOutboundCampaignOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaign o k response a status code equal to that given
func (o *GetOutboundCampaignOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOutboundCampaignOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCampaignOK) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCampaignOK) GetPayload() *models.Campaign {
	return o.Payload
}

func (o *GetOutboundCampaignOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Campaign)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignBadRequest creates a GetOutboundCampaignBadRequest with default headers values
func NewGetOutboundCampaignBadRequest() *GetOutboundCampaignBadRequest {
	return &GetOutboundCampaignBadRequest{}
}

/*
GetOutboundCampaignBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundCampaignBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign bad request response has a 2xx status code
func (o *GetOutboundCampaignBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign bad request response has a 3xx status code
func (o *GetOutboundCampaignBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign bad request response has a 4xx status code
func (o *GetOutboundCampaignBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaign bad request response has a 5xx status code
func (o *GetOutboundCampaignBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaign bad request response a status code equal to that given
func (o *GetOutboundCampaignBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOutboundCampaignBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCampaignBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCampaignBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignUnauthorized creates a GetOutboundCampaignUnauthorized with default headers values
func NewGetOutboundCampaignUnauthorized() *GetOutboundCampaignUnauthorized {
	return &GetOutboundCampaignUnauthorized{}
}

/*
GetOutboundCampaignUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundCampaignUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign unauthorized response has a 2xx status code
func (o *GetOutboundCampaignUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign unauthorized response has a 3xx status code
func (o *GetOutboundCampaignUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign unauthorized response has a 4xx status code
func (o *GetOutboundCampaignUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaign unauthorized response has a 5xx status code
func (o *GetOutboundCampaignUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaign unauthorized response a status code equal to that given
func (o *GetOutboundCampaignUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOutboundCampaignUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCampaignUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCampaignUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignForbidden creates a GetOutboundCampaignForbidden with default headers values
func NewGetOutboundCampaignForbidden() *GetOutboundCampaignForbidden {
	return &GetOutboundCampaignForbidden{}
}

/*
GetOutboundCampaignForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundCampaignForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign forbidden response has a 2xx status code
func (o *GetOutboundCampaignForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign forbidden response has a 3xx status code
func (o *GetOutboundCampaignForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign forbidden response has a 4xx status code
func (o *GetOutboundCampaignForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaign forbidden response has a 5xx status code
func (o *GetOutboundCampaignForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaign forbidden response a status code equal to that given
func (o *GetOutboundCampaignForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOutboundCampaignForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCampaignForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCampaignForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignNotFound creates a GetOutboundCampaignNotFound with default headers values
func NewGetOutboundCampaignNotFound() *GetOutboundCampaignNotFound {
	return &GetOutboundCampaignNotFound{}
}

/*
GetOutboundCampaignNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOutboundCampaignNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign not found response has a 2xx status code
func (o *GetOutboundCampaignNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign not found response has a 3xx status code
func (o *GetOutboundCampaignNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign not found response has a 4xx status code
func (o *GetOutboundCampaignNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaign not found response has a 5xx status code
func (o *GetOutboundCampaignNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaign not found response a status code equal to that given
func (o *GetOutboundCampaignNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOutboundCampaignNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCampaignNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCampaignNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignRequestTimeout creates a GetOutboundCampaignRequestTimeout with default headers values
func NewGetOutboundCampaignRequestTimeout() *GetOutboundCampaignRequestTimeout {
	return &GetOutboundCampaignRequestTimeout{}
}

/*
GetOutboundCampaignRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundCampaignRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign request timeout response has a 2xx status code
func (o *GetOutboundCampaignRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign request timeout response has a 3xx status code
func (o *GetOutboundCampaignRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign request timeout response has a 4xx status code
func (o *GetOutboundCampaignRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaign request timeout response has a 5xx status code
func (o *GetOutboundCampaignRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaign request timeout response a status code equal to that given
func (o *GetOutboundCampaignRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOutboundCampaignRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundCampaignRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundCampaignRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignRequestEntityTooLarge creates a GetOutboundCampaignRequestEntityTooLarge with default headers values
func NewGetOutboundCampaignRequestEntityTooLarge() *GetOutboundCampaignRequestEntityTooLarge {
	return &GetOutboundCampaignRequestEntityTooLarge{}
}

/*
GetOutboundCampaignRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetOutboundCampaignRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign request entity too large response has a 2xx status code
func (o *GetOutboundCampaignRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign request entity too large response has a 3xx status code
func (o *GetOutboundCampaignRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign request entity too large response has a 4xx status code
func (o *GetOutboundCampaignRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaign request entity too large response has a 5xx status code
func (o *GetOutboundCampaignRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaign request entity too large response a status code equal to that given
func (o *GetOutboundCampaignRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOutboundCampaignRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCampaignRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCampaignRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignUnsupportedMediaType creates a GetOutboundCampaignUnsupportedMediaType with default headers values
func NewGetOutboundCampaignUnsupportedMediaType() *GetOutboundCampaignUnsupportedMediaType {
	return &GetOutboundCampaignUnsupportedMediaType{}
}

/*
GetOutboundCampaignUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundCampaignUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign unsupported media type response has a 2xx status code
func (o *GetOutboundCampaignUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign unsupported media type response has a 3xx status code
func (o *GetOutboundCampaignUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign unsupported media type response has a 4xx status code
func (o *GetOutboundCampaignUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaign unsupported media type response has a 5xx status code
func (o *GetOutboundCampaignUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaign unsupported media type response a status code equal to that given
func (o *GetOutboundCampaignUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOutboundCampaignUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCampaignUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCampaignUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignTooManyRequests creates a GetOutboundCampaignTooManyRequests with default headers values
func NewGetOutboundCampaignTooManyRequests() *GetOutboundCampaignTooManyRequests {
	return &GetOutboundCampaignTooManyRequests{}
}

/*
GetOutboundCampaignTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundCampaignTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign too many requests response has a 2xx status code
func (o *GetOutboundCampaignTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign too many requests response has a 3xx status code
func (o *GetOutboundCampaignTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign too many requests response has a 4xx status code
func (o *GetOutboundCampaignTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaign too many requests response has a 5xx status code
func (o *GetOutboundCampaignTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaign too many requests response a status code equal to that given
func (o *GetOutboundCampaignTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOutboundCampaignTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCampaignTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCampaignTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignInternalServerError creates a GetOutboundCampaignInternalServerError with default headers values
func NewGetOutboundCampaignInternalServerError() *GetOutboundCampaignInternalServerError {
	return &GetOutboundCampaignInternalServerError{}
}

/*
GetOutboundCampaignInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundCampaignInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign internal server error response has a 2xx status code
func (o *GetOutboundCampaignInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign internal server error response has a 3xx status code
func (o *GetOutboundCampaignInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign internal server error response has a 4xx status code
func (o *GetOutboundCampaignInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound campaign internal server error response has a 5xx status code
func (o *GetOutboundCampaignInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound campaign internal server error response a status code equal to that given
func (o *GetOutboundCampaignInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOutboundCampaignInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCampaignInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCampaignInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignServiceUnavailable creates a GetOutboundCampaignServiceUnavailable with default headers values
func NewGetOutboundCampaignServiceUnavailable() *GetOutboundCampaignServiceUnavailable {
	return &GetOutboundCampaignServiceUnavailable{}
}

/*
GetOutboundCampaignServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundCampaignServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign service unavailable response has a 2xx status code
func (o *GetOutboundCampaignServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign service unavailable response has a 3xx status code
func (o *GetOutboundCampaignServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign service unavailable response has a 4xx status code
func (o *GetOutboundCampaignServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound campaign service unavailable response has a 5xx status code
func (o *GetOutboundCampaignServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound campaign service unavailable response a status code equal to that given
func (o *GetOutboundCampaignServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOutboundCampaignServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCampaignServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCampaignServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignGatewayTimeout creates a GetOutboundCampaignGatewayTimeout with default headers values
func NewGetOutboundCampaignGatewayTimeout() *GetOutboundCampaignGatewayTimeout {
	return &GetOutboundCampaignGatewayTimeout{}
}

/*
GetOutboundCampaignGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOutboundCampaignGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaign gateway timeout response has a 2xx status code
func (o *GetOutboundCampaignGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaign gateway timeout response has a 3xx status code
func (o *GetOutboundCampaignGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaign gateway timeout response has a 4xx status code
func (o *GetOutboundCampaignGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound campaign gateway timeout response has a 5xx status code
func (o *GetOutboundCampaignGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound campaign gateway timeout response a status code equal to that given
func (o *GetOutboundCampaignGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOutboundCampaignGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCampaignGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaigns/{campaignId}][%d] getOutboundCampaignGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCampaignGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
