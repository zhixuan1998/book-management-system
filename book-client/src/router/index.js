import { createRouter, createWebHistory } from 'vue-router'

const BookAdd = () => import('@/views/Book/BookAdd.vue')
const BookDetail = () => import('@/views/Book/BookDetail.vue')
const BookListing = () => import('@/views/Book/BookListing.vue')

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: { path: '/books'},
      children: [
        {
          path: '/books',
          name: 'BookListing',
          component: BookListing,
        },
        {
          path: '/books/add',
          name: 'BookAdd',
          component: BookAdd,
        },
        {
          path: '/books/:bookId/detail',
          name: 'BookDetail',
          component: BookDetail,
        },
      ],
    },
  ],
})

export default router
