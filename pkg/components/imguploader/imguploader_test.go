package imguploader

import (
	"testing"

	"github.com/myback/open-grafana/pkg/setting"

	. "github.com/smartystreets/goconvey/convey"
)

func TestImageUploaderFactory(t *testing.T) {
	Convey("Can create image uploader for ", t, func() {
		Convey("S3ImageUploader config", func() {
			cfg := setting.NewCfg()
			err := cfg.Load(&setting.CommandLineArgs{
				HomePath: "../../../",
			})
			So(err, ShouldBeNil)

			setting.ImageUploadProvider = "s3"

			Convey("with bucket url https://foo.bar.baz.s3-us-east-2.amazonaws.com", func() {
				s3sec, err := setting.Raw.GetSection("external_image_storage.s3")
				So(err, ShouldBeNil)
				_, err = s3sec.NewKey("bucket_url", "https://foo.bar.baz.s3-us-east-2.amazonaws.com")
				So(err, ShouldBeNil)
				_, err = s3sec.NewKey("access_key", "access_key")
				So(err, ShouldBeNil)
				_, err = s3sec.NewKey("secret_key", "secret_key")
				So(err, ShouldBeNil)

				uploader, err := NewImageUploader()
				So(err, ShouldBeNil)

				original, ok := uploader.(*S3Uploader)
				So(ok, ShouldBeTrue)
				So(original.region, ShouldEqual, "us-east-2")
				So(original.bucket, ShouldEqual, "foo.bar.baz")
				So(original.accessKey, ShouldEqual, "access_key")
				So(original.secretKey, ShouldEqual, "secret_key")
			})

			Convey("with bucket url https://s3.amazonaws.com/mybucket", func() {
				s3sec, err := setting.Raw.GetSection("external_image_storage.s3")
				So(err, ShouldBeNil)
				_, err = s3sec.NewKey("bucket_url", "https://s3.amazonaws.com/my.bucket.com")
				So(err, ShouldBeNil)
				_, err = s3sec.NewKey("access_key", "access_key")
				So(err, ShouldBeNil)
				_, err = s3sec.NewKey("secret_key", "secret_key")
				So(err, ShouldBeNil)

				uploader, err := NewImageUploader()
				So(err, ShouldBeNil)

				original, ok := uploader.(*S3Uploader)
				So(ok, ShouldBeTrue)
				So(original.region, ShouldEqual, "us-east-1")
				So(original.bucket, ShouldEqual, "my.bucket.com")
				So(original.accessKey, ShouldEqual, "access_key")
				So(original.secretKey, ShouldEqual, "secret_key")
			})

			Convey("with bucket url https://s3-us-west-2.amazonaws.com/mybucket", func() {
				s3sec, err := setting.Raw.GetSection("external_image_storage.s3")
				So(err, ShouldBeNil)
				_, err = s3sec.NewKey("bucket_url", "https://s3-us-west-2.amazonaws.com/my.bucket.com")
				So(err, ShouldBeNil)
				_, err = s3sec.NewKey("access_key", "access_key")
				So(err, ShouldBeNil)
				_, err = s3sec.NewKey("secret_key", "secret_key")
				So(err, ShouldBeNil)

				uploader, err := NewImageUploader()
				So(err, ShouldBeNil)

				original, ok := uploader.(*S3Uploader)
				So(ok, ShouldBeTrue)
				So(original.region, ShouldEqual, "us-west-2")
				So(original.bucket, ShouldEqual, "my.bucket.com")
				So(original.accessKey, ShouldEqual, "access_key")
				So(original.secretKey, ShouldEqual, "secret_key")
			})
		})

		Convey("Webdav uploader", func() {
			cfg := setting.NewCfg()
			err := cfg.Load(&setting.CommandLineArgs{
				HomePath: "../../../",
			})
			So(err, ShouldBeNil)

			setting.ImageUploadProvider = "webdav"

			webdavSec, err := cfg.Raw.GetSection("external_image_storage.webdav")
			So(err, ShouldBeNil)
			_, err = webdavSec.NewKey("url", "webdavUrl")
			So(err, ShouldBeNil)
			_, err = webdavSec.NewKey("username", "username")
			So(err, ShouldBeNil)
			_, err = webdavSec.NewKey("password", "password")
			So(err, ShouldBeNil)

			uploader, err := NewImageUploader()
			So(err, ShouldBeNil)
			original, ok := uploader.(*WebdavUploader)

			So(ok, ShouldBeTrue)
			So(original.url, ShouldEqual, "webdavUrl")
			So(original.username, ShouldEqual, "username")
			So(original.password, ShouldEqual, "password")
		})

		Convey("Local uploader", func() {
			cfg := setting.NewCfg()
			err := cfg.Load(&setting.CommandLineArgs{
				HomePath: "../../../",
			})
			So(err, ShouldBeNil)

			setting.ImageUploadProvider = "local"

			uploader, err := NewImageUploader()
			So(err, ShouldBeNil)

			original, ok := uploader.(*LocalUploader)
			So(ok, ShouldBeTrue)
			So(original, ShouldNotBeNil)
		})
	})
}
