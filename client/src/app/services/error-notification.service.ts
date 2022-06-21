import { Injectable } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ErrorShown } from 'src/common/error-shown';

@Injectable({
    providedIn: 'root',
})
export class ErrorNotificationService {
    snackDuration: number;
    constructor(private snackBar: MatSnackBar) {
        this.snackDuration = 2000;
    }

    showError(errorShown: ErrorShown) {
        this.snackBar.open(errorShown, 'Close', { horizontalPosition: 'right', duration: this.snackDuration });
    }
}
