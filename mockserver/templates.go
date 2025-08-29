package mockserver

const custom404HTML = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>404 - Endpoint Not Found</title>
  <style>
    * {
      margin: 0;
      padding: 0;
      box-sizing: border-box;
    }
    body {
      font-family: ui-monospace, monospace, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
      background: #fafafa;
      color: #111;
      padding: 2rem;
      line-height: 1.5;
    }
    .container {
      max-width: 800px;
      margin: 0 auto;
    }
    h1 {
      font-size: 2rem;
      font-weight: 600;
      margin-bottom: 0.5rem;
    }
    .code {
      font-size: 4rem;
      font-weight: bold;
      color: #d32f2f;
      margin-bottom: 1rem;
    }
    .section {
      margin: 1.5rem 0;
      padding: 1rem;
      border: 1px solid #ddd;
      border-radius: 4px;
      background: #fff;
    }
    code {
      background: #f5f5f5;
      padding: 0.2rem 0.4rem;
      border-radius: 3px;
      font-size: 0.9rem;
    }
    .endpoint {
      display: flex;
      justify-content: space-between;
      padding: 0.4rem 0;
      border-bottom: 1px solid #eee;
    }
    .endpoint:last-child {
      border-bottom: none;
    }
    .badge {
      font-size: 0.75rem;
      padding: 0.2rem 0.5rem;
      border-radius: 3px;
      font-weight: 600;
      text-transform: uppercase;
    }
    .GET { background: #e0f2f1; color: #00695c; }
    .POST { background: #e3f2fd; color: #1565c0; }
    .PUT { background: #fff8e1; color: #ef6c00; }
    .DELETE { background: #ffebee; color: #c62828; }
    .PATCH { background: #f3e5f5; color: #6a1b9a; }
    ul { margin-left: 1.2rem; }
    .muted { color: #555; font-size: 0.9rem; }
  </style>
</head>
<body>
  <div class="container">
    <div class="code">404</div>
    <h1>Endpoint Not Found</h1>
    <p>The requested endpoint could not be resolved on this mock server.</p>

    <div class="section">
      <strong>Requested:</strong> 
      <span class="badge {{.Method}}">{{.Method}}</span> 
      <code>{{.Endpoint}}</code><br>
      <strong>Port:</strong> {{.Port}}
    </div>

    {{if .AvailableEndpoints}}
    <div class="section">
      <strong>Available Endpoints (Port {{.Port}}):</strong>
      {{range .AvailableEndpoints}}
      <div class="endpoint">
        <div>
          <span class="badge {{.Method}}">{{.Method}}</span>
          <code><a href="http://localhost:{{$.Port}}{{.Endpoint}}" target="_blank">{{.Endpoint}}</a></code>
        </div>
        <div class="muted">{{.Name}}</div>
      </div>
      {{end}}
    </div>
    {{else}}
    <div class="section">
      <strong>No active endpoints</strong> on port {{.Port}}.
    </div>
    {{end}}

    <div class="section">
      <strong>Suggestions:</strong>
      <ul>
        <li>Check endpoint configuration in MockWails.</li>
        <li>Verify the server is running and active.</li>
        <li>Confirm the HTTP method matches your request.</li>
      </ul>
    </div>

    <p class="muted">MockWails Created by <a href="https://github.com/tacheraSasi" target="_blank">tacheraSasi</a></p>
  </div>
</body>
</html>`
