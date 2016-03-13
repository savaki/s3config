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
	"strings"
	"testing"
)

func TestDecoder(t *testing.T) {
	bucket := "the-bucket"
	key := "the-key"
	body := `env="local"`
	m := &Mock{
		Body: strings.NewReader(body),
	}

	config := struct {
		Env string
	}{}

	err := NewDecoder(bucket, key, HCL, Client(m)).Decode(&config)
	if err != nil {
		t.Error(err)
	}

	if m.Bucket != bucket {
		t.Errorf("expected %v; got %v", bucket, m.Bucket)
	}
	if m.Key != key {
		t.Errorf("expected %v; got %v", key, m.Key)
	}

	if config.Env != "local" {
		t.Errorf("expected %v; got %v", "local", config.Env)
	}
}
