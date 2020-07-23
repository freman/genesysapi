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

// GetContentmanagementDocumentContentReader is a Reader for the GetContentmanagementDocumentContent structure.
type GetContentmanagementDocumentContentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContentmanagementDocumentContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContentmanagementDocumentContentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewGetContentmanagementDocumentContentAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetContentmanagementDocumentContentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetContentmanagementDocumentContentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetContentmanagementDocumentContentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetContentmanagementDocumentContentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetContentmanagementDocumentContentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetContentmanagementDocumentContentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetContentmanagementDocumentContentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetContentmanagementDocumentContentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetContentmanagementDocumentContentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetContentmanagementDocumentContentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetContentmanagementDocumentContentOK creates a GetContentmanagementDocumentContentOK with default headers values
func NewGetContentmanagementDocumentContentOK() *GetContentmanagementDocumentContentOK {
	return &GetContentmanagementDocumentContentOK{}
}

/*GetContentmanagementDocumentContentOK handles this case with default header values.

Download location returned
*/
type GetContentmanagementDocumentContentOK struct {
	Payload *models.DownloadResponse
}

func (o *GetContentmanagementDocumentContentOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/documents/{documentId}/content][%d] getContentmanagementDocumentContentOK  %+v", 200, o.Payload)
}

func (o *GetContentmanagementDocumentContentOK) GetPayload() *models.DownloadResponse {
	return o.Payload
}

func (o *GetContentmanagementDocumentContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DownloadResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementDocumentContentAccepted creates a GetContentmanagementDocumentContentAccepted with default headers values
func NewGetContentmanagementDocumentContentAccepted() *GetContentmanagementDocumentContentAccepted {
	return &GetContentmanagementDocumentContentAccepted{}
}

/*GetContentmanagementDocumentContentAccepted handles this case with default header values.

Accepted - Preparing file for download - try again soon.
*/
type GetContentmanagementDocumentContentAccepted struct {
}

func (o *GetContentmanagementDocumentContentAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/documents/{documentId}/content][%d] getContentmanagementDocumentContentAccepted ", 202)
}

func (o *GetContentmanagementDocumentContentAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetContentmanagementDocumentContentBadRequest creates a GetContentmanagementDocumentContentBadRequest with default headers values
func NewGetContentmanagementDocumentContentBadRequest() *GetContentmanagementDocumentContentBadRequest {
	return &GetContentmanagementDocumentContentBadRequest{}
}

/*GetContentmanagementDocumentContentBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetContentmanagementDocumentContentBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementDocumentContentBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/documents/{documentId}/content][%d] getContentmanagementDocumentContentBadRequest  %+v", 400, o.Payload)
}

func (o *GetContentmanagementDocumentContentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementDocumentContentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementDocumentContentUnauthorized creates a GetContentmanagementDocumentContentUnauthorized with default headers values
func NewGetContentmanagementDocumentContentUnauthorized() *GetContentmanagementDocumentContentUnauthorized {
	return &GetContentmanagementDocumentContentUnauthorized{}
}

/*GetContentmanagementDocumentContentUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetContentmanagementDocumentContentUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementDocumentContentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/documents/{documentId}/content][%d] getContentmanagementDocumentContentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetContentmanagementDocumentContentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementDocumentContentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementDocumentContentForbidden creates a GetContentmanagementDocumentContentForbidden with default headers values
func NewGetContentmanagementDocumentContentForbidden() *GetContentmanagementDocumentContentForbidden {
	return &GetContentmanagementDocumentContentForbidden{}
}

/*GetContentmanagementDocumentContentForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetContentmanagementDocumentContentForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementDocumentContentForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/documents/{documentId}/content][%d] getContentmanagementDocumentContentForbidden  %+v", 403, o.Payload)
}

func (o *GetContentmanagementDocumentContentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementDocumentContentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementDocumentContentNotFound creates a GetContentmanagementDocumentContentNotFound with default headers values
func NewGetContentmanagementDocumentContentNotFound() *GetContentmanagementDocumentContentNotFound {
	return &GetContentmanagementDocumentContentNotFound{}
}

/*GetContentmanagementDocumentContentNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetContentmanagementDocumentContentNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementDocumentContentNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/documents/{documentId}/content][%d] getContentmanagementDocumentContentNotFound  %+v", 404, o.Payload)
}

func (o *GetContentmanagementDocumentContentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementDocumentContentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementDocumentContentRequestEntityTooLarge creates a GetContentmanagementDocumentContentRequestEntityTooLarge with default headers values
func NewGetContentmanagementDocumentContentRequestEntityTooLarge() *GetContentmanagementDocumentContentRequestEntityTooLarge {
	return &GetContentmanagementDocumentContentRequestEntityTooLarge{}
}

/*GetContentmanagementDocumentContentRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetContentmanagementDocumentContentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementDocumentContentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/documents/{documentId}/content][%d] getContentmanagementDocumentContentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetContentmanagementDocumentContentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementDocumentContentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementDocumentContentUnsupportedMediaType creates a GetContentmanagementDocumentContentUnsupportedMediaType with default headers values
func NewGetContentmanagementDocumentContentUnsupportedMediaType() *GetContentmanagementDocumentContentUnsupportedMediaType {
	return &GetContentmanagementDocumentContentUnsupportedMediaType{}
}

/*GetContentmanagementDocumentContentUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetContentmanagementDocumentContentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementDocumentContentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/documents/{documentId}/content][%d] getContentmanagementDocumentContentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetContentmanagementDocumentContentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementDocumentContentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementDocumentContentTooManyRequests creates a GetContentmanagementDocumentContentTooManyRequests with default headers values
func NewGetContentmanagementDocumentContentTooManyRequests() *GetContentmanagementDocumentContentTooManyRequests {
	return &GetContentmanagementDocumentContentTooManyRequests{}
}

/*GetContentmanagementDocumentContentTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetContentmanagementDocumentContentTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementDocumentContentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/documents/{documentId}/content][%d] getContentmanagementDocumentContentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetContentmanagementDocumentContentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementDocumentContentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementDocumentContentInternalServerError creates a GetContentmanagementDocumentContentInternalServerError with default headers values
func NewGetContentmanagementDocumentContentInternalServerError() *GetContentmanagementDocumentContentInternalServerError {
	return &GetContentmanagementDocumentContentInternalServerError{}
}

/*GetContentmanagementDocumentContentInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetContentmanagementDocumentContentInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementDocumentContentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/documents/{documentId}/content][%d] getContentmanagementDocumentContentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetContentmanagementDocumentContentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementDocumentContentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementDocumentContentServiceUnavailable creates a GetContentmanagementDocumentContentServiceUnavailable with default headers values
func NewGetContentmanagementDocumentContentServiceUnavailable() *GetContentmanagementDocumentContentServiceUnavailable {
	return &GetContentmanagementDocumentContentServiceUnavailable{}
}

/*GetContentmanagementDocumentContentServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetContentmanagementDocumentContentServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementDocumentContentServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/documents/{documentId}/content][%d] getContentmanagementDocumentContentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetContentmanagementDocumentContentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementDocumentContentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementDocumentContentGatewayTimeout creates a GetContentmanagementDocumentContentGatewayTimeout with default headers values
func NewGetContentmanagementDocumentContentGatewayTimeout() *GetContentmanagementDocumentContentGatewayTimeout {
	return &GetContentmanagementDocumentContentGatewayTimeout{}
}

/*GetContentmanagementDocumentContentGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetContentmanagementDocumentContentGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementDocumentContentGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/documents/{documentId}/content][%d] getContentmanagementDocumentContentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetContentmanagementDocumentContentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementDocumentContentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
