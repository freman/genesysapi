// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetUserFavoritesReader is a Reader for the GetUserFavorites structure.
type GetUserFavoritesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserFavoritesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserFavoritesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserFavoritesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserFavoritesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserFavoritesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserFavoritesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetUserFavoritesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUserFavoritesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUserFavoritesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserFavoritesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserFavoritesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUserFavoritesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserFavoritesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserFavoritesOK creates a GetUserFavoritesOK with default headers values
func NewGetUserFavoritesOK() *GetUserFavoritesOK {
	return &GetUserFavoritesOK{}
}

/*
GetUserFavoritesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetUserFavoritesOK struct {
	Payload *models.UserEntityListing
}

// IsSuccess returns true when this get user favorites o k response has a 2xx status code
func (o *GetUserFavoritesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user favorites o k response has a 3xx status code
func (o *GetUserFavoritesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user favorites o k response has a 4xx status code
func (o *GetUserFavoritesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user favorites o k response has a 5xx status code
func (o *GetUserFavoritesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user favorites o k response a status code equal to that given
func (o *GetUserFavoritesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUserFavoritesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesOK  %+v", 200, o.Payload)
}

func (o *GetUserFavoritesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesOK  %+v", 200, o.Payload)
}

func (o *GetUserFavoritesOK) GetPayload() *models.UserEntityListing {
	return o.Payload
}

func (o *GetUserFavoritesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserFavoritesBadRequest creates a GetUserFavoritesBadRequest with default headers values
func NewGetUserFavoritesBadRequest() *GetUserFavoritesBadRequest {
	return &GetUserFavoritesBadRequest{}
}

/*
GetUserFavoritesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUserFavoritesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user favorites bad request response has a 2xx status code
func (o *GetUserFavoritesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user favorites bad request response has a 3xx status code
func (o *GetUserFavoritesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user favorites bad request response has a 4xx status code
func (o *GetUserFavoritesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user favorites bad request response has a 5xx status code
func (o *GetUserFavoritesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get user favorites bad request response a status code equal to that given
func (o *GetUserFavoritesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetUserFavoritesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserFavoritesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserFavoritesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserFavoritesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserFavoritesUnauthorized creates a GetUserFavoritesUnauthorized with default headers values
func NewGetUserFavoritesUnauthorized() *GetUserFavoritesUnauthorized {
	return &GetUserFavoritesUnauthorized{}
}

/*
GetUserFavoritesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUserFavoritesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user favorites unauthorized response has a 2xx status code
func (o *GetUserFavoritesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user favorites unauthorized response has a 3xx status code
func (o *GetUserFavoritesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user favorites unauthorized response has a 4xx status code
func (o *GetUserFavoritesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user favorites unauthorized response has a 5xx status code
func (o *GetUserFavoritesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get user favorites unauthorized response a status code equal to that given
func (o *GetUserFavoritesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetUserFavoritesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserFavoritesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserFavoritesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserFavoritesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserFavoritesForbidden creates a GetUserFavoritesForbidden with default headers values
func NewGetUserFavoritesForbidden() *GetUserFavoritesForbidden {
	return &GetUserFavoritesForbidden{}
}

/*
GetUserFavoritesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetUserFavoritesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user favorites forbidden response has a 2xx status code
func (o *GetUserFavoritesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user favorites forbidden response has a 3xx status code
func (o *GetUserFavoritesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user favorites forbidden response has a 4xx status code
func (o *GetUserFavoritesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user favorites forbidden response has a 5xx status code
func (o *GetUserFavoritesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user favorites forbidden response a status code equal to that given
func (o *GetUserFavoritesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetUserFavoritesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesForbidden  %+v", 403, o.Payload)
}

func (o *GetUserFavoritesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesForbidden  %+v", 403, o.Payload)
}

func (o *GetUserFavoritesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserFavoritesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserFavoritesNotFound creates a GetUserFavoritesNotFound with default headers values
func NewGetUserFavoritesNotFound() *GetUserFavoritesNotFound {
	return &GetUserFavoritesNotFound{}
}

/*
GetUserFavoritesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetUserFavoritesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user favorites not found response has a 2xx status code
func (o *GetUserFavoritesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user favorites not found response has a 3xx status code
func (o *GetUserFavoritesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user favorites not found response has a 4xx status code
func (o *GetUserFavoritesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user favorites not found response has a 5xx status code
func (o *GetUserFavoritesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get user favorites not found response a status code equal to that given
func (o *GetUserFavoritesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetUserFavoritesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesNotFound  %+v", 404, o.Payload)
}

func (o *GetUserFavoritesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesNotFound  %+v", 404, o.Payload)
}

func (o *GetUserFavoritesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserFavoritesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserFavoritesRequestTimeout creates a GetUserFavoritesRequestTimeout with default headers values
func NewGetUserFavoritesRequestTimeout() *GetUserFavoritesRequestTimeout {
	return &GetUserFavoritesRequestTimeout{}
}

/*
GetUserFavoritesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetUserFavoritesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user favorites request timeout response has a 2xx status code
func (o *GetUserFavoritesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user favorites request timeout response has a 3xx status code
func (o *GetUserFavoritesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user favorites request timeout response has a 4xx status code
func (o *GetUserFavoritesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user favorites request timeout response has a 5xx status code
func (o *GetUserFavoritesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get user favorites request timeout response a status code equal to that given
func (o *GetUserFavoritesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetUserFavoritesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserFavoritesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserFavoritesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserFavoritesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserFavoritesRequestEntityTooLarge creates a GetUserFavoritesRequestEntityTooLarge with default headers values
func NewGetUserFavoritesRequestEntityTooLarge() *GetUserFavoritesRequestEntityTooLarge {
	return &GetUserFavoritesRequestEntityTooLarge{}
}

/*
GetUserFavoritesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetUserFavoritesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user favorites request entity too large response has a 2xx status code
func (o *GetUserFavoritesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user favorites request entity too large response has a 3xx status code
func (o *GetUserFavoritesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user favorites request entity too large response has a 4xx status code
func (o *GetUserFavoritesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user favorites request entity too large response has a 5xx status code
func (o *GetUserFavoritesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get user favorites request entity too large response a status code equal to that given
func (o *GetUserFavoritesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetUserFavoritesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserFavoritesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserFavoritesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserFavoritesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserFavoritesUnsupportedMediaType creates a GetUserFavoritesUnsupportedMediaType with default headers values
func NewGetUserFavoritesUnsupportedMediaType() *GetUserFavoritesUnsupportedMediaType {
	return &GetUserFavoritesUnsupportedMediaType{}
}

/*
GetUserFavoritesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUserFavoritesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user favorites unsupported media type response has a 2xx status code
func (o *GetUserFavoritesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user favorites unsupported media type response has a 3xx status code
func (o *GetUserFavoritesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user favorites unsupported media type response has a 4xx status code
func (o *GetUserFavoritesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user favorites unsupported media type response has a 5xx status code
func (o *GetUserFavoritesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get user favorites unsupported media type response a status code equal to that given
func (o *GetUserFavoritesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetUserFavoritesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserFavoritesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserFavoritesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserFavoritesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserFavoritesTooManyRequests creates a GetUserFavoritesTooManyRequests with default headers values
func NewGetUserFavoritesTooManyRequests() *GetUserFavoritesTooManyRequests {
	return &GetUserFavoritesTooManyRequests{}
}

/*
GetUserFavoritesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetUserFavoritesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user favorites too many requests response has a 2xx status code
func (o *GetUserFavoritesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user favorites too many requests response has a 3xx status code
func (o *GetUserFavoritesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user favorites too many requests response has a 4xx status code
func (o *GetUserFavoritesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user favorites too many requests response has a 5xx status code
func (o *GetUserFavoritesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get user favorites too many requests response a status code equal to that given
func (o *GetUserFavoritesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetUserFavoritesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserFavoritesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserFavoritesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserFavoritesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserFavoritesInternalServerError creates a GetUserFavoritesInternalServerError with default headers values
func NewGetUserFavoritesInternalServerError() *GetUserFavoritesInternalServerError {
	return &GetUserFavoritesInternalServerError{}
}

/*
GetUserFavoritesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUserFavoritesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user favorites internal server error response has a 2xx status code
func (o *GetUserFavoritesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user favorites internal server error response has a 3xx status code
func (o *GetUserFavoritesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user favorites internal server error response has a 4xx status code
func (o *GetUserFavoritesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user favorites internal server error response has a 5xx status code
func (o *GetUserFavoritesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user favorites internal server error response a status code equal to that given
func (o *GetUserFavoritesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUserFavoritesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserFavoritesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserFavoritesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserFavoritesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserFavoritesServiceUnavailable creates a GetUserFavoritesServiceUnavailable with default headers values
func NewGetUserFavoritesServiceUnavailable() *GetUserFavoritesServiceUnavailable {
	return &GetUserFavoritesServiceUnavailable{}
}

/*
GetUserFavoritesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUserFavoritesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user favorites service unavailable response has a 2xx status code
func (o *GetUserFavoritesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user favorites service unavailable response has a 3xx status code
func (o *GetUserFavoritesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user favorites service unavailable response has a 4xx status code
func (o *GetUserFavoritesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user favorites service unavailable response has a 5xx status code
func (o *GetUserFavoritesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get user favorites service unavailable response a status code equal to that given
func (o *GetUserFavoritesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetUserFavoritesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserFavoritesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserFavoritesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserFavoritesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserFavoritesGatewayTimeout creates a GetUserFavoritesGatewayTimeout with default headers values
func NewGetUserFavoritesGatewayTimeout() *GetUserFavoritesGatewayTimeout {
	return &GetUserFavoritesGatewayTimeout{}
}

/*
GetUserFavoritesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetUserFavoritesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user favorites gateway timeout response has a 2xx status code
func (o *GetUserFavoritesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user favorites gateway timeout response has a 3xx status code
func (o *GetUserFavoritesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user favorites gateway timeout response has a 4xx status code
func (o *GetUserFavoritesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user favorites gateway timeout response has a 5xx status code
func (o *GetUserFavoritesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get user favorites gateway timeout response a status code equal to that given
func (o *GetUserFavoritesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetUserFavoritesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserFavoritesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/favorites][%d] getUserFavoritesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserFavoritesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserFavoritesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
