# VRP Authors, Maintainers, and Project Identity

**Project:** Veil Routing Protocol  
**Abbreviation:** VRP  
**Canonical Public Identity:** Endless33  
**Original Author and Architect:** Vitalijus Riabovas  
**Country:** Lithuania  
**Document Status:** Public Authorship Record  
**Version:** 1.0

---

## 1. Purpose

This document identifies the original author, current maintainer, public engineering identity, project ownership boundary, and official authorship-verification methods for the Veil Routing Protocol project.

Its purpose is to help reviewers determine:

- who created VRP;
- who currently maintains the public project materials;
- which identity is authorized to represent the project;
- which repositories are considered official;
- how authorship can be independently checked;
- how unofficial claims, mirrors, forks, and redistributed artifacts should be treated;
- which responsibilities belong to the author;
- which responsibilities remain with independent reviewers and Pilot participants.

This document is a first-party project record.

It must be evaluated together with repository history, cryptographic signatures, release manifests, public records, and independently reproducible evidence.

---

## 2. Original Author

The original author and architect of the Veil Routing Protocol is:

**Vitalijus Riabovas**

Public engineering identity:

**Endless33**

Country:

**Lithuania**

Project role:

- original creator of the Veil Routing Protocol concept;
- protocol and runtime architecture author;
- primary designer of the continuity-first model;
- maintainer of the official public VRP repositories;
- maintainer of protected VRP implementation boundaries;
- author of public engineering documentation;
- author of public validation and evaluation materials;
- coordinator of controlled Pilot access;
- current authority for official project releases and public project statements.

---

## 3. Public Engineering Identity

The public engineering identity **Endless33** is used to represent the VRP project in public technical communication.

This identity may appear in:

- GitHub repository ownership;
- Git commit metadata;
- release notes;
- public engineering documentation;
- technical reports;
- validation materials;
- Pilot announcements;
- public social-media communication;
- authorship declarations;
- cryptographic manifests;
- evidence bundles;
- project contact instructions.

The use of the identity Endless33 does not replace the legal identity of the author.

Where legal, contractual, security, Pilot, or commercial verification is required, the name **Vitalijus Riabovas** must be used.

---

## 4. Author Background

Vitalijus Riabovas is an independent systems developer from Lithuania and the original creator of the Veil Routing Protocol project.

The project was developed through iterative implementation, testing, refactoring, failure analysis, adversarial validation, reproducibility work, security hardening, and public documentation.

The author’s engineering focus includes:

- networking systems;
- distributed systems;
- failure semantics;
- recovery behavior;
- session continuity;
- protocol-state transitions;
- replay rejection;
- deterministic admission;
- authority control;
- evidence verification;
- tamper detection;
- runtime boundaries;
- secure evaluation procedures;
- reproducible validation;
- security-oriented Go development.

The project must not be evaluated based only on biography.

The engineering claims of VRP must be evaluated through source inspection, reproducible tests, evidence verification, and controlled technical assessment.

---

## 5. Project Origin

VRP originated as an independently developed architecture focused on one central engineering principle:

> A transport-path failure should not automatically be treated as the permanent death of a logical session.

From this principle, the project expanded into a broader architecture involving:

- session identity independent from a single transport path;
- controlled path migration;
- recovery after temporary disruption;
- replay resistance;
- duplicate-event handling;
- deterministic commit selection;
- monotonic authority transitions;
- fail-closed admission;
- tamper-evident evidence;
- explicit verification boundaries;
- protected runtime separation;
- controlled Pilot evaluation.

The project was not created as a consumer anonymity product or as a conventional commercial VPN application.

The name “Veil Routing Protocol” identifies the architecture and project.

It does not imply that every repository contains a complete network stack or production-ready deployment.

---

## 6. Development Model

VRP has been developed primarily as an independent engineering project.

The development model includes:

1. defining protocol and runtime invariants;
2. implementing bounded components;
3. writing unit and adversarial tests;
4. running race and concurrency validation;
5. recording evidence;
6. identifying architectural conflicts;
7. refactoring when required;
8. separating protected logic from public validation materials;
9. publishing reproducible interfaces and test procedures;
10. preparing controlled Pilot evaluation.

Independent development does not remove the need for external review.

The project explicitly expects serious evaluators to:

- inspect the source;
- challenge the documentation;
- reproduce tests;
- modify evidence;
- test rejection paths;
- examine dependencies;
- monitor execution;
- document inconsistencies;
- report unresolved concerns.

---

## 7. Current Maintainer

At the time of this document’s publication, the primary maintainer of VRP is:

**Vitalijus Riabovas / Endless33**

The maintainer is responsible for:

- maintaining the canonical repositories;
- publishing official project documentation;
- identifying official release versions;
- maintaining public trust and assurance materials;
- disclosing known limitations;
- responding to valid security reports;
- maintaining project contact information;
- reviewing proposed changes;
- protecting the private runtime boundary;
- managing Pilot evaluation access;
- publishing cryptographic checksums where applicable;
- preserving the consistency of project identity.

The maintainer may delegate technical review, Pilot support, security assessment, or repository maintenance in the future.

