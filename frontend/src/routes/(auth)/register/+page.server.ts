import { redirect } from "@sveltejs/kit";

export const load = async (event) => {
  const token = event.cookies.get("token");

  if (token) throw redirect(301, "/");
};

export const actions = {
  default: async (event) => {
    const formData = await event.request.formData();
    const name = formData.get("name");
    const email = formData.get("email");
    const privy_id = formData.get("privy_id");
    const body = await JSON.stringify({ name, email, privy_id });

    const res = await fetch("http://localhost:8000/api/users/register", {
      body,
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (res.ok) {
      throw redirect(301, "/login");
    }

    return {
      error: await res.text(),
    };
  },
};
