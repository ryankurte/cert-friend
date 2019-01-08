# cert-friend

**THIS PROJECT IS AN ASPIRATIONAL WORK IN PROGRESS, ALMOST NOTHING EXISTS / WORKS YET**

cert-friend is a friendly certificate and certificate authority utility, designed to support the *internal development use of self-signed CAs* and *production deployment of Public Key Infrastructure (PKI)*.

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

## Usage

cert-friend uses a pair of files to store and maintain infrastructure state. The *configuration* file contains the static configuration such as algorithms and names, as well as optionally a set of desired certificates. The *database* file contains a list of issued certificates and certificate states for management. Both of these files can be checked into version control to provide a simple mechanism for change management and tracking.

Broadly speaking, PKI consists of a certificate authority and a number of *intermediate*, *client*, and *server* certificates and cryptographic public-private key pairs. *server* certificates allow server identities to be validated, *client* certificates allow clients to be identified and the use of mutual TLS, and *intermediate* certificates allow signing of further *client* or *server* certificates, useful for delegating certificate creation to other services. 

To get started, run `cert-friend configure`. This will walk you through the configuration of a new PKI instance and create a configuration file for future use.

Once you've created a configuration, it's time to generate a CA. Run `cert-friend new-ca` to generate a new Certificate Authority. Again this will interactively walk you through the creation of a CA. `cert-friend new-ca --help` will list the possible arguments.

### For internal / development use


### For production PKI deployment


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


