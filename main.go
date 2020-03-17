package main

import (
	"context"
	"encoding/hex"
	"fmt"
	tb "gRpcProject/types"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address		= "ac.testnet.libra.org:8000"
)

func main() {
	fmt.Println("gRPC Client Test")

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := tb.NewAdmissionControlClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 100)
	defer cancel()

	nAccount, err := hex.DecodeString("8b5e7635746e7dd17ed988cff739bfe3dea25eb783b85708cd5572ab9e94a3de")
	if err != nil {
		fmt.Println("Err:", err)
	}
	reqItems := []*tb.RequestItem{
		&tb.RequestItem{
			RequestedItems: &tb.RequestItem_GetAccountStateRequest{
				GetAccountStateRequest: &tb.GetAccountStateRequest{
					Address: nAccount }, // #end GetAccountStateRequest
			}, // #end RequestedItems
		}, // #end RequestItem
	}

	ReqData := tb.UpdateToLatestLedgerRequest{
		ClientKnownVersion:   0,
		RequestedItems:       reqItems,
	}

	ResData, err := c.UpdateToLatestLedger(ctx, &ReqData)
	if err != nil {
		fmt.Println("UpdateToLatestLedger Error:", err)
		return
	}
	fmt.Println("ResData:", ResData)

}
