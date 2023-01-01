package core

import (
	"errors"
	"fmt"
	"strings"
	"regexp"
)

func CheckIfId(possID string) (bool, error) {
	isDocId, err := regexp.MatchString("^[0-9]+.[0-9]$", possID)
	if err != nil {
		return false, err
	}
	return isDocId, nil
}

func GenerateUrl(docType, docID string) (string, error) {
	mosurl := ""

	switch docType {
	case "doc":
		mosurl = fmt.Sprintf("https://support.oracle.com/epmos/faces/DocumentDisplay?id=%s", docID)
	case "bug":
		mosurl = fmt.Sprintf("https://support.oracle.com/epmos/faces/BugMatrix?id=%s", docID)
	case "patch":
		mosurl = fmt.Sprintf("https://support.oracle.com/epmos/faces/PatchResultsNDetails?patchId=%s", docID)
	case "idea":
		mosurl = fmt.Sprintf("https://community.oracle.com/mosc/discussion/%s", docID)
	case "sr":
		mosurl = fmt.Sprintf("https://support.oracle.com/epmos/faces/SrDetail?srNumber=%s", docID)
	default:
		mosurl = fmt.Sprintf("https://support.oracle.com/epmos/faces/DocumentDisplay?id=%s", docID)
	}
	if mosurl == "" {
		return "", errors.New("Couldn't create URL.")
	}
	return mosurl, nil
}

func GeneratePublicUrl(id string) (string, error) {
	id = strings.Replace(id, ".", "_", -1)
	mospuburl := fmt.Sprintf("https://support.oracle.com/knowledge/PeopleSoft%%20Enterprise/%s.html", id)
	if mospuburl == "" {
		return "", errors.New("Couldn't create URL.")
	}
	return mospuburl, nil
}