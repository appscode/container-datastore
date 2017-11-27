---
title: Kubedb Create
menu:
  product_kubedb_0.7.1:
    identifier: kubedb-create
    name: Kubedb Create
    parent: reference
product_name: kubedb
left_menu: product_kubedb_0.7.1
section_menu_id: reference
---
## kubedb create

Create a resource by filename or stdin

### Synopsis


Create a resource by filename or stdin. 

JSON and YAML formats are accepted.

```
kubedb create [flags]
```

### Examples

```
  # Create a elastic using the data in elastic.json.
  kubedb create -f ./elastic.json
  
  # Create a elastic based on the JSON passed into stdin.
  cat elastic.json | kubedb create -f -
```

### Options

```
  -f, --filename stringSlice   Filename to use to create the resource
  -h, --help                   help for create
  -n, --namespace string       Create object(s) in this namespace. (default "default")
  -R, --recursive              Process the directory used in -f, --filename recursively.
```

### Options inherited from parent commands

```
      --analytics             Send analytical events to Google Analytics (default true)
      --kube-context string   name of the kubeconfig context to use
```

### SEE ALSO
* [kubedb](/docs/reference/kubedb.md)	 - Command line interface for KubeDB


