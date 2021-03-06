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

// GetAnalyticsConversationDetailsReader is a Reader for the GetAnalyticsConversationDetails structure.
type GetAnalyticsConversationDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnalyticsConversationDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnalyticsConversationDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAnalyticsConversationDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAnalyticsConversationDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAnalyticsConversationDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAnalyticsConversationDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAnalyticsConversationDetailsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAnalyticsConversationDetailsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAnalyticsConversationDetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAnalyticsConversationDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAnalyticsConversationDetailsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAnalyticsConversationDetailsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAnalyticsConversationDetailsOK creates a GetAnalyticsConversationDetailsOK with default headers values
func NewGetAnalyticsConversationDetailsOK() *GetAnalyticsConversationDetailsOK {
	return &GetAnalyticsConversationDetailsOK{}
}

/*GetAnalyticsConversationDetailsOK handles this case with default header values.

successful operation
*/
type GetAnalyticsConversationDetailsOK struct {
	Payload *models.AnalyticsConversationWithoutAttributes
}

func (o *GetAnalyticsConversationDetailsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/{conversationId}/details][%d] getAnalyticsConversationDetailsOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsConversationDetailsOK) GetPayload() *models.AnalyticsConversationWithoutAttributes {
	return o.Payload
}

func (o *GetAnalyticsConversationDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AnalyticsConversationWithoutAttributes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationDetailsBadRequest creates a GetAnalyticsConversationDetailsBadRequest with default headers values
func NewGetAnalyticsConversationDetailsBadRequest() *GetAnalyticsConversationDetailsBadRequest {
	return &GetAnalyticsConversationDetailsBadRequest{}
}

/*GetAnalyticsConversationDetailsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAnalyticsConversationDetailsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/{conversationId}/details][%d] getAnalyticsConversationDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsConversationDetailsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationDetailsUnauthorized creates a GetAnalyticsConversationDetailsUnauthorized with default headers values
func NewGetAnalyticsConversationDetailsUnauthorized() *GetAnalyticsConversationDetailsUnauthorized {
	return &GetAnalyticsConversationDetailsUnauthorized{}
}

/*GetAnalyticsConversationDetailsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAnalyticsConversationDetailsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/{conversationId}/details][%d] getAnalyticsConversationDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsConversationDetailsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationDetailsForbidden creates a GetAnalyticsConversationDetailsForbidden with default headers values
func NewGetAnalyticsConversationDetailsForbidden() *GetAnalyticsConversationDetailsForbidden {
	return &GetAnalyticsConversationDetailsForbidden{}
}

/*GetAnalyticsConversationDetailsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAnalyticsConversationDetailsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/{conversationId}/details][%d] getAnalyticsConversationDetailsForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsConversationDetailsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationDetailsNotFound creates a GetAnalyticsConversationDetailsNotFound with default headers values
func NewGetAnalyticsConversationDetailsNotFound() *GetAnalyticsConversationDetailsNotFound {
	return &GetAnalyticsConversationDetailsNotFound{}
}

/*GetAnalyticsConversationDetailsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAnalyticsConversationDetailsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/{conversationId}/details][%d] getAnalyticsConversationDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsConversationDetailsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationDetailsRequestEntityTooLarge creates a GetAnalyticsConversationDetailsRequestEntityTooLarge with default headers values
func NewGetAnalyticsConversationDetailsRequestEntityTooLarge() *GetAnalyticsConversationDetailsRequestEntityTooLarge {
	return &GetAnalyticsConversationDetailsRequestEntityTooLarge{}
}

/*GetAnalyticsConversationDetailsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAnalyticsConversationDetailsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationDetailsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/{conversationId}/details][%d] getAnalyticsConversationDetailsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsConversationDetailsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationDetailsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationDetailsUnsupportedMediaType creates a GetAnalyticsConversationDetailsUnsupportedMediaType with default headers values
func NewGetAnalyticsConversationDetailsUnsupportedMediaType() *GetAnalyticsConversationDetailsUnsupportedMediaType {
	return &GetAnalyticsConversationDetailsUnsupportedMediaType{}
}

/*GetAnalyticsConversationDetailsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAnalyticsConversationDetailsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationDetailsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/{conversationId}/details][%d] getAnalyticsConversationDetailsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsConversationDetailsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationDetailsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationDetailsTooManyRequests creates a GetAnalyticsConversationDetailsTooManyRequests with default headers values
func NewGetAnalyticsConversationDetailsTooManyRequests() *GetAnalyticsConversationDetailsTooManyRequests {
	return &GetAnalyticsConversationDetailsTooManyRequests{}
}

/*GetAnalyticsConversationDetailsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetAnalyticsConversationDetailsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationDetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/{conversationId}/details][%d] getAnalyticsConversationDetailsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsConversationDetailsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationDetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationDetailsInternalServerError creates a GetAnalyticsConversationDetailsInternalServerError with default headers values
func NewGetAnalyticsConversationDetailsInternalServerError() *GetAnalyticsConversationDetailsInternalServerError {
	return &GetAnalyticsConversationDetailsInternalServerError{}
}

/*GetAnalyticsConversationDetailsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAnalyticsConversationDetailsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/{conversationId}/details][%d] getAnalyticsConversationDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsConversationDetailsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationDetailsServiceUnavailable creates a GetAnalyticsConversationDetailsServiceUnavailable with default headers values
func NewGetAnalyticsConversationDetailsServiceUnavailable() *GetAnalyticsConversationDetailsServiceUnavailable {
	return &GetAnalyticsConversationDetailsServiceUnavailable{}
}

/*GetAnalyticsConversationDetailsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAnalyticsConversationDetailsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationDetailsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/{conversationId}/details][%d] getAnalyticsConversationDetailsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsConversationDetailsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationDetailsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationDetailsGatewayTimeout creates a GetAnalyticsConversationDetailsGatewayTimeout with default headers values
func NewGetAnalyticsConversationDetailsGatewayTimeout() *GetAnalyticsConversationDetailsGatewayTimeout {
	return &GetAnalyticsConversationDetailsGatewayTimeout{}
}

/*GetAnalyticsConversationDetailsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAnalyticsConversationDetailsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationDetailsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/{conversationId}/details][%d] getAnalyticsConversationDetailsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsConversationDetailsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationDetailsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
