package fbreach

import (
	_ "encoding/json"
	_ "fmt"
	_ "fullmetalgo/communicator"
	_ "fullmetalgo/fmdatabase"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetLocation(t *testing.T) {

	var v = Area(10)
	if v != 100 {
		t.Error("Expected 100, got ", v)
	}

	//var res = "{'data': {'users': 24000,'estimate_ready': true}}"
	//httpCommunicator = res
	callApi = samplehello
	FetchAndStoreReachInterest()
	//fmt.Println("The response is ::", response)

	/*	jsonV := ReadFromJson()

		cities := jsonV.Geo_locations.Cities
		fmt.Println("This is done.",len(cities))
		specs := ReadTargerSpecList(jsonV)
		//fmt.Println("This is done.",specs)
		//printAllJson(specs)
		b, _:= json.Marshal(specs[0])
		decodedUrl := formDecodedUrl(string(b))
		fmt.Println("")
		fmt.Println("Encoded url ",decodedUrl)
		_ = ReachPerInterest(decodedUrl)
	*/ //fmt.Println("the file object is ::" + sample)
}

func samplehello(url string) (string, error) {
	//var responsebody string
	//fmt.sprintf("{data:{user:232323}}")

	return "{\"data\": {\"users\":24000}}", nil
}
