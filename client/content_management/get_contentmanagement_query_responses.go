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

// GetContentmanagementQueryReader is a Reader for the GetContentmanagementQuery structure.
type GetContentmanagementQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContentmanagementQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContentmanagementQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetContentmanagementQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetContentmanagementQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetContentmanagementQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetContentmanagementQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetContentmanagementQueryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetContentmanagementQueryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetContentmanagementQueryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetContentmanagementQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetContentmanagementQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetContentmanagementQueryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetContentmanagementQueryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetContentmanagementQueryOK creates a GetContentmanagementQueryOK with default headers values
func NewGetContentmanagementQueryOK() *GetContentmanagementQueryOK {
	return &GetContentmanagementQueryOK{}
}

/*GetContentmanagementQueryOK handles this case with default header values.

successful operation
*/
type GetContentmanagementQueryOK struct {
	Payload *models.QueryResults
}

func (o *GetContentmanagementQueryOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/query][%d] getContentmanagementQueryOK  %+v", 200, o.Payload)
}

func (o *GetContentmanagementQueryOK) GetPayload() *models.QueryResults {
	return o.Payload
}

func (o *GetContentmanagementQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueryResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementQueryBadRequest creates a GetContentmanagementQueryBadRequest with default headers values
func NewGetContentmanagementQueryBadRequest() *GetContentmanagementQueryBadRequest {
	return &GetContentmanagementQueryBadRequest{}
}

/*GetContentmanagementQueryBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetContentmanagementQueryBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementQueryBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/query][%d] getContentmanagementQueryBadRequest  %+v", 400, o.Payload)
}

func (o *GetContentmanagementQueryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementQueryUnauthorized creates a GetContentmanagementQueryUnauthorized with default headers values
func NewGetContentmanagementQueryUnauthorized() *GetContentmanagementQueryUnauthorized {
	return &GetContentmanagementQueryUnauthorized{}
}

/*GetContentmanagementQueryUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetContentmanagementQueryUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementQueryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/query][%d] getContentmanagementQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetContentmanagementQueryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementQueryForbidden creates a GetContentmanagementQueryForbidden with default headers values
func NewGetContentmanagementQueryForbidden() *GetContentmanagementQueryForbidden {
	return &GetContentmanagementQueryForbidden{}
}

/*GetContentmanagementQueryForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetContentmanagementQueryForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementQueryForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/query][%d] getContentmanagementQueryForbidden  %+v", 403, o.Payload)
}

func (o *GetContentmanagementQueryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementQueryNotFound creates a GetContentmanagementQueryNotFound with default headers values
func NewGetContentmanagementQueryNotFound() *GetContentmanagementQueryNotFound {
	return &GetContentmanagementQueryNotFound{}
}

/*GetContentmanagementQueryNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetContentmanagementQueryNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementQueryNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/query][%d] getContentmanagementQueryNotFound  %+v", 404, o.Payload)
}

func (o *GetContentmanagementQueryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementQueryRequestTimeout creates a GetContentmanagementQueryRequestTimeout with default headers values
func NewGetContentmanagementQueryRequestTimeout() *GetContentmanagementQueryRequestTimeout {
	return &GetContentmanagementQueryRequestTimeout{}
}

/*GetContentmanagementQueryRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetContentmanagementQueryRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementQueryRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/query][%d] getContentmanagementQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetContentmanagementQueryRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementQueryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementQueryRequestEntityTooLarge creates a GetContentmanagementQueryRequestEntityTooLarge with default headers values
func NewGetContentmanagementQueryRequestEntityTooLarge() *GetContentmanagementQueryRequestEntityTooLarge {
	return &GetContentmanagementQueryRequestEntityTooLarge{}
}

/*GetContentmanagementQueryRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetContentmanagementQueryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementQueryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/query][%d] getContentmanagementQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetContentmanagementQueryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementQueryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementQueryUnsupportedMediaType creates a GetContentmanagementQueryUnsupportedMediaType with default headers values
func NewGetContentmanagementQueryUnsupportedMediaType() *GetContentmanagementQueryUnsupportedMediaType {
	return &GetContentmanagementQueryUnsupportedMediaType{}
}

/*GetContentmanagementQueryUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetContentmanagementQueryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementQueryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/query][%d] getContentmanagementQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetContentmanagementQueryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementQueryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementQueryTooManyRequests creates a GetContentmanagementQueryTooManyRequests with default headers values
func NewGetContentmanagementQueryTooManyRequests() *GetContentmanagementQueryTooManyRequests {
	return &GetContentmanagementQueryTooManyRequests{}
}

/*GetContentmanagementQueryTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetContentmanagementQueryTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/query][%d] getContentmanagementQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetContentmanagementQueryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementQueryInternalServerError creates a GetContentmanagementQueryInternalServerError with default headers values
func NewGetContentmanagementQueryInternalServerError() *GetContentmanagementQueryInternalServerError {
	return &GetContentmanagementQueryInternalServerError{}
}

/*GetContentmanagementQueryInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetContentmanagementQueryInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementQueryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/query][%d] getContentmanagementQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetContentmanagementQueryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementQueryServiceUnavailable creates a GetContentmanagementQueryServiceUnavailable with default headers values
func NewGetContentmanagementQueryServiceUnavailable() *GetContentmanagementQueryServiceUnavailable {
	return &GetContentmanagementQueryServiceUnavailable{}
}

/*GetContentmanagementQueryServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetContentmanagementQueryServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementQueryServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/query][%d] getContentmanagementQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetContentmanagementQueryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementQueryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementQueryGatewayTimeout creates a GetContentmanagementQueryGatewayTimeout with default headers values
func NewGetContentmanagementQueryGatewayTimeout() *GetContentmanagementQueryGatewayTimeout {
	return &GetContentmanagementQueryGatewayTimeout{}
}

/*GetContentmanagementQueryGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetContentmanagementQueryGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementQueryGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/query][%d] getContentmanagementQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetContentmanagementQueryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementQueryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
