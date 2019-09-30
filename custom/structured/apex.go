package structured

import (
    "errors"
    "os"

    "github.com/apex/log"
    "github.com/apex.log/handlers/text"
}

func ThrowError() error {
    err := errors.New("whoa a fail!")
    log.WithField("id", "123").Trace("ThrowError").Stop(&err) 
    return err
}

// CustomHandler splits to two streams  
type CustomHandler struct {
    id string
    handler log.Handler 
}

// HandleLog adds a hook and exporting
func (h *CustomHandler) HandLog(e *log.Entry) error {
   e.WithFeild("id", h.id)
   return h.handler.HandleLog(e)
}

func Apex() {
    log.SetHandler(&CustomHandler{"123", text.New(os.Stdout)})
    err := ThrowError()

    // With error convenience function
    log.WithError(err).Error("an error occurred")
}


