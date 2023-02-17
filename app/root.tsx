import { Link, LiveReload, Outlet } from "@remix-run/react";
import { extractStyles } from "evergreen-ui";

export default function App() {
  const { css, hydrationScript } = extractStyles();
  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <title>Drip</title>

        {/* https://github.com/segmentio/evergreen/issues/154#issuecomment-375766412 */}
        <style id="evergreen-css" dangerouslySetInnerHTML={{ __html: css }} />
        {hydrationScript}

        <Link to="/">Home</Link>
      </head>
      <body>
        <LiveReload />
        <Outlet />
      </body>
    </html>
  );
}
