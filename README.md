# cert-friend

cert-friend is a friendly certificate and certificate authority utility, designed to simplify the process and support people who want to use tls for stuff.

This consists of a core utility that helps create and manage local certificate infrastructure that can connect to and be run as a daemon to support remote certificate requests and signing.

cert-friend generates a human-readible database for certificate management that is suitable for storage in version control, and can be provide with a configuration file to set defaults and automate the creation of a certficate infrastructuure.

## Goals

- Encode best practices for certificate issuance and maintanence
- Simplify the process of safely creating and using certificate infrastructure
- Provide automation for inclusion with declarative infrastructure tools
- Inform users about decisions to be made in the process of deploying certificate infrastructure
- Provide simple mechanisms for managing certificates for
  - development environments and infrastructure
  - production environments
  - IoT stuff


## Status

*** Work In Progress ***

## Features

- [ ] Certificate Database
    - [ ] Add certificates
    - [ ] Add CSRs
    - [ ] Add Revocations
- [ ] Certificate Generation
    - [ ] Generate Certificate Authorities
    - [ ] Generate Intermediate Certificates
    - [ ] Generate Server Certificates
    - [ ] Generate Client Certificates
- [ ] Certificate Revocation
    - [ ] Revoke Intermediate Certificates
    - [ ] Revoke Server Certificates
    - [ ] Revoke Client Certificates
    - [ ] Generate revocation lists
- [ ] Certificate Transparency
    - [ ] Track certificate transparency
    - [ ] Generate Certificate Transparency logs
- [ ] Remote Interface
    - [ ] Generate Certificate
    - [ ] Request Signing
    - [ ] Request CSRs
    - [ ] Sign CSRs
- [ ] Yubikey support
    - [ ] Generate certs on key
    - [ ] Use on key certs for signing
- [ ] Step-by-step helper / CA creator walkthrough
- [ ] Connection tester (check cert validity, troubleshoot common TLS problems)

## Flows

### Local CAs / Dev 
Either load config file containing a list of desired certificates or run interactively (and provide option to generate config)

- [ ] Generate CA
- [ ] Generate Intermediate (optional? some things need this, others do not)
- [ ] Generate Server Cert(s)
- [ ] Generate Client Cert(s)

### Remote CAs


