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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Decoder struct {
	Options
}

func (d *Decoder) Decode(v interface{}) error {
	out, err := d.client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(d.bucket),
		Key:    aws.String(d.key),
	})
	if err != nil {
		return err
	}
	defer out.Body.Close()

	return d.decodeFunc(out.Body, v)
}

func NewDecoder(bucket, key string, configs ...func(*Options)) *Decoder {
	opts := Options{
		client: defaultClient(),
		bucket: bucket,
		key:    key,
	}
	JSON(&opts) // use JSON as the default decoder

	for _, config := range configs {
		config(&opts)
	}

	return &Decoder{Options: opts}
}
