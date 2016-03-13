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
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Encoder struct {
	Options
}

func (d *Encoder) Encode(v interface{}) error {
	buf := &bytes.Buffer{}
	err := d.encodeFunc(buf, v)
	if err != nil {
		return err
	}

	_, err = d.client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(d.bucket),
		Key:    aws.String(d.key),
		Body:   bytes.NewReader(buf.Bytes()),
	})
	return err
}

func NewEncoder(bucket, key string, configs ...func(*Options)) *Encoder {
	opts := Options{
		client: defaultClient(),
		bucket: bucket,
		key:    key,
	}
	JSON(&opts) // use JSON as the default encoder

	for _, config := range configs {
		config(&opts)
	}

	return &Encoder{Options: opts}
}
