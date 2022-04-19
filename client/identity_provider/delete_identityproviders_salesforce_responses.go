// Code generated by go-swagger; DO NOT EDIT.

package identity_provider

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteIdentityprovidersSalesforceReader is a Reader for the DeleteIdentityprovidersSalesforce structure.
type DeleteIdentityprovidersSalesforceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIdentityprovidersSalesforceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIdentityprovidersSalesforceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteIdentityprovidersSalesforceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteIdentityprovidersSalesforceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteIdentityprovidersSalesforceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteIdentityprovidersSalesforceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteIdentityprovidersSalesforceRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteIdentityprovidersSalesforceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteIdentityprovidersSalesforceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteIdentityprovidersSalesforceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteIdentityprovidersSalesforceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteIdentityprovidersSalesforceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteIdentityprovidersSalesforceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteIdentityprovidersSalesforceOK creates a DeleteIdentityprovidersSalesforceOK with default headers values
func NewDeleteIdentityprovidersSalesforceOK() *DeleteIdentityprovidersSalesforceOK {
	return &DeleteIdentityprovidersSalesforceOK{}
}

/*DeleteIdentityprovidersSalesforceOK handles this case with default header values.

successful operation
*/
type DeleteIdentityprovidersSalesforceOK struct {
	Payload models.Empty
}

func (o *DeleteIdentityprovidersSalesforceOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/salesforce][%d] deleteIdentityprovidersSalesforceOK  %+v", 200, o.Payload)
}

func (o *DeleteIdentityprovidersSalesforceOK) GetPayload() models.Empty {
	return o.Payload
}

func (o *DeleteIdentityprovidersSalesforceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersSalesforceBadRequest creates a DeleteIdentityprovidersSalesforceBadRequest with default headers values
func NewDeleteIdentityprovidersSalesforceBadRequest() *DeleteIdentityprovidersSalesforceBadRequest {
	return &DeleteIdentityprovidersSalesforceBadRequest{}
}

/*DeleteIdentityprovidersSalesforceBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteIdentityprovidersSalesforceBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteIdentityprovidersSalesforceBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/salesforce][%d] deleteIdentityprovidersSalesforceBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIdentityprovidersSalesforceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersSalesforceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersSalesforceUnauthorized creates a DeleteIdentityprovidersSalesforceUnauthorized with default headers values
func NewDeleteIdentityprovidersSalesforceUnauthorized() *DeleteIdentityprovidersSalesforceUnauthorized {
	return &DeleteIdentityprovidersSalesforceUnauthorized{}
}

/*DeleteIdentityprovidersSalesforceUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteIdentityprovidersSalesforceUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteIdentityprovidersSalesforceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/salesforce][%d] deleteIdentityprovidersSalesforceUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteIdentityprovidersSalesforceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersSalesforceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersSalesforceForbidden creates a DeleteIdentityprovidersSalesforceForbidden with default headers values
func NewDeleteIdentityprovidersSalesforceForbidden() *DeleteIdentityprovidersSalesforceForbidden {
	return &DeleteIdentityprovidersSalesforceForbidden{}
}

/*DeleteIdentityprovidersSalesforceForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteIdentityprovidersSalesforceForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteIdentityprovidersSalesforceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/salesforce][%d] deleteIdentityprovidersSalesforceForbidden  %+v", 403, o.Payload)
}

func (o *DeleteIdentityprovidersSalesforceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersSalesforceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersSalesforceNotFound creates a DeleteIdentityprovidersSalesforceNotFound with default headers values
func NewDeleteIdentityprovidersSalesforceNotFound() *DeleteIdentityprovidersSalesforceNotFound {
	return &DeleteIdentityprovidersSalesforceNotFound{}
}

/*DeleteIdentityprovidersSalesforceNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteIdentityprovidersSalesforceNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteIdentityprovidersSalesforceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/salesforce][%d] deleteIdentityprovidersSalesforceNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIdentityprovidersSalesforceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersSalesforceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersSalesforceRequestTimeout creates a DeleteIdentityprovidersSalesforceRequestTimeout with default headers values
func NewDeleteIdentityprovidersSalesforceRequestTimeout() *DeleteIdentityprovidersSalesforceRequestTimeout {
	return &DeleteIdentityprovidersSalesforceRequestTimeout{}
}

/*DeleteIdentityprovidersSalesforceRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteIdentityprovidersSalesforceRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteIdentityprovidersSalesforceRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/salesforce][%d] deleteIdentityprovidersSalesforceRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteIdentityprovidersSalesforceRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersSalesforceRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersSalesforceRequestEntityTooLarge creates a DeleteIdentityprovidersSalesforceRequestEntityTooLarge with default headers values
func NewDeleteIdentityprovidersSalesforceRequestEntityTooLarge() *DeleteIdentityprovidersSalesforceRequestEntityTooLarge {
	return &DeleteIdentityprovidersSalesforceRequestEntityTooLarge{}
}

/*DeleteIdentityprovidersSalesforceRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteIdentityprovidersSalesforceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteIdentityprovidersSalesforceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/salesforce][%d] deleteIdentityprovidersSalesforceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteIdentityprovidersSalesforceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersSalesforceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersSalesforceUnsupportedMediaType creates a DeleteIdentityprovidersSalesforceUnsupportedMediaType with default headers values
func NewDeleteIdentityprovidersSalesforceUnsupportedMediaType() *DeleteIdentityprovidersSalesforceUnsupportedMediaType {
	return &DeleteIdentityprovidersSalesforceUnsupportedMediaType{}
}

/*DeleteIdentityprovidersSalesforceUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteIdentityprovidersSalesforceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteIdentityprovidersSalesforceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/salesforce][%d] deleteIdentityprovidersSalesforceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteIdentityprovidersSalesforceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersSalesforceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersSalesforceTooManyRequests creates a DeleteIdentityprovidersSalesforceTooManyRequests with default headers values
func NewDeleteIdentityprovidersSalesforceTooManyRequests() *DeleteIdentityprovidersSalesforceTooManyRequests {
	return &DeleteIdentityprovidersSalesforceTooManyRequests{}
}

/*DeleteIdentityprovidersSalesforceTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteIdentityprovidersSalesforceTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteIdentityprovidersSalesforceTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/salesforce][%d] deleteIdentityprovidersSalesforceTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteIdentityprovidersSalesforceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersSalesforceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersSalesforceInternalServerError creates a DeleteIdentityprovidersSalesforceInternalServerError with default headers values
func NewDeleteIdentityprovidersSalesforceInternalServerError() *DeleteIdentityprovidersSalesforceInternalServerError {
	return &DeleteIdentityprovidersSalesforceInternalServerError{}
}

/*DeleteIdentityprovidersSalesforceInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteIdentityprovidersSalesforceInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteIdentityprovidersSalesforceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/salesforce][%d] deleteIdentityprovidersSalesforceInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIdentityprovidersSalesforceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersSalesforceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersSalesforceServiceUnavailable creates a DeleteIdentityprovidersSalesforceServiceUnavailable with default headers values
func NewDeleteIdentityprovidersSalesforceServiceUnavailable() *DeleteIdentityprovidersSalesforceServiceUnavailable {
	return &DeleteIdentityprovidersSalesforceServiceUnavailable{}
}

/*DeleteIdentityprovidersSalesforceServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteIdentityprovidersSalesforceServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteIdentityprovidersSalesforceServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/salesforce][%d] deleteIdentityprovidersSalesforceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteIdentityprovidersSalesforceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersSalesforceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersSalesforceGatewayTimeout creates a DeleteIdentityprovidersSalesforceGatewayTimeout with default headers values
func NewDeleteIdentityprovidersSalesforceGatewayTimeout() *DeleteIdentityprovidersSalesforceGatewayTimeout {
	return &DeleteIdentityprovidersSalesforceGatewayTimeout{}
}

/*DeleteIdentityprovidersSalesforceGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteIdentityprovidersSalesforceGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteIdentityprovidersSalesforceGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/salesforce][%d] deleteIdentityprovidersSalesforceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteIdentityprovidersSalesforceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersSalesforceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
