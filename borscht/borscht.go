package borscht

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func Diff(releasePath, fromVersion, toVersion string) (map[string]string, error) {
	jobDirs, err := ioutil.ReadDir(filepath.Join(releasePath, "jobs"))
	if err != nil {
		return nil, err
	}
	jobDiffs := map[string]string{}
	for _, jobDir := range jobDirs {
		jobDiff, err := getDiff(releasePath, jobDir.Name(), fromVersion, toVersion)
		if err != nil {
			return nil, err
		}
		jobDiffs[jobDir.Name()] = jobDiff
	}
	return jobDiffs, nil
}

func getDiff(releasePath, jobName, fromVersion, toVersion string) (string, error) {
	fromRef, err := releaseCommit(releasePath, fromVersion)
	if err != nil {
		return "", err
	}
	toRef, err := releaseCommit(releasePath, toVersion)
	if err != nil {
		return "", err
	}
	return jobSpecGitDiff(releasePath, jobName, fromRef, toRef)
}

type finalRelease struct {
	CommitHash string `yaml:"commit_hash"`
}

func releaseCommit(releasePath, releaseVersion string) (string, error) {
	errs := func(action string, err error) (string, error) {
		return "", fmt.Errorf("%s: %s", action, err)
	}

	finalReleasesDir := filepath.Join(releasePath, "releases")
	finalReleaseNames, err := ioutil.ReadDir(finalReleasesDir)
	if err != nil {
		return errs("listing releases", err)
	}
	if len(finalReleaseNames) != 1 {
		return "", fmt.Errorf("expected 1 final release directory in %s, got %d", finalReleasesDir, len(finalReleaseNames))
	}
	finalReleaseName := finalReleaseNames[0].Name()

	releaseYamlPath := filepath.Join(finalReleasesDir, finalReleaseName, fmt.Sprintf("%s-%s.yml", finalReleaseName, releaseVersion))
	releaseYaml, err := ioutil.ReadFile(releaseYamlPath)
	if err != nil {
		return errs("reading final release file", err)
	}
	var release finalRelease
	if err := yaml.Unmarshal(releaseYaml, &release); err != nil {
		return errs("parsing final release file", err)
	}
	return release.CommitHash, nil
}

func jobSpecGitDiff(releasePath, jobName, fromRef, toRef string) (string, error) {
	gitCmd := exec.Command("git", "-C", releasePath, "diff",
		fmt.Sprintf("%s..%s", fromRef, toRef), fmt.Sprintf("jobs/%s/spec", jobName))
	output, err := gitCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("running '%s': %s: %s", strings.Join(gitCmd.Args, " "), err, string(output))
	}
	return string(output), nil
}
