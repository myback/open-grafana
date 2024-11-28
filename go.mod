module github.com/myback/open-grafana

go 1.23

toolchain go1.23.2

// Override xorm's outdated go-mssqldb dependency, since we can't upgrade to current xorm (due to breaking changes).
// We need a more current go-mssqldb so we get rid of a version of apache/thrift with vulnerabilities.
// Also, use our fork with fixes for unimplemented methods (required for Go 1.16).
replace github.com/denisenkom/go-mssqldb => github.com/grafana/go-mssqldb v0.0.0-20210326084033-d0ce3c521036

require (
	cloud.google.com/go/storage v1.46.0
	github.com/BurntSushi/toml v1.4.0
	github.com/VividCortex/mysqlerr v1.0.0
	github.com/aws/aws-sdk-go v1.55.5
	github.com/beevik/etree v1.4.1
	github.com/benbjohnson/clock v1.3.5
	github.com/bradfitz/gomemcache v0.0.0-20230905024940-24af94b03874
	github.com/centrifugal/centrifuge v0.23.1
	github.com/crewjam/saml v0.4.14
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc
	github.com/denisenkom/go-mssqldb v0.12.3
	github.com/facebookgo/inject v0.0.0-20180706035515-f23751cae28b
	github.com/fatih/color v1.18.0
	github.com/gchaincl/sqlhooks v1.3.0
	github.com/getsentry/sentry-go v0.13.0
	github.com/go-jose/go-jose/v3 v3.0.3
	github.com/go-macaron/binding v1.2.0
	github.com/go-macaron/gzip v0.0.0-20200329073552-98214d7a897e
	github.com/go-sourcemap/sourcemap v2.1.4+incompatible
	github.com/go-sql-driver/mysql v1.8.1
	github.com/go-stack/stack v1.8.1
	github.com/gobwas/glob v0.2.3
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.4
	github.com/google/go-cmp v0.6.0
	github.com/gosimple/slug v1.14.0
	github.com/grafana/grafana-aws-sdk v0.31.4
	github.com/grafana/grafana-plugin-model v0.0.0-20200514130833-df1eb6bdf4c5
	github.com/grafana/grafana-plugin-sdk-go v0.259.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/hashicorp/go-hclog v1.6.3
	github.com/hashicorp/go-plugin v1.6.2
	github.com/hashicorp/go-version v1.7.0
	github.com/inconshreveable/log15 v2.16.0+incompatible
	github.com/influxdata/influxdb-client-go/v2 v2.14.0
	github.com/jmespath/go-jmespath v0.4.0
	github.com/json-iterator/go v1.1.12
	github.com/jung-kurt/gofpdf v1.16.2
	github.com/lib/pq v1.10.9
	github.com/linkedin/goavro/v2 v2.13.0
	github.com/magefile/mage v1.15.0
	github.com/mattn/go-isatty v0.0.20
	github.com/mattn/go-sqlite3 v1.14.24
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f
	github.com/opentracing/opentracing-go v1.2.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.20.5
	github.com/prometheus/client_model v0.6.1
	github.com/prometheus/common v0.60.1
	github.com/robfig/cron v1.2.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/russellhaering/goxmldsig v1.4.0
	github.com/smartystreets/goconvey v1.6.4
	github.com/stretchr/testify v1.9.0
	github.com/teris-io/shortid v0.0.0-20220617161101-71ec9f2aa569
	github.com/timberio/go-datemath v0.1.1-0.20200323150745-74ddef604fff
	github.com/ua-parser/uap-go v0.0.0-20241012191800-bbb40edc15aa
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/unknwon/com v1.0.1
	github.com/urfave/cli/v2 v2.27.5
	github.com/xorcare/pointer v1.1.0
	github.com/yudai/gojsondiff v1.0.0
	golang.org/x/crypto v0.29.0
	golang.org/x/net v0.31.0
	golang.org/x/oauth2 v0.24.0
	golang.org/x/sync v0.9.0
	golang.org/x/time v0.8.0
	gonum.org/v1/gonum v0.15.1
	google.golang.org/grpc v1.68.0
	gopkg.in/ini.v1 v1.67.0
	gopkg.in/ldap.v3 v3.1.3
	gopkg.in/macaron.v1 v1.5.0
	gopkg.in/mail.v2 v2.3.1
	gopkg.in/redis.v5 v5.2.9
	gopkg.in/yaml.v2 v2.4.0
	xorm.io/core v0.7.3
	xorm.io/xorm v0.8.2
)

