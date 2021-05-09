package topshot
import (
	"context"
	"fmt"
	"github.com/onflow/flow-go-sdk/client"
	"google.golang.org/grpc"
)

func handleErr(err error) {
	if err != nil {
		
	}
}
func Get_All_Moments() {
    // connect to flow
    flowClient, err := client.New("access.mainnet.nodes.onflow.org:9000", grpc.WithInsecure())
    handleErr(err)
    err = flowClient.Ping(context.Background())
    handleErr(err)
	   
    script := []byte(`

	import TopShot from 0x0b2a3299cc857e29

	// This script returns an array of all the plays 
	// that have ever been created for Top Shot

	// Returns: [TopShot.Play]
	// array of all plays created for Topshot

	pub fun main(): [TopShot.Play] {

		return TopShot.getAllPlays()
	}
	`)
	ctx := context.Background()
    value, err := flowClient.ExecuteScriptAtLatestBlock(ctx, script, nil)
    if err != nil {
        panic("failed to execute script")
    }

    return value
}
