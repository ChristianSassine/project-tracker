import { LogType } from "src/common/log-type";

export interface HistoryLog {
    date: Date;
    logger: string;
    type: LogType;
    arguments?: string[];
}
