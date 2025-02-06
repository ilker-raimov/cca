<script lang="ts">
    import { push } from 'svelte-spa-router';
    import { success, warning, error } from '../common/util.ts'
    import { Button, Form, FormGroup, Label, Input, Container, Card, CardBody, CardTitle } from 'sveltestrap';

    let email: string = "";
    let password: string = "";

    async function handleLogin() {
        if (!email || !password) {
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

            let ok: boolean = email !== "fail" && password !== "fail";
            let data: any = {
                token: "ahahha"
            }

            if (ok) {
                localStorage.setItem("token", data.token);

                success('Successful login!');
                push("/dashboard")
            } else {
                warning('Invalid credentials!');
            }
        } catch (err) {
            error('Error validating the credentials!');
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
            </Form>
        </CardBody>
    </Card>
</Container>
  