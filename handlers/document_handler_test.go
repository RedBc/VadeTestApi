package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	Convey("given a http request for /invalid/123", t, func() {
		req := httptest.NewRequest("get", "/invalid/123", nil)
		resp := httptest.NewRecorder()

		Convey("when the request is handled by the router", func() {
			mux.NewRouter().ServeHTTP(resp, req)

			Convey("then the response should be a 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})
}
