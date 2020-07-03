package main

import (
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	var testTags, modelTags tags

	testTags.append("tag1", "file1")
	modelTags.tf = make([]tagFiles, 0)
	modelTags.tf = append(modelTags.tf, tagFiles{"tag1", []string{"file1"}})
	if !reflect.DeepEqual(testTags, modelTags) {
		t.Error("This should be equal:", testTags, modelTags)
	}

	testTags.append("tag1", "file1")
	testTags.append("tag1", "file1")
	testTags.append("tag1", "file1")
	if !reflect.DeepEqual(testTags, modelTags) {
		t.Error("This should be equal:\n", testTags, "\n", modelTags)
	}

	testTags.append("tag1", "File1")
	modelTags.tf[0].file = append(modelTags.tf[0].file, "File1")
	if !reflect.DeepEqual(testTags, modelTags) {
		t.Error("This should be equal:\n", testTags, "\n", modelTags)
	}

	testTags.append("Tag1", "file1")
	if !reflect.DeepEqual(testTags, modelTags) {
		t.Error("This should be equal:\n", testTags, "\n", modelTags)
	}

	testTags.append("Tag2", "file1")
	modelTags.tf = append(modelTags.tf, tagFiles{"tag2", []string{"file1"}})
	if !reflect.DeepEqual(testTags, modelTags) {
		t.Error("This should be equal:\n", testTags, "\n", modelTags)
	}
}

func TestGetTags(t *testing.T) {
	getTags("asd +a +b+c_dkis_[ddd eee_fff]")
}
