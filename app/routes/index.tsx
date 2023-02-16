import { Link } from "@remix-run/react";

export default function IndexRoute() {
  return (
    <div>
      <h1>Need to drip?</h1>
      <Link to="/drip">Ready2Drip</Link>
    </div>
  );
}
