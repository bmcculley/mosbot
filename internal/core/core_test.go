package core

import "testing"

func TestGenerateDocUrl1(t *testing.T) {
	got, err := GenerateUrl("doc", "1234567.1")

	if err != nil {
		t.Errorf("got %q want nil", err)
	}

	want := "https://support.oracle.com/epmos/faces/DocumentDisplay?id=1234567.1"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGenerateDocUrl2(t *testing.T) {
	got, err := GenerateUrl("Doc", "2648807.1")

	if err != nil {
		t.Errorf("got %q want nil", err)
	}

	want := "https://support.oracle.com/epmos/faces/DocumentDisplay?id=2648807.1"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGenerateBugUrl1(t *testing.T) {
	got, err := GenerateUrl("bug", "1234567.1")

	if err != nil {
		t.Errorf("got %q want nil", err)
	}

	want := "https://support.oracle.com/epmos/faces/BugMatrix?id=1234567.1"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGenerateBugUrl2(t *testing.T) {
	got, err := GenerateUrl("Bug", "2648807.1")

	if err != nil {
		t.Errorf("got %q want nil", err)
	}

	want := "https://support.oracle.com/epmos/faces/BugMatrix?id=2648807.1"

	if got == want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGeneratePatchUrl1(t *testing.T) {
	got, err := GenerateUrl("patch", "34801725")

	if err != nil {
		t.Errorf("got %q want nil", err)
	}

	want := "https://support.oracle.com/epmos/faces/PatchResultsNDetails?patchId=34801725"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGeneratePatchUrl2(t *testing.T) {
	got, err := GenerateUrl("Patch", "34801725")

	if err != nil {
		t.Errorf("got %q want nil", err)
	}

	want := "https://support.oracle.com/epmos/faces/PatchResultsNDetails?patchId=34801725"

	if got == want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGenerateIdeaUrl1(t *testing.T) {
	got, err := GenerateUrl("idea", "34801725")

	if err != nil {
		t.Errorf("got %q want nil", err)
	}

	want := "https://community.oracle.com/mosc/discussion/34801725"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGenerateIdeaUrl2(t *testing.T) {
	got, err := GenerateUrl("Idea", "34801725")

	if err != nil {
		t.Errorf("got %q want nil", err)
	}

	want := "https://community.oracle.com/mosc/discussion/34801725"

	if got == want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGenerateSrUrl1(t *testing.T) {
	got, err := GenerateUrl("sr", "3-4801725")

	if err != nil {
		t.Errorf("got %q want nil", err)
	}

	want := "https://support.oracle.com/epmos/faces/SrDetail?srNumber=3-4801725"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGenerateSrUrl2(t *testing.T) {
	got, err := GenerateUrl("SR", "3-4801725")

	if err != nil {
		t.Errorf("got %q want nil", err)
	}

	want := "https://support.oracle.com/epmos/faces/SrDetail?srNumber=3-4801725"

	if got == want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGeneratePublicUrl(t *testing.T) {
	got, err := GeneratePublicUrl("2648807.1")

	if err != nil {
		t.Errorf("got %q want nil", err)
	}

	want := "https://support.oracle.com/knowledge/PeopleSoft%20Enterprise/2648807_1.html"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}