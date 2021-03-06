// Code generated by go-swagger; DO NOT EDIT.

package o_auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetOauthClientUsageQueryResultReader is a Reader for the GetOauthClientUsageQueryResult structure.
type GetOauthClientUsageQueryResultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOauthClientUsageQueryResultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOauthClientUsageQueryResultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOauthClientUsageQueryResultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOauthClientUsageQueryResultUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOauthClientUsageQueryResultForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOauthClientUsageQueryResultNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOauthClientUsageQueryResultRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOauthClientUsageQueryResultUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOauthClientUsageQueryResultTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOauthClientUsageQueryResultInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOauthClientUsageQueryResultServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOauthClientUsageQueryResultGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOauthClientUsageQueryResultOK creates a GetOauthClientUsageQueryResultOK with default headers values
func NewGetOauthClientUsageQueryResultOK() *GetOauthClientUsageQueryResultOK {
	return &GetOauthClientUsageQueryResultOK{}
}

/*GetOauthClientUsageQueryResultOK handles this case with default header values.

successful operation
*/
type GetOauthClientUsageQueryResultOK struct {
	Payload *models.APIUsageQueryResult
}

func (o *GetOauthClientUsageQueryResultOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients/{clientId}/usage/query/results/{executionId}][%d] getOauthClientUsageQueryResultOK  %+v", 200, o.Payload)
}

func (o *GetOauthClientUsageQueryResultOK) GetPayload() *models.APIUsageQueryResult {
	return o.Payload
}

func (o *GetOauthClientUsageQueryResultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIUsageQueryResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthClientUsageQueryResultBadRequest creates a GetOauthClientUsageQueryResultBadRequest with default headers values
func NewGetOauthClientUsageQueryResultBadRequest() *GetOauthClientUsageQueryResultBadRequest {
	return &GetOauthClientUsageQueryResultBadRequest{}
}

/*GetOauthClientUsageQueryResultBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOauthClientUsageQueryResultBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOauthClientUsageQueryResultBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients/{clientId}/usage/query/results/{executionId}][%d] getOauthClientUsageQueryResultBadRequest  %+v", 400, o.Payload)
}

func (o *GetOauthClientUsageQueryResultBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthClientUsageQueryResultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthClientUsageQueryResultUnauthorized creates a GetOauthClientUsageQueryResultUnauthorized with default headers values
func NewGetOauthClientUsageQueryResultUnauthorized() *GetOauthClientUsageQueryResultUnauthorized {
	return &GetOauthClientUsageQueryResultUnauthorized{}
}

/*GetOauthClientUsageQueryResultUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOauthClientUsageQueryResultUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOauthClientUsageQueryResultUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients/{clientId}/usage/query/results/{executionId}][%d] getOauthClientUsageQueryResultUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOauthClientUsageQueryResultUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthClientUsageQueryResultUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthClientUsageQueryResultForbidden creates a GetOauthClientUsageQueryResultForbidden with default headers values
func NewGetOauthClientUsageQueryResultForbidden() *GetOauthClientUsageQueryResultForbidden {
	return &GetOauthClientUsageQueryResultForbidden{}
}

/*GetOauthClientUsageQueryResultForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOauthClientUsageQueryResultForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOauthClientUsageQueryResultForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients/{clientId}/usage/query/results/{executionId}][%d] getOauthClientUsageQueryResultForbidden  %+v", 403, o.Payload)
}

func (o *GetOauthClientUsageQueryResultForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthClientUsageQueryResultForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthClientUsageQueryResultNotFound creates a GetOauthClientUsageQueryResultNotFound with default headers values
func NewGetOauthClientUsageQueryResultNotFound() *GetOauthClientUsageQueryResultNotFound {
	return &GetOauthClientUsageQueryResultNotFound{}
}

/*GetOauthClientUsageQueryResultNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOauthClientUsageQueryResultNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOauthClientUsageQueryResultNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients/{clientId}/usage/query/results/{executionId}][%d] getOauthClientUsageQueryResultNotFound  %+v", 404, o.Payload)
}

func (o *GetOauthClientUsageQueryResultNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthClientUsageQueryResultNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthClientUsageQueryResultRequestEntityTooLarge creates a GetOauthClientUsageQueryResultRequestEntityTooLarge with default headers values
func NewGetOauthClientUsageQueryResultRequestEntityTooLarge() *GetOauthClientUsageQueryResultRequestEntityTooLarge {
	return &GetOauthClientUsageQueryResultRequestEntityTooLarge{}
}

/*GetOauthClientUsageQueryResultRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOauthClientUsageQueryResultRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOauthClientUsageQueryResultRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients/{clientId}/usage/query/results/{executionId}][%d] getOauthClientUsageQueryResultRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOauthClientUsageQueryResultRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthClientUsageQueryResultRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthClientUsageQueryResultUnsupportedMediaType creates a GetOauthClientUsageQueryResultUnsupportedMediaType with default headers values
func NewGetOauthClientUsageQueryResultUnsupportedMediaType() *GetOauthClientUsageQueryResultUnsupportedMediaType {
	return &GetOauthClientUsageQueryResultUnsupportedMediaType{}
}

/*GetOauthClientUsageQueryResultUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOauthClientUsageQueryResultUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOauthClientUsageQueryResultUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients/{clientId}/usage/query/results/{executionId}][%d] getOauthClientUsageQueryResultUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOauthClientUsageQueryResultUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthClientUsageQueryResultUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthClientUsageQueryResultTooManyRequests creates a GetOauthClientUsageQueryResultTooManyRequests with default headers values
func NewGetOauthClientUsageQueryResultTooManyRequests() *GetOauthClientUsageQueryResultTooManyRequests {
	return &GetOauthClientUsageQueryResultTooManyRequests{}
}

/*GetOauthClientUsageQueryResultTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetOauthClientUsageQueryResultTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOauthClientUsageQueryResultTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients/{clientId}/usage/query/results/{executionId}][%d] getOauthClientUsageQueryResultTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOauthClientUsageQueryResultTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthClientUsageQueryResultTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthClientUsageQueryResultInternalServerError creates a GetOauthClientUsageQueryResultInternalServerError with default headers values
func NewGetOauthClientUsageQueryResultInternalServerError() *GetOauthClientUsageQueryResultInternalServerError {
	return &GetOauthClientUsageQueryResultInternalServerError{}
}

/*GetOauthClientUsageQueryResultInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOauthClientUsageQueryResultInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOauthClientUsageQueryResultInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients/{clientId}/usage/query/results/{executionId}][%d] getOauthClientUsageQueryResultInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOauthClientUsageQueryResultInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthClientUsageQueryResultInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthClientUsageQueryResultServiceUnavailable creates a GetOauthClientUsageQueryResultServiceUnavailable with default headers values
func NewGetOauthClientUsageQueryResultServiceUnavailable() *GetOauthClientUsageQueryResultServiceUnavailable {
	return &GetOauthClientUsageQueryResultServiceUnavailable{}
}

/*GetOauthClientUsageQueryResultServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOauthClientUsageQueryResultServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOauthClientUsageQueryResultServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients/{clientId}/usage/query/results/{executionId}][%d] getOauthClientUsageQueryResultServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOauthClientUsageQueryResultServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthClientUsageQueryResultServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthClientUsageQueryResultGatewayTimeout creates a GetOauthClientUsageQueryResultGatewayTimeout with default headers values
func NewGetOauthClientUsageQueryResultGatewayTimeout() *GetOauthClientUsageQueryResultGatewayTimeout {
	return &GetOauthClientUsageQueryResultGatewayTimeout{}
}

/*GetOauthClientUsageQueryResultGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOauthClientUsageQueryResultGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOauthClientUsageQueryResultGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/clients/{clientId}/usage/query/results/{executionId}][%d] getOauthClientUsageQueryResultGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOauthClientUsageQueryResultGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthClientUsageQueryResultGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
