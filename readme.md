# Iris Application using Goorm (from Infarma Sistemas)
### Context
This project was built in Golang, using Iris Framework and Goorm 
(a Go ORM developed by Infarma Sistema Team)

Goorm is a ORM built to work with MS SQL, so in this project i used one
database example from Microsoft (AdventureWorks2017)

#### Routes
1. Table Person

```
    /api/person                  -> Show all records from table Person
    /api/find/{BusinessEntityID} -> Find one record with BusinessEntityID
    /api/findby/{field}/{value}  -> Find one record with the combination field/value informed
  
```

