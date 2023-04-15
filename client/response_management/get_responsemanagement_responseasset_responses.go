// Code generated by go-swagger; DO NOT EDIT.

package response_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetResponsemanagementResponseassetReader is a Reader for the GetResponsemanagementResponseasset structure.
type GetResponsemanagementResponseassetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResponsemanagementResponseassetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResponsemanagementResponseassetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetResponsemanagementResponseassetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetResponsemanagementResponseassetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetResponsemanagementResponseassetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetResponsemanagementResponseassetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetResponsemanagementResponseassetRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetResponsemanagementResponseassetRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetResponsemanagementResponseassetUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetResponsemanagementResponseassetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetResponsemanagementResponseassetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetResponsemanagementResponseassetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetResponsemanagementResponseassetGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResponsemanagementResponseassetOK creates a GetResponsemanagementResponseassetOK with default headers values
func NewGetResponsemanagementResponseassetOK() *GetResponsemanagementResponseassetOK {
	return &GetResponsemanagementResponseassetOK{}
}

/*
GetResponsemanagementResponseassetOK describes a response with status code 200, with default header values.

successful operation
*/
type GetResponsemanagementResponseassetOK struct {
	Payload *models.ResponseAsset
}

// IsSuccess returns true when this get responsemanagement responseasset o k response has a 2xx status code
func (o *GetResponsemanagementResponseassetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get responsemanagement responseasset o k response has a 3xx status code
func (o *GetResponsemanagementResponseassetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responseasset o k response has a 4xx status code
func (o *GetResponsemanagementResponseassetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get responsemanagement responseasset o k response has a 5xx status code
func (o *GetResponsemanagementResponseassetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement responseasset o k response a status code equal to that given
func (o *GetResponsemanagementResponseassetOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetResponsemanagementResponseassetOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetOK  %+v", 200, o.Payload)
}

func (o *GetResponsemanagementResponseassetOK) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetOK  %+v", 200, o.Payload)
}

func (o *GetResponsemanagementResponseassetOK) GetPayload() *models.ResponseAsset {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseAsset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetBadRequest creates a GetResponsemanagementResponseassetBadRequest with default headers values
func NewGetResponsemanagementResponseassetBadRequest() *GetResponsemanagementResponseassetBadRequest {
	return &GetResponsemanagementResponseassetBadRequest{}
}

/*
GetResponsemanagementResponseassetBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetResponsemanagementResponseassetBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responseasset bad request response has a 2xx status code
func (o *GetResponsemanagementResponseassetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responseasset bad request response has a 3xx status code
func (o *GetResponsemanagementResponseassetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responseasset bad request response has a 4xx status code
func (o *GetResponsemanagementResponseassetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement responseasset bad request response has a 5xx status code
func (o *GetResponsemanagementResponseassetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement responseasset bad request response a status code equal to that given
func (o *GetResponsemanagementResponseassetBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetResponsemanagementResponseassetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetBadRequest  %+v", 400, o.Payload)
}

func (o *GetResponsemanagementResponseassetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetBadRequest  %+v", 400, o.Payload)
}

func (o *GetResponsemanagementResponseassetBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetUnauthorized creates a GetResponsemanagementResponseassetUnauthorized with default headers values
func NewGetResponsemanagementResponseassetUnauthorized() *GetResponsemanagementResponseassetUnauthorized {
	return &GetResponsemanagementResponseassetUnauthorized{}
}

/*
GetResponsemanagementResponseassetUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetResponsemanagementResponseassetUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responseasset unauthorized response has a 2xx status code
func (o *GetResponsemanagementResponseassetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responseasset unauthorized response has a 3xx status code
func (o *GetResponsemanagementResponseassetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responseasset unauthorized response has a 4xx status code
func (o *GetResponsemanagementResponseassetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement responseasset unauthorized response has a 5xx status code
func (o *GetResponsemanagementResponseassetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement responseasset unauthorized response a status code equal to that given
func (o *GetResponsemanagementResponseassetUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetResponsemanagementResponseassetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetUnauthorized  %+v", 401, o.Payload)
}

func (o *GetResponsemanagementResponseassetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetUnauthorized  %+v", 401, o.Payload)
}

func (o *GetResponsemanagementResponseassetUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetForbidden creates a GetResponsemanagementResponseassetForbidden with default headers values
func NewGetResponsemanagementResponseassetForbidden() *GetResponsemanagementResponseassetForbidden {
	return &GetResponsemanagementResponseassetForbidden{}
}

/*
GetResponsemanagementResponseassetForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetResponsemanagementResponseassetForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responseasset forbidden response has a 2xx status code
func (o *GetResponsemanagementResponseassetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responseasset forbidden response has a 3xx status code
func (o *GetResponsemanagementResponseassetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responseasset forbidden response has a 4xx status code
func (o *GetResponsemanagementResponseassetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement responseasset forbidden response has a 5xx status code
func (o *GetResponsemanagementResponseassetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement responseasset forbidden response a status code equal to that given
func (o *GetResponsemanagementResponseassetForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetResponsemanagementResponseassetForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetForbidden  %+v", 403, o.Payload)
}

func (o *GetResponsemanagementResponseassetForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetForbidden  %+v", 403, o.Payload)
}

func (o *GetResponsemanagementResponseassetForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetNotFound creates a GetResponsemanagementResponseassetNotFound with default headers values
func NewGetResponsemanagementResponseassetNotFound() *GetResponsemanagementResponseassetNotFound {
	return &GetResponsemanagementResponseassetNotFound{}
}

/*
GetResponsemanagementResponseassetNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetResponsemanagementResponseassetNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responseasset not found response has a 2xx status code
func (o *GetResponsemanagementResponseassetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responseasset not found response has a 3xx status code
func (o *GetResponsemanagementResponseassetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responseasset not found response has a 4xx status code
func (o *GetResponsemanagementResponseassetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement responseasset not found response has a 5xx status code
func (o *GetResponsemanagementResponseassetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement responseasset not found response a status code equal to that given
func (o *GetResponsemanagementResponseassetNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetResponsemanagementResponseassetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetNotFound  %+v", 404, o.Payload)
}

func (o *GetResponsemanagementResponseassetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetNotFound  %+v", 404, o.Payload)
}

func (o *GetResponsemanagementResponseassetNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetRequestTimeout creates a GetResponsemanagementResponseassetRequestTimeout with default headers values
func NewGetResponsemanagementResponseassetRequestTimeout() *GetResponsemanagementResponseassetRequestTimeout {
	return &GetResponsemanagementResponseassetRequestTimeout{}
}

/*
GetResponsemanagementResponseassetRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetResponsemanagementResponseassetRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responseasset request timeout response has a 2xx status code
func (o *GetResponsemanagementResponseassetRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responseasset request timeout response has a 3xx status code
func (o *GetResponsemanagementResponseassetRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responseasset request timeout response has a 4xx status code
func (o *GetResponsemanagementResponseassetRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement responseasset request timeout response has a 5xx status code
func (o *GetResponsemanagementResponseassetRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement responseasset request timeout response a status code equal to that given
func (o *GetResponsemanagementResponseassetRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetResponsemanagementResponseassetRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetResponsemanagementResponseassetRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetResponsemanagementResponseassetRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetRequestEntityTooLarge creates a GetResponsemanagementResponseassetRequestEntityTooLarge with default headers values
func NewGetResponsemanagementResponseassetRequestEntityTooLarge() *GetResponsemanagementResponseassetRequestEntityTooLarge {
	return &GetResponsemanagementResponseassetRequestEntityTooLarge{}
}

/*
GetResponsemanagementResponseassetRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetResponsemanagementResponseassetRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responseasset request entity too large response has a 2xx status code
func (o *GetResponsemanagementResponseassetRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responseasset request entity too large response has a 3xx status code
func (o *GetResponsemanagementResponseassetRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responseasset request entity too large response has a 4xx status code
func (o *GetResponsemanagementResponseassetRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement responseasset request entity too large response has a 5xx status code
func (o *GetResponsemanagementResponseassetRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement responseasset request entity too large response a status code equal to that given
func (o *GetResponsemanagementResponseassetRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetResponsemanagementResponseassetRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetResponsemanagementResponseassetRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetResponsemanagementResponseassetRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetUnsupportedMediaType creates a GetResponsemanagementResponseassetUnsupportedMediaType with default headers values
func NewGetResponsemanagementResponseassetUnsupportedMediaType() *GetResponsemanagementResponseassetUnsupportedMediaType {
	return &GetResponsemanagementResponseassetUnsupportedMediaType{}
}

/*
GetResponsemanagementResponseassetUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetResponsemanagementResponseassetUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responseasset unsupported media type response has a 2xx status code
func (o *GetResponsemanagementResponseassetUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responseasset unsupported media type response has a 3xx status code
func (o *GetResponsemanagementResponseassetUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responseasset unsupported media type response has a 4xx status code
func (o *GetResponsemanagementResponseassetUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement responseasset unsupported media type response has a 5xx status code
func (o *GetResponsemanagementResponseassetUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement responseasset unsupported media type response a status code equal to that given
func (o *GetResponsemanagementResponseassetUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetResponsemanagementResponseassetUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetResponsemanagementResponseassetUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetResponsemanagementResponseassetUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetTooManyRequests creates a GetResponsemanagementResponseassetTooManyRequests with default headers values
func NewGetResponsemanagementResponseassetTooManyRequests() *GetResponsemanagementResponseassetTooManyRequests {
	return &GetResponsemanagementResponseassetTooManyRequests{}
}

/*
GetResponsemanagementResponseassetTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetResponsemanagementResponseassetTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responseasset too many requests response has a 2xx status code
func (o *GetResponsemanagementResponseassetTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responseasset too many requests response has a 3xx status code
func (o *GetResponsemanagementResponseassetTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responseasset too many requests response has a 4xx status code
func (o *GetResponsemanagementResponseassetTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement responseasset too many requests response has a 5xx status code
func (o *GetResponsemanagementResponseassetTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement responseasset too many requests response a status code equal to that given
func (o *GetResponsemanagementResponseassetTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetResponsemanagementResponseassetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetResponsemanagementResponseassetTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetResponsemanagementResponseassetTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetInternalServerError creates a GetResponsemanagementResponseassetInternalServerError with default headers values
func NewGetResponsemanagementResponseassetInternalServerError() *GetResponsemanagementResponseassetInternalServerError {
	return &GetResponsemanagementResponseassetInternalServerError{}
}

/*
GetResponsemanagementResponseassetInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetResponsemanagementResponseassetInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responseasset internal server error response has a 2xx status code
func (o *GetResponsemanagementResponseassetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responseasset internal server error response has a 3xx status code
func (o *GetResponsemanagementResponseassetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responseasset internal server error response has a 4xx status code
func (o *GetResponsemanagementResponseassetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get responsemanagement responseasset internal server error response has a 5xx status code
func (o *GetResponsemanagementResponseassetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get responsemanagement responseasset internal server error response a status code equal to that given
func (o *GetResponsemanagementResponseassetInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetResponsemanagementResponseassetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetInternalServerError  %+v", 500, o.Payload)
}

func (o *GetResponsemanagementResponseassetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetInternalServerError  %+v", 500, o.Payload)
}

func (o *GetResponsemanagementResponseassetInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetServiceUnavailable creates a GetResponsemanagementResponseassetServiceUnavailable with default headers values
func NewGetResponsemanagementResponseassetServiceUnavailable() *GetResponsemanagementResponseassetServiceUnavailable {
	return &GetResponsemanagementResponseassetServiceUnavailable{}
}

/*
GetResponsemanagementResponseassetServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetResponsemanagementResponseassetServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responseasset service unavailable response has a 2xx status code
func (o *GetResponsemanagementResponseassetServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responseasset service unavailable response has a 3xx status code
func (o *GetResponsemanagementResponseassetServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responseasset service unavailable response has a 4xx status code
func (o *GetResponsemanagementResponseassetServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get responsemanagement responseasset service unavailable response has a 5xx status code
func (o *GetResponsemanagementResponseassetServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get responsemanagement responseasset service unavailable response a status code equal to that given
func (o *GetResponsemanagementResponseassetServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetResponsemanagementResponseassetServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetResponsemanagementResponseassetServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetResponsemanagementResponseassetServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetGatewayTimeout creates a GetResponsemanagementResponseassetGatewayTimeout with default headers values
func NewGetResponsemanagementResponseassetGatewayTimeout() *GetResponsemanagementResponseassetGatewayTimeout {
	return &GetResponsemanagementResponseassetGatewayTimeout{}
}

/*
GetResponsemanagementResponseassetGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetResponsemanagementResponseassetGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responseasset gateway timeout response has a 2xx status code
func (o *GetResponsemanagementResponseassetGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responseasset gateway timeout response has a 3xx status code
func (o *GetResponsemanagementResponseassetGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responseasset gateway timeout response has a 4xx status code
func (o *GetResponsemanagementResponseassetGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get responsemanagement responseasset gateway timeout response has a 5xx status code
func (o *GetResponsemanagementResponseassetGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get responsemanagement responseasset gateway timeout response a status code equal to that given
func (o *GetResponsemanagementResponseassetGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetResponsemanagementResponseassetGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetResponsemanagementResponseassetGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/{responseAssetId}][%d] getResponsemanagementResponseassetGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetResponsemanagementResponseassetGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
