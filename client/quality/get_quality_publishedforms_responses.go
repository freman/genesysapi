// Code generated by go-swagger; DO NOT EDIT.

package quality

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetQualityPublishedformsReader is a Reader for the GetQualityPublishedforms structure.
type GetQualityPublishedformsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQualityPublishedformsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQualityPublishedformsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetQualityPublishedformsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetQualityPublishedformsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetQualityPublishedformsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetQualityPublishedformsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetQualityPublishedformsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetQualityPublishedformsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetQualityPublishedformsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetQualityPublishedformsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetQualityPublishedformsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetQualityPublishedformsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetQualityPublishedformsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetQualityPublishedformsOK creates a GetQualityPublishedformsOK with default headers values
func NewGetQualityPublishedformsOK() *GetQualityPublishedformsOK {
	return &GetQualityPublishedformsOK{}
}

/*
GetQualityPublishedformsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetQualityPublishedformsOK struct {
	Payload *models.EvaluationFormEntityListing
}

// IsSuccess returns true when this get quality publishedforms o k response has a 2xx status code
func (o *GetQualityPublishedformsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get quality publishedforms o k response has a 3xx status code
func (o *GetQualityPublishedformsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms o k response has a 4xx status code
func (o *GetQualityPublishedformsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality publishedforms o k response has a 5xx status code
func (o *GetQualityPublishedformsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms o k response a status code equal to that given
func (o *GetQualityPublishedformsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetQualityPublishedformsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsOK  %+v", 200, o.Payload)
}

func (o *GetQualityPublishedformsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsOK  %+v", 200, o.Payload)
}

func (o *GetQualityPublishedformsOK) GetPayload() *models.EvaluationFormEntityListing {
	return o.Payload
}

func (o *GetQualityPublishedformsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EvaluationFormEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsBadRequest creates a GetQualityPublishedformsBadRequest with default headers values
func NewGetQualityPublishedformsBadRequest() *GetQualityPublishedformsBadRequest {
	return &GetQualityPublishedformsBadRequest{}
}

/*
GetQualityPublishedformsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetQualityPublishedformsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms bad request response has a 2xx status code
func (o *GetQualityPublishedformsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms bad request response has a 3xx status code
func (o *GetQualityPublishedformsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms bad request response has a 4xx status code
func (o *GetQualityPublishedformsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms bad request response has a 5xx status code
func (o *GetQualityPublishedformsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms bad request response a status code equal to that given
func (o *GetQualityPublishedformsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetQualityPublishedformsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityPublishedformsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityPublishedformsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsUnauthorized creates a GetQualityPublishedformsUnauthorized with default headers values
func NewGetQualityPublishedformsUnauthorized() *GetQualityPublishedformsUnauthorized {
	return &GetQualityPublishedformsUnauthorized{}
}

/*
GetQualityPublishedformsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetQualityPublishedformsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms unauthorized response has a 2xx status code
func (o *GetQualityPublishedformsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms unauthorized response has a 3xx status code
func (o *GetQualityPublishedformsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms unauthorized response has a 4xx status code
func (o *GetQualityPublishedformsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms unauthorized response has a 5xx status code
func (o *GetQualityPublishedformsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms unauthorized response a status code equal to that given
func (o *GetQualityPublishedformsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetQualityPublishedformsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityPublishedformsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityPublishedformsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsForbidden creates a GetQualityPublishedformsForbidden with default headers values
func NewGetQualityPublishedformsForbidden() *GetQualityPublishedformsForbidden {
	return &GetQualityPublishedformsForbidden{}
}

/*
GetQualityPublishedformsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetQualityPublishedformsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms forbidden response has a 2xx status code
func (o *GetQualityPublishedformsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms forbidden response has a 3xx status code
func (o *GetQualityPublishedformsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms forbidden response has a 4xx status code
func (o *GetQualityPublishedformsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms forbidden response has a 5xx status code
func (o *GetQualityPublishedformsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms forbidden response a status code equal to that given
func (o *GetQualityPublishedformsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetQualityPublishedformsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityPublishedformsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityPublishedformsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsNotFound creates a GetQualityPublishedformsNotFound with default headers values
func NewGetQualityPublishedformsNotFound() *GetQualityPublishedformsNotFound {
	return &GetQualityPublishedformsNotFound{}
}

/*
GetQualityPublishedformsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetQualityPublishedformsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms not found response has a 2xx status code
func (o *GetQualityPublishedformsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms not found response has a 3xx status code
func (o *GetQualityPublishedformsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms not found response has a 4xx status code
func (o *GetQualityPublishedformsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms not found response has a 5xx status code
func (o *GetQualityPublishedformsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms not found response a status code equal to that given
func (o *GetQualityPublishedformsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetQualityPublishedformsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityPublishedformsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityPublishedformsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsRequestTimeout creates a GetQualityPublishedformsRequestTimeout with default headers values
func NewGetQualityPublishedformsRequestTimeout() *GetQualityPublishedformsRequestTimeout {
	return &GetQualityPublishedformsRequestTimeout{}
}

/*
GetQualityPublishedformsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetQualityPublishedformsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms request timeout response has a 2xx status code
func (o *GetQualityPublishedformsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms request timeout response has a 3xx status code
func (o *GetQualityPublishedformsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms request timeout response has a 4xx status code
func (o *GetQualityPublishedformsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms request timeout response has a 5xx status code
func (o *GetQualityPublishedformsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms request timeout response a status code equal to that given
func (o *GetQualityPublishedformsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetQualityPublishedformsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityPublishedformsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityPublishedformsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsRequestEntityTooLarge creates a GetQualityPublishedformsRequestEntityTooLarge with default headers values
func NewGetQualityPublishedformsRequestEntityTooLarge() *GetQualityPublishedformsRequestEntityTooLarge {
	return &GetQualityPublishedformsRequestEntityTooLarge{}
}

/*
GetQualityPublishedformsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetQualityPublishedformsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms request entity too large response has a 2xx status code
func (o *GetQualityPublishedformsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms request entity too large response has a 3xx status code
func (o *GetQualityPublishedformsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms request entity too large response has a 4xx status code
func (o *GetQualityPublishedformsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms request entity too large response has a 5xx status code
func (o *GetQualityPublishedformsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms request entity too large response a status code equal to that given
func (o *GetQualityPublishedformsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetQualityPublishedformsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityPublishedformsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityPublishedformsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsUnsupportedMediaType creates a GetQualityPublishedformsUnsupportedMediaType with default headers values
func NewGetQualityPublishedformsUnsupportedMediaType() *GetQualityPublishedformsUnsupportedMediaType {
	return &GetQualityPublishedformsUnsupportedMediaType{}
}

/*
GetQualityPublishedformsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetQualityPublishedformsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms unsupported media type response has a 2xx status code
func (o *GetQualityPublishedformsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms unsupported media type response has a 3xx status code
func (o *GetQualityPublishedformsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms unsupported media type response has a 4xx status code
func (o *GetQualityPublishedformsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms unsupported media type response has a 5xx status code
func (o *GetQualityPublishedformsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms unsupported media type response a status code equal to that given
func (o *GetQualityPublishedformsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetQualityPublishedformsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityPublishedformsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityPublishedformsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsTooManyRequests creates a GetQualityPublishedformsTooManyRequests with default headers values
func NewGetQualityPublishedformsTooManyRequests() *GetQualityPublishedformsTooManyRequests {
	return &GetQualityPublishedformsTooManyRequests{}
}

/*
GetQualityPublishedformsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetQualityPublishedformsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms too many requests response has a 2xx status code
func (o *GetQualityPublishedformsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms too many requests response has a 3xx status code
func (o *GetQualityPublishedformsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms too many requests response has a 4xx status code
func (o *GetQualityPublishedformsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms too many requests response has a 5xx status code
func (o *GetQualityPublishedformsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms too many requests response a status code equal to that given
func (o *GetQualityPublishedformsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetQualityPublishedformsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityPublishedformsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityPublishedformsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsInternalServerError creates a GetQualityPublishedformsInternalServerError with default headers values
func NewGetQualityPublishedformsInternalServerError() *GetQualityPublishedformsInternalServerError {
	return &GetQualityPublishedformsInternalServerError{}
}

/*
GetQualityPublishedformsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetQualityPublishedformsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms internal server error response has a 2xx status code
func (o *GetQualityPublishedformsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms internal server error response has a 3xx status code
func (o *GetQualityPublishedformsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms internal server error response has a 4xx status code
func (o *GetQualityPublishedformsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality publishedforms internal server error response has a 5xx status code
func (o *GetQualityPublishedformsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality publishedforms internal server error response a status code equal to that given
func (o *GetQualityPublishedformsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetQualityPublishedformsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityPublishedformsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityPublishedformsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsServiceUnavailable creates a GetQualityPublishedformsServiceUnavailable with default headers values
func NewGetQualityPublishedformsServiceUnavailable() *GetQualityPublishedformsServiceUnavailable {
	return &GetQualityPublishedformsServiceUnavailable{}
}

/*
GetQualityPublishedformsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetQualityPublishedformsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms service unavailable response has a 2xx status code
func (o *GetQualityPublishedformsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms service unavailable response has a 3xx status code
func (o *GetQualityPublishedformsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms service unavailable response has a 4xx status code
func (o *GetQualityPublishedformsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality publishedforms service unavailable response has a 5xx status code
func (o *GetQualityPublishedformsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality publishedforms service unavailable response a status code equal to that given
func (o *GetQualityPublishedformsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetQualityPublishedformsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityPublishedformsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityPublishedformsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsGatewayTimeout creates a GetQualityPublishedformsGatewayTimeout with default headers values
func NewGetQualityPublishedformsGatewayTimeout() *GetQualityPublishedformsGatewayTimeout {
	return &GetQualityPublishedformsGatewayTimeout{}
}

/*
GetQualityPublishedformsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetQualityPublishedformsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms gateway timeout response has a 2xx status code
func (o *GetQualityPublishedformsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms gateway timeout response has a 3xx status code
func (o *GetQualityPublishedformsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms gateway timeout response has a 4xx status code
func (o *GetQualityPublishedformsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality publishedforms gateway timeout response has a 5xx status code
func (o *GetQualityPublishedformsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality publishedforms gateway timeout response a status code equal to that given
func (o *GetQualityPublishedformsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetQualityPublishedformsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityPublishedformsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms][%d] getQualityPublishedformsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityPublishedformsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
