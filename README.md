# Fitrkr Atlas

Fitrkr Atlas is an open-source exercise database providing structured exercise data with metadata like categories, muscles, equipment, and variations.

---

## Features
- Comprehensive exercise dataset (validated, structured)
- Searchable and filterable by category, muscle, or equipment
- RESTful API with CQRS architecture
- Self-hostable with PostgreSQL backend
- Clean architecture with DDD principles

---

## Usage

### Self-Hosted
Clone and run your own instance:
```bash
# Setup instructions coming soon
go run cmd/web/main.go
```

### Managed API (Coming Soon)
Use our hosted service at `api.fitrkr.com`:
- **Free tier**: Basic read access with rate limits
- **Pro/Enterprise**: Higher limits, write access, priority support

---

## Current Status
**In Development** 🚧

Completed:
- Equipment & Attachments
- Muscle Groups & Muscles  
- Categories & Subcategories
- Repository layer with CQRS

In Progress:
- Exercise entities (core feature)
- Exercise CRUD operations
- Advanced filtering & search

Planned:
- Comprehensive test coverage
- Hosted API with authentication (atlas-api)
- GraphQL support
- Community contribution system

---

## Contributing
Contribution guidelines coming after v1.0 release.

---

## Licensing
**Open Source**: MIT License for self-hosting  
**Hosted API**: Separate commercial terms (see `/LICENSE_ENTERPRISE` when available)

Similar to Supabase's model: free to self-host, paid for managed hosting.

---

## Architecture
- **Domain Layer**: Exercise entities and business rules
- **Application Layer**: CQRS commands and queries
- **Infrastructure**: PostgreSQL with separate read/write repos
- **API Layer**: Chi router with clean separation of concerns

Part of the Fitrkr ecosystem alongside:
- **fitrkr-athena**: Workout logging backend
- **fitrkr-prometheus**: Analytics & insights (planned)
- **fitrkr-hera**: Admin dashboard
