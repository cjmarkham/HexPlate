package forum

import (
	"context"
	"github.com/cjmarkham/hexplate/cmd/api"
	"github.com/cjmarkham/hexplate/internal/domain/forum"
	"github.com/cjmarkham/hexplate/internal/domain/testutils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestHandlers_CreateForum(t *testing.T) {
	tests := map[string]struct {
		mockCreate       testutils.MockParameters
		request          api.CreateForumRequestObject
		expectedError    error
		expectedResponse api.CreateForum201JSONResponse
	}{
		"it succeeds": {
			request: api.CreateForumRequestObject{
				Body: &api.CreateForumJSONRequestBody{
					Name: "Test Forum",
				},
			},
			mockCreate: testutils.MockParameters{
				Response: &forum.Forum{
					Name: "Test Forum",
				},
			},
			expectedResponse: api.CreateForum201JSONResponse{
				Id:   uuid.MustParse("8ad1ab90-3371-4085-b1ea-0b1c1a39f52a"),
				Name: "Test Forum",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockService := forum.NewMockService(ctrl)

			mockService.EXPECT().
				Create(gomock.Any(), gomock.Any()).
				Return(test.mockCreate.Response, test.mockCreate.Error).
				Times(test.mockCreate.Times)

			handler := ProvideHandlers(mockService)

			response, err := handler.CreateForum(context.Background(), test.request)
			assert.Equal(t, test.expectedResponse, response)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}
