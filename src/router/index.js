import { createRouter, createWebHistory } from 'vue-router';
import Home from '../components/Home.vue';
import Login from '../components/Login.vue';
import DisplayPage from '../components/Display/DisplayPage.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home,
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
  },
  {
    path: '/DisplayPage',
    name: 'DisplayPage',
    component: DisplayPage,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
