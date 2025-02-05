<script lang="ts">
    import Login from './Login.svelte';
    import { navigate } from 'svelte-routing';

    let token = localStorage.getItem("token");

    if (!token) {
      window.location.href = "/";
    }
</script>

<main>
<h2>Dashboard</h2>
<p>Welcome to the dashboard!</p>
<button on:click={() => {window.localStorage.removeItem("token"); navigate("/login");}}>Logout</button>
</main>
  