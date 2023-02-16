import { ActionArgs, json, LoaderArgs, redirect } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { db } from "~/backend/db.server";

export const loader = async ({ params }: LoaderArgs) => {
  return json({
    messages: await db.message.findMany({
      where: {
        room: params.id,
      },
    }),
  });
};

export const action = async ({ request, params }: ActionArgs) => {
  const form = await request.formData();
  const content = form.get("content");
  // we do this type check to be extra sure and to make TypeScript happy
  // we'll explore validation next!
  if (typeof content !== "string") {
    throw new Error(`Form not submitted correctly.`);
  }

  console.log({ params });
  const fields = { content, room: params.id ?? "hi" };

  const message = await db.message.create({ data: fields });
  console.log({ message });
  return redirect(`/drip/${message.room}`);
};

export default function DripRoute() {
  const data = useLoaderData<typeof loader>();
  return (
    <div>
      wats good playa. here's yo drip:
      <ol>
        {data.messages.map((m) => (
          <li key={m.id}>{m.content}</li>
        ))}
      </ol>
      <form method="post">
        <div>
          <label>
            Content:
            <input type="text" name="content" />
          </label>
        </div>
        <button type="submit" className="button">
          Keep Drippin
        </button>
      </form>
    </div>
  );
}
