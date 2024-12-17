package pet

import (
	"context"
	"github.com/cjmarkham/hexplate/cmd/api"
	"github.com/cjmarkham/hexplate/internal/domain/pet"
	"github.com/cjmarkham/hexplate/internal/testdata"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestHandlers_CreateForum(t *testing.T) {
	tests := map[string]struct {
		mockCreate       testdata.MockParameters
		request          api.CreatePetRequestObject
		expectedError    error
		expectedResponse api.CreatePet201JSONResponse
	}{
		"it succeeds": {
			request: api.CreatePetRequestObject{
				Body: &api.CreatePetJSONRequestBody{
					Data: api.CreatePet{
						Attributes: api.PetCreateAttributes{
							Name: "Test Pet",
						},
					},
				},
			},
			mockCreate: testdata.MockParameters{
				Response: &pet.Pet{
					Name: "Test Pet",
				},
			},
			expectedResponse: api.CreatePet201JSONResponse{
				Data: api.PetResponse{
					Id: uuid.MustParse("8ad1ab90-3371-4085-b1ea-0b1c1a39f52a"),
					Attributes: api.PetResponseAttributes{
						Name: "Test Pet",
					},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockService := pet.NewMockService(ctrl)

			mockService.EXPECT().
				Create(gomock.Any(), gomock.Any()).
				Return(test.mockCreate.Response, test.mockCreate.Error).
				Times(test.mockCreate.Times)

			handler := ProvideHandlers(mockService)

			response, err := handler.CreatePet(context.Background(), test.request)
			assert.Equal(t, test.expectedResponse, response)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}
