<script lang="ts">
  import Layout from "stwui/layout";
  import { isLoggedIn } from "$lib/stores/login-store";
  import { typeUser } from "$lib/stores/user-store";
</script>

<Layout class="min-h-screen">
  <Layout.Header class="fixed z-10 bg-red-600 bg-opacity-90 py-4 px-4">
    <h2 class="text-white">Room Reservation</h2>
    <Layout.Header.Extra>
      {#if !isLoggedIn}
        <a class="text-white font-bold text-md" href="/login">Login</a>
      {/if}
      {#if $typeUser === 1}
        <a class="text-white font-bold text-md mr-2 md:mr-4" href="/dashboard">Dashboard</a>
      {/if}
      <a class="text-white font-bold text-md" href="/logout">Logout</a>
    </Layout.Header.Extra>
  </Layout.Header>
  <Layout.Content>
    <Layout.Content.Body class="min-h-screen">
      <slot />
    </Layout.Content.Body>
  </Layout.Content>
  <Layout.Footer class="static z-0 px-4"
    >&copy; 2023 Room Reservation</Layout.Footer
  >
</Layout>
