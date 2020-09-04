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

func Test_appendSubdirTags(t *testing.T) {
	type args struct {
		source  tags
		newTags tags
		subdir  string
	}
	tests := []struct {
		name string
		args args
		want tags
	}{
		// TODO: Add test cases.
		{"test1",
			args{tags{
				[]tagFiles{{"stag1", []string{"sfile1", "sfile2"}}},
			},
				tags{
					[]tagFiles{{"ntag1", []string{"nfile1", "nfile2"}}},
				},
				"subdir"},
			tags{
				[]tagFiles{{"stag1", []string{"sfile1", "sfile2"}},
					{"ntag1", []string{"subdir/nfile1", "subdir/nfile2"}},
					{"subdir", []string{"subdir/nfile1", "subdir/nfile2"}},
				},
			},
		},
	}

	//atf := []tagFiles{{"",[]string{"",""},}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appendSubdirTags(tt.args.source, tt.args.newTags, tt.args.subdir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendSubdirTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTags(t *testing.T) {
	type args struct {
		source string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Empty tags", args{""}, []string{}},
		{"Plus tag", args{"+"}, []string{}},
		{"pluses", args{"++ ++++"}, []string{"+", "+++"}},
		{"Null tags", args{}, []string{}},
		{"Tags only", args{"+ a +a +a +b"}, []string{"a", "b"}},
		{"Generic", args{"asd +a +b+c_dkis_[ddd eee_fff]"}, []string{"a", "b+c"}},
		{"a +b.+c[+d]+e", args{"a +b.+c[+d]+e +a+b"}, []string{"b", "c", "d", "e", "a+b"}},
		{"a +b.+c[+d]+e", args{"asd +a +b+c_dkis_[ddd eee_fff]"}, []string{"a", "b+c"}},
		{"pluses", args{"++ ++++"}, []string{"+", "+++"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTags(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
