import { Form } from "@remix-run/react";

export default function PurgeRoom() {
  return (
    <Form method="post">
      <input type="hidden" name="action" value="purge-room" />
      <button type="submit">Purge messages</button>
    </Form>
  );
}
