package fbreach

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"fullmetalgo/communicator"
	model "fullmetalgo/fbreach/models"
	"fullmetalgo/fmdatabase"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/BurntSushi/toml"
)

type Int int

var httpCommunicator communicator.HttpCommunicator
var callApi = httpCommunicator.Communicate

func GetLocation() {

	loc := model.Location{Name: "Hyderabad", Country: "India"}
	fmt.Println(loc.Name)

	fmt.Println("sample")

}

type Response struct {
	Data Data
}

type Data struct {
	Users int
}

func FetchAndStoreReachInterest() {
	fmt.Println("Reading Target_spec json ")
	jsonV, err := ReadFromJson()
	checkErr(err)
	fmt.Println("Targer_spec json red successfly")
	fmt.Println("Reading target specs ")
	specs := ReadTargetSpecList(jsonV)
	db := fmdatabase.GetInstance()
	defer db.Close()
	var config reachPerInterestConfiguration

	_, err = toml.DecodeFile("/usr/share/fullmetal/fbreachconfig.toml", &config)
	checkErr(err)
	for i := len(specs) - 1; i >= 0; i-- {
		targetSpec := specs[i]
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)

			}
		}()

		/*if i%50 == 0 {
			time.Sleep(60000 * time.Millisecond)
		}
		*/
		fmt.Println("Processing location ", targetSpec.Geo_locations.Cities[0].Name)
		target_data, err := json.Marshal(targetSpec)
		if err == nil {
			encodeurl := FormEncodedUrl(string(target_data), config)
			//fmt.Println("Request = ", encodeurl)

			resonse, _ := callApi(encodeurl)
			fmt.Println("response::", resonse)
			var res Response
			if err := json.Unmarshal([]byte(resonse), &res); err != nil {
				fmt.Println("Panic error marshalling ::")
				panic(err)

			}
			data := res.Data.Users
			fmt.Println("data::", data)
			fmt.Println("Response = ", targetSpec)
			WriteToTable(targetSpec, data, db)

		}
	}
}

func ReadFromJson() (model.TargetingSpec, error) {
	var jsontype model.TargetingSpec
	err := json.Unmarshal(ReadFileFromPath("/usr/share/fullmetal/targeting_spec_all.json"), &jsontype)
	//fmt.Printf("Results: %v\n", val1)
	return jsontype, err
}

// Getting List of TargetSpecs

func ReadTargetSpecList(targetSpec model.TargetingSpec) []model.TargetingSpec {

	cities := targetSpec.Geo_locations.Cities
	totalCities := len(cities)
	interests := targetSpec.Flexible_spec[0].Interests
	totalIntrests := len(interests)

	genders := [][]int{{1}, {2}, {3}}
	genderTypes := len(genders)

	var allTargetingSpecs = make([]model.TargetingSpec, totalCities*genderTypes*totalIntrests)

	//fmt.Println(genderTypes)
	loopVar := 0
	for i := 0; i < totalCities; i++ {

		for j := 0; j < genderTypes; j++ {

			var city = make([]model.Location, 1)
			city[0] = cities[i]
			geoLocation := model.GeoLocation{Location_types: []string{"home", "recent"}, Cities: city}
			for k := totalIntrests - 1; k >= 0; k-- {
				var interset = make([]model.InterestSpec, 1)
				interset[0] = interests[k]
				var Flexible_spec = make([]model.FlexibleSpec, 1)
				Flexible_spec[0] = model.FlexibleSpec{Interests: interset}
				var trgSpec model.TargetingSpec
				if genders[j][0] != 3 {
					trgSpec = model.TargetingSpec{AgeMin: 18, AgeMax: 50, Genders: genders[j], Geo_locations: geoLocation, Flexible_spec: Flexible_spec}
				} else {
					trgSpec = model.TargetingSpec{AgeMin: 18, AgeMax: 50, Geo_locations: geoLocation, Flexible_spec: Flexible_spec}
				}

				allTargetingSpecs[loopVar] = trgSpec
				loopVar += 1
			}
		}

	}
	return allTargetingSpecs
}

type reachPerInterestConfiguration struct {
	Url, Access_token, Other_paramerters string
}

func FormEncodedUrl(targeting_spec string, config reachPerInterestConfiguration) string {
	var buffer bytes.Buffer
	buffer.WriteString(config.Url)
	buffer.WriteString("access_token=")
	buffer.WriteString(config.Access_token)
	buffer.WriteString("&")
	buffer.WriteString(config.Other_paramerters)
	buffer.WriteString("&")
	buffer.WriteString("targeting_spec=")
	buffer.WriteString(url.QueryEscape(targeting_spec))
	return buffer.String()
}

func ReadFileFromPath(filePath string) []byte {
	file, e := ioutil.ReadFile(filePath)
	if e != nil {
		fmt.Printf("Error while reading : %v\n", e)
		os.Exit(1)

	}
	return file
}

//Sample method to execute a test case.
func Area(r Int) Int {
	return r * r
}

func WriteToTable(targetSpec model.TargetingSpec, reach int, db *sql.DB) {
	var gender string
	val := targetSpec.Genders
	if val != nil {

		if val[0] == 1 {
			gender = "male"
		} else {
			gender = "female"
		}
	} else {
		gender = "both"
	}
	stmt, err := db.Prepare("INSERT fb_city_gender_interest SET gender=?, city=?, interest_id=?, interest=?, reach=?")
	checkErr(err)

	_, err = stmt.Exec(gender, targetSpec.Geo_locations.Cities[0].Name, targetSpec.Flexible_spec[0].Interests[0].Id, targetSpec.Flexible_spec[0].Interests[0].Name, reach)

	fmt.Println("Inserted into db sucessfully ")

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
