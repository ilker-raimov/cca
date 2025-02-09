<script lang="ts">
    import { push } from 'svelte-spa-router';
    import { success, warning, error, check_or_warning } from '../common/toast.ts'
    import { Button, Form, FormGroup, Label, Input, Container, Card, CardBody, CardTitle } from 'sveltestrap';

    let email: string = "";
    let password: string = "";

    async function handleLogin() {
        let check_email: boolean = check_or_warning(email, "Email is required!");
        let check_password: boolean = check_or_warning(password, "Password is required!");

        if (check_email || check_password) {
            return;
        }

        try {
            let response = await fetch("/api/auth/login", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ email, password })
            });

            if (!response.ok) {
                let data = await response.text();

                warning(data);

                return;
            }

            let data = await response.json();

            sessionStorage.setItem("username", data.username)
            sessionStorage.setItem("email", data.email)
            sessionStorage.setItem("role", data.role)
            sessionStorage.setItem("token", data.token);

            success('Successful login!');
            push("/dashboard")
        } catch (err: any) {
            error(err);
        }
    }
</script>

<Container class="d-flex justify-content-center align-items-center vh-100">
    <Card class="p-4 shadow-lg" style="max-width: 400px; width: 100%;">
        <CardBody>
            <CardTitle class="text-center mb-4">
                <h2>Login</h2>
            </CardTitle>

            <Form>
                <FormGroup>
                    <Label for="email">Email</Label>

                    <Input type="email" id="email" placeholder="Enter email" bind:value={email} />
                </FormGroup>

                <FormGroup>
                    <Label for="password">Password</Label>

                    <Input type="password" id="password" placeholder="Enter password" bind:value={password} />
                </FormGroup>

                <Button color="primary" block on:click={handleLogin}>Login</Button>
                <Button color="link" block on:click={() => push("/register")}>Don't have an account?</Button>
            </Form>
        </CardBody>
    </Card>
</Container>
  