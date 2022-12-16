package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/burkayaltuntas/go-fast-temp/pkg/common"
	"github.com/burkayaltuntas/go-fast-temp/pkg/utils"
	"github.com/gin-gonic/gin"
)

// this may be on env file
var LOGFILE = "./logs/info.log"
var ERRORLOGFILE = "./logs/error.log"
var f *os.File
var LOG *log.Logger

func init() {
	var err error

	errlogfile, err := os.OpenFile(ERRORLOGFILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panicf("Failed to create request log file: %v", err)
	}

	gin.DefaultErrorWriter = io.MultiWriter(errlogfile, os.Stderr)

	f, err = os.OpenFile(LOGFILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panicf("error opening file: %v", err)
	}

	LOG = log.New(f, "INFO ", log.LstdFlags|log.Lmicroseconds)

	LOG.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		var bs string
		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			body, _ := io.ReadAll(c.Request.Body)
			bs = string(body)
			c.Request.Body = io.NopCloser(bytes.NewReader(body))
		}

		user := common.ContextUser{}
		u, exist := c.Get("user")
		if exist {
			user = u.(common.ContextUser)
		}

		go LOG.Println(user.Email, c.Request.RemoteAddr, c.Request.Method, c.Request.RequestURI, "REQUEST", bs)

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		mdl := utils.ResponseModel{}

		if blw.Status() > 201 {
			resp, err := io.ReadAll(blw.body)
			if err == nil {
				json.Unmarshal(resp, &mdl)
			}
		}

		go LOG.Println(user.Email, c.Request.RemoteAddr, c.Request.Method, c.Request.RequestURI, "RESPONSE", blw.Status(), mdl.Message)
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
