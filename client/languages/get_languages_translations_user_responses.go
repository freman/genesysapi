// Code generated by go-swagger; DO NOT EDIT.

package languages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetLanguagesTranslationsUserReader is a Reader for the GetLanguagesTranslationsUser structure.
type GetLanguagesTranslationsUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLanguagesTranslationsUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLanguagesTranslationsUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLanguagesTranslationsUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLanguagesTranslationsUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLanguagesTranslationsUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLanguagesTranslationsUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLanguagesTranslationsUserRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLanguagesTranslationsUserUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLanguagesTranslationsUserTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLanguagesTranslationsUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLanguagesTranslationsUserServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLanguagesTranslationsUserGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLanguagesTranslationsUserOK creates a GetLanguagesTranslationsUserOK with default headers values
func NewGetLanguagesTranslationsUserOK() *GetLanguagesTranslationsUserOK {
	return &GetLanguagesTranslationsUserOK{}
}

/*GetLanguagesTranslationsUserOK handles this case with default header values.

successful operation
*/
type GetLanguagesTranslationsUserOK struct {
	Payload map[string]interface{}
}

func (o *GetLanguagesTranslationsUserOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/users/{userId}][%d] getLanguagesTranslationsUserOK  %+v", 200, o.Payload)
}

func (o *GetLanguagesTranslationsUserOK) GetPayload() map[string]interface{} {
	return o.Payload
}

func (o *GetLanguagesTranslationsUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTranslationsUserBadRequest creates a GetLanguagesTranslationsUserBadRequest with default headers values
func NewGetLanguagesTranslationsUserBadRequest() *GetLanguagesTranslationsUserBadRequest {
	return &GetLanguagesTranslationsUserBadRequest{}
}

/*GetLanguagesTranslationsUserBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLanguagesTranslationsUserBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetLanguagesTranslationsUserBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/users/{userId}][%d] getLanguagesTranslationsUserBadRequest  %+v", 400, o.Payload)
}

func (o *GetLanguagesTranslationsUserBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTranslationsUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTranslationsUserUnauthorized creates a GetLanguagesTranslationsUserUnauthorized with default headers values
func NewGetLanguagesTranslationsUserUnauthorized() *GetLanguagesTranslationsUserUnauthorized {
	return &GetLanguagesTranslationsUserUnauthorized{}
}

/*GetLanguagesTranslationsUserUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLanguagesTranslationsUserUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetLanguagesTranslationsUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/users/{userId}][%d] getLanguagesTranslationsUserUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLanguagesTranslationsUserUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTranslationsUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTranslationsUserForbidden creates a GetLanguagesTranslationsUserForbidden with default headers values
func NewGetLanguagesTranslationsUserForbidden() *GetLanguagesTranslationsUserForbidden {
	return &GetLanguagesTranslationsUserForbidden{}
}

/*GetLanguagesTranslationsUserForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetLanguagesTranslationsUserForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetLanguagesTranslationsUserForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/users/{userId}][%d] getLanguagesTranslationsUserForbidden  %+v", 403, o.Payload)
}

func (o *GetLanguagesTranslationsUserForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTranslationsUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTranslationsUserNotFound creates a GetLanguagesTranslationsUserNotFound with default headers values
func NewGetLanguagesTranslationsUserNotFound() *GetLanguagesTranslationsUserNotFound {
	return &GetLanguagesTranslationsUserNotFound{}
}

/*GetLanguagesTranslationsUserNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetLanguagesTranslationsUserNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetLanguagesTranslationsUserNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/users/{userId}][%d] getLanguagesTranslationsUserNotFound  %+v", 404, o.Payload)
}

func (o *GetLanguagesTranslationsUserNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTranslationsUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTranslationsUserRequestEntityTooLarge creates a GetLanguagesTranslationsUserRequestEntityTooLarge with default headers values
func NewGetLanguagesTranslationsUserRequestEntityTooLarge() *GetLanguagesTranslationsUserRequestEntityTooLarge {
	return &GetLanguagesTranslationsUserRequestEntityTooLarge{}
}

/*GetLanguagesTranslationsUserRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetLanguagesTranslationsUserRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetLanguagesTranslationsUserRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/users/{userId}][%d] getLanguagesTranslationsUserRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLanguagesTranslationsUserRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTranslationsUserRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTranslationsUserUnsupportedMediaType creates a GetLanguagesTranslationsUserUnsupportedMediaType with default headers values
func NewGetLanguagesTranslationsUserUnsupportedMediaType() *GetLanguagesTranslationsUserUnsupportedMediaType {
	return &GetLanguagesTranslationsUserUnsupportedMediaType{}
}

/*GetLanguagesTranslationsUserUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLanguagesTranslationsUserUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetLanguagesTranslationsUserUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/users/{userId}][%d] getLanguagesTranslationsUserUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLanguagesTranslationsUserUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTranslationsUserUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTranslationsUserTooManyRequests creates a GetLanguagesTranslationsUserTooManyRequests with default headers values
func NewGetLanguagesTranslationsUserTooManyRequests() *GetLanguagesTranslationsUserTooManyRequests {
	return &GetLanguagesTranslationsUserTooManyRequests{}
}

/*GetLanguagesTranslationsUserTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetLanguagesTranslationsUserTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetLanguagesTranslationsUserTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/users/{userId}][%d] getLanguagesTranslationsUserTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLanguagesTranslationsUserTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTranslationsUserTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTranslationsUserInternalServerError creates a GetLanguagesTranslationsUserInternalServerError with default headers values
func NewGetLanguagesTranslationsUserInternalServerError() *GetLanguagesTranslationsUserInternalServerError {
	return &GetLanguagesTranslationsUserInternalServerError{}
}

/*GetLanguagesTranslationsUserInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLanguagesTranslationsUserInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetLanguagesTranslationsUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/users/{userId}][%d] getLanguagesTranslationsUserInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLanguagesTranslationsUserInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTranslationsUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTranslationsUserServiceUnavailable creates a GetLanguagesTranslationsUserServiceUnavailable with default headers values
func NewGetLanguagesTranslationsUserServiceUnavailable() *GetLanguagesTranslationsUserServiceUnavailable {
	return &GetLanguagesTranslationsUserServiceUnavailable{}
}

/*GetLanguagesTranslationsUserServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLanguagesTranslationsUserServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetLanguagesTranslationsUserServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/users/{userId}][%d] getLanguagesTranslationsUserServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLanguagesTranslationsUserServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTranslationsUserServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTranslationsUserGatewayTimeout creates a GetLanguagesTranslationsUserGatewayTimeout with default headers values
func NewGetLanguagesTranslationsUserGatewayTimeout() *GetLanguagesTranslationsUserGatewayTimeout {
	return &GetLanguagesTranslationsUserGatewayTimeout{}
}

/*GetLanguagesTranslationsUserGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetLanguagesTranslationsUserGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLanguagesTranslationsUserGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/users/{userId}][%d] getLanguagesTranslationsUserGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLanguagesTranslationsUserGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTranslationsUserGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
