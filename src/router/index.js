import { createRouter, createWebHistory } from 'vue-router';
import Home from '../components/Home.vue';
import Login from '../components/Login.vue';
import DisplayPage from '../components/Display/DisplayPage.vue'
import RegistrationForm from '@/components/Registration/RegistrationForm.vue';
import BookList from '@/views/BookList.vue';
import Dashboard from '../components/Dashboard/Dashboard.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home,
  },
  {
    path: '/dashboard',
    name: 'dashboard',
    component: Dashboard,
  },

  {
    path: '/DisplayPage',
    name: 'DisplayPage',
    component: DisplayPage,
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
  },
  {
    path: '/registration',
    name: 'registration',
    component: RegistrationForm,
  },
  {
    path: '/books',
    name: 'BooksList',  
    component: BookList,  
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
