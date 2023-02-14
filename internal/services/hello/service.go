package hello

type ServiceWrapper interface {
	Hello()
}

type Service struct {
	// Define specific service needs such as:
	// DB *gorm.DB
}

func NewHelloService() *Service {
	return &Service{}
}
