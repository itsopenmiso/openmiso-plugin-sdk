package plugin

import (
	"github.com/itsopenmiso/openmiso-plugin-sdk/component"
)

type mix_TaskLauncher_Authenticator struct {
	component.ConfigurableNotify
	component.TaskLauncher
	component.Documented
}
