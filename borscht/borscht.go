package borscht

import (
	"io/ioutil"
	"path/filepath"
)

type Borscht struct{}

func (b *Borscht) Diff(releasePath, fromVersion, toVersion string) (map[string]string, error) {
	jobDirs, err := ioutil.ReadDir(filepath.Join(releasePath, "jobs"))
	if err != nil {
		return nil, err
	}
	jobDiffs := map[string]string{}
	for _, jobDir := range jobDirs {
		jobDiffs[jobDir.Name()] = ""
	}
	return jobDiffs, nil
}
