package program

import "github.com/glagnar/ubertest/feature"

type CoreApplication struct {
	implementation1 feature.FeaturePlace
	implementation2 feature.FeaturePlace
}

func NewCoreApplication(implementation1 feature.FeaturePlace, implentation2 feature.FeaturePlace) *CoreApplication {
	return &CoreApplication{
		implementation1: implementation1,
		implementation2: implentation2,
	}
}

func (c *CoreApplication) CallInterfaceSayHello() {
	c.implementation1.SayHello()
	c.implementation2.SayHello()
}
