package exceptions

import "context"

func getRequestId(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	if id, ok := ctx.Value("request_id").(string); ok {
		return id
	}

	return ""
}
