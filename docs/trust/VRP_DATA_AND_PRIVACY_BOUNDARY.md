# VRP Data and Privacy Boundary

**Project:** Veil Routing Protocol  
**Abbreviation:** VRP  
**Document:** Data and Privacy Boundary  
**Status:** Public  
**Version:** 1.0

---

# 1. Purpose

This document defines the intended data and privacy boundaries of the publicly available VRP engineering repositories.

Its objective is to explain:

- which categories of information are expected to be processed;
- which categories are intentionally outside the project scope;
- how reviewers should evaluate observable behavior;
- how privacy expectations relate to engineering evaluation.

This document describes engineering intent.

It does not replace independent verification.

---

# 2. Engineering Scope

The public repositories are intended to support:

- protocol evaluation;
- architecture review;
- engineering validation;
- evidence generation;
- Pilot preparation;
- reproducible testing.

They are not designed to operate as data collection platforms.

---

# 3. Data Categories

During normal evaluation, reviewers may intentionally create project artifacts such as:

- validation reports;
- evidence bundles;
- execution logs;
- test configuration files;
- deterministic outputs;
- reproducible validation artifacts.

These artifacts are expected components of the engineering workflow.

---

# 4. Information Not Intentionally Collected

The public repositories are not intended to intentionally collect unrelated personal information such as:

- personal documents;
- browser history;
- saved passwords;
- private photographs;
- personal communications;
- unrelated corporate files;
- financial information;
- unrelated credentials.

If unexpected access to such information is observed, the behavior should be investigated immediately.

---

# 5. Runtime Observation

Reviewers are encouraged to monitor:

- filesystem activity;
- process creation;
- network communication;
- generated files;
- temporary storage;
- execution logs.

Independent runtime observation is considered part of responsible engineering evaluation.

---

# 6. Network Privacy

Network activity should remain consistent with documented engineering scenarios.

Reviewers should independently verify:

- destination addresses;
- outbound connections;
- protocol usage;
- traffic volume;
- execution timing.

Undocumented network behavior should be treated as unexpected until explained.

---

# 7. Storage Expectations

Generated engineering artifacts should remain limited to files required for documented evaluation procedures.

Organizations should independently verify:

- generated directories;
- temporary files;
- report locations;
- evidence outputs;
- cleanup behavior.

---

# 8. Organizational Responsibility

Every organization evaluating VRP remains responsible for:

- applying internal security policies;
- protecting confidential information;
- selecting appropriate evaluation environments;
- determining production suitability;
- complying with applicable laws and regulations.

The project cannot determine organizational compliance requirements.

---

# 9. Privacy Principles

The public engineering repositories are developed according to several principles:

- minimum necessary processing;
- transparency;
- reproducibility;
- reviewer visibility;
- explicit documentation;
- observable behavior.

These principles are intended to support trustworthy engineering evaluation.

---

# 10. Third-Party Components

Third-party dependencies may process information according to their own documented behavior.

Organizations should independently evaluate:

- dependency origin;
- dependency licenses;
- dependency updates;
- dependency security;
- dependency privacy characteristics.

No third-party component should be trusted automatically.

---

# 11. Scope Limitations

This document applies only to the official public VRP engineering repositories.

It does not describe:

- private Pilot deployments;
- confidential runtime implementation;
- proprietary operational environments;
- third-party infrastructure.

Those environments may have additional policies.

---

# 12. Final Statement

The public VRP engineering repositories are intended to support transparent technical evaluation while minimizing unnecessary exposure of unrelated information.

Reviewers are encouraged to independently verify all observable behavior using their own tools, procedures, and organizational standards.

Privacy expectations should be validated through evidence rather than assumptions.

---

**Vitalijus Riabovas**  
Original Author and Architect of the Veil Routing Protocol  
Public Engineering Identity: Endless33  
Lithuania