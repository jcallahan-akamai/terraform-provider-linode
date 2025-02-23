package instance_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/linode/terraform-provider-linode/linode/acceptance"
	"github.com/linode/terraform-provider-linode/linode/instance/tmpl"
)

func TestAccDataSourceInstances_basic(t *testing.T) {
	t.Parallel()

	resName := "data.linode_instances.foobar"
	instanceName := acctest.RandomWithPrefix("tf_test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acceptance.PreCheck(t) },
		ProtoV5ProviderFactories: acceptance.ProtoV5ProviderFactories,
		CheckDestroy:             acceptance.CheckInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: tmpl.DataBasic(t, instanceName, testRegion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "instances.#", "1"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.id"),
					resource.TestCheckResourceAttr(resName, "instances.0.type", "g6-nanode-1"),
					resource.TestCheckResourceAttr(resName, "instances.0.tags.#", "2"),
					resource.TestCheckResourceAttr(resName, "instances.0.image", acceptance.TestImageLatest),
					resource.TestCheckResourceAttr(resName, "instances.0.region", testRegion),
					resource.TestCheckResourceAttr(resName, "instances.0.group", "tf_test"),
					resource.TestCheckResourceAttr(resName, "instances.0.swap_size", "256"),
					resource.TestCheckResourceAttr(resName, "instances.0.ipv4.#", "2"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.ipv6"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.host_uuid"),
					resource.TestCheckResourceAttr(resName, "instances.0.disk.#", "2"),
					resource.TestCheckResourceAttr(resName, "instances.0.config.#", "1"),
				),
			},
		},
	})
}

func TestAccDataSourceInstances_multipleInstances(t *testing.T) {
	resName := "data.linode_instances.foobar"
	resNameDesc := "data.linode_instances.desc"
	resNameAsc := "data.linode_instances.asc"

	instanceName := acctest.RandomWithPrefix("tf_test")
	tagName := acctest.RandomWithPrefix("tf_test")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acceptance.PreCheck(t) },
		ProtoV5ProviderFactories: acceptance.ProtoV5ProviderFactories,
		CheckDestroy:             acceptance.CheckInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: tmpl.DataMultiple(t, instanceName, tagName, testRegion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "instances.#", "3"),
				),
			},
			{
				Config: tmpl.DataMultipleOrder(t, instanceName, tagName, testRegion),
				Check: resource.ComposeTestCheckFunc(
					// Ensure order is correctly appended to filter
					resource.TestCheckResourceAttr(resNameDesc, "instances.#", "3"),
					resource.TestCheckResourceAttr(resNameAsc, "instances.#", "3"),
				),
			},
			{
				Config: tmpl.DataMultipleRegex(t, instanceName, tagName, testRegion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "instances.#", "3"),
				),
			},
			{
				Config: tmpl.DataClientFilter(t, instanceName, tagName, testRegion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "instances.#", "1"),
					resource.TestCheckResourceAttr(resName, "instances.0.status", "running"),
				),
			},
		},
	})
}
