package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceYamlConfig() *schema.Resource {

	return &schema.Resource{
		Description:   configAsCodeDescription("Resource for creating a raw YAML configuration in Harness. Note: This works for all objects EXCEPT application objects."),
		CreateContext: resourceYamlConfigCreateOrUpdate,
		ReadContext:   resourceYamlConfigRead,
		UpdateContext: resourceYamlConfigCreateOrUpdate,
		DeleteContext: resourceYamlConfigDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The unique id of the resource.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The name of the resource.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"app_id": {
				Description: "The id of the application. This is required for all resources except global ones.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"path": {
				Description: "The path of the resource.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"content": {
				Description: "The raw YAML configuration.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				// Id can either be <YAML_PATH> or <YAML_PATH>:<APP_ID>
				parts := strings.Split(d.Id(), ":")

				d.Set("path", parts[0])

				if len(parts) > 1 {
					d.Set("app_id", parts[1])
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func resourceYamlConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	app_id := d.Get("app_id").(string)
	path := cac.YamlPath(d.Get("path").(string))

	entity, err := c.ConfigAsCode().FindYamlByPath(app_id, path)
	if err != nil {
		return diag.FromErr(err)
	} else if entity == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	return readYamlConfig(d, entity)
}

func readYamlConfig(d *schema.ResourceData, entity *cac.YamlEntity) diag.Diagnostics {
	d.SetId(entity.Id)
	d.Set("name", entity.Name)
	d.Set("content", entity.Content)
	d.Set("path", entity.Path)
	d.Set("app_id", entity.ApplicationId)
	return nil
}

func resourceYamlConfigCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	path := cac.YamlPath(d.Get("path").(string))
	app_id := d.Get("app_id").(string)
	content := d.Get("content").(string)

	_, err := c.ConfigAsCode().UpsertRawYaml(path, []byte(content))
	if err != nil {
		return diag.FromErr(err)
	}

	yamlEntity, err := c.ConfigAsCode().FindYamlByPath(app_id, path)
	if err != nil {
		return diag.FromErr(err)
	}

	if yamlEntity == nil {
		return diag.FromErr(fmt.Errorf("failed to find yaml entity at %s", path))
	}

	return readYamlConfig(d, yamlEntity)
}

func resourceYamlConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	err := c.CloudProviders().DeleteCloudProvider(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
