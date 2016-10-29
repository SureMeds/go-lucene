# go-lucene
This is a go wrapper on the stratio lucene filter for cassandra. The example
below shows how to use this filter with gocql

```go
import (
  "encoding/json"
  "github.com/suremeds/go-lucene"
)

search := new(lucene.Search)
search.Filters = make([]lucene.Filter, 0, 2)
search.Filters = append(search.Filters, LuceneFilter{
  Type:  "prefix",
  Field: "first_name",
  Value: "Adi"})
}
search.Filters = append(search.Filters, LuceneFilter{
  Type:  "prefix",
  Field: "last_name",
  Value: "Suresh")
}
jsonFilter, err := json.Marshal(search)
if err != nil {
  return nil, err
}
cql += `SELECT * FROM users WHERE expr(users_index, '` + fmt.Sprintf("%s", jsonFilter) +
  `') LIMIT ` + strconv.FormatInt(limit, 10) + ``
iter = opsDb.Query(cql).Iter()
```
