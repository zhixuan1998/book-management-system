import moment from 'moment/moment';

function convertDateToFriendlyDisplay(date) {
    return moment(date).format('MMMM Do YYYY');
}

function convertDateToDatePickerFormat(date) {
    return moment(date).format('YYYY-MM-DD');
}

function convertDateToISO(date) {
    return moment(date).format();
}

export { convertDateToFriendlyDisplay, convertDateToDatePickerFormat, convertDateToISO };
