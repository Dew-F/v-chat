<script lang="ts">
  import { login } from "../lib/api/auth";

  let email = "";
  let password = "";

  let error = "";
  let loading = false;

  async function submit() {
    loading = true;
    error = "";

    try {
      const user = await login({
        email,
        password,
      });

      console.log(user);

      window.location.hash = "/app";
    } catch (e) {
      error = String(e);
    } finally {
      loading = false;
    }
  }
</script>

<h1>Login</h1>

<input bind:value={email} placeholder="email" />

<input type="password" bind:value={password} placeholder="password" />

<button onclick={submit} disabled={loading}>
  {loading ? "Loading..." : "Login"}
</button>

{#if error}
  <p>{error}</p>
{/if}
