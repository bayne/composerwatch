package packagist

import (
	"encoding/json"
	"github.com/mcuadros/go-version"
	"regexp"
)

type packageVersionJson struct {
	Name              string `json:"name"`
	VersionNormalized string `json:"version_normalized"`
	Version           string `json:"version"`
}

type PackageVersion struct {
	jsonData packageVersionJson
}

func (packageVersion *PackageVersion) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &packageVersion.jsonData)
}

func (packageVersion *PackageVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(packageVersion.jsonData)
}

func (packageVersion *PackageVersion) Name() string {
	return packageVersion.jsonData.Name
}

func (packageVersion *PackageVersion) VersionNormalized() string {
	return packageVersion.jsonData.VersionNormalized
}

func (packageVersion *PackageVersion) Compare(otherPackageVersion *PackageVersion) int {
	return version.CompareSimple(packageVersion.VersionNormalized(), otherPackageVersion.VersionNormalized())
}

func (packageVersion *PackageVersion) Version() string {
	return packageVersion.jsonData.Version
}

func (packageVersion *PackageVersion) Stable() bool {
	result, err := regexp.Match("^([0-9]+\\.?){4}$", []byte(packageVersion.VersionNormalized()))
	if err != nil {
		panic(err)
	}
	return result
}
