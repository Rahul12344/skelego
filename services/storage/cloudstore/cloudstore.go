package cloudstore

import (
	"github.com/Rahul12344/skelego"
)

//CloudService Interface dealing with cloud service bucketing systems like AWS and GCloud. Different than a traditional DB, so does not
//implement the Store service.
type CloudService interface {
	GetBucket(string) Bucket
	AddBucket(...Bucket)
}

//Bucket Buckets are the storage schema of cloud storage.
type Bucket interface {
	ModifyContents()
	Create()
}

type cloud struct {
}

func (c *cloud) Configurifier(conf skelego.Config) {
	conf.DefaultSetting("", "")
	conf.DefaultSetting("", "")
	conf.DefaultSetting("", "")
}

func (c *cloud) GetBucket(bucketName string) Bucket {
	return nil
}

func (c *cloud) AddBucket(buckets ...Bucket) {

}
