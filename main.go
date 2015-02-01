package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bayne/composerwatch/composer"
	"github.com/bayne/composerwatch/packagist"
	"io/ioutil"
)

type PackageVersionLookup interface {
	LookupAll(name string) (packageVersions map[string]*packagist.PackageVersion, err error)
}

type UpgradeReport map[string][]*packagist.PackageVersion

func NewUpgradeReport() *UpgradeReport {
	upgradeReport := make(UpgradeReport)
	return &upgradeReport
}

func (this UpgradeReport) add(current *packagist.PackageVersion, newer *packagist.PackageVersion) {
	this[current.Name()] = append(this[current.Name()], newer)
}

func CreateUpgradeReport(contents []byte, onlyStable bool) (*UpgradeReport, error) {

	upgradeReport := NewUpgradeReport()

	var lockFile composer.LockFile

	if err := json.Unmarshal(contents, &lockFile); err != nil {
		return nil, err
	}

	var packageVersionLookup PackageVersionLookup

	packageVersionLookup = packagist.NewPackageVersionLookup()

	for _, lockFilePackage := range lockFile.Packages() {
		packageVersions, err := packageVersionLookup.LookupAll(lockFilePackage.Name())
		if err != nil {
			return nil, err
		}
		lockFilePackageVersion := packageVersions[lockFilePackage.Version()]

		for _, packageVersion := range packageVersions {
			if packageVersion.Compare(lockFilePackageVersion) > 0 {
				if onlyStable {
					if packageVersion.Stable() {
						upgradeReport.add(lockFilePackageVersion, packageVersion)
					} else {
						continue
					}
				} else {
					upgradeReport.add(lockFilePackageVersion, packageVersion)
				}
			}
		}

	}

	return upgradeReport, nil
}

func main() {
	filename := flag.String("file", "composer.lock", "The path to the composer.lock file")
	onlyStable := flag.Bool("stable", true, "Only show stable packages")
	flag.Parse()

	contents, err := ioutil.ReadFile(*filename)
	if err != nil {
		panic(err)
	}

	upgradeReport, err := CreateUpgradeReport(contents, *onlyStable)

	if err != nil {
		panic(err)
	}

	output, err := json.Marshal(upgradeReport)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))
}
