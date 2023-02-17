import { ActionArgs, json, LoaderArgs, redirect } from "@remix-run/node";
import { Form, useLoaderData } from "@remix-run/react";
import {
  Button,
  Card,
  Heading,
  Label,
  ListItem,
  Pane,
  Text,
  TextInput,
} from "evergreen-ui";
import { db } from "~/backend/db.server";

export const loader = async ({ params }: LoaderArgs) => {
  return json({
    messages: await db.message.findMany({
      where: {
        room: params.id,
      },
    }),
    params,
  });
};

// TODO: auto-delete messages?
export const action = async ({ request, params }: ActionArgs) => {
  const form = await request.formData();
  const actionType = form.get("action");

  switch (actionType) {
    case "new-message":
      const content = form.get("content");
      // we do this type check to be extra sure and to make TypeScript happy
      // we'll explore validation next!
      if (typeof content !== "string") {
        throw new Error(`Form not submitted correctly.`);
      }

      if (typeof params.id !== "string" || !params.id) {
        throw new Error("Invalid room type. Needs to be non-empty string");
      }

      const fields = { content, room: params.id };

      const message = await db.message.create({ data: fields });

      return redirect(`/${message.room}`);
    case "delete-message":
      const messageID = form.get("messageID");
      if (typeof messageID !== "string" || !messageID) {
        return new Error("invalid id");
      }
      await db.message.delete({ where: { id: messageID } });
      return redirect(`/${params.id}`);

    default:
      return null;
  }
};

export default function DripRoute() {
  const data = useLoaderData<typeof loader>();
  return (
    <Pane>
      <Heading size="800">Room: {data.params.id}</Heading>
      <Pane display="flex" justifyContent="center" padding="64px">
        <Form method="post">
          <input type="hidden" name="action" value="new-message" />
          <Label marginRight="12px" htmlFor="content">
            Message
          </Label>
          <TextInput type="text" name="content" />
          <Button marginLeft="12px" type="submit">
            New message
          </Button>
        </Form>
      </Pane>
      <Pane display="flex" justifyContent="flex-start">
        <Pane margin="12px">
          {data.messages.map((m) => (
            <Pane key={m.id}>
              <Form method="delete">
                <input type="hidden" name="action" value="delete-message" />
                <input type="hidden" name="messageID" value={m.id} />
                <Pane display="flex" alignItems="center">
                  <Text>{m.content}</Text>
                  <Button marginLeft="12px" type="submit" className="button">
                    delete
                  </Button>
                </Pane>
              </Form>
            </Pane>
          ))}
        </Pane>
      </Pane>
    </Pane>
  );
}
