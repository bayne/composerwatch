package packagist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type responseJson struct {
	Package struct {
		Versions map[string]*PackageVersion `json:"versions"`
	} `json:"package"`
}

type PackageVersionLookup struct {
}

func NewPackageVersionLookup() *PackageVersionLookup {
	var packageVersionLookup PackageVersionLookup
	return &packageVersionLookup
}

func (packageVersionLookup *PackageVersionLookup) LookupAll(name string) (map[string]*PackageVersion, error) {
	url := fmt.Sprintf("https://packagist.org/packages/%s.json", name)
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)

	var responseJson responseJson

	json.Unmarshal(body, &responseJson)

	return responseJson.Package.Versions, nil

}
