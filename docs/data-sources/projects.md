---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bitbucket_projects Data Source - terraform-provider-bitbucket"
subcategory: ""
description: |-
  
---

# bitbucket_projects (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `owner` (String)

### Read-Only

- `id` (String) The ID of this resource.
- `projects` (Set of Object) (see [below for nested schema](#nestedatt--projects))

<a id="nestedatt--projects"></a>
### Nested Schema for `projects`

Read-Only:

- `description` (String)
- `has_publicly_visible_repos` (Boolean)
- `is_private` (Boolean)
- `key` (String)
- `name` (String)
- `uuid` (String)


