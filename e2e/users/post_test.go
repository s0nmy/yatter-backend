package users_test

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"
	"yatter-backend-go/app/server"
	"yatter-backend-go/e2e/testutil"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type UserResponse struct {
	ID             int     `json:"id"`
	Username       string  `json:"username"`
	DisplayName    *string `json:"display_name"`
	CreatedAt      *string `json:"created_at"`
	FollowersCount *int    `json:"followers_count"`
	FollowingCount *int    `json:"following_count"`
	Note           *string `json:"note"`
	Avatar         *string `json:"avatar"`
	Header         *string `json:"header"`
}

func Test_POST_v1_users_Success(t *testing.T) {
	client := testutil.NewTestClient(t)
	t.Setenv("PORT", strconv.Itoa(client.Port()))
	go func() { _ = server.Run(client.DB()) }()

	testCases := []struct {
		name           string
		requestBody    map[string]any
		expectedStatus int
		validateFunc   func(t *testing.T, resp *resty.Response)
	}{
		{
			name: "正常系: 正常なユーザ名とパスワードの場合、ユーザーが作成される",
			requestBody: map[string]any{
				"username": "test_user",
				"password": "P@ssw0rdああ",
			},
			expectedStatus: 201,
			validateFunc: func(t *testing.T, resp *resty.Response) {
				var user UserResponse
				err := json.Unmarshal(resp.Body(), &user)
				require.NoError(t, err)
				assert.Equal(t, "test_user", user.Username)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := client.R().
				SetBody(tc.requestBody).
				Post("/users")

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, resp.StatusCode())

			tc.validateFunc(t, resp)
		})
	}
}

func Test_POST_v1_users_ユーザー名が重複している場合409エラーが返される(t *testing.T) {
	client := testutil.NewTestClient(t)
	t.Setenv("PORT", strconv.Itoa(client.Port()))
	go func() { _ = server.Run(client.DB()) }()

	// あらかじめユーザーを作成
	resp, err := client.R().
		SetBody(map[string]any{
			"username": "duplicateuser",
			"password": "P@ssw0rd",
		}).
		Post("/users")

	require.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode())

	// 同じユーザー名で再度作成を試みる
	resp, err = client.R().
		SetBody(map[string]any{
			"username": "duplicateuser",
			"password": "P@ssw0rd",
		}).
		Post("/users")

	require.NoError(t, err)
	assert.Equal(t, 409, resp.StatusCode())
}

func Test_POST_v1_users_Error(t *testing.T) {
	client := testutil.NewTestClient(t)
	t.Setenv("PORT", strconv.Itoa(client.Port()))
	go func() { _ = server.Run(client.DB()) }()

	testCases := []struct {
		name           string
		requestBody    map[string]any
		expectedStatus int
	}{
		{
			name: "異常系: 空のユーザ名の場合、400エラーが返される",
			requestBody: map[string]any{
				"username": "",
				"password": "P@ssw0rd",
			},
			expectedStatus: 400,
		},
		{
			name: "異常系: 50文字を超えるユーザ名の場合、400エラーが返される",
			requestBody: map[string]any{
				"username": strings.Repeat("a", 51),
				"password": "P@ssw0rd",
			},
			expectedStatus: 400,
		},
		{
			name: "異常系: 半角英数字、アンダースコア以外を含むユーザ名の場合、400エラーが返される",
			requestBody: map[string]any{
				"username": "user@$name",
				"password": "P@ssw0rd",
			},
			expectedStatus: 400,
		},
		{
			name: "異常系: 空のパスワードの場合、400エラーが返される",
			requestBody: map[string]any{
				"username": "validuser",
				"password": "",
			},
			expectedStatus: 400,
		},
		{
			name: "異常系: 72byteを超えるパスワードの場合、400エラーが返される",
			requestBody: map[string]any{
				"username": "validuser",
				"password": strings.Repeat("a", 73),
			},
			expectedStatus: 400,
		},
		{
			name: "異常系: 不正なJSONの場合、400エラーが返される",
			requestBody: map[string]any{
				"invalid": "json",
			},
			expectedStatus: 400,
		},
		{
			name: "異常系: usernameが省略されている場合、400エラーが返される",
			requestBody: map[string]any{
				"password": "P@ssw0rd",
			},
			expectedStatus: 400,
		},
		{
			name: "異常系: passwordが省略されている場合、400エラーが返される",
			requestBody: map[string]any{
				"username": "testuser",
			},
			expectedStatus: 400,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := client.R().
				SetBody(tc.requestBody).
				Post("/users")

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, resp.StatusCode())
		})
	}
}
