**Project:** Veil Routing Protocol  
**Abbreviation:** VRP  
**Original Author and Architect:** Vitalijus Riabovas  
**Public Engineering Identity:** Endless33  
**Country:** Lithuania  
**Document Status:** Public Assurance Entry Point  
**Assurance Package Version:** 1.0  
**Repository:** `vrp-engineer-evaluation-kit`

---

## 1. Purpose

This document is the primary entry point for evaluating the identity, safety, engineering purpose, declared behavior, verification boundaries, and public assurance materials of the Veil Routing Protocol project.

This document does not ask reviewers to trust VRP because of marketing statements, personal claims, visual demonstrations, or the identity of its author.

It provides a structured path for independently examining:

- what VRP is;
- what VRP is not;
- who created and maintains it;
- what this public evaluation repository contains;
- what the supplied software is expected to do;
- what the supplied software is not expected to do;
- which files, processes, privileges, and network resources it may use;
- how published artifacts can be inspected;
- how engineering claims can be reproduced;
- how evidence can be independently verified;
- how suspicious behavior or security concerns can be reported;
- which limitations and unresolved risks remain.

The central principle of this assurance package is:

> Do not trust VRP because of who created it.  
> Verify VRP because its boundaries, evidence, tooling, and procedures make independent review possible.

---

## 2. Project Identity

The Veil Routing Protocol is an independently developed networking and distributed-systems architecture created by Vitalijus Riabovas under the public engineering identity Endless33.

VRP is designed around continuity-first protocol semantics.

Its engineering model treats transport paths, network interfaces, endpoints, and temporary connectivity conditions as changeable operational resources rather than as the permanent identity of a logical session.

The project focuses on such properties as:

- session continuity across transport-path changes;
- controlled recovery after temporary disruption;
- replay rejection;
- duplicate-operation resistance;
- deterministic admission behavior;
- monotonic authority transitions;
- fail-closed decision handling;
- tamper-evident evidence generation;
- independently reproducible validation;
- explicit trust-boundary separation.

VRP is not presented as a consumer VPN product.

The name “Veil Routing Protocol” does not imply that this repository contains a complete commercial networking service, a universally deployable replacement for every existing protocol, or unrestricted access to the protected VRP runtime.

---

## 3. Repository Purpose

The `vrp-engineer-evaluation-kit` repository exists to support structured engineering evaluation.

Its purpose is to help reviewers:

1. understand the public VRP evaluation boundary;
2. inspect available source code and dependencies;
3. execute declared validation scenarios;
4. review generated evidence;
5. test rejection of modified or invalid evidence;
6. compare observed behavior with published claims;
7. document independent findings;
8. identify limitations, inconsistencies, or unresolved concerns.

This repository must not be treated as evidence that every private or future VRP component has been disclosed.

Public evaluation materials and protected runtime materials are intentionally separated.

That separation is designed to permit meaningful validation without publishing proprietary implementation logic, private authority mechanisms, protected decision mathematics, confidential key material, or internal runtime algorithms.

---

## 4. What VRP Is

Within the scope of publicly documented engineering objectives, VRP is:

- a continuity-first protocol architecture;
- a distributed-systems engineering project;
- a framework for evaluating behavior under failure and recovery;
- a project with explicit admission, authority, replay, and evidence boundaries;
- a system designed to produce testable and reviewable outcomes;
- an architecture whose public claims should be tied to specific validation procedures;
- a project intended for controlled technical evaluation before production consideration.

VRP seeks to answer engineering questions such as:

- Can a logical session remain valid when its transport path changes?
- Can temporary path failure be handled without automatically treating the logical session as permanently dead?
- Can replayed operations be rejected consistently?
- Can concurrent or duplicated decisions result in one canonical accepted outcome?
- Can authority transitions be made monotonic and resistant to rollback?
- Can exported evidence be verified and tampering detected?
- Can failures produce explicit, auditable verdicts instead of ambiguous behavior?

These are engineering objectives.

They are not guarantees of universal security, universal compatibility, or suitability for every environment.

---

## 5. What VRP Is Not

VRP is not represented by this assurance package as:

