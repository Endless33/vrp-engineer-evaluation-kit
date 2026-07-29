# VRP Independent Verification Guide

**Project:** Veil Routing Protocol  
**Abbreviation:** VRP  
**Document:** Independent Verification Guide  
**Status:** Public  
**Version:** 1.0

---

# 1. Purpose

This guide explains how independent reviewers should approach the evaluation of the public VRP engineering repositories.

The objective is not to persuade reviewers that the project is trustworthy.

The objective is to provide a structured process for determining whether the published engineering claims are supported by reproducible technical evidence.

Independent verification is considered an essential part of the VRP engineering philosophy.

---

# 2. Verification Principles

Every reviewer is encouraged to rely on:

- direct observation;
- reproducible experiments;
- source-code inspection;
- deterministic validation;
- engineering evidence;
- documented assumptions.

Engineering conclusions should not depend solely on documentation or project statements.

---

# 3. Recommended Evaluation Environment

Whenever practical, evaluation should be performed using:

- isolated virtual machines;
- disposable test environments;
- dedicated engineering workstations;
- restricted user privileges;
- monitored network environments;
- reproducible build environments.

Production systems should not be used for initial evaluation.

---

# 4. Repository Verification

Before executing any component, reviewers should verify:

- repository ownership;
- repository URL;
- Git history;
- release history;
- commit consistency;
- published documentation;
- available integrity information.

Repository authenticity should always be established before technical evaluation begins.

---

# 5. Source Review

Reviewers are encouraged to inspect:

- project structure;
- implementation consistency;
- dependency usage;
- configuration files;
- exported interfaces;
- validation procedures;
- documentation alignment.

Questions identified during review should be documented for later verification.

---

# 6. Build Verification

Whenever possible, reviewers should:

- build from source;
- compare generated artifacts;
- inspect build logs;
- review dependency resolution;
- confirm successful compilation.

Successful compilation alone should not be interpreted as architectural validation.

---

# 7. Runtime Observation

Independent runtime observation is recommended.

Examples include:

- process monitoring;
- filesystem observation;
- network monitoring;
- memory observation;
- sandbox execution;
- integrity monitoring.

Observed behavior should remain consistent with published documentation.

---

# 8. Validation Procedures

Public validation procedures should be executed exactly as documented.

Reviewers are encouraged to verify:

- reproducibility;
- deterministic outputs where applicable;
- evidence generation;
- expected failure handling;
- documented limitations.

Unexpected behavior should always be investigated.

---

# 9. Evidence Review

Generated evidence should be examined for:

- internal consistency;
- reproducibility;
- integrity;
- documented structure;
- expected metadata;
- verification compatibility.

Reviewers may intentionally modify evidence to confirm that documented verification mechanisms detect tampering where applicable.

---

# 10. Documentation Review

Engineering documentation should be evaluated together with implementation.

Reviewers should compare:

- documented behavior;
- observable behavior;
- implementation;
- generated evidence;
- published limitations.

Documentation should accurately describe the public engineering scope.

---

# 11. Recording Findings

Reviewers are encouraged to maintain records including:

- tested version;
- repository commit;
- execution environment;
- executed procedures;
- observed results;
- identified limitations;
- unanswered questions;
- reproducibility notes.

Well-documented findings improve future engineering discussions.

---

# 12. Reporting

Constructive technical reports are encouraged.

Useful reports generally include:

- reproduction steps;
- expected behavior;
- observed behavior;
- supporting evidence;
- environment information;
- possible contributing factors.

Engineering discussions should remain evidence-based.

---

# 13. Limitations

Independent verification cannot prove the absence of all defects.

Instead, it increases confidence through:

- repeatability;
- transparency;
- observation;
- engineering discipline;
- technical scrutiny.

Confidence should remain proportional to available evidence.

---

# 14. Final Statement

The VRP project encourages every reviewer to independently verify published engineering claims.

Verification should be based on implementation, observable behavior, reproducible procedures, and documented evidence.

Engineering trust should be earned through independent technical evaluation rather than assumption.

---

**Vitalijus Riabovas**  
Original Author and Architect of the Veil Routing Protocol  
Public Engineering Identity: Endless33  
Lithuania