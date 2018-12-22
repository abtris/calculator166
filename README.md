# Calculator 166

This repository contains demo application using for training.


## Main parts

- Worker (generate load)
- Router (manage API)
- Services (by operations)

## API documentation

- https://calculator166cz.docs.apiary.io/

## Diagram

![](mermaid-diagram-20181220183016.svg)


## Testing

```bash
# Operations
curl -X POST -H 'Content-Type: application/json' -d "{\"numbers\": [1,2,33]}" http://localhost:8080/operation
```
