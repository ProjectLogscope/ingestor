package config

import (
	"time"

	"github.com/alexflint/go-arg"
)

type config struct {
	ServiceName            string          `arg:"--service-name,env:SERVICE_NAME" help:"Service name" default:"ingestor" placeholder:"ingestor"`
	ServiceLogLevel        ServiceLogLevel `arg:"--service-log-level,env:SERVICE_LOG_LEVEL" help:"Specify log level" default:"info" placeholder:"info"`
	ServiceLogFilepath     string          `arg:"--service-log-filepath,env:SERVICE_LOG_FILEPATH" help:"Relative filepath for logger output" default:"/service/log/ingestor.log" placeholder:"/service/log/ingestor.log"`
	ServiceLogAddSource    bool            `arg:"--service-log-add-source,env:SERVICE_LOG_ADD_SOURCE" help:"Add source info to logger output" default:"true" placeholder:"true"`
	ServiceMessageValidate bool            `arg:"--service-message-validate,env:SERVICE_MESSAGE_VALIDATE" help:"Enable message validation" default:"true" placeholder:"true"`
	ServiceGracePeriod     time.Duration   `arg:"--service-grace-period,env:SERVICE_GRACE_PERIOD" help:"Service shutdown grace period" default:"1m" placeholder:"1m"`
	ServiceEnableMocking   bool            `arg:"--service-enable-mocking,env:SERVICE_ENABLE_MOCKING" help:"Run service in mock mode" default:"false" placeholder:"false"`

	NATSServerEndpoint string `arg:"--nats-server-endpoint,env:NATS_SERVER_ENDPOINT" help:"NATS server endpoint" placeholder:"nats://localhost:4222"`
	NATSSubjectIngest  string `arg:"--nats-subject-ingest,env:NATS_SUBJECT_INGEST" help:"NATS ingest subject" placeholder:"subject-ingest"`
	NATSGroupIngest    string `arg:"--nats-group-ingest,env:NATS_GROUP_INGEST" help:"NATS ingest group" placeholder:"group-ingest"`

	ElasticsearchEndpoints []string `arg:"--elasticsearch-endpoints,env:ELASTICSEARCH_ENDPOINTS" help:"Elasticsearch server endpoints" placeholder:"http://localhost:9200"`
	ElasticsearchIndex     string   `arg:"--elasticsearch-index,env:ELASTICSEARCH_INDEX" help:"Elasticsearch primary index" placeholder:"log_index"`

	WorkerTimeout  time.Duration `arg:"--worker-timeout,env:WORKER_TIMEOUT" help:"Worker timeout" default:"1m" placeholder:"1m"`
	WorkerMaxCount int           `arg:"--worker-max-count,env:WORKER_MAX_COUNT" help:"Worker max count" default:"10" placeholder:"10"`
}

func (config) Version() string {
	return "Log Ingest - Ingestor v1.0"
}

func LoadConfig() config {
	var cfg config
	p := arg.MustParse(&cfg)
	validateStrings(cfg, p)
	return cfg
}
