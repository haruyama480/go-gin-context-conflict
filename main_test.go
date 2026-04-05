package main

import (
	"net/http"
	"net/http/httptest"
	"runtime"
	"testing"
)

func TestHelloRouteParallel(t *testing.T) {
	router := setupRouter()

	// 少なくともGOMAXPROC数と同数並列に実行すると、n*2のgoroutineが生成されることで、競合が発生しやすくなる
	n := runtime.GOMAXPROCS(0)
	for range n {
		t.Run("ParallelRequest", func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/hello", nil)

			router.ServeHTTP(w, req)
			if w.Code != http.StatusOK {
				t.Errorf("")
			}
		})
	}
}
