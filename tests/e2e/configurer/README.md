# End-to-end Tests

## The `configurer` package

The chain can either be initialized off of the current branch, or off the prior mainnet release and then upgraded to the current branch.

If current, we run chain initialization off of the current Git branch
by calling `chain.Init(...)` method in the `configurer/current.go`.

If with the upgrade, the same `chain.Init(...)` function is run inside a Docker container
of the previous Chain4Energy version, inside `configurer/upgrade.go`. This is
needed to initialize chain configs and the genesis of the previous version that
we are upgrading from.

The decision of what configuration type to use is decided by the `Configurer`.
This is an interface that has `CurrentBranchConfigurer` and `UpgradeConfigurer` implementations.
There is also a `BaseConfigurer` which is shared by the concrete implementations. However,
the user of the `configurer` package does not need to know about this detail.

When the desired configurer is created, the caller may
configure the chain in the desired way as follows:

```go
conf, _ := configurer.New(..., < isIBCEnabled bool >, < isUpgradeEnabled bool >)

conf.ConfigureChains()
```

The caller (e2e setup logic), does not need to be concerned about what type of
configurations is hapenning in the background. The appropriate logic is selected
depending on what the values of the arguments to `configurer.New(...)` are.

The configurer constructor is using a factory design pattern
to decide on what kind of configurer to return. Factory design
pattern is used to decouple the client from the initialization
details of the configurer. More on this can be found
[here](https://www.tutorialspoint.com/design_pattern/factory_pattern.htm)

The rules for deciding on the configurer type must be sent to each of the test suites.
We currently have 2 options to run the test suite

- If only `startUpgrade` is set to true only one chain using the previous version of the codebase will be started. 
Then an upgrade proposal will be sent and after reaching the appropriate height, the upgrade will be carried out,

- If only `startIbc` is set to true 2 chains will be launched, and they will be connected by the appropriate IBC relayer.

- If both `startUpgrade` and `startIbc` 2 chains will be launched, then an upgrade will be carried out and these 
2 chains will be connected using the IBC relayer

- If none are true, we only need one chain at the current branch version of the Chain4Energy code

## `upgrade` Package

The `upgrade` package starts chain initialization. In addition, there is
a Dockerfile `init-e2e.Dockerfile`. When executed, its container
produces all files necessary for starting up a new chain. These
resulting files can be mounted on a volume and propagated to our
production Chain4Energy container to start the `c4ed` service.

The decoupling between chain initialization and start-up allows to
minimize the differences between our test suite and the production
environment.