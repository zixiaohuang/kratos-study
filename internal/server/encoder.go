package server

import (

	"github.com/go-kratos/kratos/v2/transport/http"
	"kratos-realworld/internal/errors"
	nethttp "net/http"
)

func errorEncoder(w nethttp.ResponseWriter, r *nethttp.Request, err error) {
	se := errors.FromError(err)
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal((se))
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/")
	w.WriteHeader(se.Code)
	_, _ = w.Write(body)
}