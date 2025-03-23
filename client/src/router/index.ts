import { createRouter, createWebHistory } from 'vue-router';

import HomeView from '@/views/home/HomeView.vue';
import TodayOrders from '@/views/today-orders/TodayOrders.vue';

const routes = [
  { path: '/', component: HomeView },
  { path: '/today-orders', component: TodayOrders },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
