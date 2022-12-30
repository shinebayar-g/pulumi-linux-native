package main

import (
	"os"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

var Version string

func main() {
	p.RunProvider("linux", Version,
		infer.Provider(infer.Options{
			Resources: []infer.InferredResource{
				infer.Resource[File, FileArgs, FileState](),
			},
		}))
}

// Each resource has a controlling struct.
// Resource behavior is determined by implementing methods on the controlling struct.
// The `Create` method is mandatory, but other methods are optional.
// - Check: Remap inputs before they are typed.
// - Diff: Change how instances of a resource are compared.
// - Update: Mutate a resource in place.
// - Read: Get the state of a resource from the backing provider.
// - Delete: Custom logic when the resource is deleted.
// - Annotate: Describe fields and set defaults for a resource.
// - WireDependencies: Control how outputs and secrets flows through values.
type File struct{}

// Each resource has in input struct, defining what arguments it accepts.
type FileArgs struct {
	// Fields projected into Pulumi must be public and hava a `pulumi:"..."` tag.
	// The pulumi tag doesn't need to match the field name, but its generally a
	// good idea.
	Path string `pulumi:"path"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type FileState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	FileArgs
	// Here we define a required output called result.
	Result string `pulumi:"result"`
}

// All resources must implement Create at a minumum.
func (File) Create(ctx p.Context, name string, input FileArgs, preview bool) (string, FileState, error) {
	state := FileState{FileArgs: input}
	if preview {
		return name, state, nil
	}
	err := createFile(input.Path)
	state.Result = "not working"
	return name, state, err
}

func createFile(path string) error {
	_, err := os.Create(path)
	return err
}
