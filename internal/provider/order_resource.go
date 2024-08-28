package provider

import (
	"context"

	"github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource = &orderResource{}
)

// NewOrderResource is a helper function to simplify the provider implementation.
func NewOrderResource() resource.Resource {
	return &orderResource{}
}

// orderResourceModel maps the resource schema data.
type orderResourceModel struct {
	Account  types.String `tfsdk:"account"`
	Password types.String `tfsdk:"password"`
}

type readResponseModel string

// orderResource is the resource implementation.
type orderResource struct {
	client *hashicups.Client
}

// Create a new resource.
func (r *orderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
}

// Read reads the resource state.
func (r *orderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Implement the read logic here.
	credentials := get_access_credential()
	// // Retrieve values from the credentials
	// account := credentials.Account
	// password := credentials.Password

	// // Create a new readResponseModel instance
	// var ReadResponse readResponseModel
	// ReadResponse = fmt.Sprintf("account: %s, password: %s", account, password)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &credentials)...)
}

// Update updates the resource.
func (r *orderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Implement the update logic here.
}

// Delete deletes the resource.
func (r *orderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Implement the delete logic here.
}

// Metadata returns the resource type name.
func (r *orderResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_order"
}

// Schema defines the schema for the resource.
func (r *orderResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account": schema.StringAttribute{
				Required: true,
			},
			"password": schema.StringAttribute{
				Required: true,
			},
		},
	}
}
