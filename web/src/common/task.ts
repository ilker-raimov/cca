import { error, warning } from "./toast";

export type Task = {
    id: string,
    competition_id: string,
    name: string;
    description: string;
    execution_time: number;
    setup_code: string;
    user_code: string;
};

export async function getTaskIdList(competition_id: string, token: string): Promise<string[]> {
    try {
        const response: Response = await fetch(`/api/competition/${competition_id}/tasks`, {
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

export async function getTask(competition_id: string, id: string, token: string): Promise<Task | null> {
    try {
        const response: Response = await fetch(`/api/competition/${competition_id}/tasks/${id}`, {
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