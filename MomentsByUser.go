package topshot
import (
	"context"
	// "fmt"
	// "github.com/rrrkren/topshot-sales/topshot"
	"github.com/onflow/flow-go-sdk/client"
	"google.golang.org/grpc"
	"github.com/onflow/cadence"
)


func Moments_By_User(userAddress string) cadence.Value {
    // connect to flow
    flowClient, err := client.New("access.mainnet.nodes.onflow.org:9000", grpc.WithInsecure())
    handleErr(err)
    err = flowClient.Ping(context.Background())
    handleErr(err)

    // fetch latest block
    // latestBlock, err := flowClient.GetLatestBlock(context.Background(), false)
    // handleErr(err)
    
	// accounts := [0x3a62f7b838d5bd23, ]
    script := []byte(`

	import TopShot from 0x0b2a3299cc857e29

	// This is the script to get a list of all the moments' ids an account owns
	// Just change the argument to  to whatever account you want
	// and as long as they have a published Collection receiver, you can see
	// the moments they own.
	
	// Parameters:
	//
	// account: The Flow Address of the account whose moment data needs to be read
	
	// Returns: [UInt64]
	// list of all moments' ids an account owns
	
	pub fun main(): [UInt64] {
	
		let acct = getAccount(` + userAddress + `)
	
		let collectionRef = acct.getCapability(/public/MomentCollection)
								.borrow<&{TopShot.MomentCollectionPublic}>()!
	
		log(collectionRef.getIDs())
	
		return collectionRef.getIDs()
	}
	`)
	ctx := context.Background()
    value, err := flowClient.ExecuteScriptAtLatestBlock(ctx, script, nil)
    // if err != nil {
    //     panic("failed to execute script")
    // }

    return value
}