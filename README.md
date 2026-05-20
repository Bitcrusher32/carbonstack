# CarbonStack

CarbonStack is an open-source secure communications appliance project.

Its goal is not to build a secure text comms surface using a smartphone base.

CarbonStack is composed of:

- **CarbonStackOS** — a deliberately constrained Android-derived operating system for a single-purpose communications appliance.
- **CarbonStackComms** — a text-first encrypted messaging client.
- **CarbonStackCypher** — a hostile-server relay and storage stack.
- **CarbonStack Protocol Docs** — shared threat model, identity model, trust-state model, text policy, file policy, and recovery model.

Core doctrine:

> Every feature is guilty until it proves it does not add unacceptable parser, network, sensor, identity, or filesystem authority.

CarbonStack prioritizes boring text, loud trust changes, hostile-server assumptions, disposable parsers, immutable base images, and minimal ambient attack surfaces. Convenience smartphone features are deprioritized over security.
