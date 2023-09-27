import { redirect } from "@sveltejs/kit";
import { tokenAuth, isLoggedIn } from "$lib/stores/login-store.js";

export const load = async (event) => {
  const token = event.cookies.get("token");

  if (token) throw redirect(301, "/");
};

export const actions = {
  default: async (event) => {
    const formData = await event.request.formData();
    const email = formData.get("email");
    const privy_id = formData.get("privy_id");
    const body = await JSON.stringify({ email, privy_id });

    const res = await fetch("http://localhost:8000/api/users/login", {
      body,
      method: "POST",
      headers: { "Content-Type": "application/json" },
    });

    if (res.ok) {
      let resBody = await res.json();
      event.cookies.set("token", resBody.data, {
        path: "/",
      });
      tokenAuth.set(resBody.data);
      isLoggedIn.set(true);

      throw redirect(301, "/");
    }

    return {
      error: await res.text(),
    };
  },
};
