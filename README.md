# DDD Design with Golang

```mermaid
flowchart LR
    A[User Request] --> B[Controller]
    B --> C[Use Case]
    C --> D[Repository]
    D --> E[(Database)]
```