require (
	cel.dev/expr v0.18.0 // indirect
	cloud.google.com/go v0.116.0 // indirect
	cloud.google.com/go/auth v0.10.2 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.5 // indirect
	cloud.google.com/go/compute/metadata v0.5.2 // indirect
	cloud.google.com/go/iam v1.2.2 // indirect
	cloud.google.com/go/monitoring v1.21.2 // indirect
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/FZambia/eagle v0.1.0 // indirect
	github.com/FZambia/sentinel v1.1.1 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp v1.25.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric v0.49.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping v0.49.0 // indirect
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/apache/arrow/go/v15 v15.0.2 // indirect
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/census-instrumentation/opencensus-proto v0.4.1 // indirect
	github.com/centrifugal/protocol v0.8.8 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cheekybits/genny v1.0.0 // indirect
	github.com/chromedp/cdproto v0.0.0-20240810084448-b931b754e476 // indirect
	github.com/cncf/xds/go v0.0.0-20240905190251-b4127c9b8d78 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.5 // indirect
	github.com/elazarl/goproxy v0.0.0-20240726154733-8b0c20506380 // indirect
	github.com/envoyproxy/go-control-plane v0.13.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.1.0 // indirect
	github.com/facebookgo/ensure v0.0.0-20160127193407-b4ab57deab51 // indirect
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/facebookgo/structtag v0.0.0-20150214074306-217e25fb9691 // indirect
	github.com/facebookgo/subset v0.0.0-20150612182917-8dac2c3c4870 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/getkin/kin-openapi v0.127.0 // indirect
	github.com/go-asn1-ber/asn1-ber v1.5.7 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-macaron/inject v0.0.0-20200308113650-138e5925c53b // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/google/flatbuffers v24.3.25+incompatible // indirect
	github.com/google/s2a-go v0.1.8 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.4 // indirect
	github.com/googleapis/gax-go/v2 v2.14.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20191106031601-ce3c9ade29de // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/gosimple/unidecode v1.0.1 // indirect
	github.com/grafana/dataplane/sdata v0.0.9 // indirect
	github.com/grafana/otel-profiling-go v0.5.1 // indirect
	github.com/grafana/pyroscope-go/godeltaprof v0.1.8 // indirect
	github.com/grafana/sqlds/v4 v4.1.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus v1.0.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.1.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.22.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/yamux v0.1.2 // indirect
	github.com/igm/sockjs-go/v3 v3.0.3 // indirect
	github.com/influxdata/line-protocol v0.0.0-20210922203350-b1ad95c89adf // indirect
	github.com/invopop/yaml v0.3.1 // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/jszwedko/go-datemath v0.1.1-0.20230526204004-640a500621d6 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/klauspost/cpuid/v2 v2.2.8 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattermost/xml-roundtrip-validator v0.1.0 // indirect
	github.com/mattetti/filebuffer v1.0.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mithrandie/csvq v1.18.1 // indirect
	github.com/mithrandie/csvq-driver v1.7.0 // indirect
	github.com/mithrandie/go-file/v2 v2.1.0 // indirect
	github.com/mithrandie/go-text v1.6.0 // indirect
	github.com/mithrandie/ternary v1.1.1 // indirect
	github.com/mna/redisc v1.4.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nxadm/tail v1.4.11 // indirect
	github.com/oapi-codegen/runtime v1.1.1 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/onsi/ginkgo v1.16.4 // indirect
	github.com/onsi/gomega v1.34.1 // indirect
	github.com/perimeterx/marshmallow v1.1.5 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	github.com/segmentio/encoding v0.4.0 // indirect
	github.com/sergi/go-diff v1.3.1 // indirect
	github.com/smartystreets/assertions v1.0.1 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/unknwon/bra v0.0.0-20200517080246-1e3013ecaff8 // indirect
	github.com/unknwon/log v0.0.0-20200308114134-929b1006e34a // indirect
	github.com/urfave/cli v1.22.16 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/xrash/smetrics v0.0.0-20240521201337-686a1a2994c1 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/detectors/gcp v1.32.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.57.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.56.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.57.0 // indirect
	go.opentelemetry.io/contrib/propagators/jaeger v1.31.0 // indirect
	go.opentelemetry.io/contrib/samplers/jaegerremote v0.25.0 // indirect
	go.opentelemetry.io/otel v1.32.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.31.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.31.0 // indirect
	go.opentelemetry.io/otel/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/trace v1.32.0 // indirect
	go.opentelemetry.io/proto/otlp v1.3.1 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20241108190413-2d47ceb2692f // indirect
	golang.org/x/mod v0.22.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/term v0.26.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	golang.org/x/tools v0.27.0 // indirect
	golang.org/x/xerrors v0.0.0-20240903120638-7835f813f4da // indirect
	google.golang.org/api v0.205.0 // indirect
	google.golang.org/genproto v0.0.0-20241021214115-324edc3d5d38 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241113202542-65e8d215514f // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241113202542-65e8d215514f // indirect
	google.golang.org/grpc/stats/opentelemetry v0.0.0-20241028142157-ada6787961b3 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/fsnotify/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	xorm.io/builder v0.3.13 // indirect
)

replace (
	github.com/crewjam/saml => github.com/grafana/saml v0.4.15-0.20240917091248-ae3bbdad8a56
	github.com/go-macaron/binding => ./pkg/macaron/binding
	github.com/grafana/grafana-aws-sdk v0.0.0 => ./src/grafana/grafana-aws-sdk
	github.com/grafana/grafana-plugin-sdk-go v0.0.0 => ./src/grafana/grafana-plugin-sdk-go
	github.com/ugorji/go v1.1.4 => github.com/ugorji/go v1.2.11
	gopkg.in/ldap.v3 v3.1.3 => github.com/go-ldap/ldap/v3 v3.1.3
	gopkg.in/macaron.v1 v1.4.0 => ./pkg/macaron
)
