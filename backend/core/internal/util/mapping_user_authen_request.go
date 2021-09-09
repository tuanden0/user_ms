package util

import "user_ms/backend/core/api"

func MapUserLoginRequest(in *api.UserLoginRequest) (string, string) {
	return in.GetUsername(), in.GetPassword()
}

func MapUserLoginResponse(token string) *api.UserLoginResponse {
	return &api.UserLoginResponse{
		AccessToken: token,
	}
}
