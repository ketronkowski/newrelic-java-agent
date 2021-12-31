# `gcr.io/paketo-buildpacks/newrelic-java-agent`
The New Relic Java Agent Buildpack is a Cloud Native Buildpack that contributes the New Relic Java Agent and 
configures the java application to use the agent.

## Behavior
This buildpack will participate if all the following conditions are met

* An BPL_NEW_RELIC_AGENT_ENABLED environment variable exists with value of `true`

The buildpack will do the following for Java applications:

* Contributes a Java agent to a layer and configures `JAVA_TOOL_OPTIONS` to use it via `--javaagent` command line flag
* Transforms the following "BPL" environment variables to runtime environment variables  (without the BPL_ prefix) 
  if set
  * BPL_NEW_RELIC_AGENT_ENABLED
  * BPL_NEW_RELIC_APP_NAME
  * BPL_NEW_RELIC_DISTRIBUTED_TRACING_ENABLED
  * BPL_NEW_RELIC_LICENSE_KEY
  
## License
This buildpack is released under version 2.0 of the [Apache License][a].

[a]: http://www.apache.org/licenses/LICENSE-2.0

