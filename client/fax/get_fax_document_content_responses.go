// Code generated by go-swagger; DO NOT EDIT.

package fax

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetFaxDocumentContentReader is a Reader for the GetFaxDocumentContent structure.
type GetFaxDocumentContentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFaxDocumentContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFaxDocumentContentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFaxDocumentContentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFaxDocumentContentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFaxDocumentContentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFaxDocumentContentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetFaxDocumentContentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFaxDocumentContentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFaxDocumentContentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFaxDocumentContentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFaxDocumentContentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFaxDocumentContentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFaxDocumentContentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFaxDocumentContentOK creates a GetFaxDocumentContentOK with default headers values
func NewGetFaxDocumentContentOK() *GetFaxDocumentContentOK {
	return &GetFaxDocumentContentOK{}
}

/*GetFaxDocumentContentOK handles this case with default header values.

successful operation
*/
type GetFaxDocumentContentOK struct {
	Payload *models.DownloadResponse
}

func (o *GetFaxDocumentContentOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}/content][%d] getFaxDocumentContentOK  %+v", 200, o.Payload)
}

func (o *GetFaxDocumentContentOK) GetPayload() *models.DownloadResponse {
	return o.Payload
}

func (o *GetFaxDocumentContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DownloadResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentContentBadRequest creates a GetFaxDocumentContentBadRequest with default headers values
func NewGetFaxDocumentContentBadRequest() *GetFaxDocumentContentBadRequest {
	return &GetFaxDocumentContentBadRequest{}
}

/*GetFaxDocumentContentBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFaxDocumentContentBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentContentBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}/content][%d] getFaxDocumentContentBadRequest  %+v", 400, o.Payload)
}

func (o *GetFaxDocumentContentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentContentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentContentUnauthorized creates a GetFaxDocumentContentUnauthorized with default headers values
func NewGetFaxDocumentContentUnauthorized() *GetFaxDocumentContentUnauthorized {
	return &GetFaxDocumentContentUnauthorized{}
}

/*GetFaxDocumentContentUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFaxDocumentContentUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentContentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}/content][%d] getFaxDocumentContentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFaxDocumentContentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentContentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentContentForbidden creates a GetFaxDocumentContentForbidden with default headers values
func NewGetFaxDocumentContentForbidden() *GetFaxDocumentContentForbidden {
	return &GetFaxDocumentContentForbidden{}
}

/*GetFaxDocumentContentForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetFaxDocumentContentForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentContentForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}/content][%d] getFaxDocumentContentForbidden  %+v", 403, o.Payload)
}

func (o *GetFaxDocumentContentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentContentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentContentNotFound creates a GetFaxDocumentContentNotFound with default headers values
func NewGetFaxDocumentContentNotFound() *GetFaxDocumentContentNotFound {
	return &GetFaxDocumentContentNotFound{}
}

/*GetFaxDocumentContentNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetFaxDocumentContentNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentContentNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}/content][%d] getFaxDocumentContentNotFound  %+v", 404, o.Payload)
}

func (o *GetFaxDocumentContentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentContentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentContentRequestTimeout creates a GetFaxDocumentContentRequestTimeout with default headers values
func NewGetFaxDocumentContentRequestTimeout() *GetFaxDocumentContentRequestTimeout {
	return &GetFaxDocumentContentRequestTimeout{}
}

/*GetFaxDocumentContentRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetFaxDocumentContentRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentContentRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}/content][%d] getFaxDocumentContentRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFaxDocumentContentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentContentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentContentRequestEntityTooLarge creates a GetFaxDocumentContentRequestEntityTooLarge with default headers values
func NewGetFaxDocumentContentRequestEntityTooLarge() *GetFaxDocumentContentRequestEntityTooLarge {
	return &GetFaxDocumentContentRequestEntityTooLarge{}
}

/*GetFaxDocumentContentRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetFaxDocumentContentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentContentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}/content][%d] getFaxDocumentContentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFaxDocumentContentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentContentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentContentUnsupportedMediaType creates a GetFaxDocumentContentUnsupportedMediaType with default headers values
func NewGetFaxDocumentContentUnsupportedMediaType() *GetFaxDocumentContentUnsupportedMediaType {
	return &GetFaxDocumentContentUnsupportedMediaType{}
}

/*GetFaxDocumentContentUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFaxDocumentContentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentContentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}/content][%d] getFaxDocumentContentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFaxDocumentContentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentContentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentContentTooManyRequests creates a GetFaxDocumentContentTooManyRequests with default headers values
func NewGetFaxDocumentContentTooManyRequests() *GetFaxDocumentContentTooManyRequests {
	return &GetFaxDocumentContentTooManyRequests{}
}

/*GetFaxDocumentContentTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetFaxDocumentContentTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentContentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}/content][%d] getFaxDocumentContentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFaxDocumentContentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentContentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentContentInternalServerError creates a GetFaxDocumentContentInternalServerError with default headers values
func NewGetFaxDocumentContentInternalServerError() *GetFaxDocumentContentInternalServerError {
	return &GetFaxDocumentContentInternalServerError{}
}

/*GetFaxDocumentContentInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFaxDocumentContentInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentContentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}/content][%d] getFaxDocumentContentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFaxDocumentContentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentContentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentContentServiceUnavailable creates a GetFaxDocumentContentServiceUnavailable with default headers values
func NewGetFaxDocumentContentServiceUnavailable() *GetFaxDocumentContentServiceUnavailable {
	return &GetFaxDocumentContentServiceUnavailable{}
}

/*GetFaxDocumentContentServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFaxDocumentContentServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentContentServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}/content][%d] getFaxDocumentContentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFaxDocumentContentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentContentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentContentGatewayTimeout creates a GetFaxDocumentContentGatewayTimeout with default headers values
func NewGetFaxDocumentContentGatewayTimeout() *GetFaxDocumentContentGatewayTimeout {
	return &GetFaxDocumentContentGatewayTimeout{}
}

/*GetFaxDocumentContentGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetFaxDocumentContentGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentContentGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}/content][%d] getFaxDocumentContentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFaxDocumentContentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentContentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
