// Code generated by go-swagger; DO NOT EDIT.

package external_contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutExternalcontactsRelationshipReader is a Reader for the PutExternalcontactsRelationship structure.
type PutExternalcontactsRelationshipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutExternalcontactsRelationshipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutExternalcontactsRelationshipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutExternalcontactsRelationshipBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutExternalcontactsRelationshipUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutExternalcontactsRelationshipForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutExternalcontactsRelationshipNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutExternalcontactsRelationshipRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutExternalcontactsRelationshipRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutExternalcontactsRelationshipUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPutExternalcontactsRelationshipUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutExternalcontactsRelationshipTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutExternalcontactsRelationshipInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutExternalcontactsRelationshipServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutExternalcontactsRelationshipGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutExternalcontactsRelationshipOK creates a PutExternalcontactsRelationshipOK with default headers values
func NewPutExternalcontactsRelationshipOK() *PutExternalcontactsRelationshipOK {
	return &PutExternalcontactsRelationshipOK{}
}

/*
PutExternalcontactsRelationshipOK describes a response with status code 200, with default header values.

successful operation
*/
type PutExternalcontactsRelationshipOK struct {
	Payload *models.Relationship
}

// IsSuccess returns true when this put externalcontacts relationship o k response has a 2xx status code
func (o *PutExternalcontactsRelationshipOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put externalcontacts relationship o k response has a 3xx status code
func (o *PutExternalcontactsRelationshipOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put externalcontacts relationship o k response has a 4xx status code
func (o *PutExternalcontactsRelationshipOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put externalcontacts relationship o k response has a 5xx status code
func (o *PutExternalcontactsRelationshipOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put externalcontacts relationship o k response a status code equal to that given
func (o *PutExternalcontactsRelationshipOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutExternalcontactsRelationshipOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipOK  %+v", 200, o.Payload)
}

func (o *PutExternalcontactsRelationshipOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipOK  %+v", 200, o.Payload)
}

func (o *PutExternalcontactsRelationshipOK) GetPayload() *models.Relationship {
	return o.Payload
}

func (o *PutExternalcontactsRelationshipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Relationship)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsRelationshipBadRequest creates a PutExternalcontactsRelationshipBadRequest with default headers values
func NewPutExternalcontactsRelationshipBadRequest() *PutExternalcontactsRelationshipBadRequest {
	return &PutExternalcontactsRelationshipBadRequest{}
}

/*
PutExternalcontactsRelationshipBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutExternalcontactsRelationshipBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put externalcontacts relationship bad request response has a 2xx status code
func (o *PutExternalcontactsRelationshipBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put externalcontacts relationship bad request response has a 3xx status code
func (o *PutExternalcontactsRelationshipBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put externalcontacts relationship bad request response has a 4xx status code
func (o *PutExternalcontactsRelationshipBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put externalcontacts relationship bad request response has a 5xx status code
func (o *PutExternalcontactsRelationshipBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put externalcontacts relationship bad request response a status code equal to that given
func (o *PutExternalcontactsRelationshipBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutExternalcontactsRelationshipBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipBadRequest  %+v", 400, o.Payload)
}

func (o *PutExternalcontactsRelationshipBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipBadRequest  %+v", 400, o.Payload)
}

func (o *PutExternalcontactsRelationshipBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsRelationshipBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsRelationshipUnauthorized creates a PutExternalcontactsRelationshipUnauthorized with default headers values
func NewPutExternalcontactsRelationshipUnauthorized() *PutExternalcontactsRelationshipUnauthorized {
	return &PutExternalcontactsRelationshipUnauthorized{}
}

/*
PutExternalcontactsRelationshipUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutExternalcontactsRelationshipUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put externalcontacts relationship unauthorized response has a 2xx status code
func (o *PutExternalcontactsRelationshipUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put externalcontacts relationship unauthorized response has a 3xx status code
func (o *PutExternalcontactsRelationshipUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put externalcontacts relationship unauthorized response has a 4xx status code
func (o *PutExternalcontactsRelationshipUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put externalcontacts relationship unauthorized response has a 5xx status code
func (o *PutExternalcontactsRelationshipUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put externalcontacts relationship unauthorized response a status code equal to that given
func (o *PutExternalcontactsRelationshipUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutExternalcontactsRelationshipUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipUnauthorized  %+v", 401, o.Payload)
}

func (o *PutExternalcontactsRelationshipUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipUnauthorized  %+v", 401, o.Payload)
}

func (o *PutExternalcontactsRelationshipUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsRelationshipUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsRelationshipForbidden creates a PutExternalcontactsRelationshipForbidden with default headers values
func NewPutExternalcontactsRelationshipForbidden() *PutExternalcontactsRelationshipForbidden {
	return &PutExternalcontactsRelationshipForbidden{}
}

/*
PutExternalcontactsRelationshipForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutExternalcontactsRelationshipForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put externalcontacts relationship forbidden response has a 2xx status code
func (o *PutExternalcontactsRelationshipForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put externalcontacts relationship forbidden response has a 3xx status code
func (o *PutExternalcontactsRelationshipForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put externalcontacts relationship forbidden response has a 4xx status code
func (o *PutExternalcontactsRelationshipForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put externalcontacts relationship forbidden response has a 5xx status code
func (o *PutExternalcontactsRelationshipForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put externalcontacts relationship forbidden response a status code equal to that given
func (o *PutExternalcontactsRelationshipForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutExternalcontactsRelationshipForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipForbidden  %+v", 403, o.Payload)
}

func (o *PutExternalcontactsRelationshipForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipForbidden  %+v", 403, o.Payload)
}

func (o *PutExternalcontactsRelationshipForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsRelationshipForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsRelationshipNotFound creates a PutExternalcontactsRelationshipNotFound with default headers values
func NewPutExternalcontactsRelationshipNotFound() *PutExternalcontactsRelationshipNotFound {
	return &PutExternalcontactsRelationshipNotFound{}
}

/*
PutExternalcontactsRelationshipNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutExternalcontactsRelationshipNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put externalcontacts relationship not found response has a 2xx status code
func (o *PutExternalcontactsRelationshipNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put externalcontacts relationship not found response has a 3xx status code
func (o *PutExternalcontactsRelationshipNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put externalcontacts relationship not found response has a 4xx status code
func (o *PutExternalcontactsRelationshipNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put externalcontacts relationship not found response has a 5xx status code
func (o *PutExternalcontactsRelationshipNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put externalcontacts relationship not found response a status code equal to that given
func (o *PutExternalcontactsRelationshipNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutExternalcontactsRelationshipNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipNotFound  %+v", 404, o.Payload)
}

func (o *PutExternalcontactsRelationshipNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipNotFound  %+v", 404, o.Payload)
}

func (o *PutExternalcontactsRelationshipNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsRelationshipNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsRelationshipRequestTimeout creates a PutExternalcontactsRelationshipRequestTimeout with default headers values
func NewPutExternalcontactsRelationshipRequestTimeout() *PutExternalcontactsRelationshipRequestTimeout {
	return &PutExternalcontactsRelationshipRequestTimeout{}
}

/*
PutExternalcontactsRelationshipRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutExternalcontactsRelationshipRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put externalcontacts relationship request timeout response has a 2xx status code
func (o *PutExternalcontactsRelationshipRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put externalcontacts relationship request timeout response has a 3xx status code
func (o *PutExternalcontactsRelationshipRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put externalcontacts relationship request timeout response has a 4xx status code
func (o *PutExternalcontactsRelationshipRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put externalcontacts relationship request timeout response has a 5xx status code
func (o *PutExternalcontactsRelationshipRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put externalcontacts relationship request timeout response a status code equal to that given
func (o *PutExternalcontactsRelationshipRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutExternalcontactsRelationshipRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutExternalcontactsRelationshipRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutExternalcontactsRelationshipRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsRelationshipRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsRelationshipRequestEntityTooLarge creates a PutExternalcontactsRelationshipRequestEntityTooLarge with default headers values
func NewPutExternalcontactsRelationshipRequestEntityTooLarge() *PutExternalcontactsRelationshipRequestEntityTooLarge {
	return &PutExternalcontactsRelationshipRequestEntityTooLarge{}
}

/*
PutExternalcontactsRelationshipRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutExternalcontactsRelationshipRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put externalcontacts relationship request entity too large response has a 2xx status code
func (o *PutExternalcontactsRelationshipRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put externalcontacts relationship request entity too large response has a 3xx status code
func (o *PutExternalcontactsRelationshipRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put externalcontacts relationship request entity too large response has a 4xx status code
func (o *PutExternalcontactsRelationshipRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put externalcontacts relationship request entity too large response has a 5xx status code
func (o *PutExternalcontactsRelationshipRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put externalcontacts relationship request entity too large response a status code equal to that given
func (o *PutExternalcontactsRelationshipRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutExternalcontactsRelationshipRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutExternalcontactsRelationshipRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutExternalcontactsRelationshipRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsRelationshipRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsRelationshipUnsupportedMediaType creates a PutExternalcontactsRelationshipUnsupportedMediaType with default headers values
func NewPutExternalcontactsRelationshipUnsupportedMediaType() *PutExternalcontactsRelationshipUnsupportedMediaType {
	return &PutExternalcontactsRelationshipUnsupportedMediaType{}
}

/*
PutExternalcontactsRelationshipUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutExternalcontactsRelationshipUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put externalcontacts relationship unsupported media type response has a 2xx status code
func (o *PutExternalcontactsRelationshipUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put externalcontacts relationship unsupported media type response has a 3xx status code
func (o *PutExternalcontactsRelationshipUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put externalcontacts relationship unsupported media type response has a 4xx status code
func (o *PutExternalcontactsRelationshipUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put externalcontacts relationship unsupported media type response has a 5xx status code
func (o *PutExternalcontactsRelationshipUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put externalcontacts relationship unsupported media type response a status code equal to that given
func (o *PutExternalcontactsRelationshipUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutExternalcontactsRelationshipUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutExternalcontactsRelationshipUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutExternalcontactsRelationshipUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsRelationshipUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsRelationshipUnprocessableEntity creates a PutExternalcontactsRelationshipUnprocessableEntity with default headers values
func NewPutExternalcontactsRelationshipUnprocessableEntity() *PutExternalcontactsRelationshipUnprocessableEntity {
	return &PutExternalcontactsRelationshipUnprocessableEntity{}
}

/*
PutExternalcontactsRelationshipUnprocessableEntity describes a response with status code 422, with default header values.

PutExternalcontactsRelationshipUnprocessableEntity put externalcontacts relationship unprocessable entity
*/
type PutExternalcontactsRelationshipUnprocessableEntity struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put externalcontacts relationship unprocessable entity response has a 2xx status code
func (o *PutExternalcontactsRelationshipUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put externalcontacts relationship unprocessable entity response has a 3xx status code
func (o *PutExternalcontactsRelationshipUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put externalcontacts relationship unprocessable entity response has a 4xx status code
func (o *PutExternalcontactsRelationshipUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this put externalcontacts relationship unprocessable entity response has a 5xx status code
func (o *PutExternalcontactsRelationshipUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this put externalcontacts relationship unprocessable entity response a status code equal to that given
func (o *PutExternalcontactsRelationshipUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *PutExternalcontactsRelationshipUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PutExternalcontactsRelationshipUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PutExternalcontactsRelationshipUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsRelationshipUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsRelationshipTooManyRequests creates a PutExternalcontactsRelationshipTooManyRequests with default headers values
func NewPutExternalcontactsRelationshipTooManyRequests() *PutExternalcontactsRelationshipTooManyRequests {
	return &PutExternalcontactsRelationshipTooManyRequests{}
}

/*
PutExternalcontactsRelationshipTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutExternalcontactsRelationshipTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put externalcontacts relationship too many requests response has a 2xx status code
func (o *PutExternalcontactsRelationshipTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put externalcontacts relationship too many requests response has a 3xx status code
func (o *PutExternalcontactsRelationshipTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put externalcontacts relationship too many requests response has a 4xx status code
func (o *PutExternalcontactsRelationshipTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put externalcontacts relationship too many requests response has a 5xx status code
func (o *PutExternalcontactsRelationshipTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put externalcontacts relationship too many requests response a status code equal to that given
func (o *PutExternalcontactsRelationshipTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutExternalcontactsRelationshipTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutExternalcontactsRelationshipTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutExternalcontactsRelationshipTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsRelationshipTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsRelationshipInternalServerError creates a PutExternalcontactsRelationshipInternalServerError with default headers values
func NewPutExternalcontactsRelationshipInternalServerError() *PutExternalcontactsRelationshipInternalServerError {
	return &PutExternalcontactsRelationshipInternalServerError{}
}

/*
PutExternalcontactsRelationshipInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutExternalcontactsRelationshipInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put externalcontacts relationship internal server error response has a 2xx status code
func (o *PutExternalcontactsRelationshipInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put externalcontacts relationship internal server error response has a 3xx status code
func (o *PutExternalcontactsRelationshipInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put externalcontacts relationship internal server error response has a 4xx status code
func (o *PutExternalcontactsRelationshipInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put externalcontacts relationship internal server error response has a 5xx status code
func (o *PutExternalcontactsRelationshipInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put externalcontacts relationship internal server error response a status code equal to that given
func (o *PutExternalcontactsRelationshipInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutExternalcontactsRelationshipInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipInternalServerError  %+v", 500, o.Payload)
}

func (o *PutExternalcontactsRelationshipInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipInternalServerError  %+v", 500, o.Payload)
}

func (o *PutExternalcontactsRelationshipInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsRelationshipInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsRelationshipServiceUnavailable creates a PutExternalcontactsRelationshipServiceUnavailable with default headers values
func NewPutExternalcontactsRelationshipServiceUnavailable() *PutExternalcontactsRelationshipServiceUnavailable {
	return &PutExternalcontactsRelationshipServiceUnavailable{}
}

/*
PutExternalcontactsRelationshipServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutExternalcontactsRelationshipServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put externalcontacts relationship service unavailable response has a 2xx status code
func (o *PutExternalcontactsRelationshipServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put externalcontacts relationship service unavailable response has a 3xx status code
func (o *PutExternalcontactsRelationshipServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put externalcontacts relationship service unavailable response has a 4xx status code
func (o *PutExternalcontactsRelationshipServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put externalcontacts relationship service unavailable response has a 5xx status code
func (o *PutExternalcontactsRelationshipServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put externalcontacts relationship service unavailable response a status code equal to that given
func (o *PutExternalcontactsRelationshipServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutExternalcontactsRelationshipServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutExternalcontactsRelationshipServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutExternalcontactsRelationshipServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsRelationshipServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutExternalcontactsRelationshipGatewayTimeout creates a PutExternalcontactsRelationshipGatewayTimeout with default headers values
func NewPutExternalcontactsRelationshipGatewayTimeout() *PutExternalcontactsRelationshipGatewayTimeout {
	return &PutExternalcontactsRelationshipGatewayTimeout{}
}

/*
PutExternalcontactsRelationshipGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutExternalcontactsRelationshipGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put externalcontacts relationship gateway timeout response has a 2xx status code
func (o *PutExternalcontactsRelationshipGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put externalcontacts relationship gateway timeout response has a 3xx status code
func (o *PutExternalcontactsRelationshipGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put externalcontacts relationship gateway timeout response has a 4xx status code
func (o *PutExternalcontactsRelationshipGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put externalcontacts relationship gateway timeout response has a 5xx status code
func (o *PutExternalcontactsRelationshipGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put externalcontacts relationship gateway timeout response a status code equal to that given
func (o *PutExternalcontactsRelationshipGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutExternalcontactsRelationshipGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutExternalcontactsRelationshipGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/externalcontacts/relationships/{relationshipId}][%d] putExternalcontactsRelationshipGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutExternalcontactsRelationshipGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutExternalcontactsRelationshipGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
