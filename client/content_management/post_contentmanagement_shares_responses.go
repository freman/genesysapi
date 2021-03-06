// Code generated by go-swagger; DO NOT EDIT.

package content_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostContentmanagementSharesReader is a Reader for the PostContentmanagementShares structure.
type PostContentmanagementSharesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostContentmanagementSharesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostContentmanagementSharesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostContentmanagementSharesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostContentmanagementSharesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostContentmanagementSharesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostContentmanagementSharesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostContentmanagementSharesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostContentmanagementSharesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostContentmanagementSharesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostContentmanagementSharesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostContentmanagementSharesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostContentmanagementSharesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostContentmanagementSharesOK creates a PostContentmanagementSharesOK with default headers values
func NewPostContentmanagementSharesOK() *PostContentmanagementSharesOK {
	return &PostContentmanagementSharesOK{}
}

/*PostContentmanagementSharesOK handles this case with default header values.

successful operation
*/
type PostContentmanagementSharesOK struct {
	Payload *models.CreateShareResponse
}

func (o *PostContentmanagementSharesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/shares][%d] postContentmanagementSharesOK  %+v", 200, o.Payload)
}

func (o *PostContentmanagementSharesOK) GetPayload() *models.CreateShareResponse {
	return o.Payload
}

func (o *PostContentmanagementSharesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateShareResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementSharesBadRequest creates a PostContentmanagementSharesBadRequest with default headers values
func NewPostContentmanagementSharesBadRequest() *PostContentmanagementSharesBadRequest {
	return &PostContentmanagementSharesBadRequest{}
}

/*PostContentmanagementSharesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostContentmanagementSharesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementSharesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/shares][%d] postContentmanagementSharesBadRequest  %+v", 400, o.Payload)
}

func (o *PostContentmanagementSharesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementSharesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementSharesUnauthorized creates a PostContentmanagementSharesUnauthorized with default headers values
func NewPostContentmanagementSharesUnauthorized() *PostContentmanagementSharesUnauthorized {
	return &PostContentmanagementSharesUnauthorized{}
}

/*PostContentmanagementSharesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostContentmanagementSharesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementSharesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/shares][%d] postContentmanagementSharesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostContentmanagementSharesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementSharesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementSharesForbidden creates a PostContentmanagementSharesForbidden with default headers values
func NewPostContentmanagementSharesForbidden() *PostContentmanagementSharesForbidden {
	return &PostContentmanagementSharesForbidden{}
}

/*PostContentmanagementSharesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostContentmanagementSharesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementSharesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/shares][%d] postContentmanagementSharesForbidden  %+v", 403, o.Payload)
}

func (o *PostContentmanagementSharesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementSharesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementSharesNotFound creates a PostContentmanagementSharesNotFound with default headers values
func NewPostContentmanagementSharesNotFound() *PostContentmanagementSharesNotFound {
	return &PostContentmanagementSharesNotFound{}
}

/*PostContentmanagementSharesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostContentmanagementSharesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementSharesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/shares][%d] postContentmanagementSharesNotFound  %+v", 404, o.Payload)
}

func (o *PostContentmanagementSharesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementSharesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementSharesRequestEntityTooLarge creates a PostContentmanagementSharesRequestEntityTooLarge with default headers values
func NewPostContentmanagementSharesRequestEntityTooLarge() *PostContentmanagementSharesRequestEntityTooLarge {
	return &PostContentmanagementSharesRequestEntityTooLarge{}
}

/*PostContentmanagementSharesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostContentmanagementSharesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementSharesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/shares][%d] postContentmanagementSharesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostContentmanagementSharesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementSharesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementSharesUnsupportedMediaType creates a PostContentmanagementSharesUnsupportedMediaType with default headers values
func NewPostContentmanagementSharesUnsupportedMediaType() *PostContentmanagementSharesUnsupportedMediaType {
	return &PostContentmanagementSharesUnsupportedMediaType{}
}

/*PostContentmanagementSharesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostContentmanagementSharesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementSharesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/shares][%d] postContentmanagementSharesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostContentmanagementSharesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementSharesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementSharesTooManyRequests creates a PostContentmanagementSharesTooManyRequests with default headers values
func NewPostContentmanagementSharesTooManyRequests() *PostContentmanagementSharesTooManyRequests {
	return &PostContentmanagementSharesTooManyRequests{}
}

/*PostContentmanagementSharesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostContentmanagementSharesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementSharesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/shares][%d] postContentmanagementSharesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostContentmanagementSharesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementSharesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementSharesInternalServerError creates a PostContentmanagementSharesInternalServerError with default headers values
func NewPostContentmanagementSharesInternalServerError() *PostContentmanagementSharesInternalServerError {
	return &PostContentmanagementSharesInternalServerError{}
}

/*PostContentmanagementSharesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostContentmanagementSharesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementSharesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/shares][%d] postContentmanagementSharesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostContentmanagementSharesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementSharesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementSharesServiceUnavailable creates a PostContentmanagementSharesServiceUnavailable with default headers values
func NewPostContentmanagementSharesServiceUnavailable() *PostContentmanagementSharesServiceUnavailable {
	return &PostContentmanagementSharesServiceUnavailable{}
}

/*PostContentmanagementSharesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostContentmanagementSharesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementSharesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/shares][%d] postContentmanagementSharesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostContentmanagementSharesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementSharesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementSharesGatewayTimeout creates a PostContentmanagementSharesGatewayTimeout with default headers values
func NewPostContentmanagementSharesGatewayTimeout() *PostContentmanagementSharesGatewayTimeout {
	return &PostContentmanagementSharesGatewayTimeout{}
}

/*PostContentmanagementSharesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostContentmanagementSharesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementSharesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/shares][%d] postContentmanagementSharesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostContentmanagementSharesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementSharesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