- malware;
- spyware;
- ransomware;
- a credential-harvesting tool;
- a remote-access trojan;
- a persistence mechanism;
- a cryptocurrency miner;
- a botnet component;
- a destructive payload;
- an endpoint-surveillance product;
- a tool for bypassing organizational authorization;
- a mechanism for unauthorized access to third-party systems;
- a guaranteed replacement for TCP, UDP, TLS, QUIC, WireGuard, IPsec, or every other networking technology;
- an absolutely secure or “unbreakable” system;
- formally verified unless a specific component is explicitly documented as formally verified;
- independently certified unless a named independent organization issues such certification;
- production-ready solely because public tests pass;
- trustworthy solely because the author declares it trustworthy.

No reviewer should accept a VRP component merely because it carries the VRP name.

Every supplied artifact should be inspected, hashed, isolated, monitored, and validated according to the organization’s own security procedures.

---

## 6. Public Evaluation Safety Objective

The public evaluation kit is intended to support controlled inspection and testing.

Unless explicitly documented for a particular scenario, public evaluation components are not designed to:

- establish operating-system persistence;
- modify boot configuration;
- create hidden administrator accounts;
- install kernel modules;
- disable antivirus or endpoint security products;
- disable host firewalls;
- capture keyboard input;
- access browser password databases;
- collect unrelated personal documents;
- harvest authentication credentials;
- scan unrelated private directories;
- silently exfiltrate files;
- open undeclared remote-control channels;
- propagate to other systems;
- conceal running processes;
- mine cryptocurrency;
- destroy user data;
- encrypt unrelated files;
- modify unrelated repositories;
- execute self-deleting payloads;
- evade organizational monitoring controls.

Any observed behavior inconsistent with the declared scope must be treated as a security concern and investigated before further execution.

This statement is a declaration of intended behavior, not a substitute for independent inspection.

---

## 7. Required Reviewer Mindset

Organizations evaluating VRP should apply a zero-trust review process.

Recommended assumptions include:

- the author’s statements may be incomplete;
- documentation may contain errors;
- source code may contain defects;
- dependencies may introduce risk;
- build infrastructure may be compromised;
- release artifacts may be replaced or modified;
- runtime behavior may differ from documentation;
- test success may cover only the tested scope;
- a compromised host may invalidate security assumptions;
- private components require separate contractual and technical review.

Reviewers should independently establish:

- the exact repository origin;
- the exact commit being evaluated;
- the dependency set;
- the build toolchain;
- the generated binary hashes;
- the required privileges;
- the filesystem changes;
- the spawned processes;
- the listening sockets;
- the outbound connections;
- the generated evidence;
- the cleanup behavior;
- the difference between expected and observed behavior.

---

## 8. Verification Before Execution

Before executing any VRP evaluation component, reviewers should:

1. confirm that the repository owner is the expected canonical owner;
2. record the complete Git commit identifier;
3. inspect the repository status;
4. inspect recent commits and changed files;
5. review build scripts and workflow files;
6. inspect direct and transitive dependencies;
7. search for embedded credentials or secrets;
8. inspect process-execution calls;
9. inspect filesystem-write operations;
10. inspect network-dialing and listening operations;
11. inspect privilege requirements;
12. inspect cleanup and shutdown paths;
13. build inside an isolated environment;
14. calculate cryptographic hashes;
15. retain an evaluation log.

Example repository inspection commands may include:

