package composer

import "encoding/json"

type lockPackageJson struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type LockPackage struct {
	jsonData lockPackageJson
}

func (lockPackage *LockPackage) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &lockPackage.jsonData)
}

func (lockPackage *LockPackage) Name() string {
	return lockPackage.jsonData.Name
}

func (lockPackage *LockPackage) Version() string {
	return lockPackage.jsonData.Version
}
