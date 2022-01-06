# `ghcr.io/ketronkowski/newrelic-java-agent`
The New Relic Java Agent Buildpack is a Cloud Native Buildpack that contributes the New Relic Java Agent and 
configures the java application to use the agent.

## Behavior
This buildpack will participate if all the following conditions are met

* An BP_NEW_RELIC_AGENT_ENABLED environment variable exists with value of `true`

The buildpack will do the following for Java applications:

* Contributes a Java agent to a layer and configures `JAVA_TOOL_OPTIONS` to use it via `--javaagent` command line flag
  
## License
This buildpack is released under version 2.0 of the [Apache License][a].

[a]: http://www.apache.org/licenses/LICENSE-2.0

