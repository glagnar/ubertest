package cmd

import (
	"github.com/glagnar/ubertest/app"
	"github.com/glagnar/ubertest/program"
	"go.uber.org/fx"
)

func Runcode() {
	fx.New(
		app.BaseModule,
		fx.Invoke(func(c *program.CoreApplication) {
			c.CallInterfaceSayHello()
		}),
	).Run()
}
