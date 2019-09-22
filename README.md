# HCEN: Hyper-Converged Edge Node

A Hyper-Converged Edge Node processes business transactional data
and censor data at the location where the data is created or captured.

## Components
This software comprises:
1. An event store to capture and distribute events (implemented with NATS Streaming)
1. A transactional database (implemented with DB2)
1. An analytics database (implemented with DB2)
1. Application skeleton code:
  1. Point of Sale transactional application
  1. Self-service price check and recommendations kiosk
