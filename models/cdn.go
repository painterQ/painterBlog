package models

var CDNSingleCase *CDN

type CDN struct {
	UseCDN 	  bool
	Bucket    string
	Domain    string
	AccessKey string
	SecretKey string
}
