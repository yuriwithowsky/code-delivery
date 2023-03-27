# DevTools no Kibana

## Criar indices via DevTools

- route.new-direction

```json
PUT route.new-direction
{
  "mappings": {
    "properties": {
      "clientId": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword"
          }
        }
      },
      "routeId": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword"
          }
        }
      },
      "timestamp": {
        "type": "date"
      }
    }
  }
}
```

- route.new-position

```json
PUT route.new-position
{
  "mappings": {
    "properties": {
      "clientId": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword"
          }
        }
      },
      "routeId": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword"
          }
        }
      },
      "timestamp": {
        "type": "date"
      },
      "finished": {
        "type": "boolean"
      },
      "position": {
        "type": "geo_point"
      }
    }
  }
}
```

## Remover dados de um indice

```json
POST route.new-position/_delete_by_query?conflicts=proceed
{
  "query": {
    "match_all": {}
  }
}

POST route.new-direction/_delete_by_query?conflicts=proceed
{
  "query": {
    "match_all": {}
  }
}
```
