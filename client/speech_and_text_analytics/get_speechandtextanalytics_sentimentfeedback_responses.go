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

// GetSpeechandtextanalyticsSentimentfeedbackReader is a Reader for the GetSpeechandtextanalyticsSentimentfeedback structure.
type GetSpeechandtextanalyticsSentimentfeedbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpeechandtextanalyticsSentimentfeedbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpeechandtextanalyticsSentimentfeedbackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSpeechandtextanalyticsSentimentfeedbackBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSpeechandtextanalyticsSentimentfeedbackUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSpeechandtextanalyticsSentimentfeedbackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSpeechandtextanalyticsSentimentfeedbackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetSpeechandtextanalyticsSentimentfeedbackRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSpeechandtextanalyticsSentimentfeedbackTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSpeechandtextanalyticsSentimentfeedbackInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSpeechandtextanalyticsSentimentfeedbackServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetSpeechandtextanalyticsSentimentfeedbackGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSpeechandtextanalyticsSentimentfeedbackOK creates a GetSpeechandtextanalyticsSentimentfeedbackOK with default headers values
func NewGetSpeechandtextanalyticsSentimentfeedbackOK() *GetSpeechandtextanalyticsSentimentfeedbackOK {
	return &GetSpeechandtextanalyticsSentimentfeedbackOK{}
}

