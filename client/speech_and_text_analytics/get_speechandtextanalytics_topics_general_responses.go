// Code generated by go-swagger; DO NOT EDIT.

package speech_and_text_analytics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetSpeechandtextanalyticsTopicsGeneralReader is a Reader for the GetSpeechandtextanalyticsTopicsGeneral structure.
type GetSpeechandtextanalyticsTopicsGeneralReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpeechandtextanalyticsTopicsGeneralReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpeechandtextanalyticsTopicsGeneralOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSpeechandtextanalyticsTopicsGeneralBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSpeechandtextanalyticsTopicsGeneralUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSpeechandtextanalyticsTopicsGeneralForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSpeechandtextanalyticsTopicsGeneralNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetSpeechandtextanalyticsTopicsGeneralRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSpeechandtextanalyticsTopicsGeneralTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSpeechandtextanalyticsTopicsGeneralInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSpeechandtextanalyticsTopicsGeneralServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetSpeechandtextanalyticsTopicsGeneralGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSpeechandtextanalyticsTopicsGeneralOK creates a GetSpeechandtextanalyticsTopicsGeneralOK with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralOK() *GetSpeechandtextanalyticsTopicsGeneralOK {
	return &GetSpeechandtextanalyticsTopicsGeneralOK{}
}

