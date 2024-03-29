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

// GetExternalcontactsScanRelationshipsReader is a Reader for the GetExternalcontactsScanRelationships structure.
type GetExternalcontactsScanRelationshipsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalcontactsScanRelationshipsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalcontactsScanRelationshipsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExternalcontactsScanRelationshipsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExternalcontactsScanRelationshipsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExternalcontactsScanRelationshipsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExternalcontactsScanRelationshipsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetExternalcontactsScanRelationshipsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetExternalcontactsScanRelationshipsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetExternalcontactsScanRelationshipsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetExternalcontactsScanRelationshipsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetExternalcontactsScanRelationshipsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExternalcontactsScanRelationshipsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetExternalcontactsScanRelationshipsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetExternalcontactsScanRelationshipsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExternalcontactsScanRelationshipsOK creates a GetExternalcontactsScanRelationshipsOK with default headers values
func NewGetExternalcontactsScanRelationshipsOK() *GetExternalcontactsScanRelationshipsOK {
	return &GetExternalcontactsScanRelationshipsOK{}
}

/*
GetExternalcontactsScanRelationshipsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetExternalcontactsScanRelationshipsOK struct {
	Payload *models.CursorRelationshipListing
}

// IsSuccess returns true when this get externalcontacts scan relationships o k response has a 2xx status code
func (o *GetExternalcontactsScanRelationshipsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get externalcontacts scan relationships o k response has a 3xx status code
func (o *GetExternalcontactsScanRelationshipsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan relationships o k response has a 4xx status code
func (o *GetExternalcontactsScanRelationshipsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts scan relationships o k response has a 5xx status code
func (o *GetExternalcontactsScanRelationshipsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts scan relationships o k response a status code equal to that given
func (o *GetExternalcontactsScanRelationshipsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetExternalcontactsScanRelationshipsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsOK) GetPayload() *models.CursorRelationshipListing {
	return o.Payload
}

func (o *GetExternalcontactsScanRelationshipsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CursorRelationshipListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanRelationshipsBadRequest creates a GetExternalcontactsScanRelationshipsBadRequest with default headers values
func NewGetExternalcontactsScanRelationshipsBadRequest() *GetExternalcontactsScanRelationshipsBadRequest {
	return &GetExternalcontactsScanRelationshipsBadRequest{}
}

/*
GetExternalcontactsScanRelationshipsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetExternalcontactsScanRelationshipsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan relationships bad request response has a 2xx status code
func (o *GetExternalcontactsScanRelationshipsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan relationships bad request response has a 3xx status code
func (o *GetExternalcontactsScanRelationshipsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan relationships bad request response has a 4xx status code
func (o *GetExternalcontactsScanRelationshipsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts scan relationships bad request response has a 5xx status code
func (o *GetExternalcontactsScanRelationshipsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts scan relationships bad request response a status code equal to that given
func (o *GetExternalcontactsScanRelationshipsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetExternalcontactsScanRelationshipsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanRelationshipsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanRelationshipsUnauthorized creates a GetExternalcontactsScanRelationshipsUnauthorized with default headers values
func NewGetExternalcontactsScanRelationshipsUnauthorized() *GetExternalcontactsScanRelationshipsUnauthorized {
	return &GetExternalcontactsScanRelationshipsUnauthorized{}
}

/*
GetExternalcontactsScanRelationshipsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetExternalcontactsScanRelationshipsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan relationships unauthorized response has a 2xx status code
func (o *GetExternalcontactsScanRelationshipsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan relationships unauthorized response has a 3xx status code
func (o *GetExternalcontactsScanRelationshipsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan relationships unauthorized response has a 4xx status code
func (o *GetExternalcontactsScanRelationshipsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts scan relationships unauthorized response has a 5xx status code
func (o *GetExternalcontactsScanRelationshipsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts scan relationships unauthorized response a status code equal to that given
func (o *GetExternalcontactsScanRelationshipsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetExternalcontactsScanRelationshipsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanRelationshipsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanRelationshipsForbidden creates a GetExternalcontactsScanRelationshipsForbidden with default headers values
func NewGetExternalcontactsScanRelationshipsForbidden() *GetExternalcontactsScanRelationshipsForbidden {
	return &GetExternalcontactsScanRelationshipsForbidden{}
}

/*
GetExternalcontactsScanRelationshipsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetExternalcontactsScanRelationshipsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan relationships forbidden response has a 2xx status code
func (o *GetExternalcontactsScanRelationshipsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan relationships forbidden response has a 3xx status code
func (o *GetExternalcontactsScanRelationshipsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan relationships forbidden response has a 4xx status code
func (o *GetExternalcontactsScanRelationshipsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts scan relationships forbidden response has a 5xx status code
func (o *GetExternalcontactsScanRelationshipsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts scan relationships forbidden response a status code equal to that given
func (o *GetExternalcontactsScanRelationshipsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetExternalcontactsScanRelationshipsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanRelationshipsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanRelationshipsNotFound creates a GetExternalcontactsScanRelationshipsNotFound with default headers values
func NewGetExternalcontactsScanRelationshipsNotFound() *GetExternalcontactsScanRelationshipsNotFound {
	return &GetExternalcontactsScanRelationshipsNotFound{}
}

/*
GetExternalcontactsScanRelationshipsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetExternalcontactsScanRelationshipsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan relationships not found response has a 2xx status code
func (o *GetExternalcontactsScanRelationshipsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan relationships not found response has a 3xx status code
func (o *GetExternalcontactsScanRelationshipsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan relationships not found response has a 4xx status code
func (o *GetExternalcontactsScanRelationshipsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts scan relationships not found response has a 5xx status code
func (o *GetExternalcontactsScanRelationshipsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts scan relationships not found response a status code equal to that given
func (o *GetExternalcontactsScanRelationshipsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetExternalcontactsScanRelationshipsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanRelationshipsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanRelationshipsRequestTimeout creates a GetExternalcontactsScanRelationshipsRequestTimeout with default headers values
func NewGetExternalcontactsScanRelationshipsRequestTimeout() *GetExternalcontactsScanRelationshipsRequestTimeout {
	return &GetExternalcontactsScanRelationshipsRequestTimeout{}
}

/*
GetExternalcontactsScanRelationshipsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetExternalcontactsScanRelationshipsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan relationships request timeout response has a 2xx status code
func (o *GetExternalcontactsScanRelationshipsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan relationships request timeout response has a 3xx status code
func (o *GetExternalcontactsScanRelationshipsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan relationships request timeout response has a 4xx status code
func (o *GetExternalcontactsScanRelationshipsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts scan relationships request timeout response has a 5xx status code
func (o *GetExternalcontactsScanRelationshipsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts scan relationships request timeout response a status code equal to that given
func (o *GetExternalcontactsScanRelationshipsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetExternalcontactsScanRelationshipsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanRelationshipsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanRelationshipsRequestEntityTooLarge creates a GetExternalcontactsScanRelationshipsRequestEntityTooLarge with default headers values
func NewGetExternalcontactsScanRelationshipsRequestEntityTooLarge() *GetExternalcontactsScanRelationshipsRequestEntityTooLarge {
	return &GetExternalcontactsScanRelationshipsRequestEntityTooLarge{}
}

/*
GetExternalcontactsScanRelationshipsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetExternalcontactsScanRelationshipsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan relationships request entity too large response has a 2xx status code
func (o *GetExternalcontactsScanRelationshipsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan relationships request entity too large response has a 3xx status code
func (o *GetExternalcontactsScanRelationshipsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan relationships request entity too large response has a 4xx status code
func (o *GetExternalcontactsScanRelationshipsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts scan relationships request entity too large response has a 5xx status code
func (o *GetExternalcontactsScanRelationshipsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts scan relationships request entity too large response a status code equal to that given
func (o *GetExternalcontactsScanRelationshipsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetExternalcontactsScanRelationshipsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanRelationshipsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanRelationshipsUnsupportedMediaType creates a GetExternalcontactsScanRelationshipsUnsupportedMediaType with default headers values
func NewGetExternalcontactsScanRelationshipsUnsupportedMediaType() *GetExternalcontactsScanRelationshipsUnsupportedMediaType {
	return &GetExternalcontactsScanRelationshipsUnsupportedMediaType{}
}

/*
GetExternalcontactsScanRelationshipsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetExternalcontactsScanRelationshipsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan relationships unsupported media type response has a 2xx status code
func (o *GetExternalcontactsScanRelationshipsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan relationships unsupported media type response has a 3xx status code
func (o *GetExternalcontactsScanRelationshipsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan relationships unsupported media type response has a 4xx status code
func (o *GetExternalcontactsScanRelationshipsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts scan relationships unsupported media type response has a 5xx status code
func (o *GetExternalcontactsScanRelationshipsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts scan relationships unsupported media type response a status code equal to that given
func (o *GetExternalcontactsScanRelationshipsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetExternalcontactsScanRelationshipsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanRelationshipsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanRelationshipsUnprocessableEntity creates a GetExternalcontactsScanRelationshipsUnprocessableEntity with default headers values
func NewGetExternalcontactsScanRelationshipsUnprocessableEntity() *GetExternalcontactsScanRelationshipsUnprocessableEntity {
	return &GetExternalcontactsScanRelationshipsUnprocessableEntity{}
}

/*
GetExternalcontactsScanRelationshipsUnprocessableEntity describes a response with status code 422, with default header values.

GetExternalcontactsScanRelationshipsUnprocessableEntity get externalcontacts scan relationships unprocessable entity
*/
type GetExternalcontactsScanRelationshipsUnprocessableEntity struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan relationships unprocessable entity response has a 2xx status code
func (o *GetExternalcontactsScanRelationshipsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan relationships unprocessable entity response has a 3xx status code
func (o *GetExternalcontactsScanRelationshipsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan relationships unprocessable entity response has a 4xx status code
func (o *GetExternalcontactsScanRelationshipsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts scan relationships unprocessable entity response has a 5xx status code
func (o *GetExternalcontactsScanRelationshipsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts scan relationships unprocessable entity response a status code equal to that given
func (o *GetExternalcontactsScanRelationshipsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *GetExternalcontactsScanRelationshipsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanRelationshipsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanRelationshipsTooManyRequests creates a GetExternalcontactsScanRelationshipsTooManyRequests with default headers values
func NewGetExternalcontactsScanRelationshipsTooManyRequests() *GetExternalcontactsScanRelationshipsTooManyRequests {
	return &GetExternalcontactsScanRelationshipsTooManyRequests{}
}

/*
GetExternalcontactsScanRelationshipsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetExternalcontactsScanRelationshipsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan relationships too many requests response has a 2xx status code
func (o *GetExternalcontactsScanRelationshipsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan relationships too many requests response has a 3xx status code
func (o *GetExternalcontactsScanRelationshipsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan relationships too many requests response has a 4xx status code
func (o *GetExternalcontactsScanRelationshipsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts scan relationships too many requests response has a 5xx status code
func (o *GetExternalcontactsScanRelationshipsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts scan relationships too many requests response a status code equal to that given
func (o *GetExternalcontactsScanRelationshipsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetExternalcontactsScanRelationshipsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanRelationshipsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanRelationshipsInternalServerError creates a GetExternalcontactsScanRelationshipsInternalServerError with default headers values
func NewGetExternalcontactsScanRelationshipsInternalServerError() *GetExternalcontactsScanRelationshipsInternalServerError {
	return &GetExternalcontactsScanRelationshipsInternalServerError{}
}

/*
GetExternalcontactsScanRelationshipsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetExternalcontactsScanRelationshipsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan relationships internal server error response has a 2xx status code
func (o *GetExternalcontactsScanRelationshipsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan relationships internal server error response has a 3xx status code
func (o *GetExternalcontactsScanRelationshipsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan relationships internal server error response has a 4xx status code
func (o *GetExternalcontactsScanRelationshipsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts scan relationships internal server error response has a 5xx status code
func (o *GetExternalcontactsScanRelationshipsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts scan relationships internal server error response a status code equal to that given
func (o *GetExternalcontactsScanRelationshipsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetExternalcontactsScanRelationshipsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanRelationshipsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanRelationshipsServiceUnavailable creates a GetExternalcontactsScanRelationshipsServiceUnavailable with default headers values
func NewGetExternalcontactsScanRelationshipsServiceUnavailable() *GetExternalcontactsScanRelationshipsServiceUnavailable {
	return &GetExternalcontactsScanRelationshipsServiceUnavailable{}
}

/*
GetExternalcontactsScanRelationshipsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetExternalcontactsScanRelationshipsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan relationships service unavailable response has a 2xx status code
func (o *GetExternalcontactsScanRelationshipsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan relationships service unavailable response has a 3xx status code
func (o *GetExternalcontactsScanRelationshipsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan relationships service unavailable response has a 4xx status code
func (o *GetExternalcontactsScanRelationshipsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts scan relationships service unavailable response has a 5xx status code
func (o *GetExternalcontactsScanRelationshipsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts scan relationships service unavailable response a status code equal to that given
func (o *GetExternalcontactsScanRelationshipsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetExternalcontactsScanRelationshipsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanRelationshipsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsScanRelationshipsGatewayTimeout creates a GetExternalcontactsScanRelationshipsGatewayTimeout with default headers values
func NewGetExternalcontactsScanRelationshipsGatewayTimeout() *GetExternalcontactsScanRelationshipsGatewayTimeout {
	return &GetExternalcontactsScanRelationshipsGatewayTimeout{}
}

/*
GetExternalcontactsScanRelationshipsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetExternalcontactsScanRelationshipsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts scan relationships gateway timeout response has a 2xx status code
func (o *GetExternalcontactsScanRelationshipsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts scan relationships gateway timeout response has a 3xx status code
func (o *GetExternalcontactsScanRelationshipsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts scan relationships gateway timeout response has a 4xx status code
func (o *GetExternalcontactsScanRelationshipsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts scan relationships gateway timeout response has a 5xx status code
func (o *GetExternalcontactsScanRelationshipsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts scan relationships gateway timeout response a status code equal to that given
func (o *GetExternalcontactsScanRelationshipsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetExternalcontactsScanRelationshipsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/scan/relationships][%d] getExternalcontactsScanRelationshipsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsScanRelationshipsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsScanRelationshipsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
