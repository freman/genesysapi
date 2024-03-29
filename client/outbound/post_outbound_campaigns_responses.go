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

// PostOutboundCampaignsReader is a Reader for the PostOutboundCampaigns structure.
type PostOutboundCampaignsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOutboundCampaignsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOutboundCampaignsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOutboundCampaignsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOutboundCampaignsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOutboundCampaignsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOutboundCampaignsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostOutboundCampaignsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOutboundCampaignsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOutboundCampaignsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOutboundCampaignsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOutboundCampaignsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOutboundCampaignsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOutboundCampaignsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOutboundCampaignsOK creates a PostOutboundCampaignsOK with default headers values
func NewPostOutboundCampaignsOK() *PostOutboundCampaignsOK {
	return &PostOutboundCampaignsOK{}
}

/*
PostOutboundCampaignsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostOutboundCampaignsOK struct {
	Payload *models.Campaign
}

// IsSuccess returns true when this post outbound campaigns o k response has a 2xx status code
func (o *PostOutboundCampaignsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post outbound campaigns o k response has a 3xx status code
func (o *PostOutboundCampaignsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound campaigns o k response has a 4xx status code
func (o *PostOutboundCampaignsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound campaigns o k response has a 5xx status code
func (o *PostOutboundCampaignsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound campaigns o k response a status code equal to that given
func (o *PostOutboundCampaignsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostOutboundCampaignsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsOK  %+v", 200, o.Payload)
}

func (o *PostOutboundCampaignsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsOK  %+v", 200, o.Payload)
}

func (o *PostOutboundCampaignsOK) GetPayload() *models.Campaign {
	return o.Payload
}

func (o *PostOutboundCampaignsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Campaign)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignsBadRequest creates a PostOutboundCampaignsBadRequest with default headers values
func NewPostOutboundCampaignsBadRequest() *PostOutboundCampaignsBadRequest {
	return &PostOutboundCampaignsBadRequest{}
}

/*
PostOutboundCampaignsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOutboundCampaignsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound campaigns bad request response has a 2xx status code
func (o *PostOutboundCampaignsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound campaigns bad request response has a 3xx status code
func (o *PostOutboundCampaignsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound campaigns bad request response has a 4xx status code
func (o *PostOutboundCampaignsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound campaigns bad request response has a 5xx status code
func (o *PostOutboundCampaignsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound campaigns bad request response a status code equal to that given
func (o *PostOutboundCampaignsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostOutboundCampaignsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsBadRequest  %+v", 400, o.Payload)
}

func (o *PostOutboundCampaignsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsBadRequest  %+v", 400, o.Payload)
}

func (o *PostOutboundCampaignsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignsUnauthorized creates a PostOutboundCampaignsUnauthorized with default headers values
func NewPostOutboundCampaignsUnauthorized() *PostOutboundCampaignsUnauthorized {
	return &PostOutboundCampaignsUnauthorized{}
}

/*
PostOutboundCampaignsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOutboundCampaignsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound campaigns unauthorized response has a 2xx status code
func (o *PostOutboundCampaignsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound campaigns unauthorized response has a 3xx status code
func (o *PostOutboundCampaignsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound campaigns unauthorized response has a 4xx status code
func (o *PostOutboundCampaignsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound campaigns unauthorized response has a 5xx status code
func (o *PostOutboundCampaignsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound campaigns unauthorized response a status code equal to that given
func (o *PostOutboundCampaignsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostOutboundCampaignsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOutboundCampaignsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOutboundCampaignsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignsForbidden creates a PostOutboundCampaignsForbidden with default headers values
func NewPostOutboundCampaignsForbidden() *PostOutboundCampaignsForbidden {
	return &PostOutboundCampaignsForbidden{}
}

/*
PostOutboundCampaignsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostOutboundCampaignsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound campaigns forbidden response has a 2xx status code
func (o *PostOutboundCampaignsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound campaigns forbidden response has a 3xx status code
func (o *PostOutboundCampaignsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound campaigns forbidden response has a 4xx status code
func (o *PostOutboundCampaignsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound campaigns forbidden response has a 5xx status code
func (o *PostOutboundCampaignsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound campaigns forbidden response a status code equal to that given
func (o *PostOutboundCampaignsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostOutboundCampaignsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsForbidden  %+v", 403, o.Payload)
}

func (o *PostOutboundCampaignsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsForbidden  %+v", 403, o.Payload)
}

func (o *PostOutboundCampaignsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignsNotFound creates a PostOutboundCampaignsNotFound with default headers values
func NewPostOutboundCampaignsNotFound() *PostOutboundCampaignsNotFound {
	return &PostOutboundCampaignsNotFound{}
}

/*
PostOutboundCampaignsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostOutboundCampaignsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound campaigns not found response has a 2xx status code
func (o *PostOutboundCampaignsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound campaigns not found response has a 3xx status code
func (o *PostOutboundCampaignsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound campaigns not found response has a 4xx status code
func (o *PostOutboundCampaignsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound campaigns not found response has a 5xx status code
func (o *PostOutboundCampaignsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound campaigns not found response a status code equal to that given
func (o *PostOutboundCampaignsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostOutboundCampaignsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsNotFound  %+v", 404, o.Payload)
}

func (o *PostOutboundCampaignsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsNotFound  %+v", 404, o.Payload)
}

func (o *PostOutboundCampaignsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignsRequestTimeout creates a PostOutboundCampaignsRequestTimeout with default headers values
func NewPostOutboundCampaignsRequestTimeout() *PostOutboundCampaignsRequestTimeout {
	return &PostOutboundCampaignsRequestTimeout{}
}

/*
PostOutboundCampaignsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostOutboundCampaignsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound campaigns request timeout response has a 2xx status code
func (o *PostOutboundCampaignsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound campaigns request timeout response has a 3xx status code
func (o *PostOutboundCampaignsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound campaigns request timeout response has a 4xx status code
func (o *PostOutboundCampaignsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound campaigns request timeout response has a 5xx status code
func (o *PostOutboundCampaignsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound campaigns request timeout response a status code equal to that given
func (o *PostOutboundCampaignsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostOutboundCampaignsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOutboundCampaignsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOutboundCampaignsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignsRequestEntityTooLarge creates a PostOutboundCampaignsRequestEntityTooLarge with default headers values
func NewPostOutboundCampaignsRequestEntityTooLarge() *PostOutboundCampaignsRequestEntityTooLarge {
	return &PostOutboundCampaignsRequestEntityTooLarge{}
}

/*
PostOutboundCampaignsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostOutboundCampaignsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound campaigns request entity too large response has a 2xx status code
func (o *PostOutboundCampaignsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound campaigns request entity too large response has a 3xx status code
func (o *PostOutboundCampaignsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound campaigns request entity too large response has a 4xx status code
func (o *PostOutboundCampaignsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound campaigns request entity too large response has a 5xx status code
func (o *PostOutboundCampaignsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound campaigns request entity too large response a status code equal to that given
func (o *PostOutboundCampaignsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostOutboundCampaignsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOutboundCampaignsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOutboundCampaignsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignsUnsupportedMediaType creates a PostOutboundCampaignsUnsupportedMediaType with default headers values
func NewPostOutboundCampaignsUnsupportedMediaType() *PostOutboundCampaignsUnsupportedMediaType {
	return &PostOutboundCampaignsUnsupportedMediaType{}
}

/*
PostOutboundCampaignsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOutboundCampaignsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound campaigns unsupported media type response has a 2xx status code
func (o *PostOutboundCampaignsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound campaigns unsupported media type response has a 3xx status code
func (o *PostOutboundCampaignsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound campaigns unsupported media type response has a 4xx status code
func (o *PostOutboundCampaignsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound campaigns unsupported media type response has a 5xx status code
func (o *PostOutboundCampaignsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound campaigns unsupported media type response a status code equal to that given
func (o *PostOutboundCampaignsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostOutboundCampaignsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOutboundCampaignsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOutboundCampaignsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignsTooManyRequests creates a PostOutboundCampaignsTooManyRequests with default headers values
func NewPostOutboundCampaignsTooManyRequests() *PostOutboundCampaignsTooManyRequests {
	return &PostOutboundCampaignsTooManyRequests{}
}

/*
PostOutboundCampaignsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostOutboundCampaignsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound campaigns too many requests response has a 2xx status code
func (o *PostOutboundCampaignsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound campaigns too many requests response has a 3xx status code
func (o *PostOutboundCampaignsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound campaigns too many requests response has a 4xx status code
func (o *PostOutboundCampaignsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound campaigns too many requests response has a 5xx status code
func (o *PostOutboundCampaignsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound campaigns too many requests response a status code equal to that given
func (o *PostOutboundCampaignsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostOutboundCampaignsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOutboundCampaignsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOutboundCampaignsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignsInternalServerError creates a PostOutboundCampaignsInternalServerError with default headers values
func NewPostOutboundCampaignsInternalServerError() *PostOutboundCampaignsInternalServerError {
	return &PostOutboundCampaignsInternalServerError{}
}

/*
PostOutboundCampaignsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOutboundCampaignsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound campaigns internal server error response has a 2xx status code
func (o *PostOutboundCampaignsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound campaigns internal server error response has a 3xx status code
func (o *PostOutboundCampaignsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound campaigns internal server error response has a 4xx status code
func (o *PostOutboundCampaignsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound campaigns internal server error response has a 5xx status code
func (o *PostOutboundCampaignsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post outbound campaigns internal server error response a status code equal to that given
func (o *PostOutboundCampaignsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostOutboundCampaignsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOutboundCampaignsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOutboundCampaignsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignsServiceUnavailable creates a PostOutboundCampaignsServiceUnavailable with default headers values
func NewPostOutboundCampaignsServiceUnavailable() *PostOutboundCampaignsServiceUnavailable {
	return &PostOutboundCampaignsServiceUnavailable{}
}

/*
PostOutboundCampaignsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOutboundCampaignsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound campaigns service unavailable response has a 2xx status code
func (o *PostOutboundCampaignsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound campaigns service unavailable response has a 3xx status code
func (o *PostOutboundCampaignsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound campaigns service unavailable response has a 4xx status code
func (o *PostOutboundCampaignsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound campaigns service unavailable response has a 5xx status code
func (o *PostOutboundCampaignsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post outbound campaigns service unavailable response a status code equal to that given
func (o *PostOutboundCampaignsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostOutboundCampaignsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOutboundCampaignsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOutboundCampaignsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCampaignsGatewayTimeout creates a PostOutboundCampaignsGatewayTimeout with default headers values
func NewPostOutboundCampaignsGatewayTimeout() *PostOutboundCampaignsGatewayTimeout {
	return &PostOutboundCampaignsGatewayTimeout{}
}

/*
PostOutboundCampaignsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostOutboundCampaignsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound campaigns gateway timeout response has a 2xx status code
func (o *PostOutboundCampaignsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound campaigns gateway timeout response has a 3xx status code
func (o *PostOutboundCampaignsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound campaigns gateway timeout response has a 4xx status code
func (o *PostOutboundCampaignsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound campaigns gateway timeout response has a 5xx status code
func (o *PostOutboundCampaignsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post outbound campaigns gateway timeout response a status code equal to that given
func (o *PostOutboundCampaignsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostOutboundCampaignsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOutboundCampaignsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/campaigns][%d] postOutboundCampaignsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOutboundCampaignsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCampaignsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
