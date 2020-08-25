// Code generated by go-swagger; DO NOT EDIT.

package response_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetResponsemanagementResponseReader is a Reader for the GetResponsemanagementResponse structure.
type GetResponsemanagementResponseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResponsemanagementResponseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResponsemanagementResponseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetResponsemanagementResponseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetResponsemanagementResponseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetResponsemanagementResponseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetResponsemanagementResponseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetResponsemanagementResponseRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetResponsemanagementResponseUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetResponsemanagementResponseTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetResponsemanagementResponseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetResponsemanagementResponseServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetResponsemanagementResponseGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResponsemanagementResponseOK creates a GetResponsemanagementResponseOK with default headers values
func NewGetResponsemanagementResponseOK() *GetResponsemanagementResponseOK {
	return &GetResponsemanagementResponseOK{}
}

/*GetResponsemanagementResponseOK handles this case with default header values.

successful operation
*/
type GetResponsemanagementResponseOK struct {
	Payload *models.Response
}

func (o *GetResponsemanagementResponseOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseOK  %+v", 200, o.Payload)
}

func (o *GetResponsemanagementResponseOK) GetPayload() *models.Response {
	return o.Payload
}

func (o *GetResponsemanagementResponseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseBadRequest creates a GetResponsemanagementResponseBadRequest with default headers values
func NewGetResponsemanagementResponseBadRequest() *GetResponsemanagementResponseBadRequest {
	return &GetResponsemanagementResponseBadRequest{}
}

/*GetResponsemanagementResponseBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetResponsemanagementResponseBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementResponseBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseBadRequest  %+v", 400, o.Payload)
}

func (o *GetResponsemanagementResponseBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseUnauthorized creates a GetResponsemanagementResponseUnauthorized with default headers values
func NewGetResponsemanagementResponseUnauthorized() *GetResponsemanagementResponseUnauthorized {
	return &GetResponsemanagementResponseUnauthorized{}
}

/*GetResponsemanagementResponseUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetResponsemanagementResponseUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementResponseUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseUnauthorized  %+v", 401, o.Payload)
}

func (o *GetResponsemanagementResponseUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseForbidden creates a GetResponsemanagementResponseForbidden with default headers values
func NewGetResponsemanagementResponseForbidden() *GetResponsemanagementResponseForbidden {
	return &GetResponsemanagementResponseForbidden{}
}

/*GetResponsemanagementResponseForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetResponsemanagementResponseForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementResponseForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseForbidden  %+v", 403, o.Payload)
}

func (o *GetResponsemanagementResponseForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseNotFound creates a GetResponsemanagementResponseNotFound with default headers values
func NewGetResponsemanagementResponseNotFound() *GetResponsemanagementResponseNotFound {
	return &GetResponsemanagementResponseNotFound{}
}

/*GetResponsemanagementResponseNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetResponsemanagementResponseNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementResponseNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseNotFound  %+v", 404, o.Payload)
}

func (o *GetResponsemanagementResponseNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseRequestEntityTooLarge creates a GetResponsemanagementResponseRequestEntityTooLarge with default headers values
func NewGetResponsemanagementResponseRequestEntityTooLarge() *GetResponsemanagementResponseRequestEntityTooLarge {
	return &GetResponsemanagementResponseRequestEntityTooLarge{}
}

/*GetResponsemanagementResponseRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetResponsemanagementResponseRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementResponseRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetResponsemanagementResponseRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseUnsupportedMediaType creates a GetResponsemanagementResponseUnsupportedMediaType with default headers values
func NewGetResponsemanagementResponseUnsupportedMediaType() *GetResponsemanagementResponseUnsupportedMediaType {
	return &GetResponsemanagementResponseUnsupportedMediaType{}
}

/*GetResponsemanagementResponseUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetResponsemanagementResponseUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementResponseUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetResponsemanagementResponseUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseTooManyRequests creates a GetResponsemanagementResponseTooManyRequests with default headers values
func NewGetResponsemanagementResponseTooManyRequests() *GetResponsemanagementResponseTooManyRequests {
	return &GetResponsemanagementResponseTooManyRequests{}
}

/*GetResponsemanagementResponseTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetResponsemanagementResponseTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementResponseTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetResponsemanagementResponseTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseInternalServerError creates a GetResponsemanagementResponseInternalServerError with default headers values
func NewGetResponsemanagementResponseInternalServerError() *GetResponsemanagementResponseInternalServerError {
	return &GetResponsemanagementResponseInternalServerError{}
}

/*GetResponsemanagementResponseInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetResponsemanagementResponseInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementResponseInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseInternalServerError  %+v", 500, o.Payload)
}

func (o *GetResponsemanagementResponseInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseServiceUnavailable creates a GetResponsemanagementResponseServiceUnavailable with default headers values
func NewGetResponsemanagementResponseServiceUnavailable() *GetResponsemanagementResponseServiceUnavailable {
	return &GetResponsemanagementResponseServiceUnavailable{}
}

/*GetResponsemanagementResponseServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetResponsemanagementResponseServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementResponseServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetResponsemanagementResponseServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseGatewayTimeout creates a GetResponsemanagementResponseGatewayTimeout with default headers values
func NewGetResponsemanagementResponseGatewayTimeout() *GetResponsemanagementResponseGatewayTimeout {
	return &GetResponsemanagementResponseGatewayTimeout{}
}

/*GetResponsemanagementResponseGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetResponsemanagementResponseGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementResponseGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetResponsemanagementResponseGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}