package s3config

//	Copyright 2016 Matt Ho
//
//	Licensed under the Apache License, Version 2.0 (the "License");
//	you may not use this file except in compliance with the License.
//	You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
//	Unless required by applicable law or agreed to in writing, software
//	distributed under the License is distributed on an "AS IS" BASIS,
//	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	See the License for the specific language governing permissions and
//	limitations under the License.

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestEncoder(t *testing.T) {
	bucket := "the-bucket"
	key := "the-key"
	m := &Mock{}

	err := NewEncoder(bucket, key, JSON, Client(m)).Encode(map[string]string{"hello": "world"})
	if err != nil {
		t.Error(err)
	}

	if m.Bucket != bucket {
		t.Errorf("expected %v; got %v", bucket, m.Bucket)
	}
	if m.Key != key {
		t.Errorf("expected %v; got %v", key, m.Key)
	}

	data, err := ioutil.ReadAll(m.PutBody)
	if err != nil {
		t.Error(err)
	}
	if v := strings.TrimSpace(string(data)); v != `{"hello":"world"}` {
		t.Errorf("expected %v; got %v", `{"hello":"world"}`, v)
	}
}
