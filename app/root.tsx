import { Link, LiveReload, Outlet } from "@remix-run/react";
import { extractStyles } from "evergreen-ui";

export default function App() {
  const { css, hydrationScript } = extractStyles();
  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <title>Drip</title>
        <style id="evergreen-css" dangerouslySetInnerHTML={{ __html: css }} />
      </head>
      <body>
        {hydrationScript}
        <Link to="/">Home</Link>
        <LiveReload />
        <Outlet />
      </body>
    </html>
  );
}
