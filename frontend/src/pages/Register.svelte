<script lang="ts">
  import { register } from "../lib/api/auth";

  let username = "";
  let email = "";
  let password = "";

  let error = "";
  let loading = false;

  async function submit() {
    loading = true;
    error = "";

    try {
      const user = await register({
        username,
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

<h1>Register</h1>

<input bind:value={username} placeholder="Username" />

<input bind:value={email} placeholder="Email" />

<input type="password" bind:value={password} placeholder="Password" />

<button onclick={submit} disabled={loading}>
  {loading ? "Loading..." : "Register"}
</button>

{#if error}
  <p>{error}</p>
{/if}
