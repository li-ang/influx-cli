## v2.1.0 [unreleased]

### New Repository

This is the initial release of the `influx` CLI from the [`influxdata/influx-cli`](https://github.com/influxdata/influx-cli/)
GitHub repository.

### Breaking Changes

#### `influx write` skip-header parsing

The `write` command no longer supports `--skipHeader` as short-hand for `--skipHeader 1`. This change was made to
simplify our CLI parser.

#### Stricter input validation for template-related commands

The `apply`, `export`, and `stacks` commands now raise errors when CLI options fail to parse  instead of silently
discarding bad inputs. This change was made to help users debug when their commands fail to execute as expected.

#### Server-side template summarization & validation

The `template` and `template validate` commands now use an API request to the server to perform their logic,
instead of performing the work on the client-side. Offline summarization & validation is no longer supported.
This change was made to avoid significant code duplication between `influxdb` and `influx-cli`, and to allow server-
side template logic to evolve without requiring coordinated CLI changes.

### Features

1. [33](https://github.com/influxdata/influx-cli/pull/33): Add global `--http-debug` flag to help inspect communication with InfluxDB servers.
1. [52](https://github.com/influxdata/influx-cli/pull/52): Add `bucket-schema` commands to manage explicit measurement schemas in InfluxDB Cloud.
1. [52](https://github.com/influxdata/influx-cli/pull/52): Update `bucket create` to allow setting a schema type.
1. [52](https://github.com/influxdata/influx-cli/pull/52): Update `bucket list` to display schema types.
1. [116](https://github.com/influxdata/influx-cli/pull/116): Reimplement `backup` to support downloading embedded SQL store from InfluxDB v2.1.x.
1. [116](https://github.com/influxdata/influx-cli/pull/116): Add `--compression` flag to `backup` to support enabling/disabling GZIP compression of downloaded files.
1. [121](https://github.com/influxdata/influx-cli/pull/121): Reimplement `restore` to support uploading embedded SQL store from InfluxDB v2.1.x.
1. [191](https://github.com/influxdata/influx-cli/pull/191): Add `--password` flag to `user password` command to allow bypassing interactive prompt.

### Bug Fixes

1. [35](https://github.com/influxdata/influx-cli/pull/35): Fix interactive password collection & color rendering in PowerShell.
1. [97](https://github.com/influxdata/influx-cli/pull/97): `org members list` no longer hangs on organizations with more than 10 members.
1. [109](https://github.com/influxdata/influx-cli/pull/109): Detect & warn when inputs to `write` contain standalone CR characters.
1. [122](https://github.com/influxdata/influx-cli/pull/122): `dashboards` command now accepts `--org` flag, or falls back to default org in config.
1. [140](https://github.com/influxdata/influx-cli/pull/140): Return a consistent error when responses fail to decode, with hints for OSS- our Cloud-only commands.