Any delegation should be documented publicly before it is treated as authoritative.

---

## 8. Canonical Repository Ownership

The canonical public repository owner is:

**GitHub account:** `Endless33`

Official repositories must be confirmed directly through that account or through an official project document that identifies the repository.

Reviewers must not assume that a repository is official merely because:

- it contains the letters “VRP”;
- it uses the project logo;
- it copies public documentation;
- it references the author;
- it contains similar source code;
- it claims to be an official mirror;
- it distributes a binary with the VRP name.

Before relying on a repository, confirm:

- repository owner;
- repository URL;
- Git remote;
- commit history;
- release history;
- tags;
- signatures where available;
- references from canonical project documentation.

---

## 9. Official Repository Categories

The VRP public ecosystem may include repositories with different purposes.

These categories may include:

### 9.1 Engineer Evaluation Repository

Purpose:

- structured engineering review;
- reproducible evaluation procedures;
- trust and assurance documentation;
- reviewer guidance;
- evidence inspection;
- independent reporting.

### 9.2 Validation Repository

Purpose:

- execution of public validation scenarios;
- evidence verification;
- tamper-rejection testing;
- deterministic test procedures;
- public verdict generation.

### 9.3 Pilot Repository

Purpose:

- Pilot onboarding;
- evaluation requirements;
- Pilot security expectations;
- deployment preparation;
- Pilot success criteria;
- participant responsibilities.

### 9.4 Public Boundary Repository

Purpose:

- public interfaces;
- schemas;
- integration boundaries;
- non-sensitive contracts;
- protected-runtime separation.

### 9.5 Protected Runtime Repository

Purpose:

- private implementation;
- protected authority logic;
- proprietary runtime behavior;
- confidential security mechanisms;
- non-public algorithms;
- private test materials.

A public repository must not be assumed to contain the complete protected runtime.

---

## 10. Authorship Verification

Authorship should be verified using multiple independent signals.

Recommended signals include:

- canonical repository ownership;
- Git commit history;
- commit author metadata;
- signed commits where available;
- signed release tags where available;
- release manifests;
- published SHA-256 checksums;
- machine-readable trust manifests;
- public project announcements;
- archived documentation;
- external timestamps;
- consistent official contact information;
- continuity of project history across releases.

No single signal should be treated as absolute proof.

For example:

- Git author metadata can be imitated;
- unsigned commits do not provide cryptographic identity proof;
- repository ownership can change;
- screenshots can be edited;
- copied documentation can appear authentic;
- unofficial archives can preserve old material without authorization.

The strongest available authorship record combines:

1. canonical repository control;
2. cryptographic signatures;
3. reproducible release manifests;
4. public timestamps;
5. consistent project history;
6. independent references.

---

## 11. Signed Commits and Tags

Where supported and operationally practical, official VRP releases should use:

- signed Git commits;
- signed Git tags;
- cryptographic checksums;
- release manifests;
- build provenance;
- artifact attestations.

A valid signature confirms that a specific key signed specific content.

It does not prove:

- that the signed code is free from defects;
- that the author’s host was uncompromised;
- that the build environment was trustworthy;
- that the artifact is suitable for production;
- that the signer personally reviewed every line.

Signatures establish provenance and integrity within their documented scope.

---

## 12. Release Authorship

An official VRP release should identify:

- project name;
- repository;
- release version;
- Git commit;
- release date;
- author or release authority;
- build toolchain;
- artifact list;
- SHA-256 checksums;
- signature information where available;
- validation status;
- known limitations.

Reviewers should reject or investigate releases that cannot establish these minimum identity properties.

---

## 13. Unofficial Forks and Mirrors

The VRP project may be forked, copied, archived, translated, or mirrored by third parties.

Unless explicitly approved, such copies are unofficial.

The author is not responsible for:

- modifications made by third parties;
- binaries produced by third parties;
- removed security controls;
- changed documentation;
- added dependencies;
- hidden payloads;
- altered release manifests;
- misleading claims;
- expired or abandoned mirrors;
- incompatible forks;
- malicious redistribution.

Third-party forks should identify themselves clearly and must not imply official endorsement without permission.

---

## 14. Unauthorized Representation

No individual or organization is authorized to claim that it:

- owns VRP;
- represents the author;
- controls official Pilot admission;
- distributes official private runtime access;
- has received exclusive rights;
- can issue official releases;
- can approve production deployment;
- can speak for the project;

unless such authority is documented through an official and verifiable project statement.

Suspicious representation should be reported through the official project contact.

---

## 15. Contribution Attribution

Future contributors should be credited according to the scope of their work.

Possible contribution categories include:

- code contribution;
- documentation;
- testing;
- independent validation;
- security research;
- architecture review;
- reproducibility work;
- Pilot engineering;
- translation;
- release engineering;
- vulnerability disclosure.

Contribution does not automatically grant:

- project ownership;
- authority over the protected runtime;
- release authority;
- Pilot approval authority;
- rights to represent the project;
- access to confidential materials.

Attribution should remain accurate, specific, and verifiable.

---

