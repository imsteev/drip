import { ActionArgs, redirect } from "@remix-run/node";
import { db } from "~/backend/db.server";

import randomWords from "random-words";
import { Form } from "@remix-run/react";
import { Button, Label, Pane, TextInput } from "evergreen-ui";

export const action = async ({ request }: ActionArgs) => {
  const form = await request.formData();

  const content = form.get("content");
  // we do this type check to be extra sure and to make TypeScript happy
  // we'll explore validation next!
  if (typeof content !== "string") {
    throw new Error(`Form not submitted correctly.`);
  }

  const message = await db.message.create({
    data: { content, room: randomWords(2).join("-") },
  });

  return redirect(`/drip/${message.room}`);
};

export default function IndexRoute() {
  return (
    <Pane width="100%">
      <h1>Drip</h1>
      <Pane display="flex" justifyContent="center">
        <Form method="post">
          <Label marginRight="12px" htmlFor="content">
            Message
          </Label>
          <TextInput type="text" name="content" />
          <Button marginLeft="12px" type="submit">
            Send message to start a new room
          </Button>
        </Form>
      </Pane>
    </Pane>
  );
}
