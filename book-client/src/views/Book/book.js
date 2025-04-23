import { convertDateToISO } from '@/utils/dateFormatter';

function getDefaultData() {
    const data = {};

    data.id = '';
    data.title = '';
    data.author = '';
    data.genre = '';
    data.description = '';
    data.isbn = '';
    data.publisher = '';
    data.publishedAt = null;

    return data;
}

function generateRequestData(book) {
    book.publishedAt = convertDateToISO(book.publishedAt);
    return book;
}

export { getDefaultData, generateRequestData };
