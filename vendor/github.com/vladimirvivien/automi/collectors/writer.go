package collectors

import (
	"context"
	"fmt"
	"io"

	"github.com/vladimirvivien/automi/api"
	autoctx "github.com/vladimirvivien/automi/api/context"
	"github.com/vladimirvivien/automi/util"
)

type WriterCollector struct {
	writer io.Writer
	input  <-chan interface{}
	logf   api.LogFunc
	errf   api.ErrorFunc
}

func Writer(writer io.Writer) *WriterCollector {
	return &WriterCollector{
		writer: writer,
	}
}

func (c *WriterCollector) SetInput(in <-chan interface{}) {
	c.input = in
}

func (c *WriterCollector) Open(ctx context.Context) <-chan error {
	c.logf = autoctx.GetLogFunc(ctx)
	c.errf = autoctx.GetErrFunc(ctx)

	util.Logfn(c.logf, "Opening io.Writer collector")
	result := make(chan error)

	go func() {
		defer func() {
			close(result)
			util.Logfn(c.logf, "Closing io.Writer collector")
		}()

		for {
			select {
			case val, opened := <-c.input:
				if !opened {
					return
				}
				switch data := val.(type) {
				case string:
					_, err := fmt.Fprint(c.writer, data)
					if err != nil {
						util.Logfn(c.logf, err)
						autoctx.Err(c.errf, api.Error(err.Error()))
						continue
					}
				case []byte:
					if _, err := c.writer.Write(data); err != nil {
						util.Logfn(c.logf, err)
						autoctx.Err(c.errf, api.Error(err.Error()))
						continue
					}
				default:
					// other types are serialized using string representation
					// extracted by fmt
					_, err := fmt.Fprintf(c.writer, "%v", data)
					if err != nil {
						util.Logfn(c.logf, err)
						autoctx.Err(c.errf, api.Error(err.Error()))
						continue
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return result
}
