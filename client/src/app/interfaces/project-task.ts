import { TaskState } from "src/common/task-state";

export interface ProjectTask {
    id: number;
    title: string;
    description: string;
    state: TaskState;
    creationDate: Date;
}
