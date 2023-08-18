package user

import (
	"context"
	"fmt"

	"gitlab.com/willysihombing/task-c3/pkg/httpclient"
	"gitlab.com/willysihombing/task-c3/pkg/logger"
	"gitlab.com/willysihombing/task-c3/pkg/tracer"
)

func (c *client) Send(ctx context.Context, reqOpt *RequestOptionsSupabase, response interface{}) error {
	var (
		lf = logger.NewFields(
			logger.EventName("supabase"),
		)
	)

	ctx = tracer.SpanStart(ctx, "supabase")
	defer tracer.SpanFinish(ctx)

	h := httpclient.Headers{}
	h.Add(httpclient.ContentType, httpclient.MediaTypeJSON)
	h.Add(httpclient.ApiKey, c.cfg.Depedency.Supabase.Token)
	h.Add(httpclient.Authorization, fmt.Sprintf(`Bearer %s`, c.cfg.Depedency.Supabase.Token))

	resp, err := httpclient.Request(httpclient.RequestOptions{
		Payload:       reqOpt.Payload,
		URL:           fmt.Sprintf("%s%s", c.cfg.Depedency.Supabase.BaseURL, reqOpt.URL),
		Header:        h,
		Method:        reqOpt.Method,
		TimeoutSecond: c.cfg.Depedency.Supabase.Timeout,
		Context:       ctx,
	})

	if err != nil {
		tracer.SpanError(ctx, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("Send Document Sevice API error %v", err), lf...)
		return err
	}

	lf.Append(logger.Any("response_status_code", resp.Status()))
	if resp.Header().Get("Content-Type") == httpclient.MediaTypeJSONSupabase {
		err = resp.DecodeJSON(&response)
		if err != nil {
			tracer.SpanError(ctx, err)
			logger.ErrorWithContext(ctx, fmt.Sprintf("Decode Cocument Service API error %v, err "), lf...)
			return err
		}
	}

	logger.Info(fmt.Sprintf("Success send API call, response http status %v", resp.Status()), lf...)
	return nil
}
