// Example assertion response header
func AssertResponseHeader(t *testing.T, res *http.Response, code int) {
    t.Helper()
	if code != res.StatusCode {
		t.Errorf("expected status is '%#v',\n but actual is '%#v'", code, res.StatusCode)
	}
	if expected := "application/json; charset=utf-8"; res.Header.Get("Content-Type") != expected {
		t.Errorf("expected Content-Type is '%#v',\n but given '%#v'",
            expected, res.Header.Get("Content-Type"))
	}
}

