package purchase

import (
	profileProto "asanpardakht/profile/api/proto"
	"google.golang.org/grpc"
	"log"
)

func NewMessagingClient(address string) profileProto.ProfileClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("message rpc client did not connect: %v", err)
	}
	return profileProto.NewProfileClient(conn)
}

func RunPurchaseService() (func(), int, error) {
	return nil, 0, nil
}
