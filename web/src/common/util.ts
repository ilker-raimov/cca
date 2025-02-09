import { push } from "svelte-spa-router";
import CompetitorDashboard from "../page/CompetitorDashboard.svelte";
import OrganizerDashboard from "../page/OrganizerDashboard.svelte";
import AdministratorDashboard from "../page/AdministratorDashboard.svelte";
import { warning } from "./toast";

function check_token(redirect: string) {
    let token = sessionStorage.getItem("token");

    if (!token) {
        warning(`Missing token! Redirecting to: ${redirect}`)

        push(redirect)
    }
}

export const check_token_or_login = () => check_token("/login");

export function get_dashboard(role: number) {
    switch(role) {
        case 0: return CompetitorDashboard;
        case 1: return OrganizerDashboard;
        case 2: return AdministratorDashboard;
    }
}