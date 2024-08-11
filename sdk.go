package streamdk

import (
	"context"
	"github.com/usedatabrew/golang-stream-sdk/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type StreamSdk struct {
	StreamClient gen.StreamClient
	conn         *grpc.ClientConn
	opts         *Options
}

func NewStreamSdk(options ...Option) *StreamSdk {
	opts := NewOptions(options...)

	return &StreamSdk{
		opts: opts,
	}
}

func (s *StreamSdk) Connect() error {
	conn, err := grpc.NewClient(s.opts.StreamHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	s.conn = conn

	s.StreamClient = gen.NewStreamClient(conn)
	return nil
}

func (s *StreamSdk) Subscribe(ctx context.Context, pipelineId string, response chan gen.StreamResponse) error {
	authCtx := s.getRequestCtx(ctx)
	stream, err := s.StreamClient.GetStream(authCtx, &gen.StreamRequest{
		Id:      pipelineId,
		AutoAck: true,
	})
	if err != nil {
		return err
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}

		if msg == nil {
			break
		} else {
			response <- *msg
		}
	}

	return nil
}

func (s *StreamSdk) getRequestCtx(ctx context.Context) context.Context {
	return addCustomHeader(ctx, "custom-header-key", "custom-header-value")
}

func (s *StreamSdk) Close() error {
	return s.conn.Close()
}

func addCustomHeader(ctx context.Context, key, value string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, key, value)
}
