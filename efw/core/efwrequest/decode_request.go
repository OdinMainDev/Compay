package efwrequest

import (
	"compress/gzip"
	"compress/zlib"
	"efw/utils"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/andybalholm/brotli"
)

var ErrorDecodeGzip = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: REQUEST] DecodeGzip: ", err)
}

var ErrorDecodeDeflate = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: REQUEST] DecodeDeflate: ", err)
}

var ErrorDecodeUnsupportedType = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: REQUEST] DecodeUnsupportedType: ", err)
}

var ErrorDecodeCloseAll = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: REQUEST] DecodeCloseAll: ", err)
}

// TODO: добавить zstd и другие типы, когда будут нужны
func DecodeRequest(resp *http.Response) ([]byte, error) {
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	encHeader := strings.TrimSpace(resp.Header.Get("Content-Encoding"))
	var r io.Reader = resp.Body

	var closers []io.Closer
	closers = append(closers, resp.Body)

	if encHeader != "" {
		encs := strings.Split(encHeader, ",")
		for i := len(encs) - 1; i >= 0; i-- {
			switch strings.TrimSpace(utils.ToLower(encs[i])) {
			case "", "identity":
				// ничего
			case "gzip":
				gz, err := gzip.NewReader(r)
				if err != nil {
					closeAll(closers)
					return nil, ErrorDecodeGzip(err.Error())
				}
				r = gz
				closers = append(closers, gz)
			case "deflate":
				zr, err := zlib.NewReader(r)
				if err != nil {
					closeAll(closers)
					return nil, ErrorDecodeDeflate(err.Error())
				}
				r = zr
				closers = append(closers, zr)
			case "br":
				r = brotli.NewReader(r)
			default:
				closeAll(closers)
				return nil, ErrorDecodeUnsupportedType(encs[i])
			}
		}
	}

	data, err := io.ReadAll(r)
	if err := closeAll(closers); err != nil {
		return nil, ErrorDecodeCloseAll(err.Error())
	}
	return data, err
}

func closeAll(cs []io.Closer) error {
	var firstErr error
	for i := len(cs) - 1; i >= 0; i-- {
		if cs[i] == nil {
			continue
		}
		if err := cs[i].Close(); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}
