<script lang="ts">
    import { push } from 'svelte-spa-router';
    import { success, warning, error, check_or_warning } from '../common/toast.ts'
    import { Button, Form, FormGroup, Label, Input, Container, Card, CardBody, CardTitle, Row } from 'sveltestrap';

    let email: string = "";
    let username: string = "";
    let password: string = "";
    let password_repeat: string = "";

    async function handleRegister() {
        let check_email: boolean = check_or_warning(email, "Email is required!");
        let check_username: boolean = check_or_warning(username, "Username is required!");
        let check_password: boolean = check_or_warning(password, "Password is required!");
        let check_password_repeat: boolean = check_or_warning(password_repeat, "Password repeat is required!");

        if (check_email || check_username || check_password || check_password_repeat) {
            return;
        }

        let check_passwords_match: boolean = check_or_warning(password === password_repeat, "Passwords must match!");

        if (check_passwords_match) {
            return;
        }

        try {
            let response = await fetch("/api/auth/register", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ email, username, password })
            });

            if (!response.ok) {
                let data = await response.text();

                warning(data);

                return;
            }

            let data = await response.json();

            localStorage.setItem("token", data.token);

            success('Successful register!');
            push("/login")
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
                    <Label for="username">Email</Label>

                    <Input type="text" id="username" placeholder="Enter username" bind:value={username} />
                </FormGroup>

                <FormGroup>
                    <Label for="password">Password</Label>

                    <Input type="password" id="password" placeholder="Enter password" bind:value={password} />
                </FormGroup>

                <FormGroup>
                    <Label for="password_repeat">Repeat the password</Label>

                    <Input type="password" id="password_repeat" placeholder="Repeat the password" bind:value={password_repeat} />
                </FormGroup>
            </Form>

            <Button color="primary" block on:click={handleRegister}>Register</Button>
            <Button color="link" block on:click={() => push("/login")}>Already have an account?</Button>
        </CardBody>
    </Card>
</Container>
  