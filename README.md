## Blog Journal

<p> A simple blogging application built to explore web development, REST APIs, authentication, and database integration. It featured with signup, login, and share content, with work-in-progress frontend. </p>

## Tech Stack

<ul>
<li><strong>Language:</strong> Go</li>
<li><strong>Router:</strong> Chi</li>
<li><strong>Database:</strong> PostgreSQL</li>
<li><strong>ORM:</strong> GORM</li>
<li><strong>Authentication:</strong> JWT</li>
<li><strong>Password Hashing:</strong> bcrypt</li>
<li><strong>Configuration:</strong> Environment variables with godotenv</li>
</ul>

## Project Structure (BE)
<pre>
├── main.go
├── internal/
│   └── routes/      # Register API endpoints
│   ├── middleware/  # Authenticate and authorize requests
│   ├── handlers/    # Handle HTTP requests and responses
│   ├── services/    # Business logic and validation
│   ├── models/      # Handle database operations
│   ├── entities/    # Define application entities
├── .env
├── go.mod
└── go.sum
</pre>

## Application Startup

<p>When the application starts, <code>main.go</code>:</p>
- Loads environment variables.
- Connects to PostgreSQL using GORM.
- Runs database migrations using <code>AutoMigrate</code>.
- Initializes models, services, and handlers.
- Registers routes and middleware.
- Starts the HTTP server.

<p>Go does not load <code>.env</code> files automatically. A package such as <code>github.com/joho/godotenv</code> can load them at startup.

## Password Security

<p> During registration, passwords are hashed using bcrypt<code>(GenerateFromPassword)</code> before being stored. During login, bcrypt<code>(CompareHashAndPassword)</code> verifies the submitted password against the stored hash, and a JWT is issued upon successful authentication. </p>

## Authentication Middleware
<pre> Client → Authentication Middleware → JWT Verification → User Context  → Handler → Service → Database </pre>

<p> Middleware extracts & verifies the JWT from the <code>Authorization</code> header, validates its signature, extracts the user's claims, and adds the authenticated user information to the request context. Downstream handlers can use this information to identify the user and enforce access rules. </p>

## Go Concepts Used

<ul>
<li><strong>Structs:</strong> Define data structures and hold fields.</li>
<li><strong>Receiver methods:</strong> Associate behavior with a type, similar to methods in Python classes.</li>
<li><strong>Pointers:</strong> Pass references to values, such as the GORM database handle.</li>
<li><strong>Interfaces:</strong> Define behavior contracts and support interchangeable implementations.</li>
<li><strong>Constructors:</strong> Functions such as <code>NewUserRepository()</code> initialize and return repository instances.</li>
<li><strong>Dependency inversion:</strong> Allow higher-level business logic to depend on interfaces rather than concrete implementations.</li>
<li><strong>Context:</strong> Carry request-scoped data, such as authenticated user claims.</li>
<li><strong>Error handling:</strong> Return and handle errors explicitly using Go's <code>error</code> type.</li>
</ul>

## Getting Started
### Prerequisites

<ul>
<li>Go</li>
<li>PostgreSQL</li>
</ul>

### Setup


<ol>
<li>Clone the repository:</li>
</ol>


<pre><code>git clone &lt;repository-url&gt;
cd &lt;project-directory&gt;</code></pre>


<ol start="2">
<li>Create a PostgreSQL database.</li>
<li>Configure the environment variables.</li>
</ol>

ol start="4">
<li>Install dependencies:</li>
</ol>

<pre><code>go mod tidy</code></pre>


<ol start="5">
<li>Start PostgreSQL.</li>
<li>Run the backend:</li>
</ol>


<pre><code>go run main.go</code></pre>

<p>The server will start on the configured port.</p>

<p>Frontend setup instructions will be added as the frontend development progresses.</p>

## Notes

<ul>
<li>Keep <code>.env</code> out of version control.</li>
<li>Use a strong, private JWT secret.</li>
<li>Validate incoming requests.</li>
<li>Never store plain-text passwords.</li>
<li>Verify JWT signatures and expiration on protected routes.</li>
</ul>

## Project Status
<p> Work in progress. Frontend is not yet complete. </p>