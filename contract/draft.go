package contract

import (
	"io/ioutil"

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type Draft struct {
	contract *Contract
}

func NewContract(service, consumer string) *Draft {
	return &Draft{&Contract{
		Service:      service,
		Consumer:     consumer,
		Interactions: make([]*Interaction, 0),
	}}
}

func (d *Draft) AddInteraction(method string, req proto.Message, response proto.Message) error {
	reqpb, err := anypb.New(req)
	if err != nil {
		return err
	}
	responsepb, err := anypb.New(response)
	if err != nil {
		return err
	}
	d.contract.Interactions = append(d.contract.Interactions, &Interaction{
		Method:   method,
		Request:  reqpb,
		Response: responsepb,
	})
	return nil
}

func (d *Draft) Commit() error {
	content, err := prototext.MarshalOptions{Multiline: true}.Marshal(d.contract)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(d.contract.Consumer+".textproto", content, 0777)
}

// TODO: save to Cloud
func (d *Draft) CommitTo(url string) error {
	return nil
}
