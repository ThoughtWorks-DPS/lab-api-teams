## Notes

TODO - come up with a migrations pattern

TODO - The build step in circle could also at the end deploy to a sandbox or kind cluster
and validate that the release IS deployable with the e2e tests.

Openapi spec 2.0/3.0
Gin, fast api 
Harvest, structurally what has happened at VA, RB
Lambda, pign it interval, interaction. Or just build in your own interval
Github Webhooks - might be noisy
Talk to a Notifications API as opposed to notifying everything
 - abstraction between the api(s) and event subscriptions
 - at least once message queue vs once and only once
 - take sns out of the mix, not a durable queue
 - we're trying to create an at least once queue



 Team api CREATE TEAM -> Sync Teams -> Notifications api subscribed, receives -> Notifications API delegates to API

 We don't need all the same levels of durability of once and only once.
 If it fails that moment, ok, no big deal. Next one will get it. Time is also a trigger.
 CloudEvents/NATS/etc.

 Put it in line with, complies w/ CE/redis/whatever
 still has repository/service pattern so it's somewhat easy to change


 Define the Sync Team Event
    What happens when a team is created?
        - gh team
        - namespace records in the namespace api/data
        - does not call a remote cluster, remote clusters will reconcile by checking the team/ns api
        - etc.
    How do we make the sync fast?
        - the thing sync'ing at remote cluster level, is a custom operator
        - public event api (public to the remote clusters, not PUBLIC public) that cluster operator subscribes to

Starting out from scratch
- deploy the teams api, masters are empty
- PO, creates dev/test/prod to start (master records)
- now you create a team or teams
- remote clusters start reconciling and get events right away, or from the next cron event

DI managed namespaces didn't have resource limits

GIN
- Liveness and readiness response, include TIME
- Gin good practices
- engineering practices for liveness/readiness

Backstage
- Plugin
- create teams, ns's, envs, etc.

DPS Styra Instance, for Authz on actions of this API
- OPA for delete team
- OPA for master ns records
- OPA for team update functions
- etc.

Platforms support
- Additional layer of abstraction
- Potentially treat downstream platforms (Humantec as an example) as the source of truth in the future
- Allow us to plugin to whatever a client has (AD, Humanitec, etc.)

Backend store
- Consider using standard protocol for repository interactions
- meaning, SQL can be ANY sql. Repository implements talking to A sql database (whether that's cosmo, sql server, etc.)
- similarly, mongodb, etc. Pretty much every cloud has some flavor of Mongo backend, for example
