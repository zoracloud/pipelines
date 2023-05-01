package config

import "github.com/zoracloud/pipelines/backend/src/apiserver/launcher"

type WoorkflowConfig struct {
	// Supported Launchers  List sequence determines the order in
    // which launchers are chosen for each component being run
	SupportedLaunchers launcher.Launcher
}