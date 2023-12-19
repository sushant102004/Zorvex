package agent

import (
	"github.com/rs/zerolog/log"
	"github.com/sushant102004/zorvex/internal/db"
	loadbalancer "github.com/sushant102004/zorvex/internal/load-balancer"
	"github.com/sushant102004/zorvex/internal/types"
)

type Agent interface {
	RegisterService(types.Service) error
	// Get all the services and send them to load balancer
	GetServicesData(string) ([]types.Service, error)

	// This function will serve client
	ServeClient(string) (string, error)
}

type ServiceAgent struct {
	db db.DBClient
	lb *loadbalancer.LoadBalancer
}

func NewServiceAgent(lb *loadbalancer.LoadBalancer, db db.DBClient) (*ServiceAgent, error) {
	return &ServiceAgent{
		db: db,
		lb: lb,
	}, nil
}

func (sa *ServiceAgent) RegisterService(data types.Service) error {
	if err := sa.db.AddNewServiceToDB(data); err != nil {
		log.Err(err).Msgf("unable to register new service")
		return err
	}
	return nil
}

func (sa *ServiceAgent) GetServicesData(name string) ([]types.Service, error) {
	svcInstances, err := sa.db.GetServiceInstances(name)
	if err != nil {
		log.Err(err).Msgf("unable to get all services")
		return nil, err
	}
	return svcInstances, nil
}

func (sa *ServiceAgent) ServeClient(service string) (string, error) {
	targetURL, err := sa.lb.RoundRobin(service)
	if err != nil {
		return "", err
	}
	return targetURL, nil
}
