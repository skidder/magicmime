// Copyright 2013 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package magicmime

import (
	"encoding/base64"
	"testing"
)

var (
	m *Magic
)

func TestNew(t *testing.T) {
	var err error

	m, err = New()
	if err != nil {
		t.Fatal(err)
	}
}

// Tests a gif file.
func TestGifFile(t *testing.T) {
	testFile(t, "./testdata/sample.gif", "image/gif")
}

// Tests a jpeg file.
func TestJpegFile(t *testing.T) {
	testFile(t, "./testdata/sample.jpg", "image/jpeg")
}

// Tests a png file.
func TestPngFile(t *testing.T) {
	testFile(t, "./testdata/sample.png", "image/png")
}

// Tests a pdf file.
func TestPdfFile(t *testing.T) {
	testFile(t, "./testdata/sample.pdf", "application/pdf")
}

// Tests a plain text file.
func TestTextFile(t *testing.T) {
	testFile(t, "./testdata/sample.txt", "text/plain")
}

// Tests a gzipped tar file.
func TestGzippedTarFile(t *testing.T) {
	testFile(t, "./testdata/sample.tar.gz", "application/x-gzip")
}

// Tests a zip file.
func TestZipFile(t *testing.T) {
	testFile(t, "./testdata/sample.zip", "application/zip")
}

// Tests a gif buffer.
func TestGifBuffer(t *testing.T) {
	b64Gif := "R0lGODlhAQABAIAAAAUEBAAAACwAAAAAAQABAAACAkQBADs="
	expected := "image/gif"
	gif, err := base64.StdEncoding.DecodeString(b64Gif)
	if err != nil {
		panic(err)
	}
	mimetype, err := m.TypeByBuffer(gif)
	if err != nil {
		panic(err)
	}
	if mimetype != expected {
		t.Errorf("expected %s; got %s.", expected, mimetype)
	}
}

func testFile(tb testing.TB, path string, expected string) {
	mimetype, err := m.TypeByFile(path)
	if err != nil {
		panic(err)
	}
	if mimetype != expected {
		tb.Errorf("expected %s; got %s.", expected, mimetype)
	}
}

func TestMissingFile(t *testing.T) {
	_, err := m.TypeByFile("missingFile.txt")
	if err == nil {
		t.Error("no error for missing file")
	}
}

func BenchmarkZipFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testFile(b, "./testdata/sample.zip", "application/zip")
	}
}
