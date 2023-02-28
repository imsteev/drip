import { ActionArgs, redirect } from "@remix-run/node";
import { db } from "~/backend/db.server";

import randomWords from "random-words";
import { Form } from "@remix-run/react";

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

  return redirect(`/rooms/${message.room}`);
};

export default function IndexRoute() {
  return (
    <div
      style={{
        width: "100%",
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
      }}
    >
      <h1>Drip</h1>
      <div style={{ marginTop: "24px", width: "40%" }}>
        Share messages between your devices. Send a message and we'll create a
        room for you. Messages will only last for 10 minutes.
      </div>
      <div
        style={{ display: "flex", flexDirection: "column", marginTop: "32px" }}
      >
        <Form method="post">
          <input type="text" name="content" />
          <button style={{ marginLeft: "12px" }} type="submit">
            Send message to start a new room
          </button>
        </Form>
      </div>
    </div>
  );
}
