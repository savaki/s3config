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
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hashicorp/hcl"
	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
)

type S3 interface {
	GetObject(*s3.GetObjectInput) (*s3.GetObjectOutput, error)
	PutObject(*s3.PutObjectInput) (*s3.PutObjectOutput, error)
}

func lookupAvailabilityZone(ctx context.Context) (io.ReadCloser, error) {
	resp, err := ctxhttp.Get(ctx, http.DefaultClient, "http://169.254.169.254/latest/meta-data/placement/availability-zone")
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func region(lookupAZ func(context.Context) (io.ReadCloser, error)) string {
	region := os.Getenv("AWS_REGION")

	if region == "" {
		region = os.Getenv("AWS_DEFAULT_REGION")
	}

	if region == "" {
		ctx, _ := context.WithTimeout(context.Background(), time.Second)
		body, err := lookupAZ(ctx)
		if err == nil {
			defer body.Close()

			data, err := ioutil.ReadAll(body)
			if err == nil {
				region = strings.TrimSpace(string(data))
				if len(region) > 0 {
					region = region[0 : len(region)-1]
				}
			}
		}
	}

	if region == "" {
		region = "us-east-1"
	}

	return region
}

func defaultClient() S3 {
	cfg := &aws.Config{Region: aws.String(region(lookupAvailabilityZone))}

	return s3.New(session.New(cfg))
}

type Options struct {
	client     S3
	bucket     string
	key        string
	encodeFunc func(w io.Writer, v interface{}) error
	decodeFunc func(r io.Reader, v interface{}) error
}

func JSON(opts *Options) {
	opts.decodeFunc = func(r io.Reader, v interface{}) error {
		return json.NewDecoder(r).Decode(v)
	}
	opts.encodeFunc = func(r io.Writer, v interface{}) error {
		return json.NewEncoder(r).Encode(v)
	}
}

func HCL(opts *Options) {
	opts.decodeFunc = func(r io.Reader, v interface{}) error {
		data, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}

		return hcl.Decode(v, string(data))
	}
	opts.encodeFunc = func(r io.Writer, v interface{}) error {
		return errors.New("Cannot encode HCL")
	}
}

func XML(d *Options) {
	d.decodeFunc = func(r io.Reader, v interface{}) error {
		return xml.NewDecoder(r).Decode(v)
	}
	d.encodeFunc = func(r io.Writer, v interface{}) error {
		return xml.NewEncoder(r).Encode(v)
	}
}

func Client(client S3) func(*Options) {
	return func(d *Options) {
		d.client = client
	}
}
