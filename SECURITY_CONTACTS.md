# VRP Security Contacts and Disclosure Policy

**Project:** Veil Routing Protocol  
**Abbreviation:** VRP  
**Document:** Security Contacts and Coordinated Disclosure Policy  
**Status:** Public  
**Version:** 1.0

---

# 1. Purpose

This document defines the official security communication process for the Veil Routing Protocol (VRP) project.

Its objectives are to:

- provide a trusted communication channel;
- define responsible disclosure expectations;
- establish reporting requirements;
- reduce the risk of misinformation;
- improve response quality;
- protect both researchers and project participants;
- encourage evidence-based communication.

This document applies to all official public VRP repositories.

---

# 2. Official Security Contact

Security reports should be submitted only through the official project security contact published in the canonical VRP repositories.

Reviewers should always verify that the contact information originates from an official repository before sending confidential information.

Do not rely on screenshots, reposts, social-media messages, or unofficial mirrors.

---

# 3. What Should Be Reported

Examples include:

- unexpected network communication;
- unexpected filesystem modifications;
- privilege escalation;
- integrity failures;
- evidence verification failures;
- replay protection failures;
- deterministic validation inconsistencies;
- undocumented behavior;
- dependency security concerns;
- reproducibility failures;
- build integrity concerns;
- documentation inaccuracies affecting security;
- supply-chain observations.

Reports do not need to prove exploitation.

Reasonable evidence of unexpected behavior is sufficient.

---

# 4. Report Quality

A high-quality report should include:

- affected repository;
- repository URL;
- commit hash;
- operating system;
- architecture;
- Go version;
- executed command;
- expected behavior;
- observed behavior;
- exact reproduction steps;
- logs;
- evidence files;
- screenshots where appropriate;
- packet captures if relevant;
- hashes of modified artifacts.

Incomplete reports may require additional clarification before investigation can begin.

---

# 5. Coordinated Disclosure

The preferred process is:

1. identify the issue;
2. reproduce the issue;
3. collect technical evidence;
4. privately report the issue;
5. allow reasonable investigation time;
6. coordinate public disclosure when appropriate.

The objective is responsible engineering rather than publicity.

---

# 6. Confidential Information

Please avoid including:

- unrelated customer information;
- personal credentials;
- private cryptographic keys;
- confidential corporate documents;
- unrelated proprietary source code.

Only include information necessary to reproduce the reported issue.

---

# 7. Good Faith

The project welcomes responsible security research.

Good-faith researchers acting to improve security are appreciated.

Reports should remain factual, reproducible, and technically supported.

Speculation without technical evidence is significantly less useful than reproducible observations.

---

# 8. Investigation Process

Each report may undergo:

- initial review;
- reproduction attempts;
- engineering analysis;
- scope determination;
- severity assessment;
- documentation review;
- correction planning;
- validation of the proposed fix;
- publication where appropriate.

Not every report will necessarily result in a code change.

---

# 9. Possible Outcomes

Following investigation, a report may be classified as:

- confirmed vulnerability;
- implementation defect;
- documentation issue;
- expected behavior;
- duplicate report;
- unsupported configuration;
- out-of-scope observation;
- insufficient evidence.

Every classification should be supported by technical reasoning whenever possible.

---

# 10. Response Expectations

While every reasonable effort will be made to review reports, response times cannot be guaranteed.

Complex architectural investigations may require additional analysis before conclusions can be reached.

Multiple reports concerning the same issue may be consolidated.

---

# 11. Public Discussion

Researchers are encouraged to avoid publishing full technical details before the project has had a reasonable opportunity to investigate.

Responsible disclosure benefits both reviewers and future users.

---

# 12. Legal Notice

Submitting a report does not create:

- employment;
- partnership;
- consulting agreements;
- ownership rights;
- licensing rights;
- financial obligations.

Unless explicitly announced, the project does not currently operate a bug bounty or financial reward program.

---

# 13. Final Statement

The VRP project values independent review, technical criticism, and reproducible security research.

Strong engineering grows through verification, transparency, and continuous improvement.

Security should never depend on assumptions.

Security should be demonstrated through evidence.

---

**Vitalijus Riabovas**  
Original Author and Architect of the Veil Routing Protocol  
Public Engineering Identity: Endless33  
Lithuania