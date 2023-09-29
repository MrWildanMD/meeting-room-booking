<script lang="ts">
    import Layout from "stwui/layout";
  </script>
  
  <Layout class="min-h-screen">
    <Layout.Header class="fixed z-10 bg-red-600 bg-opacity-90 py-4 px-4">
      <h2 class="text-white">Room Reservation</h2>
      <Layout.Header.Extra>
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
  