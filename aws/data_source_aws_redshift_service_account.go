package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// See http://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#db-auditing-bucket-permissions
// See https://docs.aws.amazon.com/govcloud-us/latest/UserGuide/govcloud-redshift.html
// See https://docs.amazonaws.cn/en_us/redshift/latest/mgmt/db-auditing.html#db-auditing-bucket-permissions
var redshiftServiceAccountPerRegionMap = map[string]string{
	"us-east-1":      "193672423079",
	"us-east-2":      "391106570357",
	"us-west-1":      "262260360010",
	"us-west-2":      "902366379725",
	"af-south-1":     "365689465814",
	"ap-east-1":      "313564881002",
	"ap-south-1":     "865932855811",
	"ap-northeast-3": "090321488786",
	"ap-northeast-2": "760740231472",
	"ap-southeast-1": "361669875840",
	"ap-southeast-2": "762762565011",
	"ap-northeast-1": "404641285394",
	"ca-central-1":   "907379612154",
	"cn-north-1":     "111890595117",
	"cn-northwest-1": "660998842044",
	"eu-central-1":   "053454850223",
	"eu-west-1":      "210876761215",
	"eu-west-2":      "307160386991",
	"eu-west-3":      "915173422425",
	"eu-north-1":     "729911121831",
	"eu-south-1":     "945612479654",
	"me-south-1":     "013126148197",
	"sa-east-1":      "075028567923",
	"us-gov-east-1":  "665727464434",
	"us-gov-west-1":  "665727464434",
}

func dataSourceAwsRedshiftServiceAccount() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsRedshiftServiceAccountRead,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAwsRedshiftServiceAccountRead(d *schema.ResourceData, meta interface{}) error {
	region := meta.(*AWSClient).region
	if v, ok := d.GetOk("region"); ok {
		region = v.(string)
	}

	if accid, ok := redshiftServiceAccountPerRegionMap[region]; ok {
		d.SetId(accid)
		arn := arn.ARN{
			Partition: meta.(*AWSClient).partition,
			Service:   "iam",
			AccountID: accid,
			Resource:  "user/logs",
		}.String()
		d.Set("arn", arn)

		return nil
	}

	return fmt.Errorf("Unknown region (%q)", region)
}