/*
GetSpeechandtextanalyticsTopicsGeneralOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSpeechandtextanalyticsTopicsGeneralOK struct {
	Payload *models.GeneralTopicsEntityListing
}

// IsSuccess returns true when this get speechandtextanalytics topics general o k response has a 2xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get speechandtextanalytics topics general o k response has a 3xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics topics general o k response has a 4xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get speechandtextanalytics topics general o k response has a 5xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics topics general o k response a status code equal to that given
func (o *GetSpeechandtextanalyticsTopicsGeneralOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSpeechandtextanalyticsTopicsGeneralOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralOK  %+v", 200, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralOK) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralOK  %+v", 200, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralOK) GetPayload() *models.GeneralTopicsEntityListing {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralTopicsEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicsGeneralBadRequest creates a GetSpeechandtextanalyticsTopicsGeneralBadRequest with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralBadRequest() *GetSpeechandtextanalyticsTopicsGeneralBadRequest {
	return &GetSpeechandtextanalyticsTopicsGeneralBadRequest{}
}

/*
GetSpeechandtextanalyticsTopicsGeneralBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetSpeechandtextanalyticsTopicsGeneralBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics topics general bad request response has a 2xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics topics general bad request response has a 3xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics topics general bad request response has a 4xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics topics general bad request response has a 5xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics topics general bad request response a status code equal to that given
func (o *GetSpeechandtextanalyticsTopicsGeneralBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetSpeechandtextanalyticsTopicsGeneralBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralBadRequest  %+v", 400, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralBadRequest  %+v", 400, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicsGeneralUnauthorized creates a GetSpeechandtextanalyticsTopicsGeneralUnauthorized with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralUnauthorized() *GetSpeechandtextanalyticsTopicsGeneralUnauthorized {
	return &GetSpeechandtextanalyticsTopicsGeneralUnauthorized{}
}

/*
GetSpeechandtextanalyticsTopicsGeneralUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetSpeechandtextanalyticsTopicsGeneralUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics topics general unauthorized response has a 2xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics topics general unauthorized response has a 3xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics topics general unauthorized response has a 4xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics topics general unauthorized response has a 5xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics topics general unauthorized response a status code equal to that given
func (o *GetSpeechandtextanalyticsTopicsGeneralUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetSpeechandtextanalyticsTopicsGeneralUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicsGeneralForbidden creates a GetSpeechandtextanalyticsTopicsGeneralForbidden with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralForbidden() *GetSpeechandtextanalyticsTopicsGeneralForbidden {
	return &GetSpeechandtextanalyticsTopicsGeneralForbidden{}
}

/*
GetSpeechandtextanalyticsTopicsGeneralForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetSpeechandtextanalyticsTopicsGeneralForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics topics general forbidden response has a 2xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics topics general forbidden response has a 3xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics topics general forbidden response has a 4xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics topics general forbidden response has a 5xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics topics general forbidden response a status code equal to that given
func (o *GetSpeechandtextanalyticsTopicsGeneralForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetSpeechandtextanalyticsTopicsGeneralForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralForbidden  %+v", 403, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralForbidden  %+v", 403, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicsGeneralNotFound creates a GetSpeechandtextanalyticsTopicsGeneralNotFound with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralNotFound() *GetSpeechandtextanalyticsTopicsGeneralNotFound {
	return &GetSpeechandtextanalyticsTopicsGeneralNotFound{}
}

/*
GetSpeechandtextanalyticsTopicsGeneralNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetSpeechandtextanalyticsTopicsGeneralNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics topics general not found response has a 2xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics topics general not found response has a 3xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics topics general not found response has a 4xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics topics general not found response has a 5xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics topics general not found response a status code equal to that given
func (o *GetSpeechandtextanalyticsTopicsGeneralNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetSpeechandtextanalyticsTopicsGeneralNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralNotFound  %+v", 404, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralNotFound  %+v", 404, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicsGeneralRequestTimeout creates a GetSpeechandtextanalyticsTopicsGeneralRequestTimeout with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralRequestTimeout() *GetSpeechandtextanalyticsTopicsGeneralRequestTimeout {
	return &GetSpeechandtextanalyticsTopicsGeneralRequestTimeout{}
}

/*
GetSpeechandtextanalyticsTopicsGeneralRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetSpeechandtextanalyticsTopicsGeneralRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics topics general request timeout response has a 2xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics topics general request timeout response has a 3xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics topics general request timeout response has a 4xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics topics general request timeout response has a 5xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics topics general request timeout response a status code equal to that given
func (o *GetSpeechandtextanalyticsTopicsGeneralRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetSpeechandtextanalyticsTopicsGeneralRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge creates a GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge() *GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge {
	return &GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge{}
}

/*
GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics topics general request entity too large response has a 2xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics topics general request entity too large response has a 3xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics topics general request entity too large response has a 4xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics topics general request entity too large response has a 5xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics topics general request entity too large response a status code equal to that given
func (o *GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType creates a GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType() *GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType {
	return &GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType{}
}

/*
GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics topics general unsupported media type response has a 2xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics topics general unsupported media type response has a 3xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics topics general unsupported media type response has a 4xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics topics general unsupported media type response has a 5xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics topics general unsupported media type response a status code equal to that given
func (o *GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicsGeneralTooManyRequests creates a GetSpeechandtextanalyticsTopicsGeneralTooManyRequests with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralTooManyRequests() *GetSpeechandtextanalyticsTopicsGeneralTooManyRequests {
	return &GetSpeechandtextanalyticsTopicsGeneralTooManyRequests{}
}

/*
GetSpeechandtextanalyticsTopicsGeneralTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetSpeechandtextanalyticsTopicsGeneralTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics topics general too many requests response has a 2xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics topics general too many requests response has a 3xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics topics general too many requests response has a 4xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics topics general too many requests response has a 5xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics topics general too many requests response a status code equal to that given
func (o *GetSpeechandtextanalyticsTopicsGeneralTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetSpeechandtextanalyticsTopicsGeneralTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicsGeneralInternalServerError creates a GetSpeechandtextanalyticsTopicsGeneralInternalServerError with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralInternalServerError() *GetSpeechandtextanalyticsTopicsGeneralInternalServerError {
	return &GetSpeechandtextanalyticsTopicsGeneralInternalServerError{}
}

/*
GetSpeechandtextanalyticsTopicsGeneralInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetSpeechandtextanalyticsTopicsGeneralInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics topics general internal server error response has a 2xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics topics general internal server error response has a 3xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics topics general internal server error response has a 4xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get speechandtextanalytics topics general internal server error response has a 5xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get speechandtextanalytics topics general internal server error response a status code equal to that given
func (o *GetSpeechandtextanalyticsTopicsGeneralInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetSpeechandtextanalyticsTopicsGeneralInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicsGeneralServiceUnavailable creates a GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralServiceUnavailable() *GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable {
	return &GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable{}
}

/*
GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics topics general service unavailable response has a 2xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics topics general service unavailable response has a 3xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics topics general service unavailable response has a 4xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get speechandtextanalytics topics general service unavailable response has a 5xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get speechandtextanalytics topics general service unavailable response a status code equal to that given
func (o *GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicsGeneralGatewayTimeout creates a GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout with default headers values
func NewGetSpeechandtextanalyticsTopicsGeneralGatewayTimeout() *GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout {
	return &GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout{}
}

/*
GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics topics general gateway timeout response has a 2xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics topics general gateway timeout response has a 3xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics topics general gateway timeout response has a 4xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get speechandtextanalytics topics general gateway timeout response has a 5xx status code
func (o *GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get speechandtextanalytics topics general gateway timeout response a status code equal to that given
func (o *GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/general][%d] getSpeechandtextanalyticsTopicsGeneralGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicsGeneralGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
