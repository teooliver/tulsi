package manualtester

import (
	"net/http"
)

func test() {
	r, _ := http.NewRequest("GET", "/users?is_member=1&age_range=18&age_range=60", nil)
	r.Header.Set("Authorization", "my-secret-here")

	// rw := httptest.NewRecorder()
	// router.ServeHTTP(rw, r)
}
