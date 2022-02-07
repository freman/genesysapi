// Code generated by go-swagger; DO NOT EDIT.

package organization_authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutOrgauthorizationTrustorCloneduserReader is a Reader for the PutOrgauthorizationTrustorCloneduser structure.
type PutOrgauthorizationTrustorCloneduserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOrgauthorizationTrustorCloneduserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOrgauthorizationTrustorCloneduserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutOrgauthorizationTrustorCloneduserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutOrgauthorizationTrustorCloneduserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutOrgauthorizationTrustorCloneduserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutOrgauthorizationTrustorCloneduserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutOrgauthorizationTrustorCloneduserRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutOrgauthorizationTrustorCloneduserRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutOrgauthorizationTrustorCloneduserUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutOrgauthorizationTrustorCloneduserTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutOrgauthorizationTrustorCloneduserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutOrgauthorizationTrustorCloneduserServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutOrgauthorizationTrustorCloneduserGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutOrgauthorizationTrustorCloneduserOK creates a PutOrgauthorizationTrustorCloneduserOK with default headers values
func NewPutOrgauthorizationTrustorCloneduserOK() *PutOrgauthorizationTrustorCloneduserOK {
	return &PutOrgauthorizationTrustorCloneduserOK{}
}

/*PutOrgauthorizationTrustorCloneduserOK handles this case with default header values.

successful operation
*/
type PutOrgauthorizationTrustorCloneduserOK struct {
	Payload *models.ClonedUser
}

func (o *PutOrgauthorizationTrustorCloneduserOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] putOrgauthorizationTrustorCloneduserOK  %+v", 200, o.Payload)
}

