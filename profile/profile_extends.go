package profile

import (
	"github.com/ycxwi/micro/v3/service/auth"
	"github.com/ycxwi/micro/v3/service/client"
	"github.com/ycxwi/micro/v3/service/config"
	"github.com/ycxwi/micro/v3/service/events"
	"github.com/ycxwi/micro/v3/service/logger"
	"github.com/ycxwi/micro/v3/service/registry"
	"github.com/ycxwi/micro/v3/service/router"
	"github.com/ycxwi/micro/v3/service/runtime"
	"github.com/ycxwi/micro/v3/service/server"
	"github.com/ycxwi/micro/v3/service/store"

	mAuthNoop "github.com/ycxwi/micro/v3/service/auth/noop"
	mBrokerMemory "github.com/ycxwi/micro/v3/service/broker/memory"
	mConfigEnv "github.com/ycxwi/micro/v3/service/config/env"
	mEventStream "github.com/ycxwi/micro/v3/service/events/stream/memory"
	mRegistryNoop "github.com/ycxwi/micro/v3/service/registry/noop"
	mRouterStatic "github.com/ycxwi/micro/v3/service/router/static"
	mRuntimeLocal "github.com/ycxwi/micro/v3/service/runtime/local"
	mStoreMemory "github.com/ycxwi/micro/v3/service/store/memory"
	mStoreNoop "github.com/ycxwi/micro/v3/service/store/noop"

	"github.com/urfave/cli/v2"
)

// Dev profile to run service in simple config
var Dev = &Profile{
	Name: "dev",
	Setup: func(ctx *cli.Context) error {
		auth.DefaultAuth = mAuthNoop.NewAuth()
		runtime.DefaultRuntime = mRuntimeLocal.NewRuntime()
		//store.DefaultStore = fstore.NewStore()
		store.DefaultStore = mStoreNoop.NewStore()
		config.DefaultConfig, _ = mConfigEnv.NewConfig()
		var err error
		events.DefaultStream, err = mEventStream.NewStream()
		if err != nil {
			logger.Fatalf("Error configuring stream for simple profile: %v", err)
		}

		SetupBroker(mBrokerMemory.NewBroker())
		//turn off Registry
		setupRegistry(mRegistryNoop.NewRegistry())
		// store.DefaultBlobStore, err = fstore.NewBlobStore()
		// if err != nil {
		// 	logger.Fatalf("Error configuring file blob store: %v", err)
		// }

		return nil
	},
}

// Simple profile to run service in simple config
var Simple = &Profile{
	Name: "simple",
	Setup: func(ctx *cli.Context) error {
		auth.DefaultAuth = mAuthNoop.NewAuth()
		runtime.DefaultRuntime = mRuntimeLocal.NewRuntime()
		//store.DefaultStore = fstore.NewStore()
		store.DefaultStore = mStoreMemory.NewStore()
		config.DefaultConfig, _ = mConfigEnv.NewConfig()
		var err error
		events.DefaultStream, err = mEventStream.NewStream()
		if err != nil {
			logger.Fatalf("Error configuring stream for simple profile: %v", err)
		}

		SetupBroker(mBrokerMemory.NewBroker())
		//turn off Registry
		setupRegistry(mRegistryNoop.NewRegistry())
		// store.DefaultBlobStore, err = fstore.NewBlobStore()
		// if err != nil {
		// 	logger.Fatalf("Error configuring file blob store: %v", err)
		// }

		return nil
	},
}

// setupRegistry configures the registry
func setupRegistry(reg registry.Registry) {
	registry.DefaultRegistry = reg
	router.DefaultRouter = mRouterStatic.NewRouter(router.Registry(reg))
	_ = client.DefaultClient.Init(client.Registry(reg))
	_ = server.DefaultServer.Init(server.Registry(reg))
}
