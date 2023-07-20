package newrelicconfig

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"os"
)

func Initialize() (*newrelic.Application, error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(os.Getenv("NEW_RELIC_NAME")),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE")),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		return nil, err
	}
	return app, nil
}
