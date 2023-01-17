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

// GetQualityPublishedformsSurveysReader is a Reader for the GetQualityPublishedformsSurveys structure.
type GetQualityPublishedformsSurveysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQualityPublishedformsSurveysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQualityPublishedformsSurveysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetQualityPublishedformsSurveysBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetQualityPublishedformsSurveysUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetQualityPublishedformsSurveysForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetQualityPublishedformsSurveysNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetQualityPublishedformsSurveysRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetQualityPublishedformsSurveysRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetQualityPublishedformsSurveysUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetQualityPublishedformsSurveysTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetQualityPublishedformsSurveysInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetQualityPublishedformsSurveysServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetQualityPublishedformsSurveysGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetQualityPublishedformsSurveysOK creates a GetQualityPublishedformsSurveysOK with default headers values
func NewGetQualityPublishedformsSurveysOK() *GetQualityPublishedformsSurveysOK {
	return &GetQualityPublishedformsSurveysOK{}
}

/*
GetQualityPublishedformsSurveysOK describes a response with status code 200, with default header values.

successful operation
*/
type GetQualityPublishedformsSurveysOK struct {
	Payload *models.SurveyFormEntityListing
}

// IsSuccess returns true when this get quality publishedforms surveys o k response has a 2xx status code
func (o *GetQualityPublishedformsSurveysOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get quality publishedforms surveys o k response has a 3xx status code
func (o *GetQualityPublishedformsSurveysOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms surveys o k response has a 4xx status code
func (o *GetQualityPublishedformsSurveysOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality publishedforms surveys o k response has a 5xx status code
func (o *GetQualityPublishedformsSurveysOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms surveys o k response a status code equal to that given
func (o *GetQualityPublishedformsSurveysOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetQualityPublishedformsSurveysOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysOK  %+v", 200, o.Payload)
}

func (o *GetQualityPublishedformsSurveysOK) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysOK  %+v", 200, o.Payload)
}

func (o *GetQualityPublishedformsSurveysOK) GetPayload() *models.SurveyFormEntityListing {
	return o.Payload
}

func (o *GetQualityPublishedformsSurveysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SurveyFormEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsSurveysBadRequest creates a GetQualityPublishedformsSurveysBadRequest with default headers values
func NewGetQualityPublishedformsSurveysBadRequest() *GetQualityPublishedformsSurveysBadRequest {
	return &GetQualityPublishedformsSurveysBadRequest{}
}

/*
GetQualityPublishedformsSurveysBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetQualityPublishedformsSurveysBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms surveys bad request response has a 2xx status code
func (o *GetQualityPublishedformsSurveysBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms surveys bad request response has a 3xx status code
func (o *GetQualityPublishedformsSurveysBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms surveys bad request response has a 4xx status code
func (o *GetQualityPublishedformsSurveysBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms surveys bad request response has a 5xx status code
func (o *GetQualityPublishedformsSurveysBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms surveys bad request response a status code equal to that given
func (o *GetQualityPublishedformsSurveysBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetQualityPublishedformsSurveysBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityPublishedformsSurveysBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityPublishedformsSurveysBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsSurveysBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsSurveysUnauthorized creates a GetQualityPublishedformsSurveysUnauthorized with default headers values
func NewGetQualityPublishedformsSurveysUnauthorized() *GetQualityPublishedformsSurveysUnauthorized {
	return &GetQualityPublishedformsSurveysUnauthorized{}
}

/*
GetQualityPublishedformsSurveysUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetQualityPublishedformsSurveysUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms surveys unauthorized response has a 2xx status code
func (o *GetQualityPublishedformsSurveysUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms surveys unauthorized response has a 3xx status code
func (o *GetQualityPublishedformsSurveysUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms surveys unauthorized response has a 4xx status code
func (o *GetQualityPublishedformsSurveysUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms surveys unauthorized response has a 5xx status code
func (o *GetQualityPublishedformsSurveysUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms surveys unauthorized response a status code equal to that given
func (o *GetQualityPublishedformsSurveysUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetQualityPublishedformsSurveysUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityPublishedformsSurveysUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityPublishedformsSurveysUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsSurveysUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsSurveysForbidden creates a GetQualityPublishedformsSurveysForbidden with default headers values
func NewGetQualityPublishedformsSurveysForbidden() *GetQualityPublishedformsSurveysForbidden {
	return &GetQualityPublishedformsSurveysForbidden{}
}

/*
GetQualityPublishedformsSurveysForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetQualityPublishedformsSurveysForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms surveys forbidden response has a 2xx status code
func (o *GetQualityPublishedformsSurveysForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms surveys forbidden response has a 3xx status code
func (o *GetQualityPublishedformsSurveysForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms surveys forbidden response has a 4xx status code
func (o *GetQualityPublishedformsSurveysForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms surveys forbidden response has a 5xx status code
func (o *GetQualityPublishedformsSurveysForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms surveys forbidden response a status code equal to that given
func (o *GetQualityPublishedformsSurveysForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetQualityPublishedformsSurveysForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityPublishedformsSurveysForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityPublishedformsSurveysForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsSurveysForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsSurveysNotFound creates a GetQualityPublishedformsSurveysNotFound with default headers values
func NewGetQualityPublishedformsSurveysNotFound() *GetQualityPublishedformsSurveysNotFound {
	return &GetQualityPublishedformsSurveysNotFound{}
}

/*
GetQualityPublishedformsSurveysNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetQualityPublishedformsSurveysNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms surveys not found response has a 2xx status code
func (o *GetQualityPublishedformsSurveysNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms surveys not found response has a 3xx status code
func (o *GetQualityPublishedformsSurveysNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms surveys not found response has a 4xx status code
func (o *GetQualityPublishedformsSurveysNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms surveys not found response has a 5xx status code
func (o *GetQualityPublishedformsSurveysNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms surveys not found response a status code equal to that given
func (o *GetQualityPublishedformsSurveysNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetQualityPublishedformsSurveysNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityPublishedformsSurveysNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityPublishedformsSurveysNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsSurveysNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsSurveysRequestTimeout creates a GetQualityPublishedformsSurveysRequestTimeout with default headers values
func NewGetQualityPublishedformsSurveysRequestTimeout() *GetQualityPublishedformsSurveysRequestTimeout {
	return &GetQualityPublishedformsSurveysRequestTimeout{}
}

/*
GetQualityPublishedformsSurveysRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetQualityPublishedformsSurveysRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms surveys request timeout response has a 2xx status code
func (o *GetQualityPublishedformsSurveysRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms surveys request timeout response has a 3xx status code
func (o *GetQualityPublishedformsSurveysRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms surveys request timeout response has a 4xx status code
func (o *GetQualityPublishedformsSurveysRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms surveys request timeout response has a 5xx status code
func (o *GetQualityPublishedformsSurveysRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms surveys request timeout response a status code equal to that given
func (o *GetQualityPublishedformsSurveysRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetQualityPublishedformsSurveysRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityPublishedformsSurveysRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityPublishedformsSurveysRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsSurveysRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsSurveysRequestEntityTooLarge creates a GetQualityPublishedformsSurveysRequestEntityTooLarge with default headers values
func NewGetQualityPublishedformsSurveysRequestEntityTooLarge() *GetQualityPublishedformsSurveysRequestEntityTooLarge {
	return &GetQualityPublishedformsSurveysRequestEntityTooLarge{}
}

/*
GetQualityPublishedformsSurveysRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetQualityPublishedformsSurveysRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms surveys request entity too large response has a 2xx status code
func (o *GetQualityPublishedformsSurveysRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms surveys request entity too large response has a 3xx status code
func (o *GetQualityPublishedformsSurveysRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms surveys request entity too large response has a 4xx status code
func (o *GetQualityPublishedformsSurveysRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms surveys request entity too large response has a 5xx status code
func (o *GetQualityPublishedformsSurveysRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms surveys request entity too large response a status code equal to that given
func (o *GetQualityPublishedformsSurveysRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetQualityPublishedformsSurveysRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityPublishedformsSurveysRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityPublishedformsSurveysRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsSurveysRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsSurveysUnsupportedMediaType creates a GetQualityPublishedformsSurveysUnsupportedMediaType with default headers values
func NewGetQualityPublishedformsSurveysUnsupportedMediaType() *GetQualityPublishedformsSurveysUnsupportedMediaType {
	return &GetQualityPublishedformsSurveysUnsupportedMediaType{}
}

/*
GetQualityPublishedformsSurveysUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetQualityPublishedformsSurveysUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms surveys unsupported media type response has a 2xx status code
func (o *GetQualityPublishedformsSurveysUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms surveys unsupported media type response has a 3xx status code
func (o *GetQualityPublishedformsSurveysUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms surveys unsupported media type response has a 4xx status code
func (o *GetQualityPublishedformsSurveysUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms surveys unsupported media type response has a 5xx status code
func (o *GetQualityPublishedformsSurveysUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms surveys unsupported media type response a status code equal to that given
func (o *GetQualityPublishedformsSurveysUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetQualityPublishedformsSurveysUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityPublishedformsSurveysUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityPublishedformsSurveysUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsSurveysUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsSurveysTooManyRequests creates a GetQualityPublishedformsSurveysTooManyRequests with default headers values
func NewGetQualityPublishedformsSurveysTooManyRequests() *GetQualityPublishedformsSurveysTooManyRequests {
	return &GetQualityPublishedformsSurveysTooManyRequests{}
}

/*
GetQualityPublishedformsSurveysTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetQualityPublishedformsSurveysTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms surveys too many requests response has a 2xx status code
func (o *GetQualityPublishedformsSurveysTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms surveys too many requests response has a 3xx status code
func (o *GetQualityPublishedformsSurveysTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms surveys too many requests response has a 4xx status code
func (o *GetQualityPublishedformsSurveysTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms surveys too many requests response has a 5xx status code
func (o *GetQualityPublishedformsSurveysTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms surveys too many requests response a status code equal to that given
func (o *GetQualityPublishedformsSurveysTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetQualityPublishedformsSurveysTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityPublishedformsSurveysTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityPublishedformsSurveysTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsSurveysTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsSurveysInternalServerError creates a GetQualityPublishedformsSurveysInternalServerError with default headers values
func NewGetQualityPublishedformsSurveysInternalServerError() *GetQualityPublishedformsSurveysInternalServerError {
	return &GetQualityPublishedformsSurveysInternalServerError{}
}

/*
GetQualityPublishedformsSurveysInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetQualityPublishedformsSurveysInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms surveys internal server error response has a 2xx status code
func (o *GetQualityPublishedformsSurveysInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms surveys internal server error response has a 3xx status code
func (o *GetQualityPublishedformsSurveysInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms surveys internal server error response has a 4xx status code
func (o *GetQualityPublishedformsSurveysInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality publishedforms surveys internal server error response has a 5xx status code
func (o *GetQualityPublishedformsSurveysInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality publishedforms surveys internal server error response a status code equal to that given
func (o *GetQualityPublishedformsSurveysInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetQualityPublishedformsSurveysInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityPublishedformsSurveysInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityPublishedformsSurveysInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsSurveysInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsSurveysServiceUnavailable creates a GetQualityPublishedformsSurveysServiceUnavailable with default headers values
func NewGetQualityPublishedformsSurveysServiceUnavailable() *GetQualityPublishedformsSurveysServiceUnavailable {
	return &GetQualityPublishedformsSurveysServiceUnavailable{}
}

/*
GetQualityPublishedformsSurveysServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetQualityPublishedformsSurveysServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms surveys service unavailable response has a 2xx status code
func (o *GetQualityPublishedformsSurveysServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms surveys service unavailable response has a 3xx status code
func (o *GetQualityPublishedformsSurveysServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms surveys service unavailable response has a 4xx status code
func (o *GetQualityPublishedformsSurveysServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality publishedforms surveys service unavailable response has a 5xx status code
func (o *GetQualityPublishedformsSurveysServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality publishedforms surveys service unavailable response a status code equal to that given
func (o *GetQualityPublishedformsSurveysServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetQualityPublishedformsSurveysServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityPublishedformsSurveysServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityPublishedformsSurveysServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsSurveysServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsSurveysGatewayTimeout creates a GetQualityPublishedformsSurveysGatewayTimeout with default headers values
func NewGetQualityPublishedformsSurveysGatewayTimeout() *GetQualityPublishedformsSurveysGatewayTimeout {
	return &GetQualityPublishedformsSurveysGatewayTimeout{}
}

/*
GetQualityPublishedformsSurveysGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetQualityPublishedformsSurveysGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms surveys gateway timeout response has a 2xx status code
func (o *GetQualityPublishedformsSurveysGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms surveys gateway timeout response has a 3xx status code
func (o *GetQualityPublishedformsSurveysGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms surveys gateway timeout response has a 4xx status code
func (o *GetQualityPublishedformsSurveysGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality publishedforms surveys gateway timeout response has a 5xx status code
func (o *GetQualityPublishedformsSurveysGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality publishedforms surveys gateway timeout response a status code equal to that given
func (o *GetQualityPublishedformsSurveysGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetQualityPublishedformsSurveysGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityPublishedformsSurveysGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/surveys][%d] getQualityPublishedformsSurveysGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityPublishedformsSurveysGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsSurveysGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
