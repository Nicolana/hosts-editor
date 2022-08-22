import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

export const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "/",
    meta: {
      keepAlive: true,
      requireAuth: true,
    },
    component: () => import("@/components/Layout.vue"),
    children: [
      {
        path: "/",
        name: "Hosts",
        component: () => import("@/components/Hosts/Editor.vue"),
        meta: {
          title: "Hosts管理",
          keepAlive: true,
        },
      },
      {
        path: "/frp",
        name: "Projects",
        component: () => import("@/components/FrpClient/FrpList.vue"),
        meta: {
          title: "Frp端口映射管理",
          keepAlive: true,
        },
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  document.title = to.meta.title as string;
  next();
});

export default router;
