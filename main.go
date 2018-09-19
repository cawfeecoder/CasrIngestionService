package main

import (
"time"
"github.com/olivere/elastic"
"log"
"github.com/kataras/iris"
"github.com/google/uuid"
"net/http"
"github.com/google/go-tika/tika"
	"CasrIngestionService/models"
)

const (
	elasticIndexName = "documents"
	elasticTypeName = "document"
)


type DocumentResponse struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
}

type SearchResponse struct {
	Time string `json:"time"`
	Hits string `json:"hits"`
	Documents []DocumentResponse `json:"documents"`
}

var (
	elasticClient *elastic.Client
	tikaClient *tika.Client
)

func errorResponse(ctx iris.Context, code int, err string) {
	ctx.StatusCode(code)
	ctx.JSON(iris.Map{
		"error": err,
	})
}

func createDocumentsEndpoint(ctx iris.Context) {
	var docs []models.DocumentRequest
	if err := ctx.ReadJSON(&docs); err != nil {
		errorResponse(ctx, iris.StatusBadRequest, "Malformed request body")
		return
	}
	ids := make([]string, 0)
	bulk := elasticClient.
		Bulk().
		Index(elasticIndexName).
		Type(elasticTypeName)
	for _, d := range docs {
		doc := models.Document{
			ID: uuid.New().String(),
			Title: d.Title,
			Type: d.Type,
			IsHosted: d.IsHosted,
			CreatedAt: time.Now().UTC(),
			Content: d.Content,
			ExternalLink: d.ExternalLink,
			Tags: d.Tags,
		}
		ids = append(ids, doc.ID)
		bulk.Add(elastic.NewBulkIndexRequest().Id(doc.ID).Doc(doc))
	}
	if _, err := bulk.Do(ctx.Request().Context()); err != nil {
		log.Println(err)
		errorResponse(ctx, http.StatusInternalServerError, "failed to create documents")
		return
	}
	ctx.StatusCode(200)
	ctx.JSON(iris.Map{
		"ids": ids,
	})
	return
}

//Thjs is an internal service to ingest already parsed data into ElasticSearch
func main() {
	var err error
	for {
		elasticClient, err = elastic.NewClient(
			elastic.SetURL("https://bd7e24e33d7c4da4a47e6a7ef3298619.us-east-1.aws.found.io:9243"),
			elastic.SetBasicAuth("elastic", "AXfgOez9MmnF9Jq3j4iCJnnT"),
			elastic.SetSniff(false),
		)
		if err != nil {
			log.Println(err)
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}
	app := iris.Default()
	//TODO: Convert to GRPC interface or NATS interface
	app.Post("/documents", createDocumentsEndpoint)
	app.Run(iris.Addr(":8081"))
}