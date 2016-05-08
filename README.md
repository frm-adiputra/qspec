# QSpec

QSpec generates source code for data access.


## Features

- [x] Using YAML

## TODO

- [x] Shared result struct
- [x] Imports additional package
- [ ] Query parameters validation
- [ ] Struct as query parameters
- [ ] Customizable query result (partially select columns)

## Code Generation

Each package spec (YAML file) will generate a package with the same name as the
YAML file name.
Inside a package spec can contains multiple queries spec and also structs spec
related to those queries.

## QSpec specification

### PackageSpec

This is the root document object for the specification.

#### Fixed fields

Name | Type | Description
---|---|---
`description` | `string` | The package description.
`imports` | `[]string` | Additional package imports.
`structs` | [`map[string]StructSpec`](#structspec) | Structs specification.
`queries` | [`map[string]QuerySpec`](#queryspec) | Queries specification.

### StructSpec

Each direct key descendant will be the name of the generated struct.

#### Fixed fields

Name | Type | Description
---|---|---
`description` | `string` | The struct description.
`fields` | [`[]FieldSpec`](#fieldspec) | List of the struct's fields.

#### Examples

```yaml
structs:
  Blog:
    description: represents a blog.
    fields:
      # ...
  Comment:
    description: represents a comment.
    fields:
      # ...
```

### QuerySpec

Each direct key descendant will be the name of the generated query.
For each QuerySpec, QSpec will generate code that provides two way to execute the
query, directly or using a prepared statement.

#### Fixed fields

Name | Type | Description
---|---|---
`description` | `string` | The query description.
`statement` | `string`  | List of the struct's fields.
`prepared` | `bool`  | If true then the prepared version of the query will also be generated.
`result` | [ResultSpec](#resultspec) | Specifies the result of the query.

### ResultSpec

TODO

### ResultStructSpec

TODO

### FieldSpec

TODO
