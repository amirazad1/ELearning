package services

import (
	"github.com/amirazad1/ELearning/api/dto"
	"github.com/amirazad1/ELearning/config"
	"github.com/amirazad1/ELearning/pkg/logging"
)

type UserService struct {
	logger     logging.Logger
	cfg        *config.Config
	otpService *OtpService
}

func NewUserService(cfg *config.Config) *UserService {
	logger := logging.NewLogger(cfg)
	return &UserService{
		cfg:        cfg,
		logger:     logger,
		otpService: NewOtpService(cfg),
	}
}

func (s *UserService) SendOtp(req *dto.GetOtpRequest) error {
	err := s.otpService.SendOtp(req.MobileNumber)
	if err != nil {
		return err
	}
	return err
}
