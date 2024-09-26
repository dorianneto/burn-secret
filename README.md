# burn-secret

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/dorianneto/burn-secret/cicd.yml)
![GitHub License](https://img.shields.io/github/license/dorianneto/burn-secret)


Web application open-source created for generating secrets from sensitive information.

## Requires

- Docker
- Docker compose
- OpenSSL
- Node

## Installation

1. Run docker
```
docker compose up -d
```

2. Generate the certificates
```
mkdir certs && make cert
```

3. You must be able to access `https://localhost`

## Usage

### Backend

The package Air is watching any changes made into GO files. Check `.air.toml` for further information.

### Frontend

You can find the React files into `web/src`.

The following command will watch any changes made into React files.

```
npm run dev
```
