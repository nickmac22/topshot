package topshot
import (
	"context"
// 	"fmt"
	// "github.com/rrrkren/topshot-sales/topshot"
	"github.com/onflow/flow-go-sdk/client"
	// "github.com/onflow/flow-go-sdk"
	"google.golang.org/grpc"
	// "github.com/onflow/cadence"
)

// func handleErr(err error) {
// 	if err != nil {
		
// 	}
// }
func Get_Moment_PlayID(account string, momentId string) int{
    // connect to flow
    flowClient, err := client.New("access.mainnet.nodes.onflow.org:9000", grpc.WithInsecure())
    handleErr(err)
    err = flowClient.Ping(context.Background())
    handleErr(err)

    // fetch latest block
    // latestBlock, err := flowClient.GetLatestBlock(context.Background(), false)
    // handleErr(err)
    
    // account := "0x3a62f7b838d5bd23"
	// momentId := "5921810"
    script := []byte(`

	import TopShot from 0x0b2a3299cc857e29

	// This script gets the metadata associated with a moment
	// in a collection by looking up its playID and then searching
	// for that play's metadata in the TopShot contract
	
	// Parameters:
	//
	// account: The Flow Address of the account whose moment data needs to be read
	// id: The unique ID for the moment whose data needs to be read
	
	// Returns: {String: String} 
	// A dictionary of all the play metadata associated
	// with the specified moment
	
	
	pub fun main(): UInt32 {

		let collectionRef = getAccount(`+account+`).getCapability(/public/MomentCollection)
			.borrow<&{TopShot.MomentCollectionPublic}>()
			?? panic("Could not get public moment collection reference")
	
		let token = collectionRef.borrowMoment(id: `+momentId + `)
			?? panic("Could not borrow a reference to the specified moment")
	
		let data = token.data
	
		return data.playID
	}
	`)

	// id := 5921810


	// address := cadence.Value(address1)
	ctx := context.Background()
    value, err := flowClient.ExecuteScriptAtLatestBlock(ctx, script, nil)
    // if err != nil {
    //     panic("failed to execute script")
    // }

    return value
}
