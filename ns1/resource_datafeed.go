package ns1

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"strconv"

	ns1 "gopkg.in/ns1/ns1-go.v2/rest"
	"gopkg.in/ns1/ns1-go.v2/rest/model/data"
)

func dataFeedResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"source_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"config": {
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
		Create:   DataFeedCreate,
		Read:     DataFeedRead,
		Update:   DataFeedUpdate,
		Delete:   DataFeedDelete,
		Importer: &schema.ResourceImporter{State: dataFeedStateFunc},
	}
}

func dataFeedToResourceData(d *schema.ResourceData, f *data.Feed) {
	configAdapterOut(f)
	d.SetId(f.ID)
	d.Set("name", f.Name)
	d.Set("config", f.Config)
}

func resourceDataToDataFeed(d *schema.ResourceData) (feed *data.Feed, e error) {
	config := d.Get("config").(map[string]interface{})
	if config != nil {
		test_id := config["test_id"]
		if test_id != nil {
			test_id_int, err := strconv.Atoi(test_id.(string))
			if err != nil {
				return &data.Feed{}, fmt.Errorf("could not convert %v as int %w", test_id, err)
			}
			config["test_id"] = test_id_int
		}
	}

	return &data.Feed{
		Name:     d.Get("name").(string),
		SourceID: d.Get("source_id").(string),
		Config:   config,
	}, nil
}

// DataFeedCreate creates an ns1 datafeed
func DataFeedCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ns1.Client)
	f, err := resourceDataToDataFeed(d)
	if err != nil {
		return err
	}
	if resp, err := client.DataFeeds.Create(d.Get("source_id").(string), f); err != nil {
		return ConvertToNs1Error(resp, err)
	}
	dataFeedToResourceData(d, f)
	return nil
}

// DataFeedRead reads the datafeed for the given ID from ns1
func DataFeedRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ns1.Client)
	f, resp, err := client.DataFeeds.Get(d.Get("source_id").(string), d.Id())
	if err != nil {
		// No custom error type is currently defined in the SDK for a data feed.
		if strings.Contains(err.Error(), "not found") {
			log.Printf("[DEBUG] NS1 data source (%s) not found", d.Id())
			d.SetId("")
			return nil
		}

		return ConvertToNs1Error(resp, err)
	}
	dataFeedToResourceData(d, f)
	return nil
}

// DataFeedDelete delets the given datafeed from ns1
func DataFeedDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ns1.Client)
	resp, err := client.DataFeeds.Delete(d.Get("source_id").(string), d.Id())
	d.SetId("")
	return ConvertToNs1Error(resp, err)
}

// DataFeedUpdate updates the given datafeed with modified parameters
func DataFeedUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ns1.Client)
	f, err := resourceDataToDataFeed(d)
	if err != nil {
		return err
	}
	f.ID = d.Id()
	if resp, err := client.DataFeeds.Update(d.Get("source_id").(string), f); err != nil {
		return ConvertToNs1Error(resp, err)
	}
	dataFeedToResourceData(d, f)
	return nil
}

// configAdapterOut back the original configuration types
func configAdapterOut(f *data.Feed) {
	config := f.Config
	if config != nil {
		test_id := config["test_id"]
		if test_id != nil {
			test_id_str := strconv.Itoa(int(test_id.(float64)))
			config["test_id"] = test_id_str
			f.Config = config
		}
	}
}

func dataFeedStateFunc(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid datafeed specifier.  Expecting 1 slashe (\"datasource_id/datafeed_id\"), got %d", len(parts)-1)
	}

	d.SetId(parts[1])
	d.Set("source_id", parts[0])

	return []*schema.ResourceData{d}, nil
}
