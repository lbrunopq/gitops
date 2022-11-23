# criar credencicias Azure
```
az ad sp create-for-rbac \
    --name "SERVICE_NAME" \
    --scope /subscriptions/<SUBSCRIPTION_ID>/resourceGroups/<RESOURCE_GROUP> \
    --role Contributor \
    --sdk-auth
```

### Exemplo
```
az ad sp create-for-rbac \
    --name "k8s-goserver" \
    --scope /subscriptions/a3cc3bc8-4d7b-4d40-8684-7825b6f8ee9d/resourceGroups/goserv_group \
    --role Contributor \
    --sdk-auth


1 - saída:
{
  "clientId": "7204ed46-0802-4ed2-b54b-f59cc6e61235",
  "clientSecret": "iqi8Q~yPWs7V23~G6pgbBjVM7HSMoQCn3L3HxaUC",
  "subscriptionId": "a3cc3bc8-4d7b-4d40-8684-7825b6f8ee9d",
  "tenantId": "88b40e93-9291-4105-87e6-5cca5456da75",
  "activeDirectoryEndpointUrl": "https://login.microsoftonline.com",
  "resourceManagerEndpointUrl": "https://management.azure.com/",
  "activeDirectoryGraphResourceId": "https://graph.windows.net/",
  "sqlManagementEndpointUrl": "https://management.core.windows.net:8443/",
  "galleryEndpointUrl": "https://gallery.azure.com/",
  "managementEndpointUrl": "https://management.core.windows.net/"
}

2 - saída:
{
  "clientId": "7bdd8a53-86dc-48ff-83a1-31c4cd446148",
  "clientSecret": "IUw8Q~DzY41zQZacZQzSywNXZvQ.zCXLoF54hcNy",
  "subscriptionId": "a3cc3bc8-4d7b-4d40-8684-7825b6f8ee9d",
  "tenantId": "88b40e93-9291-4105-87e6-5cca5456da75",
  "activeDirectoryEndpointUrl": "https://login.microsoftonline.com",
  "resourceManagerEndpointUrl": "https://management.azure.com/",
  "activeDirectoryGraphResourceId": "https://graph.windows.net/",
  "sqlManagementEndpointUrl": "https://management.core.windows.net:8443/",
  "galleryEndpointUrl": "https://gallery.azure.com/",
  "managementEndpointUrl": "https://management.core.windows.net/"
}
```

