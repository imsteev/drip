import { Link, LiveReload, Outlet } from "@remix-run/react";

export default function App() {
  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <title>Remix: So great, it's funny!</title>
      </head>
      <body>
        <Link to="/">Home</Link>

        <LiveReload />
        {/* this lays out pages in /routes */}
        <Outlet />
      </body>
    </html>
  );
}