```bash
git remote -v
git rev-parse HEAD
git status --short
git log --oneline --decorate -n 20
git diff --check
go env
go version
go list -m all
go mod verify
go test ./...
go test -race ./...
go vet ./...

Reviewers remain responsible for determining whether these commands are sufficient for their environment.

---

## 9. Recommended Evaluation Environment

Initial evaluation should occur inside an isolated and disposable environment.

Recommended controls include:

- a dedicated virtual machine;
- no production credentials;
- no personal files;
- no access to sensitive internal networks;
- no mounted cloud credentials;
- restricted outbound connectivity;
- host-level process monitoring;
- filesystem monitoring;
- packet capture;
- resource limits;
- reproducible snapshots;
- complete command logging;
- destruction of the environment after testing.

Production deployment should not be considered until the organization has completed its own security, legal, privacy, architecture, and operational reviews.

---

## 10. Declared Trust Boundaries

The VRP project uses explicit separation between several conceptual boundaries.

### 10.1 Participant-Controlled Environment

The participant controls:

- the evaluation host;
- operating-system configuration;
- network access;
- test data;
- monitoring tools;
- credentials;
- deployment decisions;
- production authorization.

### 10.2 Public Evaluation Boundary

The public evaluation boundary may expose:

- interfaces;
- schemas;
- validation commands;
- deterministic scenarios;
- evidence formats;
- expected verdicts;
- public documentation;
- integration examples;
- non-sensitive test implementations.

### 10.3 Protected Runtime Boundary

The protected runtime boundary may contain non-public material, including:

- proprietary algorithms;
- private decision logic;
- protected authority mechanisms;
- confidential runtime state;
- private cryptographic operations;
- internal recovery logic;
- protected source code;
- commercial or security-sensitive implementation details.

The existence of a public evaluation kit does not grant access to the protected runtime.

### 10.4 Evidence Boundary

The evidence boundary exports reviewable results without necessarily exposing protected internal implementation details.

Evidence should be evaluated for:

- schema validity;
- provenance;
- canonical representation;
- cryptographic integrity;
- deterministic recomputation where applicable;
- timestamp validity;
- tamper detection;
- downgrade resistance;
- consistency with the executed scenario.

---

## 11. Claims Must Be Scoped

Every meaningful engineering claim should identify:

- the exact component under evaluation;
- the exact version or commit;
- the environment;
- the command;
- the scenario;
- the expected result;
- the observed result;
- the generated evidence;
- known limitations;
- whether the result was independently reproduced.

A successful test does not prove more than the test actually covers.

For example:

- replay-rejection tests do not prove protection against every possible attack;
- evidence-tampering tests do not prove that the host was uncompromised;
- race tests do not prove correctness under every distributed deployment;
- a successful laboratory migration does not prove production suitability;
- passing static analysis does not prove the absence of vulnerabilities;
- reproducible output does not prove that the architecture is universally correct.

The assurance package must distinguish between:

- implemented behavior;
- tested behavior;
- observed behavior;
- claimed behavior;
- planned behavior;
- unsupported assumptions.

---

## 12. Author Identity and Responsibility

The original author and architect of VRP is:

**Vitalijus Riabovas**  
**Public engineering identity:** Endless33  
**Country:** Lithuania

The author is responsible for accurately representing:

- the project’s origin;
- the published engineering scope;
- the official repositories;
- the official communication channels;
- the intended behavior of supplied components;
- known limitations;
- known security concerns;
- release and provenance information under the author’s control.

The author cannot guarantee:

- the security of third-party mirrors;
- the integrity of unofficial binaries;
- the behavior of modified forks;
- the security of compromised evaluation hosts;
- the behavior of undisclosed third-party integrations;
- the accuracy of claims made by unaffiliated parties;
- the absence of every software defect.

Reviewers should use repository history, signed releases where available, cryptographic hashes, public records, and reproducible procedures to assess authorship and artifact integrity.

---

## 13. Official Artifact Policy

Reviewers should not execute VRP binaries or archives received through unsolicited private messages.

Preferred artifact sources are:

- the canonical public repository;
- an official tagged release;
- an artifact with a published SHA-256 hash;
- an artifact with verifiable build provenance;
- an artifact signed through a documented signing process.

Before execution, confirm:

- repository ownership;
- release tag;
- commit identifier;
- checksum;
- signature where available;
- build provenance;
- dependency inventory;
- expected file list.

Any mismatch must be investigated.

---

## 14. Security and Privacy Review Questions

Before approving further evaluation, organizations should be able to answer:

- What exact code will run?
- Which user account will run it?
- Does it require elevated privileges?
- Which files will it read?
- Which files will it write?
- Which processes will it start?
- Which ports will it open?
- Which remote addresses will it contact?
- Does it require DNS?
- Does it collect telemetry?
- Does it transmit machine identifiers?
- Does it retain packet contents?
- Does it access credentials?
- Can it run completely offline?
- How is it stopped?
- How is it removed?
- Which logs and evidence remain afterward?
- How are release artifacts authenticated?
- Which behaviors remain outside the current validation scope?

Detailed answers must be maintained in the supporting trust documents.

---

## 15. Security Concerns and Unexpected Behavior

Stop execution and preserve evidence if any component appears to:

- contact an undeclared external destination;
- request unexplained privileges;
- access unrelated user files;
- modify operating-system security configuration;
- create persistence;
- launch an undocumented executable;
- conceal a process;
- write outside declared directories;
- transmit unexpected data;
- alter unrelated repositories;
- continue running after documented shutdown;
- produce inconsistent or unverifiable evidence.

A useful report should include:

- repository URL;
- commit identifier;
- operating system;
- toolchain version;
- exact command;
- complete output;
- process list;
- network capture;
- filesystem changes;
- generated hashes;
- reproduction steps;
- relevant timestamps.

Security concerns should be reported through the official security contact published by the project.

---

## 16. Assurance Documents

This trust package is intended to include dedicated documents covering:

- project and system identity;
- software safety;
- non-malicious behavior declarations;
- architecture assurance;
- threat modeling;
- privacy and data boundaries;
- network behavior;
- build and release provenance;
- independent verification;
- authorship and project history;
- claims and evidence mapping;
- limitations and non-claims;
- Pilot trust requirements;
- responsible disclosure;
- maintainership and security contacts;
- machine-readable attestations;
- cryptographic checksums.

Each document has a limited scope.

No single document should be treated as complete proof of system safety.

---

## 17. Independent Review Outcomes

An independent reviewer should avoid reducing the result to “trusted” or “not trusted” without context.

Recommended outcomes are:

### VERIFIED

The reviewed behavior matched the declared scope within the documented environment, and the relevant evidence was successfully verified.

### VERIFIED WITH LIMITATIONS

The reviewed behavior matched the declared scope, but important environmental, operational, implementation, or coverage limitations remain.

### INCONCLUSIVE

The available materials, access, evidence, or test coverage were insufficient to reach a reliable conclusion.

### REJECTED

The observed behavior contradicted declared behavior, evidence verification failed, provenance could not be established, or unacceptable security concerns were identified.

Every outcome should identify:

- the tested commit;
- the environment;
- the commands;
- the evidence;
- the limitations.

---

## 18. Limitations of This Assurance Package

This assurance package cannot prove:

- that the software contains no vulnerabilities;
- that every future release will behave identically;
- that every dependency is trustworthy;
- that every build environment is uncompromised;
- that a private runtime is safe without reviewing it;
- that a Pilot deployment is appropriate for every organization;
- that a compromised host can be made trustworthy;
- that all attacks are prevented;
- that public evidence represents events occurring outside the evaluated environment;
- that organizational procedures were followed correctly.

Documents created and signed by the project author are first-party declarations.

They become more meaningful when combined with:

- reproducible builds;
- independent source review;
- dependency analysis;
- isolated execution;
- network monitoring;
- cryptographic checksums;
- signed release tags;
- machine-readable provenance;
- external testing;
- independently published findings.

---

## 19. Engineering Principle

VRP should not be accepted because it appears advanced.

It should not be rejected because it is independently developed.

It should be evaluated through evidence.

The relevant questions are:

- Is the architecture clearly defined?
- Are trust boundaries explicit?
- Are claims limited to observable scope?
- Can the behavior be reproduced?
- Can evidence be independently verified?
- Are tampered artifacts rejected?
- Are limitations acknowledged?
- Are suspicious behaviors absent from the inspected implementation?
- Can the organization control the evaluation environment?
- Does the observed result justify continuing to the next stage?

The final decision belongs to the evaluating organization.

---

## 20. Final Statement

VRP is presented as a serious engineering architecture whose claims must survive inspection, testing, adversarial review, and independent verification.

No organization is expected to rely solely on the author’s confidence.

No reviewer is expected to disable normal security controls.

No Pilot participant is expected to accept undeclared behavior.

The correct path is:

**Identify. Inspect. Isolate. Build. Monitor. Test. Verify. Document. Decide.**

Trust is not requested in advance.

Trust must be earned through reproducible evidence.

---

**Vitalijus Riabovas**  
Original Author and Architect of the Veil Routing Protocol  
Public Engineering Identity: Endless33  
Lithuania
