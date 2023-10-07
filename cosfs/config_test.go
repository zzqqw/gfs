package cosfs

import (
	"github.com/zzqqw/gfs"
	"net/url"
	"reflect"
	"testing"
)

func TestConfig_BucketUrl(t *testing.T) {
	type fields struct {
		CDN       string
		BucketURL string
		SecretID  string
		SecretKey string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *url.URL
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				CDN:       tt.fields.CDN,
				BucketURL: tt.fields.BucketURL,
				SecretID:  tt.fields.SecretID,
				SecretKey: tt.fields.SecretKey,
			}
			got, err := c.BucketUrl()
			if (err != nil) != tt.wantErr {
				t.Errorf("BucketUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BucketUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_GetBucket(t *testing.T) {
	type fields struct {
		CDN       string
		BucketURL string
		SecretID  string
		SecretKey string
	}
	type args struct {
		bucket string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				CDN:       tt.fields.CDN,
				BucketURL: tt.fields.BucketURL,
				SecretID:  tt.fields.SecretID,
				SecretKey: tt.fields.SecretKey,
			}
			if got := c.GetBucket(tt.args.bucket); got != tt.want {
				t.Errorf("GetBucket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_New(t *testing.T) {
	type fields struct {
		CDN       string
		BucketURL string
		SecretID  string
		SecretKey string
	}
	tests := []struct {
		name   string
		fields fields
		want   gfs.IAdapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				CDN:       tt.fields.CDN,
				BucketURL: tt.fields.BucketURL,
				SecretID:  tt.fields.SecretID,
				SecretKey: tt.fields.SecretKey,
			}
			if got := c.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_URL(t *testing.T) {
	type fields struct {
		CDN       string
		BucketURL string
		SecretID  string
		SecretKey string
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *url.URL
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				CDN:       tt.fields.CDN,
				BucketURL: tt.fields.BucketURL,
				SecretID:  tt.fields.SecretID,
				SecretKey: tt.fields.SecretKey,
			}
			got, err := c.URL(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("URL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URL() got = %v, want %v", got, tt.want)
			}
		})
	}
}