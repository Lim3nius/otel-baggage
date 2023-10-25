module bag-test

go 1.21.3

require go.opentelemetry.io/otel v1.19.0

require go.opentelemetry.io/otel/trace v1.19.0 // indirect

replace go.opentelemetry.io/otel => github.com/odenio/opentelemetry-go v0.0.0-20231024195813-f22a4f2c6187

replace go.opentelemetry.io/otel/trace => github.com/odenio/opentelemetry-go/trace v0.0.0-20231024195813-f22a4f2c6187
