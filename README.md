# Perry (Periskopein)

Peri - skopein: Ancient Greek for to look over or guard

Perry is a deterministic engine for monitoring for alerts and sending out notifications

## Quickstart

```
docker compose up -d --build
```

View Workflow Status UI: 127.0.0.1:8080

## Overview

Deterministic, horizontally scalable workflow engine for checking alerts and taking action on them.

After a request to scan for alerts and process them is requested, the deterministic workflow engine takes over and follows it through to completion with explicit retry and crash logic.

The system is robust to system failures at any point in the workflow process.

### Core Components
- App (Request) Server. Initiates alert scanning on a regular interval.

- Workflow Server(s) - Processes Workflows and acitivities. Handles all logic related to checking for alerts and sending out notifications. Able to be scaled horizontally as workload increases