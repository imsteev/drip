import { Form } from "@remix-run/react";
import { Button } from "evergreen-ui";

export default function PurgeRoom() {
  return (
    <Form method="post">
      <input type="hidden" name="action" value="purge-room" />
      <Button type="submit">Purge messages</Button>
    </Form>
  );
}
