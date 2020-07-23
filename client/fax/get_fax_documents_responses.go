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

// GetFaxDocumentsReader is a Reader for the GetFaxDocuments structure.
type GetFaxDocumentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFaxDocumentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFaxDocumentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFaxDocumentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFaxDocumentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFaxDocumentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFaxDocumentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFaxDocumentsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFaxDocumentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFaxDocumentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFaxDocumentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFaxDocumentsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFaxDocumentsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFaxDocumentsOK creates a GetFaxDocumentsOK with default headers values
func NewGetFaxDocumentsOK() *GetFaxDocumentsOK {
	return &GetFaxDocumentsOK{}
}

/*GetFaxDocumentsOK handles this case with default header values.

successful operation
*/
type GetFaxDocumentsOK struct {
	Payload *models.FaxDocumentEntityListing
}

func (o *GetFaxDocumentsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsOK  %+v", 200, o.Payload)
}

func (o *GetFaxDocumentsOK) GetPayload() *models.FaxDocumentEntityListing {
	return o.Payload
}

func (o *GetFaxDocumentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FaxDocumentEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentsBadRequest creates a GetFaxDocumentsBadRequest with default headers values
func NewGetFaxDocumentsBadRequest() *GetFaxDocumentsBadRequest {
	return &GetFaxDocumentsBadRequest{}
}

/*GetFaxDocumentsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFaxDocumentsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsBadRequest  %+v", 400, o.Payload)
}

func (o *GetFaxDocumentsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentsUnauthorized creates a GetFaxDocumentsUnauthorized with default headers values
func NewGetFaxDocumentsUnauthorized() *GetFaxDocumentsUnauthorized {
	return &GetFaxDocumentsUnauthorized{}
}

/*GetFaxDocumentsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFaxDocumentsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFaxDocumentsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentsForbidden creates a GetFaxDocumentsForbidden with default headers values
func NewGetFaxDocumentsForbidden() *GetFaxDocumentsForbidden {
	return &GetFaxDocumentsForbidden{}
}

/*GetFaxDocumentsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetFaxDocumentsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsForbidden  %+v", 403, o.Payload)
}

func (o *GetFaxDocumentsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentsNotFound creates a GetFaxDocumentsNotFound with default headers values
func NewGetFaxDocumentsNotFound() *GetFaxDocumentsNotFound {
	return &GetFaxDocumentsNotFound{}
}

/*GetFaxDocumentsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetFaxDocumentsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsNotFound  %+v", 404, o.Payload)
}

func (o *GetFaxDocumentsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentsRequestEntityTooLarge creates a GetFaxDocumentsRequestEntityTooLarge with default headers values
func NewGetFaxDocumentsRequestEntityTooLarge() *GetFaxDocumentsRequestEntityTooLarge {
	return &GetFaxDocumentsRequestEntityTooLarge{}
}

/*GetFaxDocumentsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetFaxDocumentsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFaxDocumentsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentsUnsupportedMediaType creates a GetFaxDocumentsUnsupportedMediaType with default headers values
func NewGetFaxDocumentsUnsupportedMediaType() *GetFaxDocumentsUnsupportedMediaType {
	return &GetFaxDocumentsUnsupportedMediaType{}
}

/*GetFaxDocumentsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFaxDocumentsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFaxDocumentsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentsTooManyRequests creates a GetFaxDocumentsTooManyRequests with default headers values
func NewGetFaxDocumentsTooManyRequests() *GetFaxDocumentsTooManyRequests {
	return &GetFaxDocumentsTooManyRequests{}
}

/*GetFaxDocumentsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetFaxDocumentsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFaxDocumentsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentsInternalServerError creates a GetFaxDocumentsInternalServerError with default headers values
func NewGetFaxDocumentsInternalServerError() *GetFaxDocumentsInternalServerError {
	return &GetFaxDocumentsInternalServerError{}
}

/*GetFaxDocumentsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFaxDocumentsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFaxDocumentsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentsServiceUnavailable creates a GetFaxDocumentsServiceUnavailable with default headers values
func NewGetFaxDocumentsServiceUnavailable() *GetFaxDocumentsServiceUnavailable {
	return &GetFaxDocumentsServiceUnavailable{}
}

/*GetFaxDocumentsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFaxDocumentsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFaxDocumentsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFaxDocumentsGatewayTimeout creates a GetFaxDocumentsGatewayTimeout with default headers values
func NewGetFaxDocumentsGatewayTimeout() *GetFaxDocumentsGatewayTimeout {
	return &GetFaxDocumentsGatewayTimeout{}
}

/*GetFaxDocumentsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetFaxDocumentsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFaxDocumentsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/fax/documents][%d] getFaxDocumentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFaxDocumentsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFaxDocumentsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
