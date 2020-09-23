package index

import (
	"context"
	"encoding/json"

	"strings"
	"sync"

	"github.com/Rahul12344/skelego"
	es8 "github.com/elastic/go-elasticsearch/v8"
)

//Index Interface for dealing with an indexing system. An example of an indexing system is Elasticsearch.
type Index interface {
	skelego.Service
	CreateIndex(string, ...Document)
	SearchIndex(context.Context, string, *strings.Reader, skelego.Logging) []Document
	ElasticSearch() *es8.Client
	Query(string, skelego.Logging) *strings.Reader
}

//Document Index document
type Document interface {
}

//indexer ES compatible search/storage index
type indexer struct {
	clientInit sync.Once
	client     *es8.Client
	conn       chan error
}

//NewIndex new index
func NewIndex() Index {
	return &indexer{}
}

//Connect connects to an Elasticsearch client
func (ses *indexer) Configurifier(conf skelego.Config) {
	conf.DefaultSetting("index.port", 9201)
}

//Connect connects to an Elasticsearch client
func (ses *indexer) Connect(ctx context.Context, config skelego.Config, logger skelego.Logging) {
	ses.clientInit.Do(func() {
		c, err := es8.NewDefaultClient()
		if err != nil {
			panic(err)
		}
		ses.client = c
		logger.LogEvent("Connecting to ES", c)
	})

}

//Start Logs the start of the Index service
func (ses *indexer) Start(ctx context.Context, logger skelego.Logging) {
	logger.LogEvent("Starting elasticsearch index service...")
	return
}

//Stop Stops the execution of the Index service
func (ses *indexer) Stop(ctx context.Context, logger skelego.Logging) {
	logger.LogEvent("Stopping elasticsearch index service...")
	return
}

//CreateIndex Creates an Elasticsearch index using a provided JSON file or a list of entries that follow the JSON format
func (ses *indexer) CreateIndex(JSONfile string, entries ...Document) {
}

//ElasticSearch Returns client
func (ses *indexer) ElasticSearch() *es8.Client {
	return ses.client
}

//Query Creates query from map
func (ses *indexer) Query(conditional string, logger skelego.Logging) *strings.Reader {
	query := `{"query": {`
	query = query + conditional
	query = query + `}`

	// Check for JSON errors
	isValid := json.Valid([]byte(query)) // returns bool

	// Default query is "{}" if JSON is invalid
	if isValid == false {
		logger.LogError("constructQuery() ERROR: query string not valid:", query)
		logger.LogEvent("Using default match_all query")
		query = "{}"
	} else {
		logger.LogEvent("constructQuery() valid JSON:", isValid)
	}

	// Build a new string from JSON query
	var b strings.Builder
	b.WriteString(query)

	// Instantiate a *strings.Reader object from string
	read := strings.NewReader(b.String())

	// Return a *strings.Reader object
	return read
}

//SearchIndex Searches an index on some query
func (ses *indexer) SearchIndex(ctx context.Context, index string, query *strings.Reader, logger skelego.Logging) []Document {
	res, err := ses.client.Search(
		ses.client.Search.WithContext(ctx),
		ses.client.Search.WithIndex(index),
		ses.client.Search.WithBody(query),
		ses.client.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		logger.LogFatal("Error with ES response: %s", err.Error())
		return nil
	}
	if res == nil {
		return nil
	}
	return nil
}
