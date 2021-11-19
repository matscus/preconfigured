# preconfigured
Preconfigured files of various applications, to make my life easier=)

### Quick start 

#### Fill in the .env file with the required values.

```sh
#Build from your OS
make build

# Configured all app
make configure

#Configured single app
env APPLICATION=<set app dir name> make configure
```

### Result: 

a directory result will be created containing all files with set variables from .env

you can add your dir/files and variables for them, they will be processed.

### Currently implemented files for:

#### Prometheus
```sh
- Rules for different exporters/

- Daemon file with flags from master and scraper nodes/

- Coufiguration files (with dockerswarm_sd_configs for swarm )
```

#### Grafana
```sh
- Dashboards for different exporters, configured to use rules.

- Pre-congigered datasource.
```

#### Jenkins
```sh
- Contain three job from deploy Jmeter to swarm

- Files to jobs, for easy edit.

- Pre-congigered casc config

- plugin list

- scriptApproval list
```


#### Registry
```sh
- Pre-congigered  config
```

#### Postrgesql
```sh
- Statement con for pg_exporter version 12 and >.
```

#### Ammunition
```sh
- Daemon file with limins and flags
```

#### Alertmanager
```sh
- Pre-congigered  config

- Alertmanager bot configuration
```

#### Consul
```sh
- Pre-congigered  config
```

