// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetAnalyticsConversationsDetailsReader is a Reader for the GetAnalyticsConversationsDetails structure.
type GetAnalyticsConversationsDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnalyticsConversationsDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnalyticsConversationsDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAnalyticsConversationsDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAnalyticsConversationsDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAnalyticsConversationsDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAnalyticsConversationsDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAnalyticsConversationsDetailsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAnalyticsConversationsDetailsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAnalyticsConversationsDetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAnalyticsConversationsDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAnalyticsConversationsDetailsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAnalyticsConversationsDetailsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAnalyticsConversationsDetailsOK creates a GetAnalyticsConversationsDetailsOK with default headers values
func NewGetAnalyticsConversationsDetailsOK() *GetAnalyticsConversationsDetailsOK {
	return &GetAnalyticsConversationsDetailsOK{}
}

/*GetAnalyticsConversationsDetailsOK handles this case with default header values.

successful operation
*/
type GetAnalyticsConversationsDetailsOK struct {
	Payload *models.AnalyticsConversationWithoutAttributesMultiGetResponse
}

func (o *GetAnalyticsConversationsDetailsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details][%d] getAnalyticsConversationsDetailsOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsOK) GetPayload() *models.AnalyticsConversationWithoutAttributesMultiGetResponse {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AnalyticsConversationWithoutAttributesMultiGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsBadRequest creates a GetAnalyticsConversationsDetailsBadRequest with default headers values
func NewGetAnalyticsConversationsDetailsBadRequest() *GetAnalyticsConversationsDetailsBadRequest {
	return &GetAnalyticsConversationsDetailsBadRequest{}
}

/*GetAnalyticsConversationsDetailsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAnalyticsConversationsDetailsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details][%d] getAnalyticsConversationsDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsUnauthorized creates a GetAnalyticsConversationsDetailsUnauthorized with default headers values
func NewGetAnalyticsConversationsDetailsUnauthorized() *GetAnalyticsConversationsDetailsUnauthorized {
	return &GetAnalyticsConversationsDetailsUnauthorized{}
}

/*GetAnalyticsConversationsDetailsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAnalyticsConversationsDetailsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details][%d] getAnalyticsConversationsDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsForbidden creates a GetAnalyticsConversationsDetailsForbidden with default headers values
func NewGetAnalyticsConversationsDetailsForbidden() *GetAnalyticsConversationsDetailsForbidden {
	return &GetAnalyticsConversationsDetailsForbidden{}
}

/*GetAnalyticsConversationsDetailsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAnalyticsConversationsDetailsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details][%d] getAnalyticsConversationsDetailsForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsNotFound creates a GetAnalyticsConversationsDetailsNotFound with default headers values
func NewGetAnalyticsConversationsDetailsNotFound() *GetAnalyticsConversationsDetailsNotFound {
	return &GetAnalyticsConversationsDetailsNotFound{}
}

/*GetAnalyticsConversationsDetailsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAnalyticsConversationsDetailsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details][%d] getAnalyticsConversationsDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsRequestEntityTooLarge creates a GetAnalyticsConversationsDetailsRequestEntityTooLarge with default headers values
func NewGetAnalyticsConversationsDetailsRequestEntityTooLarge() *GetAnalyticsConversationsDetailsRequestEntityTooLarge {
	return &GetAnalyticsConversationsDetailsRequestEntityTooLarge{}
}

/*GetAnalyticsConversationsDetailsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAnalyticsConversationsDetailsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details][%d] getAnalyticsConversationsDetailsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsUnsupportedMediaType creates a GetAnalyticsConversationsDetailsUnsupportedMediaType with default headers values
func NewGetAnalyticsConversationsDetailsUnsupportedMediaType() *GetAnalyticsConversationsDetailsUnsupportedMediaType {
	return &GetAnalyticsConversationsDetailsUnsupportedMediaType{}
}

/*GetAnalyticsConversationsDetailsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAnalyticsConversationsDetailsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details][%d] getAnalyticsConversationsDetailsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsTooManyRequests creates a GetAnalyticsConversationsDetailsTooManyRequests with default headers values
func NewGetAnalyticsConversationsDetailsTooManyRequests() *GetAnalyticsConversationsDetailsTooManyRequests {
	return &GetAnalyticsConversationsDetailsTooManyRequests{}
}

/*GetAnalyticsConversationsDetailsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetAnalyticsConversationsDetailsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details][%d] getAnalyticsConversationsDetailsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsInternalServerError creates a GetAnalyticsConversationsDetailsInternalServerError with default headers values
func NewGetAnalyticsConversationsDetailsInternalServerError() *GetAnalyticsConversationsDetailsInternalServerError {
	return &GetAnalyticsConversationsDetailsInternalServerError{}
}

/*GetAnalyticsConversationsDetailsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAnalyticsConversationsDetailsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details][%d] getAnalyticsConversationsDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsServiceUnavailable creates a GetAnalyticsConversationsDetailsServiceUnavailable with default headers values
func NewGetAnalyticsConversationsDetailsServiceUnavailable() *GetAnalyticsConversationsDetailsServiceUnavailable {
	return &GetAnalyticsConversationsDetailsServiceUnavailable{}
}

/*GetAnalyticsConversationsDetailsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAnalyticsConversationsDetailsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details][%d] getAnalyticsConversationsDetailsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsGatewayTimeout creates a GetAnalyticsConversationsDetailsGatewayTimeout with default headers values
func NewGetAnalyticsConversationsDetailsGatewayTimeout() *GetAnalyticsConversationsDetailsGatewayTimeout {
	return &GetAnalyticsConversationsDetailsGatewayTimeout{}
}

/*GetAnalyticsConversationsDetailsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAnalyticsConversationsDetailsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details][%d] getAnalyticsConversationsDetailsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
