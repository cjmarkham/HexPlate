package pet

import (
	"context"
	"github.com/cjmarkham/hexplate/internal/testdata"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestService_Create(t *testing.T) {
	tests := map[string]struct {
		dto            DTO
		expectedCreate testdata.MockParameters
		expectedResult *Pet
		expectedError  error
	}{
		"it creates a pet": {
			dto: DTO{
				Name: "Lucy",
			},
			expectedCreate: testdata.MockParameters{
				Response: &Pet{
					Name: "Lucy",
				},
				Times: 1,
			},
			expectedResult: &Pet{
				Name: "Lucy",
			},
		},
		"it validates name min length": {
			dto: DTO{
				Name: "A",
			},
			expectedCreate: testdata.MockParameters{
				Times: 0,
			},
			expectedResult: nil,
			expectedError:  ErrNameMinLength,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockRepository := NewMockRepository(ctrl)

			mockRepository.EXPECT().
				Create(gomock.Any(), gomock.Any()).
				Return(test.expectedCreate.Response, test.expectedCreate.Error).
				Times(test.expectedCreate.Times)

			svc := ProvideService(mockRepository)
			result, err := svc.Create(context.Background(), test.dto)

			assert.Equal(t, test.expectedResult, result)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}
