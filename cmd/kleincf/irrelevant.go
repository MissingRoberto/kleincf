package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func UAALogin(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"token_type":   "bearer",
		"access_token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiJhYmIzZjczYmZmZGM0Zjg2ODc4OGQwNWM4YzNkMDhjZSIsInN1YiI6IjA4OTFlODBmLTNkZjItNDVlYi04NzFmLWE4MDlmY2M5MGU3OCIsInNjb3BlIjpbIm9wZW5pZCJdLCJjbGllbnRfaWQiOiJ1c2VyaWQiLCJjaWQiOiJ0ZXN0Y2xpZW50SDRuZlQ5IiwiYXpwIjoidGVzdGNsaWVudEg0bmZUOSIsImdyYW50X3R5cGUiOiJ1cm46aWV0ZjpwYXJhbXM6b2F1dGg6Z3JhbnQtdHlwZTpzYW1sMi1iZWFyZXIiLCJ1c2VyX2lkIjoiMDg5MWU4MGYtM2RmMi00NWViLTg3MWYtYTgwOWZjYzkwZTc4Iiwib3JpZ2luIjoieGdtNHpnLmNsb3VkZm91bmRyeS1zYW1sLWxvZ2luIiwidXNlcl9uYW1lIjoidXNlciIsImVtYWlsIjoidXNlckBlbWFpbC5jb20iLCJyZXZfc2lnIjoiY2ZiZGNjMTgiLCJpYXQiOjE1Mjk2OTA1MTYsImV4cCI6MTUzMzMwMzQ4MCwiaXNzIjoiaHR0cDovL3hnbTR6Zy5sb2NhbGhvc3Q6ODA4MC91YWEvb2F1dGgvdG9rZW4iLCJ6aWQiOiJ4Z200emciLCJhdWQiOltdfQ.vFGmavVDJPXmOpdmUlVUsNTLnwa9roUckqN53oPuCgY",
	}
	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

func (cc *CloudController) Logs(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Second)
}
