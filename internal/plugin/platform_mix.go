package plugin

import (
	"github.com/itsopenmiso/openmiso-plugin-sdk/component"
)

type mix_Platform_Authenticator struct {
	component.Authenticator
	component.ConfigurableNotify
	component.Documented
	component.Platform
	component.PlatformReleaser
	component.WorkspaceDestroyer
	component.LogPlatform
	component.Generation
	component.Status
}

type mix_Platform_Destroy struct {
	component.Authenticator
	component.ConfigurableNotify
	component.Documented
	component.Platform
	component.PlatformReleaser
	component.Execer
	component.LogPlatform
	component.Destroyer
	component.WorkspaceDestroyer
	component.Generation
	component.Status
}

type mix_Platform_Exec struct {
	component.Authenticator
	component.ConfigurableNotify
	component.Documented
	component.Platform
	component.PlatformReleaser
	component.LogPlatform
	component.Execer
	component.Generation
	component.Status
}
