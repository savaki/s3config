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
	"io"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/service/s3"
)

type Mock struct {
	Bucket  string
	Key     string
	Body    io.Reader
	PutBody io.Reader
}

func (m *Mock) GetObject(in *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	m.Bucket = *in.Bucket
	m.Key = *in.Key

	return &s3.GetObjectOutput{
		Body: ioutil.NopCloser(m.Body),
	}, nil
}

func (m *Mock) PutObject(in *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	m.Bucket = *in.Bucket
	m.Key = *in.Key
	m.PutBody = in.Body

	return &s3.PutObjectOutput{}, nil
}
