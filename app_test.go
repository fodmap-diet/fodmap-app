// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	sdk "github.com/fodmap-diet/go-sdk"
)

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/search/?item=mango", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(searchHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}

	expected := sdk.Properties{
		Category: "fruit",
		Fodmap:   "high",
	}

	var returned map[string]sdk.Properties

	readBuf, _ := ioutil.ReadAll(rr.Body)

	err = json.Unmarshal(readBuf, &returned)
	if err != nil {
		t.Errorf(
			"unexpected body: got (%v) want (%v)",
			string(readBuf),
			map[string]sdk.Properties{"mango": expected},
		)
	}

	item, found := returned["mango"]
	if !found {
		t.Errorf(
			"unexpected body: got (%v) want (%v)",
			string(readBuf),
			map[string]sdk.Properties{"mango": expected},
		)
	}

	if item.Category != expected.Category || item.Fodmap != expected.Fodmap ||
		item.Condition != expected.Condition || item.Note != expected.Note {
		t.Errorf(
			"unexpected body: got (%v) want (%v)",
			item,
			expected,
		)
	}
}

func TestIndexHandlerNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/404", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(searchHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}
}
