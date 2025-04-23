import httpClient from '@/utils/http';

const bookRepository = {
    async getAll({ page = 1, limit = 10 }) {
        const queryString = { page, limit };
        return await httpClient.get('/books', { params: queryString });
    },
    async get(bookId) {
        return await httpClient.get(`/books/${bookId}`);
    },
    async add({ title, author, genre, description, isbn, publisher, publishedAt }) {
        return await httpClient.post('/books', { title, author, genre, description, isbn, publisher, publishedAt });
    },
    async update({ bookId, title, author, genre, description, isbn, publisher, publishedAt }) {
        return await httpClient.put(`/books/${bookId}`, { title, author, genre, description, genre, description, isbn, publisher, publishedAt });
    },
    async delete(bookId) {
        return await httpClient.delete(`/books/${bookId}`);
    },
    async reset() {
        return await httpClient.post(`/books/reset`);
    }
};

export default bookRepository;
