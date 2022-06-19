package contract

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/google/uuid"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type Draft struct {
	contract *Contract // the reference to the contract under draft.
}

func NewContract(service, consumer string) *Draft {
	return &Draft{&Contract{
		Service:      service,
		Consumer:     consumer,
		Interactions: make([]*Interaction, 0),
	}}
}

func (d *Draft) AddInteraction(
	name, method string, req proto.Message, response proto.Message,
	wantError bool, rules *CompositeRule, preconditions []string) error {
	// cast request message into *anypb.Any.
	reqpb, err := anypb.New(req)
	if err != nil {
		return err
	}
	// cast expected response message into *anypb.Any.
	responsepb, err := anypb.New(response)
	if err != nil {
		return err
	}
	d.contract.Interactions = append(d.contract.Interactions, &Interaction{
		Name:          name,
		Method:        method,
		Request:       reqpb,
		Response:      responsepb,
		WantError:     wantError,
		Rules:         rules,
		Preconditions: preconditions,
	})
	return nil
}

func (d *Draft) PublishLocal(path string) error {
	content, err := prototext.MarshalOptions{Multiline: true}.Marshal(d.contract)
	if err != nil {
		return err
	}
	os.MkdirAll(path, os.ModePerm)
	fName := filepath.Join(path, fmt.Sprintf("%s.prototxt", d.contract.Consumer))
	if err := ioutil.WriteFile(fName, content, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func (d *Draft) PublishRemote(dryrun bool) error {
	content, err := prototext.MarshalOptions{Multiline: true}.Marshal(d.contract)
	if err != nil {
		return err
	}
	if dryrun {
		log.Printf("The prototxt to be published is: %v", string(content))
		return nil
	}
	ctx := context.Background()
	ops := option.WithCredentialsFile("cred.json")
	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: "grpc-contract-test",
	}, ops)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)
	}
	version := fmt.Sprintf("%s_%s", time.Now().Format("2006-01-02"), uuid.New().String())
	doc := client.Collection("services").Doc(d.contract.Service).Collection("consumers").Doc(d.contract.Consumer)
	if _, err := doc.Set(ctx,
		map[string]string{
			"latest": version,
			version:  string(content),
		}, firestore.MergeAll); err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		return fmt.Errorf("error has occurred while saving contract to Firestore: %v", err)
	}
	log.Printf("Contract successfully saved to Firestore for %v", d.contract.Consumer)
	return nil
}

// Obtains a connection to the contract as server.
func NewContractConn(path string) (grpc.ClientConnInterface, error) {
	contract, err := ReadConctract(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read the contract: %w", err)
	}
	return &ContractConn{contract}, nil
}

type ContractConn struct {
	contract *Contract
}

func (c *ContractConn) Invoke(ctx context.Context, method string,
	req interface{}, resp interface{}, opts ...grpc.CallOption) error {
	// method is formatted as "/pkgName.ServiceName/MethodName".
	methodName := method[strings.LastIndex(method, "/")+1:]
	// Pack the req into Anypb for equality test with interaction.Request.
	reqpb, err := anypb.New(req.(proto.Message))
	if err != nil {
		return fmt.Errorf("unable to convert req %v into Anypb: %w", req, err)
	}
	for _, interaction := range c.contract.Interactions {
		if interaction.Method == methodName &&
			proto.Equal(reqpb, interaction.Request) {
			if interaction.WantError {
				return errors.New("expected error from contract")
			}
			err := interaction.Response.UnmarshalTo(resp.(proto.Message))
			if err != nil {
				return fmt.Errorf("unable to unmarshall response %v: %w",
					interaction.Response, err)
			}
			return nil
		}
	}
	return fmt.Errorf("interaction not found for method %q and req %v",
		methodName, req)
}

func (c *ContractConn) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string,
	opts ...grpc.CallOption) (grpc.ClientStream, error) {
	// Unimplemented.
	return nil, nil
}
