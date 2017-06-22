## kubedb init

Create or upgrade unified operator

### Synopsis


Create or upgrade unified operator for kubedb databases.

```
kubedb init [flags]
```

### Examples

```
  # Create operator with version canary.
  kubedb init --version=0.1.0
  
  # Upgrade operator to use another version.
  kubedb init --version=0.1.0 --upgrade
```

### Options

```
  -h, --help                        help for init
      --operator-namespace string   Name of namespace where operator will be deployed. (default "kube-system")
      --upgrade                     If present, Upgrade operator to use provided version
      --version string              Operator version (default "0.1.0")
```

### Options inherited from parent commands

```
      --analytics             Send events to Google Analytics (default true)
      --kube-context string   name of the kubeconfig context to use
```

### SEE ALSO
* [kubedb](kubedb.md)	 - Controls kubedb objects


