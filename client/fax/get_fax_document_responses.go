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

// GetFaxDocumentReader is a Reader for the GetFaxDocument structure.
type GetFaxDocumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFaxDocumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFaxDocumentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFaxDocumentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFaxDocumentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFaxDocumentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFaxDocumentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFaxDocumentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFaxDocumentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFaxDocumentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFaxDocumentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFaxDocumentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFaxDocumentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFaxDocumentOK creates a GetFaxDocumentOK with default headers values
func NewGetFaxDocumentOK() *GetFaxDocumentOK {
	return &GetFaxDocumentOK{}
}

/*GetFaxDocumentOK handles this case with default header values.

successful operation
*/
type GetFaxDocumentOK struct {
	Payload *models.FaxDocument
}

func (o *GetFaxDocumentOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentOK  %+v", 200, o.Payload)
}

func (o *GetFaxDocumentOK) GetPayload() *models.FaxDocument {
	return o.Payload
}

func (o *GetFaxDocumentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FaxDocument)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentBadRequest creates a GetFaxDocumentBadRequest with default headers values
func NewGetFaxDocumentBadRequest() *GetFaxDocumentBadRequest {
	return &GetFaxDocumentBadRequest{}
}

/*GetFaxDocumentBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFaxDocumentBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentBadRequest  %+v", 400, o.Payload)
}

func (o *GetFaxDocumentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentUnauthorized creates a GetFaxDocumentUnauthorized with default headers values
func NewGetFaxDocumentUnauthorized() *GetFaxDocumentUnauthorized {
	return &GetFaxDocumentUnauthorized{}
}

/*GetFaxDocumentUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFaxDocumentUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFaxDocumentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentForbidden creates a GetFaxDocumentForbidden with default headers values
func NewGetFaxDocumentForbidden() *GetFaxDocumentForbidden {
	return &GetFaxDocumentForbidden{}
}

/*GetFaxDocumentForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetFaxDocumentForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentForbidden  %+v", 403, o.Payload)
}

func (o *GetFaxDocumentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentNotFound creates a GetFaxDocumentNotFound with default headers values
func NewGetFaxDocumentNotFound() *GetFaxDocumentNotFound {
	return &GetFaxDocumentNotFound{}
}

/*GetFaxDocumentNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetFaxDocumentNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentNotFound  %+v", 404, o.Payload)
}

func (o *GetFaxDocumentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentRequestEntityTooLarge creates a GetFaxDocumentRequestEntityTooLarge with default headers values
func NewGetFaxDocumentRequestEntityTooLarge() *GetFaxDocumentRequestEntityTooLarge {
	return &GetFaxDocumentRequestEntityTooLarge{}
}

/*GetFaxDocumentRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetFaxDocumentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFaxDocumentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentUnsupportedMediaType creates a GetFaxDocumentUnsupportedMediaType with default headers values
func NewGetFaxDocumentUnsupportedMediaType() *GetFaxDocumentUnsupportedMediaType {
	return &GetFaxDocumentUnsupportedMediaType{}
}

/*GetFaxDocumentUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFaxDocumentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFaxDocumentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentTooManyRequests creates a GetFaxDocumentTooManyRequests with default headers values
func NewGetFaxDocumentTooManyRequests() *GetFaxDocumentTooManyRequests {
	return &GetFaxDocumentTooManyRequests{}
}

/*GetFaxDocumentTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetFaxDocumentTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFaxDocumentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentInternalServerError creates a GetFaxDocumentInternalServerError with default headers values
func NewGetFaxDocumentInternalServerError() *GetFaxDocumentInternalServerError {
	return &GetFaxDocumentInternalServerError{}
}

/*GetFaxDocumentInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFaxDocumentInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFaxDocumentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentServiceUnavailable creates a GetFaxDocumentServiceUnavailable with default headers values
func NewGetFaxDocumentServiceUnavailable() *GetFaxDocumentServiceUnavailable {
	return &GetFaxDocumentServiceUnavailable{}
}

/*GetFaxDocumentServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFaxDocumentServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFaxDocumentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentGatewayTimeout creates a GetFaxDocumentGatewayTimeout with default headers values
func NewGetFaxDocumentGatewayTimeout() *GetFaxDocumentGatewayTimeout {
	return &GetFaxDocumentGatewayTimeout{}
}

/*GetFaxDocumentGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetFaxDocumentGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents/{documentId}][%d] getFaxDocumentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFaxDocumentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}