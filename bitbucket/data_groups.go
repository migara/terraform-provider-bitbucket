package bitbucket

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataReadGroups,

		Schema: map[string]*schema.Schema{
			"workspace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"groups": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"slug": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"auto_add": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"permission": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataReadGroups(d *schema.ResourceData, m interface{}) error {
	client := m.(Clients).httpClient

	workspace := d.Get("workspace").(string)

	groupsReq, _ := client.Get(fmt.Sprintf("1.0/groups/%s", workspace))

	if groupsReq.Body == nil {
		return fmt.Errorf("error reading Groups (%s): empty response", d.Id())
	}

	var grps []*UserGroup

	body, readerr := ioutil.ReadAll(groupsReq.Body)
	if readerr != nil {
		return readerr
	}

	log.Printf("[DEBUG] Groups Response JSON: %v", string(body))

	decodeerr := json.Unmarshal(body, &grps)
	if decodeerr != nil {
		return decodeerr
	}

	log.Printf("[DEBUG] Groups Response Decoded: %#v", grps)

	d.SetId(workspace)
	d.Set("groups", flattenUserGroups(grps))

	return nil
}

func flattenUserGroups(groups []*UserGroup) []interface{} {
	if len(groups) == 0 {
		return nil
	}

	var tfList []interface{}

	for _, btRaw := range groups {
		log.Printf("[DEBUG] User Group Response Decoded: %#v", btRaw)

		if btRaw == nil {
			continue
		}

		group := map[string]interface{}{
			"name":                      btRaw.Name,
			"permission":                btRaw.Permission,
			"slug":                      btRaw.Slug,
			"auto_add":                  btRaw.AutoAdd,
			"email_forwarding_disabled": btRaw.EmailForwardingDisabled,
		}

		tfList = append(tfList, group)
	}

	return tfList
}
