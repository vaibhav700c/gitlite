// Copyright 2026 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package integration

import (
	"net/http"
	"strconv"
	"testing"

	auth_model "gitea.dev/models/auth"
	"gitea.dev/modules/setting"
	api "gitea.dev/modules/structs"
	"gitea.dev/modules/test"
	"gitea.dev/tests"

	"github.com/stretchr/testify/assert"
)

func TestAPIAuthorizationScheme(t *testing.T) {
	defer tests.PrepareTestEnv(t)()
	token := getUserToken(t, "user2", auth_model.AccessTokenScopeReadUser)

	t.Run("BearerIsCanonical", func(t *testing.T) {
		req := NewRequest(t, "GET", "/api/v1/user").SetHeader("Authorization", "Bearer "+token)
		resp := MakeRequest(t, req, http.StatusOK)
		assert.Empty(t, resp.Header().Get("Deprecation"))
		assert.Empty(t, resp.Header().Get("Sunset"))
		assert.Empty(t, resp.Header().Get("Warning"))
	})

	t.Run("LegacyTokenSchemeIsDeprecated", func(t *testing.T) {
		req := NewRequest(t, "GET", "/api/v1/user").SetHeader("Authorization", "token "+token)
		resp := MakeRequest(t, req, http.StatusOK)
		assert.Equal(t, "user2", DecodeJSON(t, resp, &api.User{}).UserName)
		assert.Equal(t, "true", resp.Header().Get("Deprecation"))
		assert.Equal(t, "Wed, 01 Apr 2026 00:00:00 GMT", resp.Header().Get("Sunset"))
		assert.Equal(t, `299 - "The 'token' authorization scheme is deprecated; use 'Authorization: Bearer <token>'"`, resp.Header().Get("Warning"))
	})

	t.Run("LegacyTokenSchemeDisabled", func(t *testing.T) {
		defer test.MockVariableValue(&setting.API.AllowLegacyTokenScheme, false)()
		req := NewRequest(t, "GET", "/api/v1/user").SetHeader("Authorization", "token "+token)
		resp := MakeRequest(t, req, http.StatusUnauthorized)
		apiErr := DecodeJSON(t, resp, &struct {
			Message string `json:"message"`
			URL     string `json:"url"`
		}{})
		assert.Equal(t, "the 'token' authorization scheme is disabled; use 'Authorization: Bearer <token>'", apiErr.Message)
		assert.NotEmpty(t, apiErr.URL)
	})
}

func TestAPIPaginationDefaults(t *testing.T) {
	defer tests.PrepareTestEnv(t)()

	req := NewRequest(t, "GET", "/api/v1/repos/search?private=false")
	resp := MakeRequest(t, req, http.StatusOK)
	var body api.SearchResults
	DecodeJSON(t, resp, &body)

	total, err := strconv.Atoi(resp.Header().Get("X-Total-Count"))
	assert.NoError(t, err)
	assert.Greater(t, total, 20, "fixtures must have more public repos than one default page")
	assert.Len(t, body.Data, 20)
	assert.Equal(t, "1", resp.Header().Get("X-Page"))
	assert.Equal(t, "20", resp.Header().Get("X-PerPage"))
	assert.Equal(t, "true", resp.Header().Get("X-HasMore"))

	lastPage := (total + 19) / 20
	req = NewRequest(t, "GET", "/api/v1/repos/search?private=false&page="+strconv.Itoa(lastPage))
	resp = MakeRequest(t, req, http.StatusOK)
	assert.Equal(t, "false", resp.Header().Get("X-HasMore"))
}
