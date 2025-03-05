package adapter

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Adapter struct {
	// * adapter list
	Postgres    *gorm.DB
	GrcpClient  *grpc.ClientConn
	KafkaGo     *kafka.Writer
	MinioClient *minio.Client
}

type Option interface {
	Start(a *Adapter) error
	Stop() error
}

// var Adapters *Adapter

func NewAdapter(opts ...Option) (*Adapter, error) {
	a := &Adapter{}
	var errs []error

	for _, opt := range opts {
		if err := opt.Start(a); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return nil, fmt.Errorf("failed to start adapter: %v", errs)
	}

	return a, nil
}

func (a *Adapter) Close(opts ...Option) error {
	var errs []error

	for _, opt := range opts {
		if err := opt.Stop(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("failed to stop adapter: %v", errs)
	}
	return nil
}
