import { ActionArgs, json, LoaderArgs, redirect } from "@remix-run/node";
import { Form, useLoaderData } from "@remix-run/react";
import {
  Button,
  ClipboardIcon,
  Heading,
  IconButton,
  Label,
  Pane,
  Paragraph,
  Text,
  TextInput,
} from "evergreen-ui";
import { db } from "~/backend/db.server";
import PurgeRoom from "~/routes/rooms/PurgeRoom";

export const loader = async ({ params }: LoaderArgs) => {
  return json({
    messages: await db.message.findMany({
      where: {
        room: params.roomID,
      },
    }),
    params,
  });
};

// TODO: auto-delete messages?
export const action = async ({ request, params }: ActionArgs) => {
  const form = await request.formData();
  const actionType = form.get("action");

  const roomID = params.roomID;

  switch (actionType) {
    case "new-message":
      const content = form.get("content");
      // we do this type check to be extra sure and to make TypeScript happy
      // we'll explore validation next!
      if (typeof content !== "string") {
        throw new Error(`Form not submitted correctly.`);
      }
      if (typeof roomID !== "string" || !roomID) {
        throw new Error("Invalid room type. Needs to be non-empty string");
      }
      await db.message.create({ data: { content, room: roomID } });
      break;
    case "delete-message":
      const messageID = form.get("messageID");
      if (typeof messageID !== "string" || !messageID) {
        return new Error("invalid id");
      }
      await db.message.delete({ where: { id: messageID } });
      break;
    case "purge-room":
      if (typeof roomID !== "string" || !roomID) {
        return new Error("invalid room ID");
      }
      await db.message.deleteMany({ where: { room: roomID } });
      break;
  }

  return redirect(`/rooms/${roomID}`);
};

export default function RoomRoute() {
  const data = useLoaderData<typeof loader>();
  return (
    <Pane
      display="flex"
      flexDirection="column"
      alignItems="center"
      width="100%"
    >
      <Pane display="flex" alignItems="flex-end" gap="24px">
        <Heading size={900}>Room: {data.params.roomID}</Heading>
        <PurgeRoom />
      </Pane>
      <Pane display="flex" justifyContent="center" padding="64px">
        <Form method="post">
          <input type="hidden" name="action" value="new-message" />
          <TextInput type="text" name="content" />
          <Button marginLeft="12px" type="submit">
            New message
          </Button>
        </Form>
      </Pane>
      <Pane width="50%">
        {data.messages.map((m) => (
          <Pane
            key={m.id}
            width="100%"
            borderRadius="25px"
            border="solid 1px #222222"
            padding="24px"
            marginY="24px"
          >
            <Form method="delete">
              <input type="hidden" name="action" value="delete-message" />
              <input type="hidden" name="messageID" value={m.id} />
              <Pane
                display="flex"
                alignItems="center"
                justifyContent="space-between"
              >
                <Text>{m.content}</Text>
                <Button type="submit" className="button">
                  delete
                </Button>
              </Pane>
            </Form>
          </Pane>
        ))}
      </Pane>
    </Pane>
  );
}
