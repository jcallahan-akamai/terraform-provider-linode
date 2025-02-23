package region_test

import (
	"context"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/linode/terraform-provider-linode/linode/acceptance"
	"github.com/linode/terraform-provider-linode/linode/region/tmpl"
)

var (
	testRegion string
	testLabel  string
)

func init() {
	region, err := acceptance.GetRandomRegionWithCaps([]string{"linodes"})
	if err != nil {
		log.Fatal(err)
	}

	testRegion = region

	client, err := acceptance.GetClientForSweepers()
	if err != nil {
		log.Fatal(err)
	}

	r, err := client.GetRegion(context.Background(), testRegion)
	if err != nil {
		log.Fatal(err)
	}

	testLabel = r.Label
}

func TestAccDataSourceRegion_basic(t *testing.T) {
	t.Parallel()

	regionID := testRegion
	label := testLabel
	resourceName := "data.linode_region.foobar"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acceptance.PreCheck(t) },
		ProtoV5ProviderFactories: acceptance.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: tmpl.DataBasic(t, regionID, label),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "country"),
					resource.TestCheckResourceAttr(resourceName, "id", regionID),
					resource.TestCheckResourceAttr(resourceName, "label", label),
					resource.TestCheckResourceAttrSet(resourceName, "status"),
					resource.TestCheckResourceAttrSet(resourceName, "resolvers.0.ipv4"),
					resource.TestCheckResourceAttrSet(resourceName, "resolvers.0.ipv6"),
					acceptance.CheckResourceAttrGreaterThan(resourceName, "capabilities.#", 0),
				),
			},
		},
	})
}
