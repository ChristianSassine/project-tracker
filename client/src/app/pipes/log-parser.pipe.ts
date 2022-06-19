import { Pipe, PipeTransform } from '@angular/core';
import { LogType } from 'src/common/log-type';

@Pipe({
    name: 'logParser',
})
export class LogParserPipe implements PipeTransform {
    transform(value: LogType, args: string[]): string {
        return this.parseLog(value, args);
    }

    parseLog(type: LogType, args: string[]): string {
        switch (type) {
            case LogType.ProjectCreated:
                return `Created the project <b>${args[0]}</b>`;
            case LogType.ProjectDeleted:
                return `Deleted the project <b>${args[0]}</b>`;
            case LogType.ProjectJoined:
                return `Joined the Project!`;
            case LogType.TaskCreation:
                return `Created a task with the title <b>${args[0]}</b>`;
            case LogType.TaskDeleted:
                return `Deleted a task with the title <b>${args[0]}</b>`;
            case LogType.TaskDescriptionModification:
                return `Modified the description of <b>${args[0]}</b> titled task`;
            case LogType.TaskStateModification:
                return `Modified the state of the task <b>${args[0]}</b> from <b>${args[1]}</b> to <b>${args[2]}</b>`;
            case LogType.TaskTitleModification:
                return ' <b> </b>';
            default:
                return type;
        }
    }
}
