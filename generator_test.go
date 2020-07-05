package drivinglicence_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type DrivingLicenceSuite struct {
	suite.Suite
}

func TestDrivingLicenceSuite(t *testing.T) {
	suite.Run(t, new(DrivingLicenceSuite))
}
