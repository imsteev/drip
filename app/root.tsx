import { Link, LiveReload, Outlet, Scripts } from "@remix-run/react";
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
      </head>
      <body>
        <Link to="/">Home</Link>
        <LiveReload />
        <Outlet />
      </body>
    </html>
  );
}
