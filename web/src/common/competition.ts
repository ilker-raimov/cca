import { error, warning } from "./toast";

export type Competition = {
    id: string;
    title: string;
    public: boolean;
    description: string;
    language: number;
    use_overall_time: boolean;
    use_execution_time: boolean;
};

export async function getCompetitionIdList(token: string): Promise<string[]> {
    try {
        const response: Response = await fetch("/api/competitions?all", {
            headers: { "Authorization": token }
        });

        if (!response.ok) {
            let data: string = await response.text();

            warning(data);

            return [];
        }

        return await response.json();
    } catch(err: any) {
        error(err);

        return [];
    }
}

export async function getCompetition(id: string, token: string): Promise<Competition | null> {
    try {
        const response: Response = await fetch(`/api/competitions/${id}`, {
            headers: { "Authorization": token }
        });

        if (!response.ok) {
            let data: string = await response.text();

            warning(data);

            return null;
        }

        return await response.json();
    } catch(err: any) {
        error(err);

        return null;
    }
}