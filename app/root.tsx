import { Link, LiveReload, Outlet } from "@remix-run/react";

export default function App() {
  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <title>Drip</title>
      </head>
      <body>
        <Link to="/">Home</Link>
        <LiveReload />
        <Outlet />
      </body>
    </html>
  );
}
