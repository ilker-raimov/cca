<script lang="ts">
    import { push, replace } from 'svelte-spa-router';

    let token = localStorage.getItem("token");

    if (!token) {
      replace("/login")
    }
</script>

<main>
<h2>Dashboard</h2>
<p>Welcome to the dashboard!</p>
<button on:click={() => {window.localStorage.removeItem("token"); push("/login");}}>Logout</button>
<button on:click={() => {push("/task");}}>Logout</button>
</main>
  