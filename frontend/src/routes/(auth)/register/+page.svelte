<script lang="ts">
  import { Button, Card } from "stwui";
  import type { ActionData } from "./$types";
  export let form: ActionData;
</script>

<main>
  <!-- <div class="w-full max-w-sm p-4 bg-white rounded-lg shadow-md">
      <h2 class="text-2xl font-bold text-center text-red-600">Login</h2>
      <form>
        <input type="email" name="email" placeholder="Email" class="w-full px-3 py-2 mb-2 border rounded-lg focus:outline-none border-red-400">
        <input type="password" name="password" placeholder="Password" class="w-full px-3 py-2 mb-2 border rounded-lg focus:outline-none border-red-400">
        <Button on:click={(e) => e.preventDefault()} htmlType="submit" type="danger" class="w-full text-white font-bold py-2 rounded-lg">Login</Button>
        <p class="text-center text-sm mt-4">Don't have an account? <a href="/register" class="text-red-600">Register here.</a></p>
      </form>
    </div> -->
  <Card bordered elevation="xl" class="rounded-xl p-4">
    <Card.Content class="flex flex-col gap-4">
      <form method="POST">
        <h2 class="text-2xl font-bold text-center text-red-600">Register</h2>
        {#if form}
          <h4 class="text-red-600 font-light text-md text-center">
            {form.error}
          </h4>
        {/if}
        <div class="flex flex-col gap-4">
          <input
            type="text"
            name="name"
            placeholder="Name"
            class="w-full px-3 py-2 mb-2 border rounded-lg focus:outline-none border-red-400"
          />
          <input
            type="email"
            name="email"
            placeholder="Email"
            class="w-full px-3 py-2 mb-2 border rounded-lg focus:outline-none border-red-400"
          />
          <input
            type="text"
            name="privy_id"
            placeholder="PrivyID"
            class="w-full px-3 py-2 mb-2 border rounded-lg focus:outline-none border-red-400"
          />
          <!-- on:click={(e) => e.preventDefault()} -->
          <Button
            htmlType="submit"
            type="danger"
            class="w-full text-white font-bold py-2 rounded-lg">Register</Button
          >
          <p class="text-center text-sm mt-4">
          Already have an account? <a href="/login" class="text-red-600"
            >Login now.</a
          >
        </p>
        </div>
      </form>
    </Card.Content>
  </Card>
</main>
