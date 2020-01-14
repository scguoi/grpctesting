package impl

import (
	"context"
	"log"

	"github.com/scguoi/grpctesting/example"
)

type Handler struct {
}

func (h *Handler) Stream(stream example.H3Wrapper_StreamServer) error {
	return (&realHandler{}).Service(stream)
}

func (h *Handler) OneWay(ctx context.Context, req *example.H3Req) (*example.H3Resp, error) {
	return (&realHandler{}).OneWay(ctx, req)
}

type realHandler struct {
}

func (h *realHandler) Service(stream example.H3Wrapper_StreamServer) error {
	for {
		req, err := stream.Recv()

		if err != nil {
			log.Printf("got err %v", err)
			return nil
		}

		switch req.MsgType {
		case "header":
			log.Printf("got header frame %v", req)
			stream.Send(&example.H3Resp{MsgType: "header", Fin: false})
		case "data":
			log.Printf("got data frame %v", req)
			stream.Send(&example.H3Resp{MsgType: "data", Body: []byte("server return"), Fin: false})
		case "finish":
			log.Printf("got finish frame %v", req)
			stream.Send(&example.H3Resp{MsgType: "data", Body: []byte("server return"), Fin: true})
		}
	}
}

func (h *realHandler) OneWay(ctx context.Context, req *example.H3Req) (*example.H3Resp, error) {
	log.Printf("got one way request %v", req)
	return &example.H3Resp{MsgType: "data", Body: []byte("server return"), Fin: true}, nil
}
