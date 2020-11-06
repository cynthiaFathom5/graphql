package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// BatchHandler implements the POST side of the default HTTP transport
// It will handle apollo array batching
type BatchHandler struct{}

var _ graphql.Transport = BatchHandler{}

func (h BatchHandler) Supports(r *http.Request) bool {
	if r.Header.Get("Upgrade") != "" {
		return false
	}

	mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return false
	}

	batching := r.Header.Get("X-Array-Batching")

	return r.Method == "POST" && mediaType == "application/json" && batching == "true"
}

func (h BatchHandler) Do(w http.ResponseWriter, r *http.Request, exec graphql.GraphExecutor) {
	w.Header().Set("Content-Type", "application/json")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJsonErrorf(w, "json body could not read: "+err.Error())
		return
	}

	var params []*graphql.RawParams
	start := graphql.Now()
	if err := jsonDecode(bytes.NewBuffer(data), &params); err != nil {
		var param *graphql.RawParams
		if err := jsonDecode(bytes.NewBuffer(data), &param); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			writeJsonErrorf(w, "json body could not be decoded: "+err.Error())
			return
		}
		params = append(params, param)
	}

	for i := range params {
		params[i].ReadTime = graphql.TraceTiming{
			Start: start,
			End:   graphql.Now(),
		}
	}

	output := make([]json.RawMessage, len(params))

	for i := range params {
		rc, err := exec.CreateOperationContext(r.Context(), params[i])
		if err != nil {
			w.WriteHeader(statusFor(err))
			resp := exec.DispatchError(graphql.WithOperationContext(r.Context(), rc), err)
			writeJson(w, resp)
			return
		}
		ctx := graphql.WithOperationContext(r.Context(), rc)
		responses, ctx := exec.DispatchOperation(ctx, rc)
		x := responses(ctx)
		output[i] = getJson(x)
	}

	outputData, _ := json.Marshal(output)
	w.Write(outputData)
}

func jsonDecode(r io.Reader, val interface{}) error {
	dec := json.NewDecoder(r)
	dec.UseNumber()
	return dec.Decode(val)
}

func statusFor(errs gqlerror.List) int {
	switch errcode.GetErrorKind(errs) {
	case errcode.KindProtocol:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusOK
	}
}

func writeJson(w io.Writer, response *graphql.Response) {
	b, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	w.Write(b)
}

func getJson(response *graphql.Response) []byte {
	b, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	return b
}

func writeJsonError(w io.Writer, msg string) {
	writeJson(w, &graphql.Response{Errors: gqlerror.List{{Message: msg}}})
}

func writeJsonErrorf(w io.Writer, format string, args ...interface{}) {
	writeJson(w, &graphql.Response{Errors: gqlerror.List{{Message: fmt.Sprintf(format, args...)}}})
}

func writeJsonGraphqlError(w io.Writer, err ...*gqlerror.Error) {
	writeJson(w, &graphql.Response{Errors: err})
}