## 16. Use of Artificial Intelligence Tools

Artificial intelligence tools may be used during parts of the engineering workflow, including:

- code review assistance;
- documentation drafting;
- test-case generation;
- threat-model brainstorming;
- refactoring suggestions;
- command preparation;
- formatting;
- language translation;
- consistency analysis.

The use of AI tools does not transfer authorship, project ownership, or engineering responsibility to the tool provider.

The human maintainer remains responsible for:

- accepting or rejecting generated material;
- reviewing code before publication;
- running tests;
- validating behavior;
- correcting errors;
- ensuring claims remain accurate;
- protecting confidential material;
- approving releases.

AI-generated or AI-assisted material must be treated as untrusted until reviewed and validated.

---

## 17. Author Responsibility

The author is responsible for making reasonable efforts to:

- represent the project accurately;
- distinguish implemented features from planned features;
- distinguish test results from universal guarantees;
- disclose known limitations;
- avoid knowingly distributing malicious components;
- provide verification procedures;
- identify official artifacts;
- correct significant documentation errors;
- respond to credible security concerns;
- preserve protected information responsibly.

The author is not capable of guaranteeing absolute security or universal correctness.

---

## 18. Reviewer Responsibility

Reviewers remain responsible for:

- using an isolated environment;
- confirming repository identity;
- reviewing source code;
- examining dependencies;
- verifying checksums;
- monitoring runtime behavior;
- applying organizational policy;
- protecting credentials;
- controlling network access;
- documenting findings;
- deciding whether further evaluation is justified.

The presence of an authorship document does not remove the need for independent verification.

---

## 19. Pilot Participant Responsibility

Pilot participants must independently verify:

- the identity of the project contact;
- the scope of the Pilot;
- supplied artifacts;
- artifact hashes;
- expected privileges;
- network behavior;
- storage behavior;
- removal procedures;
- confidentiality requirements;
- legal authorization;
- internal approval.

A Pilot invitation must not be treated as permission to bypass the participant’s internal security procedures.

---

## 20. Identity Disputes

If project authorship, release authority, or repository identity is disputed, reviewers should preserve:

- repository URLs;
- commit hashes;
- tags;
- signatures;
- release manifests;
- email headers;
- public statements;
- timestamps;
- archived pages;
- artifact hashes;
- relevant correspondence.

Identity disputes should be resolved using verifiable records, not screenshots or unsupported claims alone.

---

## 21. Project Continuity

The long-term continuity of VRP should not depend solely on one person’s availability.

Future continuity controls may include:

- documented release procedures;
- protected signing keys;
- key-rotation procedures;
- recovery contacts;
- maintainership succession rules;
- archived technical specifications;
- immutable release records;
- emergency revocation procedures;
- trusted contributor roles;
- controlled repository recovery.

Any future transfer of maintainership or release authority must be publicly documented and cryptographically verifiable where possible.

---

## 22. Official Contact

Official project contact information must be published through canonical VRP repositories.

Reviewers should obtain the current address from the repository itself rather than relying on copied posts, screenshots, or third-party messages.

Sensitive reports should not be posted publicly before the maintainer has had a reasonable opportunity to investigate them.

The project does not guarantee compensation for reports unless a formal reward program is explicitly announced.

---

## 23. Non-Endorsement Statement

Reference to a company, technology, country, platform, protocol, cloud provider, hardware vendor, or operating system does not imply endorsement, partnership, certification, sponsorship, or commercial agreement.

The presence of public testing against a technology does not mean its owner has evaluated or approved VRP.

Any partnership or independent certification must be confirmed through a separate official statement from all relevant parties.

---

## 24. No Unsupported Credentials

The author does not rely on invented titles or unverifiable credentials to establish trust.

Project materials should not describe the author as:

- a government representative;
- an intelligence officer;
- a military engineer;
- a certified security authority;
- an employee of a named company;
- an officially recognized protocol standards body member;
- the creator of an “unbreakable” system;

unless such a statement can be independently verified and is relevant to the document.

The validity of VRP must rest on engineering evidence rather than status claims.

---

## 25. Authorship Statement

The following statement represents the current public authorship record:

> The Veil Routing Protocol was created by Vitalijus Riabovas, an independent systems developer from Lithuania, publicly operating under the engineering identity Endless33.
>
> The project has been developed through iterative implementation, adversarial testing, reproducible validation, security hardening, and public technical documentation.
>
> Project authorship does not replace independent review. Every technical claim must be evaluated according to its evidence, scope, environment, and limitations.

---

## 26. Final Declaration

The purpose of publishing authorship information is not to demand trust.

It is to establish accountability.

Reviewers should be able to determine:

- who created the project;
- who published a release;
- who controls the canonical repository;
- who is responsible for the project’s public claims;
- where concerns should be reported;
- whether an artifact came from an official source.

Project identity should be transparent.

Artifact integrity should be verifiable.

Technical claims should remain reproducible.

Trust must be earned through evidence.

---

**Vitalijus Riabovas**  
Original Author and Architect of the Veil Routing Protocol  
Public Engineering Identity: Endless33  
Lithuania