func (o *PutOrgauthorizationTrustorCloneduserOK) GetPayload() *models.ClonedUser {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorCloneduserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClonedUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrustorCloneduserBadRequest creates a PutOrgauthorizationTrustorCloneduserBadRequest with default headers values
func NewPutOrgauthorizationTrustorCloneduserBadRequest() *PutOrgauthorizationTrustorCloneduserBadRequest {
	return &PutOrgauthorizationTrustorCloneduserBadRequest{}
}

/*PutOrgauthorizationTrustorCloneduserBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutOrgauthorizationTrustorCloneduserBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutOrgauthorizationTrustorCloneduserBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] putOrgauthorizationTrustorCloneduserBadRequest  %+v", 400, o.Payload)
}

func (o *PutOrgauthorizationTrustorCloneduserBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorCloneduserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrustorCloneduserUnauthorized creates a PutOrgauthorizationTrustorCloneduserUnauthorized with default headers values
func NewPutOrgauthorizationTrustorCloneduserUnauthorized() *PutOrgauthorizationTrustorCloneduserUnauthorized {
	return &PutOrgauthorizationTrustorCloneduserUnauthorized{}
}

/*PutOrgauthorizationTrustorCloneduserUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutOrgauthorizationTrustorCloneduserUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutOrgauthorizationTrustorCloneduserUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] putOrgauthorizationTrustorCloneduserUnauthorized  %+v", 401, o.Payload)
}

func (o *PutOrgauthorizationTrustorCloneduserUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorCloneduserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrustorCloneduserForbidden creates a PutOrgauthorizationTrustorCloneduserForbidden with default headers values
func NewPutOrgauthorizationTrustorCloneduserForbidden() *PutOrgauthorizationTrustorCloneduserForbidden {
	return &PutOrgauthorizationTrustorCloneduserForbidden{}
}

/*PutOrgauthorizationTrustorCloneduserForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutOrgauthorizationTrustorCloneduserForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutOrgauthorizationTrustorCloneduserForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] putOrgauthorizationTrustorCloneduserForbidden  %+v", 403, o.Payload)
}

func (o *PutOrgauthorizationTrustorCloneduserForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorCloneduserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrustorCloneduserNotFound creates a PutOrgauthorizationTrustorCloneduserNotFound with default headers values
func NewPutOrgauthorizationTrustorCloneduserNotFound() *PutOrgauthorizationTrustorCloneduserNotFound {
	return &PutOrgauthorizationTrustorCloneduserNotFound{}
}

/*PutOrgauthorizationTrustorCloneduserNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutOrgauthorizationTrustorCloneduserNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutOrgauthorizationTrustorCloneduserNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] putOrgauthorizationTrustorCloneduserNotFound  %+v", 404, o.Payload)
}

func (o *PutOrgauthorizationTrustorCloneduserNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorCloneduserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrustorCloneduserRequestTimeout creates a PutOrgauthorizationTrustorCloneduserRequestTimeout with default headers values
func NewPutOrgauthorizationTrustorCloneduserRequestTimeout() *PutOrgauthorizationTrustorCloneduserRequestTimeout {
	return &PutOrgauthorizationTrustorCloneduserRequestTimeout{}
}

/*PutOrgauthorizationTrustorCloneduserRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutOrgauthorizationTrustorCloneduserRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutOrgauthorizationTrustorCloneduserRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] putOrgauthorizationTrustorCloneduserRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutOrgauthorizationTrustorCloneduserRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorCloneduserRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrustorCloneduserRequestEntityTooLarge creates a PutOrgauthorizationTrustorCloneduserRequestEntityTooLarge with default headers values
func NewPutOrgauthorizationTrustorCloneduserRequestEntityTooLarge() *PutOrgauthorizationTrustorCloneduserRequestEntityTooLarge {
	return &PutOrgauthorizationTrustorCloneduserRequestEntityTooLarge{}
}

/*PutOrgauthorizationTrustorCloneduserRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutOrgauthorizationTrustorCloneduserRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutOrgauthorizationTrustorCloneduserRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] putOrgauthorizationTrustorCloneduserRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutOrgauthorizationTrustorCloneduserRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorCloneduserRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrustorCloneduserUnsupportedMediaType creates a PutOrgauthorizationTrustorCloneduserUnsupportedMediaType with default headers values
func NewPutOrgauthorizationTrustorCloneduserUnsupportedMediaType() *PutOrgauthorizationTrustorCloneduserUnsupportedMediaType {
	return &PutOrgauthorizationTrustorCloneduserUnsupportedMediaType{}
}

/*PutOrgauthorizationTrustorCloneduserUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutOrgauthorizationTrustorCloneduserUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutOrgauthorizationTrustorCloneduserUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] putOrgauthorizationTrustorCloneduserUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutOrgauthorizationTrustorCloneduserUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorCloneduserUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrustorCloneduserTooManyRequests creates a PutOrgauthorizationTrustorCloneduserTooManyRequests with default headers values
func NewPutOrgauthorizationTrustorCloneduserTooManyRequests() *PutOrgauthorizationTrustorCloneduserTooManyRequests {
	return &PutOrgauthorizationTrustorCloneduserTooManyRequests{}
}

/*PutOrgauthorizationTrustorCloneduserTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutOrgauthorizationTrustorCloneduserTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutOrgauthorizationTrustorCloneduserTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] putOrgauthorizationTrustorCloneduserTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutOrgauthorizationTrustorCloneduserTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorCloneduserTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrustorCloneduserInternalServerError creates a PutOrgauthorizationTrustorCloneduserInternalServerError with default headers values
func NewPutOrgauthorizationTrustorCloneduserInternalServerError() *PutOrgauthorizationTrustorCloneduserInternalServerError {
	return &PutOrgauthorizationTrustorCloneduserInternalServerError{}
}

/*PutOrgauthorizationTrustorCloneduserInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutOrgauthorizationTrustorCloneduserInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutOrgauthorizationTrustorCloneduserInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] putOrgauthorizationTrustorCloneduserInternalServerError  %+v", 500, o.Payload)
}

func (o *PutOrgauthorizationTrustorCloneduserInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorCloneduserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrustorCloneduserServiceUnavailable creates a PutOrgauthorizationTrustorCloneduserServiceUnavailable with default headers values
func NewPutOrgauthorizationTrustorCloneduserServiceUnavailable() *PutOrgauthorizationTrustorCloneduserServiceUnavailable {
	return &PutOrgauthorizationTrustorCloneduserServiceUnavailable{}
}

/*PutOrgauthorizationTrustorCloneduserServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutOrgauthorizationTrustorCloneduserServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutOrgauthorizationTrustorCloneduserServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] putOrgauthorizationTrustorCloneduserServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutOrgauthorizationTrustorCloneduserServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorCloneduserServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrustorCloneduserGatewayTimeout creates a PutOrgauthorizationTrustorCloneduserGatewayTimeout with default headers values
func NewPutOrgauthorizationTrustorCloneduserGatewayTimeout() *PutOrgauthorizationTrustorCloneduserGatewayTimeout {
	return &PutOrgauthorizationTrustorCloneduserGatewayTimeout{}
}

/*PutOrgauthorizationTrustorCloneduserGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutOrgauthorizationTrustorCloneduserGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutOrgauthorizationTrustorCloneduserGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] putOrgauthorizationTrustorCloneduserGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutOrgauthorizationTrustorCloneduserGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrustorCloneduserGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}