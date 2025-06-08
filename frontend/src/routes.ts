import type { RouteDefinition } from '@solidjs/router';


const routes: RouteDefinition[] = [
  {
    path: '/login',
    component: Login,
  },
  {
    path: '/admin',
    component: LayoutAdmin,
    children: [
      {
        path: '/',
        component: Dashboard,
      }
    ]
  }
];

export default routes;
