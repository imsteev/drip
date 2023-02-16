import { ActionArgs, redirect } from "@remix-run/node";
import { db } from "~/backend/db.server";

export const action = async ({ request }: ActionArgs) => {
  const form = await request.formData();
  const content = form.get("content");
  // we do this type check to be extra sure and to make TypeScript happy
  // we'll explore validation next!
  if (typeof content !== "string") {
    throw new Error(`Form not submitted correctly.`);
  }

  const fields = { content, room: "hi" };

  const message = await db.message.create({ data: fields });
  console.log({ message });
  return redirect(`/drip/${message.room}`);
};

export default function DripRoute() {
  return (
    <div>
      <form method="post">
        <div>
          <label>
            Content:
            <input type="text" name="content" />
          </label>
        </div>
        <button type="submit" className="button">
          New Drip
        </button>
      </form>
    </div>
  );
}
