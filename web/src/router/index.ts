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
      {
        path: "/base64",
        name: "Base64",
        component: () => import("@/components/Base64/index.vue"),
        meta: {
          title: "Base64加解码",
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
  document.title = `Nic的私有工具 - ${to.meta.title}`;
  next();
});

export default router;