/*
GetSpeechandtextanalyticsSentimentfeedbackOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSpeechandtextanalyticsSentimentfeedbackOK struct {
	Payload *models.SentimentFeedbackEntityListing
}

// IsSuccess returns true when this get speechandtextanalytics sentimentfeedback o k response has a 2xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get speechandtextanalytics sentimentfeedback o k response has a 3xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics sentimentfeedback o k response has a 4xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get speechandtextanalytics sentimentfeedback o k response has a 5xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics sentimentfeedback o k response a status code equal to that given
func (o *GetSpeechandtextanalyticsSentimentfeedbackOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackOK  %+v", 200, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackOK) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackOK  %+v", 200, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackOK) GetPayload() *models.SentimentFeedbackEntityListing {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SentimentFeedbackEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentfeedbackBadRequest creates a GetSpeechandtextanalyticsSentimentfeedbackBadRequest with default headers values
func NewGetSpeechandtextanalyticsSentimentfeedbackBadRequest() *GetSpeechandtextanalyticsSentimentfeedbackBadRequest {
	return &GetSpeechandtextanalyticsSentimentfeedbackBadRequest{}
}

/*
GetSpeechandtextanalyticsSentimentfeedbackBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetSpeechandtextanalyticsSentimentfeedbackBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics sentimentfeedback bad request response has a 2xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics sentimentfeedback bad request response has a 3xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics sentimentfeedback bad request response has a 4xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics sentimentfeedback bad request response has a 5xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics sentimentfeedback bad request response a status code equal to that given
func (o *GetSpeechandtextanalyticsSentimentfeedbackBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackBadRequest  %+v", 400, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackBadRequest  %+v", 400, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentfeedbackUnauthorized creates a GetSpeechandtextanalyticsSentimentfeedbackUnauthorized with default headers values
func NewGetSpeechandtextanalyticsSentimentfeedbackUnauthorized() *GetSpeechandtextanalyticsSentimentfeedbackUnauthorized {
	return &GetSpeechandtextanalyticsSentimentfeedbackUnauthorized{}
}

/*
GetSpeechandtextanalyticsSentimentfeedbackUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetSpeechandtextanalyticsSentimentfeedbackUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics sentimentfeedback unauthorized response has a 2xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics sentimentfeedback unauthorized response has a 3xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics sentimentfeedback unauthorized response has a 4xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics sentimentfeedback unauthorized response has a 5xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics sentimentfeedback unauthorized response a status code equal to that given
func (o *GetSpeechandtextanalyticsSentimentfeedbackUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentfeedbackForbidden creates a GetSpeechandtextanalyticsSentimentfeedbackForbidden with default headers values
func NewGetSpeechandtextanalyticsSentimentfeedbackForbidden() *GetSpeechandtextanalyticsSentimentfeedbackForbidden {
	return &GetSpeechandtextanalyticsSentimentfeedbackForbidden{}
}

/*
GetSpeechandtextanalyticsSentimentfeedbackForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetSpeechandtextanalyticsSentimentfeedbackForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics sentimentfeedback forbidden response has a 2xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics sentimentfeedback forbidden response has a 3xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics sentimentfeedback forbidden response has a 4xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics sentimentfeedback forbidden response has a 5xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics sentimentfeedback forbidden response a status code equal to that given
func (o *GetSpeechandtextanalyticsSentimentfeedbackForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackForbidden  %+v", 403, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackForbidden  %+v", 403, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentfeedbackNotFound creates a GetSpeechandtextanalyticsSentimentfeedbackNotFound with default headers values
func NewGetSpeechandtextanalyticsSentimentfeedbackNotFound() *GetSpeechandtextanalyticsSentimentfeedbackNotFound {
	return &GetSpeechandtextanalyticsSentimentfeedbackNotFound{}
}

/*
GetSpeechandtextanalyticsSentimentfeedbackNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetSpeechandtextanalyticsSentimentfeedbackNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics sentimentfeedback not found response has a 2xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics sentimentfeedback not found response has a 3xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics sentimentfeedback not found response has a 4xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics sentimentfeedback not found response has a 5xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics sentimentfeedback not found response a status code equal to that given
func (o *GetSpeechandtextanalyticsSentimentfeedbackNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackNotFound  %+v", 404, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackNotFound  %+v", 404, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentfeedbackRequestTimeout creates a GetSpeechandtextanalyticsSentimentfeedbackRequestTimeout with default headers values
func NewGetSpeechandtextanalyticsSentimentfeedbackRequestTimeout() *GetSpeechandtextanalyticsSentimentfeedbackRequestTimeout {
	return &GetSpeechandtextanalyticsSentimentfeedbackRequestTimeout{}
}

/*
GetSpeechandtextanalyticsSentimentfeedbackRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetSpeechandtextanalyticsSentimentfeedbackRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics sentimentfeedback request timeout response has a 2xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics sentimentfeedback request timeout response has a 3xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics sentimentfeedback request timeout response has a 4xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics sentimentfeedback request timeout response has a 5xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics sentimentfeedback request timeout response a status code equal to that given
func (o *GetSpeechandtextanalyticsSentimentfeedbackRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge creates a GetSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge with default headers values
func NewGetSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge() *GetSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge {
	return &GetSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge{}
}

/*
GetSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics sentimentfeedback request entity too large response has a 2xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics sentimentfeedback request entity too large response has a 3xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics sentimentfeedback request entity too large response has a 4xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics sentimentfeedback request entity too large response has a 5xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics sentimentfeedback request entity too large response a status code equal to that given
func (o *GetSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType creates a GetSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType with default headers values
func NewGetSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType() *GetSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType {
	return &GetSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType{}
}

/*
GetSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics sentimentfeedback unsupported media type response has a 2xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics sentimentfeedback unsupported media type response has a 3xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics sentimentfeedback unsupported media type response has a 4xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics sentimentfeedback unsupported media type response has a 5xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics sentimentfeedback unsupported media type response a status code equal to that given
func (o *GetSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentfeedbackTooManyRequests creates a GetSpeechandtextanalyticsSentimentfeedbackTooManyRequests with default headers values
func NewGetSpeechandtextanalyticsSentimentfeedbackTooManyRequests() *GetSpeechandtextanalyticsSentimentfeedbackTooManyRequests {
	return &GetSpeechandtextanalyticsSentimentfeedbackTooManyRequests{}
}

/*
GetSpeechandtextanalyticsSentimentfeedbackTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetSpeechandtextanalyticsSentimentfeedbackTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics sentimentfeedback too many requests response has a 2xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics sentimentfeedback too many requests response has a 3xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics sentimentfeedback too many requests response has a 4xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics sentimentfeedback too many requests response has a 5xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics sentimentfeedback too many requests response a status code equal to that given
func (o *GetSpeechandtextanalyticsSentimentfeedbackTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentfeedbackInternalServerError creates a GetSpeechandtextanalyticsSentimentfeedbackInternalServerError with default headers values
func NewGetSpeechandtextanalyticsSentimentfeedbackInternalServerError() *GetSpeechandtextanalyticsSentimentfeedbackInternalServerError {
	return &GetSpeechandtextanalyticsSentimentfeedbackInternalServerError{}
}

/*
GetSpeechandtextanalyticsSentimentfeedbackInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetSpeechandtextanalyticsSentimentfeedbackInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics sentimentfeedback internal server error response has a 2xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics sentimentfeedback internal server error response has a 3xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics sentimentfeedback internal server error response has a 4xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get speechandtextanalytics sentimentfeedback internal server error response has a 5xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get speechandtextanalytics sentimentfeedback internal server error response a status code equal to that given
func (o *GetSpeechandtextanalyticsSentimentfeedbackInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentfeedbackServiceUnavailable creates a GetSpeechandtextanalyticsSentimentfeedbackServiceUnavailable with default headers values
func NewGetSpeechandtextanalyticsSentimentfeedbackServiceUnavailable() *GetSpeechandtextanalyticsSentimentfeedbackServiceUnavailable {
	return &GetSpeechandtextanalyticsSentimentfeedbackServiceUnavailable{}
}

/*
GetSpeechandtextanalyticsSentimentfeedbackServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetSpeechandtextanalyticsSentimentfeedbackServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics sentimentfeedback service unavailable response has a 2xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics sentimentfeedback service unavailable response has a 3xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics sentimentfeedback service unavailable response has a 4xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get speechandtextanalytics sentimentfeedback service unavailable response has a 5xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get speechandtextanalytics sentimentfeedback service unavailable response a status code equal to that given
func (o *GetSpeechandtextanalyticsSentimentfeedbackServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentfeedbackGatewayTimeout creates a GetSpeechandtextanalyticsSentimentfeedbackGatewayTimeout with default headers values
func NewGetSpeechandtextanalyticsSentimentfeedbackGatewayTimeout() *GetSpeechandtextanalyticsSentimentfeedbackGatewayTimeout {
	return &GetSpeechandtextanalyticsSentimentfeedbackGatewayTimeout{}
}

/*
GetSpeechandtextanalyticsSentimentfeedbackGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetSpeechandtextanalyticsSentimentfeedbackGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics sentimentfeedback gateway timeout response has a 2xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics sentimentfeedback gateway timeout response has a 3xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics sentimentfeedback gateway timeout response has a 4xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get speechandtextanalytics sentimentfeedback gateway timeout response has a 5xx status code
func (o *GetSpeechandtextanalyticsSentimentfeedbackGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get speechandtextanalytics sentimentfeedback gateway timeout response a status code equal to that given
func (o *GetSpeechandtextanalyticsSentimentfeedbackGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentimentfeedback][%d] getSpeechandtextanalyticsSentimentfeedbackGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentfeedbackGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
