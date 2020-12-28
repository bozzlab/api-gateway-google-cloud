# API Gateway services on Google Cloud Platform 
The tutorial, Cloud Endpoints and API Gateway (Beta)

## Medium (Thai Lanauge)
https://p-srinikorn.medium.com/%E0%B9%80%E0%B8%A3%E0%B8%B4%E0%B9%88%E0%B8%A1%E0%B8%95%E0%B9%89%E0%B8%99%E0%B9%83%E0%B8%8A%E0%B9%89-api-gateway-service-%E0%B8%9A%E0%B8%99-google-cloud-platform-af5f33cd4abc

# Example

## Cloud Endpoints 
![Diagram](https://miro.medium.com/max/700/1*6Gru73r0LL9bOzRgwjSs7A.png)

## Cloud Endpoints 
![Diagram](https://miro.medium.com/max/700/1*qB_7fVmbf-_IF7ACd1dJPg.png)


# Repo structure 

## API services 
```
├── golang_webservice
|   ├── Dockerfile
|   ├── go.mod
|   ├── go.sum
|   └── main.go
├── nodejs_webservice
|   ├── data.js
|   ├── Dockerfile
|   ├── main.js
|   ├── middleware
|   │   └── logger.js
|   ├── package.json
|   ├── package-lock.json
|   └── routes
|      ├── index.js
|      └── stock.js
└── python_webservice
    ├── Dockerfile
    ├── main.py
    ├── Pipfile
    ├── Pipfile.lock
    └── requirements.txt

```
These folders are the application and Dockerfile, each language and services.    
Please build and deploy to your compute instance 

```
docker build -t <YOUR_IMAGE_REGISTRY_URL> . 
docker push <YOUR_IMAGE_REGISTRY_URL>
```

## API Gateway Config (OpenAPI) 
```
├── api_gateway_config
    ├── api-config-endpoints.yaml
    ├── api-config-gateway.yaml
    └── build_espv2

```
The OpenAPI Docs, Configurations for Cloud Endpoints and API Gateway (Beta).   
Please replace the <URL_HOSTNAME> to your URL domain before uses it.
