package services_test

import (
	"context"
	"testing"

	"github.com/amerikarno/exam7/services"
	"github.com/amerikarno/exam7/services/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite

	ctx context.Context

	controller *gomock.Controller
	mDomain    *mock.MockIDomain
	service    *services.Service
}

func (s *ServiceTestSuite) SetupTest() {
	s.controller = gomock.NewController(s.T())

	s.mDomain = mock.NewMockIDomain(s.controller)
	s.service = services.New(s.mDomain)
}

func TestIntegratedTest(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (s *ServiceTestSuite) TestService() {
	expect := map[string]int{
		"test": 1,
	}
	s.mDomain.EXPECT().GetResponseBody().Return([]byte("test"))

	actual := s.service.BeefSummaryService()
	s.Equal(actual, expect)
}


