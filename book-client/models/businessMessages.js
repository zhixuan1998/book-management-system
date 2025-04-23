export const messages = {
    general: {
        actions: () => 'Actions',
        successful: () => 'Successful',
        addSuccessful: () => 'Successfully added!',
        updateSuccessful: () => 'Successfully updated!',
        deleteConfirmation: () => 'Delete Confirmation'
    },
    book: {
        header: () => 'Book Management System',
        bookId: () => 'Book ID',
        title: () => 'Title',
        author: () => 'Author',
        genre: () => 'Genre',
        description: () => 'Description',
        isbn: () => 'ISBN',
        publishedAt: () => 'Published At',
        publisher: () => 'Publisher',
        actions: () => 'Actions',
        deleteMessage: (bookId = 'this') => `Are you sure you want to delete ${bookId} book?`,
        resetWarningMessage: () => "Are you sure you want reset the book list to it's initial state? This action is irreversible!"
    },
    button: {
        ok: () => 'Ok',
        back: () => 'Back',
        edit: () => 'Edit',
        view: () => 'View',
        save: () => 'Save',
        reset: () => 'Reset',
        cancel: () => 'Cancel',
        delete: () => 'Delete',
        addNew: () => 'Add New',
        confirm: () => 'Confirm',
        default: () => 'Button'
    },
    error: {
        oops: () => 'Oops!',
        beCareful: () => 'Be Careful!',
        general: () => 'Something went wrong, please try again later.',
        allFieldRequired: () => 'Please fill in all the fields!'
    }
};
