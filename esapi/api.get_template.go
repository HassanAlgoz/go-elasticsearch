// Code generated from specification version 5.6.15: DO NOT EDIT

package esapi

import (
	"context"
	"strings"
)

func newGetTemplateFunc(t Transport) GetTemplate {
	return func(id string, o ...func(*GetTemplateRequest)) (*Response, error) {
		var r = GetTemplateRequest{DocumentID: id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/5.x/search-template.html.
//
type GetTemplate func(id string, o ...func(*GetTemplateRequest)) (*Response, error)

// GetTemplateRequest configures the Get Template API request.
//
type GetTemplateRequest struct {
	DocumentID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r GetTemplateRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_search") + 1 + len("template") + 1 + len(r.DocumentID))
	path.WriteString("/")
	path.WriteString("_search")
	path.WriteString("/")
	path.WriteString("template")
	path.WriteString("/")
	path.WriteString(r.DocumentID)

	params = make(map[string]string)

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, _ := newRequest(method, path.String(), nil)

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f GetTemplate) WithContext(v context.Context) func(*GetTemplateRequest) {
	return func(r *GetTemplateRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f GetTemplate) WithPretty() func(*GetTemplateRequest) {
	return func(r *GetTemplateRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f GetTemplate) WithHuman() func(*GetTemplateRequest) {
	return func(r *GetTemplateRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f GetTemplate) WithErrorTrace() func(*GetTemplateRequest) {
	return func(r *GetTemplateRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f GetTemplate) WithFilterPath(v ...string) func(*GetTemplateRequest) {
	return func(r *GetTemplateRequest) {
		r.FilterPath = v
	}
}
