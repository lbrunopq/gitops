# criar credencicias Azure
```
az ad sp create-for-rbac \
    --name "k8s-goserver" \
    --scope /subscriptions/<SUBSCRIPTION_ID>/resourceGroups/<RESOURCE_GROUP> \
    --role Contributor \
    --sdk-auth
```

### Exemplo
```
az ad sp create-for-rbac \
    --name "k8s-goserver" \
    --scope /subscriptions/a3cc3bc8-4d7b-4d40-8684-7825b6f8ee9d/resourceGroups/goserver_group \
    --role Contributor \
    --sdk-auth


saída:
{
  "clientId": "7204ed46-0802-4ed2-b54b-f59cc6e61235",
  "clientSecret": "JxX8Q~FToTjHX7p3F6a3vpM1FbfjrbcClAaRycT1",
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