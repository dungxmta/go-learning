package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// setup server
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// setup router
	e.GET("/demo", demo)

	// setup background task
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()

	// run server...
	e.Logger.Fatal(e.Start("0.0.0.0:5000"))
}

func demo(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func payloadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	/*
		// Read the body into a string for json decoding
		var content = &PayloadCollection{}
		err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)).Decode(&content)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Go through each payload and queue items individually to be posted to S3
		for _, payload := range content.Payloads {

			// let's create a job with the payload
			work := Job{Payload: payload}

			// Push the work onto the queue.
			JobQueue <- work
		}
	*/
	w.WriteHeader(http.StatusOK)
}
