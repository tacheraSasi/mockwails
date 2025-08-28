package mockserver

const custom404HTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Endpoint Not Found - MockWails</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
            color: #333;
            padding: 1rem;
        }
        
        .container {
            background: white;
            border-radius: 12px;
            padding: 3rem;
            box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
            text-align: center;
            max-width: 600px;
            width: 100%;
        }
        
        .error-code {
            font-size: 6rem;
            font-weight: bold;
            color: #667eea;
            margin-bottom: 1rem;
            line-height: 1;
        }
        
        .error-title {
            font-size: 1.5rem;
            margin-bottom: 1rem;
            color: #2d3748;
        }
        
        .error-description {
            color: #718096;
            margin-bottom: 2rem;
            line-height: 1.6;
        }
        
        .endpoint-info {
            background: #f7fafc;
            border-radius: 8px;
            padding: 1rem;
            margin: 1.5rem 0;
            border-left: 4px solid #e53e3e;
            text-align: left;
        }
        
        .endpoint-info strong {
            color: #2d3748;
        }
        
        .endpoint-info .method {
            background: #e53e3e;
            color: white;
            padding: 0.25rem 0.5rem;
            border-radius: 4px;
            font-size: 0.8rem;
            font-weight: bold;
        }
        
        .available-endpoints {
            background: #f0fff4;
            border-radius: 8px;
            padding: 1.5rem;
            margin: 1.5rem 0;
            border-left: 4px solid #38a169;
            text-align: left;
        }
        
        .available-endpoints h4 {
            color: #2f855a;
            margin-bottom: 1rem;
            text-align: center;
        }
        
        .endpoint-list {
            list-style: none;
            padding: 0;
        }
        
        .endpoint-item {
            background: white;
            border-radius: 6px;
            padding: 0.75rem;
            margin: 0.5rem 0;
            border: 1px solid #e2e8f0;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }
        
        .endpoint-item .method-badge {
            background: #38a169;
            color: white;
            padding: 0.25rem 0.5rem;
            border-radius: 4px;
            font-size: 0.75rem;
            font-weight: bold;
            min-width: 60px;
            text-align: center;
        }
        
        .endpoint-item .method-badge.POST { background: #3182ce; }
        .endpoint-item .method-badge.PUT { background: #d69e2e; }
        .endpoint-item .method-badge.DELETE { background: #e53e3e; }
        .endpoint-item .method-badge.PATCH { background: #805ad5; }
        
        .suggestion {
            background: #fef5e7;
            border-radius: 8px;
            padding: 1rem;
            margin: 1rem 0;
            border-left: 4px solid #f6ad55;
            text-align: left;
        }
        
        .suggestion h4 {
            color: #c05621;
            margin-bottom: 0.5rem;
        }
        
        .suggestion ul {
            color: #744210;
            margin-left: 1rem;
        }
        
        .brand {
            margin-top: 2rem;
            color: #a0aec0;
            font-size: 0.9rem;
        }
        
        .brand strong {
            color: #667eea;
        }
        
        .no-endpoints {
            color: #718096;
            font-style: italic;
            text-align: center;
            padding: 1rem;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="error-code">404</div>
        <h1 class="error-title">Endpoint Not Found</h1>
        <p class="error-description">
            The requested endpoint is not available on this mock server.
        </p>
        
        <div class="endpoint-info">
            <div style="margin-bottom: 0.5rem;">
                <strong>Requested:</strong> 
                <span class="method">{{.Method}}</span> 
                <code>{{.Endpoint}}</code>
            </div>
            <div><strong>Port:</strong> {{.Port}}</div>
        </div>
        
        {{if .AvailableEndpoints}}
        <div class="available-endpoints">
            <h4>🛡️ Available Endpoints on Port {{.Port}}</h4>
            <ul class="endpoint-list">
                {{range .AvailableEndpoints}}
                <li class="endpoint-item">
                    <div>
                        <span class="method-badge {{.Method}}">{{.Method}}</span>
                        <code style="margin-left: 0.5rem;">{{.Endpoint}}</code>
                    </div>
                    <small style="color: #718096;">{{.Name}}</small>
                </li>
                {{end}}
            </ul>
        </div>
        {{else}}
        <div class="available-endpoints">
            <h4>⚠️ No Active Endpoints</h4>
            <p class="no-endpoints">There are no active endpoints configured for port {{.Port}}</p>
        </div>
        {{end}}
        
        <div class="suggestion">
            <h4>💡 Suggestions:</h4>
            <ul>
                <li>Check if the endpoint is correctly configured in MockWails</li>
                <li>Verify the server is running and set to "active" status</li>
                <li>Ensure the HTTP method matches your request</li>
                <li>Check the MockWails dashboard for available endpoints</li>
            </ul>
        </div>
        
        <div class="brand">
            Powered by <strong>MockWails</strong>
        </div>
    </div>
</body>
</html>`
