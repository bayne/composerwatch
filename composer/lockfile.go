package composer

import "encoding/json"

type lockFileJson struct {
	Readme   []string      `json:"_readme"`
	Hash     string        `json:"hash"`
	Packages []LockPackage `json:"packages"`
	Aliases  interface{}   `json:"aliases"`
}

type LockFile struct {
	jsonData lockFileJson
}

func (lockFile *LockFile) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &lockFile.jsonData)
}

func (lockFile *LockFile) Hash() string {
	return lockFile.jsonData.Hash
}

func (lockFile *LockFile) Packages() []LockPackage {
	return lockFile.jsonData.Packages
}
