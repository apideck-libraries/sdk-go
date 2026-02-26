# DataScopesResources

Data scopes resource configuration that can be either detailed field permissions or a wildcard


## Supported Types

### 

```go
dataScopesResources := components.CreateDataScopesResourcesMapOfMapOfDataScopesResources1(map[string]map[string]components.DataScopesResources1{/* values here */})
```

### DataScopesResources2

```go
dataScopesResources := components.CreateDataScopesResourcesDataScopesResources2(components.DataScopesResources2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch dataScopesResources.Type {
	case components.DataScopesResourcesTypeMapOfMapOfDataScopesResources1:
		// dataScopesResources.MapOfMapOfDataScopesResources1 is populated
	case components.DataScopesResourcesTypeDataScopesResources2:
		// dataScopesResources.DataScopesResources2 is populated
}
```
