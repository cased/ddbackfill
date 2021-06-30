# ddbackfill

Publish your Cased audit events to DataDog logs.

## Usage

1. Reach out to your Cased account representative to obtain an export of your Cased Audit events.
1. Download the [latest release](https://github.com/cased/ddbackfill/releases) of `ddbackfill` for your operating system.
1. [Setup DataDog Logs](https://www.notion.so/cased/How-to-build-an-Audit-trail-with-Datadog-592d50b4fea14857a82bdf683c91b27e)
1. Run `ddbackfill`

```sh
export DD_SITE="datadoghq.com" DD_API_KEY="<API-KEY>" DD_APP_KEY="<APP-KEY>"
ddbackfill ~/path-to-audit-events
```

If `ddbackfill` stops for any reason, it will resume where it left off.

## Debug

If any issues arise while running `ddbackfill`, you can enable debug mode by running:

```sh
DEBUG=1 ddbackfill ~/path-to-audit-events
```
