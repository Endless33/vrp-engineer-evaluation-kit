# VRP Claims and Evidence Matrix

**Project:** Veil Routing Protocol  
**Abbreviation:** VRP  
**Document:** Claims and Evidence Matrix  
**Status:** Public  
**Version:** 1.0

---

# 1. Purpose

This document defines how engineering claims made by the VRP project should be evaluated.

Its objective is to distinguish between:

- documented engineering goals;
- implemented functionality;
- reproducible observations;
- engineering assumptions;
- future work.

Every technical claim should be traceable to supporting engineering evidence.

---

# 2. Engineering Principle

The VRP project follows a simple rule:

> Claims should never exceed available evidence.

Documentation should accurately represent the implementation that is publicly available.

Where evidence is incomplete, the limitation should be stated explicitly.

---

# 3. Claim Categories

Engineering statements generally fall into one of the following categories.

### A. Design Objective

Describes an intended architectural goal.

Examples:

- continuity-first networking;
- deterministic behavior;
- explicit trust boundaries.

A design objective is not proof of implementation.

---

### B. Implemented Feature

Describes behavior implemented within the published engineering scope.

Implementation should be reviewable through:

- source code;
- validation procedures;
- documentation;
- observable behavior.

---

### C. Validation Result

Describes behavior that has been observed during testing.

Validation should include:

- reproducible procedure;
- execution environment;
- expected outcome;
- observed outcome;
- generated evidence.

---

### D. Engineering Assumption

Documents an assumption required for evaluation.

Examples include:

- supported operating system;
- supported toolchain;
- reviewer-controlled execution;
- expected configuration.

Assumptions should remain visible.

---

### E. Future Work

Identifies engineering work that has not yet been completed.

Future plans should never be presented as implemented functionality.

---

# 4. Evidence Sources

Engineering evidence may include:

- source code;
- repository history;
- validation reports;
- execution logs;
- evidence bundles;
- release documentation;
- reproducible builds;
- integrity verification;
- independent review.

Multiple independent sources provide stronger confidence than any individual artifact.

---

# 5. Reviewer Guidance

Reviewers are encouraged to ask the following questions for every significant claim.

### Is the claim implemented?

If yes, identify the implementation.

---

### Can it be reproduced?

If yes, reproduce the result independently.

---

### Is the behavior observable?

If yes, compare implementation with documentation.

---

### Is evidence available?

Review the supporting evidence.

---

### Are limitations documented?

Engineering limitations should always remain visible.

---

# 6. Unsupported Claims

Reviewers should treat unsupported claims cautiously.

Examples include claims that cannot be connected to:

- implementation;
- documentation;
- reproducible validation;
- observable behavior;
- published evidence.

Unsupported claims require additional investigation before being accepted.

---

# 7. Independent Verification

Engineering confidence should increase when multiple reviewers independently reproduce the same result.

Independent verification remains one of the strongest forms of engineering evidence.

---

# 8. Continuous Revision

As the project evolves:

- implementations improve;
- documentation expands;
- validation grows;
- evidence becomes stronger.

Engineering claims should be updated whenever new evidence becomes available.

---

# 9. Scope Limitations

This matrix applies only to the official public engineering scope.

It does not describe:

- confidential implementation;
- protected runtime mechanisms;
- proprietary operational procedures;
- non-public Pilot materials.

---

# 10. Final Statement

The purpose of this document is to ensure that every published engineering claim can be evaluated using objective technical evidence.

Engineering credibility should grow from:

- implementation;
- reproducibility;
- transparency;
- independent verification.

Evidence should always come before conclusions.

---

**Vitalijus Riabovas**  
Original Author and Architect of the Veil Routing Protocol  
Public Engineering Identity: Endless33  
Lithuania