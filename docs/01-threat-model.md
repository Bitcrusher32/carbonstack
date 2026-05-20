# Threat Model

CarbonStack assumes hostile conditions by default.

## Assumed Hostile

- the network may be hostile
- the server may be hostile
- inbound files may be hostile
- other users' devices may become hostile
- parsers may contain vulnerabilities
- metadata may be observable
- convenience features may silently expand authority

## Protected Against

CarbonStack aims to protect against:

- passive network observers
- malicious or compromised relay servers
- silent server-side identity replacement
- plaintext message exposure on the server
- casual device misuse
- broad parser compromise from non-essential features
- accidental trust changes
- ordinary user workflows that expand attack surface

## Not Fully Protected Against

CarbonStack does not fully solve:

- compromised endpoints
- malicious signed updates
- kernel vulnerabilities
- hardware implants
- supply-chain compromise
- physical coercion
- recipient betrayal
- advanced traffic analysis
- compromised build systems

## Core Truth

Encrypted messaging cannot protect plaintext once it is displayed on a compromised endpoint.
