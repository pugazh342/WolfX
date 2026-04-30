## 1. SYSTEM ARCHITECTURE (High-Level)
```mermaid
flowchart TD
    GUI["GUI (Wails)\nDashboard / Proxy Viewer"]
    CLI["CLI (Cobra)\nCI/CD Integration"]
    API["API Layer\nREST + WebSocket"]

    GUI --> API
    CLI --> API

    API --> Proxy["Proxy Layer (MITM)"]
    API --> Engine["Scan Engine"]
    API --> Manager["Scan Manager"]

    Proxy --> Target["Target Web Applications"]
    Engine --> Target

    Engine --> Detection["Detection Engine"]
    Detection --> Plugins["Plugin System"]
    Plugins --> Payload["Payload Engine"]
```
## 2. DETAILED ARCHITECTURE (Internal Components)
```mermaid
flowchart TD
    Core["Core Engine"]

    HTTP["HTTP Engine\n(HTTP/1.1, HTTP/2, WebSocket)"]
    Session["Session Manager\nCookies / Tokens"]
    Auth["Auth Engine\nJWT / OAuth / SSO"]

    Crawler["Crawler Engine\nJS + API Discovery"]
    Proxy["Proxy Engine\nMITM Traffic"]

    Orchestrator["Scan Orchestrator\nTask Scheduling"]

    Plugins["Plugin Execution Engine\n(XSS, SQLi, SSRF, IDOR)"]
    Detection["Detection Intelligence\nDiffing / Timing / Heuristics"]

    Reporting["Reporting Engine\nCVSS / Evidence"]
    Storage["Storage Layer\nSQLite / PostgreSQL"]

    Core --> HTTP
    Core --> Session
    Core --> Auth

    HTTP --> Crawler
    HTTP --> Proxy

    Crawler --> Orchestrator
    Proxy --> Orchestrator

    Orchestrator --> Plugins
    Plugins --> Detection
    Detection --> Reporting
    Reporting --> Storage
```

## 3. WORKFLOW DIAGRAM (End-to-End)
```mermaid
flowchart TD
    User["User / Developer / Pentester"]

    Mode["Select Mode\nProxy / Active / Hybrid"]
    Traffic["Traffic / Crawling\nBrowser / API"]
    Mapping["Endpoint Mapping"]

    Injection["Payload Injection\n(Plugins)"]
    Detection["Detection Engine\nDiff / Timing / Heuristics"]

    Vulnerability["Vulnerability Identified\n+ Confidence Score"]
    Report["Report Generation\nCVSS + Evidence"]

    Output["Output\nCLI / GUI / CI / Slack"]

    User --> Mode
    Mode --> Traffic
    Traffic --> Mapping
    Mapping --> Injection
    Injection --> Detection
    Detection --> Vulnerability
    Vulnerability --> Report
    Report --> Output
```
## 4. REAL-TIME PROXY FLOW (Your Key Differentiator)
```mermaid
sequenceDiagram
    participant User
    participant Browser
    participant Proxy
    participant Engine
    participant Detection

    User->>Browser: Access Web App
    Browser->>Proxy: HTTP Request
    Proxy->>Engine: Forward + Log Request
    Engine->>Detection: Passive Analysis

    Detection-->>Engine: Findings (if any)

    Engine->>Proxy: Inject Payload (Active Mode)
    Proxy->>Browser: Modified Request
    Browser->>Proxy: Response

    Proxy->>Engine: Response Data
    Engine->>Detection: Analyze Response
    Detection-->>Engine: Vulnerability Result

```
## 5. PLUGIN EXECUTION FLOW
```mermaid
flowchart TD
    Request["Captured Request"]

    PluginSelect["Select Relevant Plugins"]
    PayloadGen["Generate Payloads"]

    Injection["Inject Payload into Request"]
    Response["Receive Response"]

    Analysis["Analyze Response\nDiff / Timing / Context"]

    Result["Generate Finding\n+ Confidence Score"]

    Request --> PluginSelect
    PluginSelect --> PayloadGen
    PayloadGen --> Injection
    Injection --> Response
    Response --> Analysis
    Analysis --> Result
```
## 6. DETECTION ENGINE LOGIC FLOW
```mermaid
flowchart TD
    Input["Request + Response Data"]

    Diff["Response Diffing"]
    Timing["Timing Analysis"]
    Heuristic["Heuristic Checks"]
    Context["Context Validation"]

    Score["Confidence Scoring Engine"]

    Output["Final Vulnerability Decision"]

    Input --> Diff
    Input --> Timing
    Input --> Heuristic
    Input --> Context

    Diff --> Score
    Timing --> Score
    Heuristic --> Score
    Context --> Score

    Score --> Output
```
