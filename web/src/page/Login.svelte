<script lang="ts">
    import { navigate } from 'svelte-routing';
    import { success, warning, error } from '../common/util.ts'

    let username: string = "";
    let password: string = "";

    async function handleLogin() {
        if (!username || !password) {
            error("Username and password are required!")

            return;
        }

        try {
            // let response = await fetch("http://localhost:8081/api/auth/login", {
            //     method: "POST",
            //     headers: { "Content-Type": "application/json" },
            //     body: JSON.stringify({ username, password })
            // });

            // let data = await response.json();

            let ok: bool = username !== "fail" && password !== "fail";
            let data: any = {
                token: "ahahha"
            }

            if (ok) {
                localStorage.setItem("token", data.token);

                success('Successful login!');
                navigate("/dashboard")
            } else {
                warning('Invalid credentials!');
            }
        } catch (err) {
            error('Error validating the credentials!');
        }
    }
</script>

<h2>Login</h2>

<input type="text" placeholder="Username" bind:value={username} />
<input type="password" placeholder="Password" bind:value={password} />

<button on:click={handleLogin}>Login</button>
  