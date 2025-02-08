import { push } from "svelte-spa-router";

function check_token(redirect: string) {
    let token = localStorage.getItem("token");

    if (!token) {
        push(redirect)
    }
}

export const check_token_or_login = () => check_token("/login");