package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"go-url-shortner/handlers"
	"go-url-shortner/mocks"
	"go-url-shortner/models"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func TestGetShortURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUrlStore := mocks.NewMockURL(ctrl)
	handler := handlers.New(mockUrlStore)

	tests := []struct {
		name        string
		reqBody     string
		respCode    int
		storeResp   *gomock.Call
		handlerResp interface{}
	}{
		{
			name:        "success",
			reqBody:     `{"longUrl": "http://www.google.com"}`,
			respCode:    http.StatusOK,
			storeResp:   mockUrlStore.EXPECT().GetShortURL(models.MyURL{LongURL: "http://www.google.com"}).Return("http://localhost:8080/jvxG5k5", nil),
			handlerResp: "http://localhost:8080/jvxG5k5",
		},
		{
			name:        "error: empty body",
			reqBody:     ``,
			respCode:    http.StatusBadRequest,
			storeResp:   nil,
			handlerResp: "bad request",
		},
		{
			name:        "error: invalid body",
			reqBody:     `{"longUrl": "wwwgooglecom"}`,
			respCode:    http.StatusBadRequest,
			storeResp:   nil,
			handlerResp: "invalid url",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := mux.NewRouter()
			r.HandleFunc("/short", handler.GetShortURL)

			req := httptest.NewRequest(http.MethodPost, "/short", strings.NewReader(tc.reqBody))
			req.Header.Set("Content-Type", "application/json")
			recorder := httptest.NewRecorder()

			r.ServeHTTP(recorder, req)

			if recorder.Code != tc.respCode {
				t.Errorf("Test failed, expected %d, got %d", tc.respCode, recorder.Code)
			}

			if !reflect.DeepEqual(recorder.Body.String(), tc.handlerResp) {
				t.Errorf("Test failed, expected %s, \n got %s", tc.handlerResp, recorder.Body)
			}
		})
	}
}
