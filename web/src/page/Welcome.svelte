<script lang="ts">
    import { push } from "svelte-spa-router";
    import { Container, Card, CardBody, CardTitle, Button } from "sveltestrap";

    let token_present: boolean = sessionStorage.getItem("token") !== null;

    if (token_present) {
        push("/dashboard")
    }
</script>

<Container class="d-flex justify-content-center align-items-center vh-100">
    <Card class="p-4 shadow-lg text-center" style="max-width: 500px; width: 100%;">
        <CardBody>
            <CardTitle>
                <h2>Welcome to CodeArena</h2>
            </CardTitle>
            <p>
                Here, you can <strong>organize</strong> coding challenges, <strong>compete</strong> with others,  
                and <strong>learn</strong> from a growing community of developers.
            </p>
            <p>
                Whether you're a beginner looking to improve your skills or an expert ready to compete,  
                CodeArena is the place for you.
            </p>

            <h4 class="mt-4">Wanna join?</h4>
            <div class="d-flex justify-content-center gap-3 mt-3">
                <Button color="primary" on:click={() => push("/login")}>Login</Button>
                <Button color="success" on:click={() => push("/register")}>Register</Button>
            </div>
        </CardBody>
    </Card>
</Container>
