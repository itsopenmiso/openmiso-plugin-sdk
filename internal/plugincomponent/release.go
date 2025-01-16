package plugincomponent

import (
	"encoding/json"

	"github.com/hashicorp/opaqueany"
	"google.golang.org/protobuf/proto"

	"github.com/itsopenmiso/openmiso-plugin-sdk/component"
	pb "github.com/itsopenmiso/openmiso-plugin-sdk/proto/gen"
)

// Release implements component.Release.
type Release struct {
	Any         *opaqueany.Any
	AnyJson     string
	Release     *pb.Release
	TemplateVal map[string]interface{}
}

func (c *Release) Proto() proto.Message                 { return c.Any }
func (c *Release) URL() string                          { return c.Release.Url }
func (c *Release) TemplateData() map[string]interface{} { return c.TemplateVal }
func (c *Release) MarshalJSON() ([]byte, error)         { return []byte(c.AnyJson), nil }

var (
	_ component.Release  = (*Release)(nil)
	_ component.Template = (*Release)(nil)
	_ json.Marshaler     = (*Release)(nil)
)
